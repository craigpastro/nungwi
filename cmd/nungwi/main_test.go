package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/cenkalti/backoff/v4"
	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha/nungwiv1alphaconnect"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type checkTests struct {
	Tests []checkTest
}

type checkTest struct {
	Name       string
	Schema     []*pb.RelationConfig
	Tuples     []*pb.Tuple
	Assertions []assertion
}

type assertion struct {
	Tuple       *pb.Tuple
	Expectation bool
}

func TestCheck(t *testing.T) {
	f, _ := filepath.Abs(filepath.Join("..", "..", "testdata", "check.yaml"))
	data, err := os.ReadFile(f)
	require.NoError(t, err)

	var tests checkTests
	err = yaml.Unmarshal(data, &tests)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get a random port
	listener, _ := net.Listen("tcp", "")
	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	cfg := &config{Addr: fmt.Sprintf(":%d", port)}

	// Start the server
	go run(ctx, cfg, zap.NewNop())

	client := nungwiv1alphaconnect.NewNungwiServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://localhost:%d", port),
	)

	policy := backoff.NewExponentialBackOff()
	policy.MaxElapsedTime = 10 * time.Second
	err = backoff.Retry(func() error {
		_, err = client.GetSchema(ctx, &connect.Request[pb.GetSchemaRequest]{
			Msg: &pb.GetSchemaRequest{},
		})
		return err
	}, policy)
	if err != nil {
		panic(err)
	}

	for _, test := range tests.Tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err = client.WriteSchema(ctx, &connect.Request[pb.WriteSchemaRequest]{
				Msg: &pb.WriteSchemaRequest{
					Configs: test.Schema,
				},
			})
			require.NoError(t, err)

			_, err = client.WriteTuples(ctx, &connect.Request[pb.WriteTuplesRequest]{
				Msg: &pb.WriteTuplesRequest{
					Tuples: test.Tuples,
				},
			})
			require.NoError(t, err)

			for _, assertion := range test.Assertions {
				resp, err := client.Check(ctx, &connect.Request[pb.CheckRequest]{
					Msg: &pb.CheckRequest{
						Namespace: assertion.Tuple.Namespace,
						Id:        assertion.Tuple.Id,
						Relation:  assertion.Tuple.Relation,
						User:      assertion.Tuple.User,
					},
				})
				require.NoError(t, err)
				require.Equal(t, assertion.Expectation, resp.Msg.Allowed)
			}

			for _, tuple := range test.Tuples {
				_, err = client.DeleteTuples(ctx, &connect.Request[pb.DeleteTuplesRequest]{
					Msg: &pb.DeleteTuplesRequest{
						Tuples: []*pb.Tuple{tuple},
					},
				})
				require.NoError(t, err)
			}

			_, err = client.DeleteSchema(ctx, &connect.Request[pb.DeleteSchemaRequest]{
				Msg: &pb.DeleteSchemaRequest{},
			})
			require.NoError(t, err)
		})
	}
}
