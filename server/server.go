package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha/nungwiv1alphaconnect"
	"github.com/craigpastro/nungwi/internal/interpreter"
	"go.uber.org/zap"
)

type server struct {
	nungwiv1alphaconnect.UnimplementedNungwiServiceHandler

	Interpreter *interpreter.Interpreter
	Logger      *zap.Logger
}

func NewServer(interpreter *interpreter.Interpreter, logger *zap.Logger) *server {
	return &server{
		Interpreter: interpreter,
		Logger:      logger,
	}
}

func (s *server) WriteSchema(ctx context.Context, req *connect.Request[pb.WriteSchemaRequest]) (*connect.Response[pb.WriteSchemaResponse], error) {
	return connect.NewResponse(&pb.WriteSchemaResponse{}), nil
}

func (s *server) GetSchema(ctx context.Context, req *connect.Request[pb.GetSchemaRequest]) (*connect.Response[pb.GetSchemaResponse], error) {
	return connect.NewResponse(&pb.GetSchemaResponse{
		Namespaces: nil,
	}), nil
}

func (s *server) DeleteSchema(ctx context.Context, req *connect.Request[pb.DeleteSchemaRequest]) (*connect.Response[pb.DeleteSchemaResponse], error) {
	return connect.NewResponse(&pb.DeleteSchemaResponse{}), nil
}

func (s *server) AddTuples(ctx context.Context, req *connect.Request[pb.AddTuplesRequest]) (*connect.Response[pb.AddTuplesResponse], error) {
	if err := s.Interpreter.AddTuples(ctx, req.Msg.GetTuples()); err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.AddTuplesResponse{}), nil
}

func (s *server) DeleteTuples(ctx context.Context, req *connect.Request[pb.DeleteTuplesRequest]) (*connect.Response[pb.DeleteTuplesResponse], error) {
	if err := s.Interpreter.DeleteTuples(ctx, req.Msg.GetTuples()); err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.DeleteTuplesResponse{}), nil
}

func (s *server) Checks(ctx context.Context, req *connect.Request[pb.ChecksRequest]) (*connect.Response[pb.ChecksResponse], error) {
	return connect.NewResponse(&pb.ChecksResponse{
		Results: nil,
	}), nil
}
