package server

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha/nungwiv1alphaconnect"
	"github.com/craigpastro/nungwi/internal/prolog"
	"go.uber.org/zap"
)

type server struct {
	nungwiv1alphaconnect.UnimplementedNungwiServiceHandler

	prolog *prolog.Prolog
	logger *zap.Logger
}

func NewServer(prolog *prolog.Prolog, logger *zap.Logger) *server {
	return &server{
		prolog: prolog,
		logger: logger,
	}
}

func (s *server) WriteSchema(ctx context.Context, req *connect.Request[pb.WriteSchemaRequest]) (*connect.Response[pb.WriteSchemaResponse], error) {
	schema, err := s.prolog.GetSchema(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if len(schema) > 0 {
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("schema already exists"))
	}

	if err := s.prolog.WriteSchema(ctx, req.Msg.GetConfigs()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.WriteSchemaResponse{}), nil
}

func (s *server) GetSchema(ctx context.Context, req *connect.Request[pb.GetSchemaRequest]) (*connect.Response[pb.GetSchemaResponse], error) {
	configs, err := s.prolog.GetSchema(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetSchemaResponse{
		Configs: configs,
	}), nil
}

func (s *server) DeleteSchema(ctx context.Context, req *connect.Request[pb.DeleteSchemaRequest]) (*connect.Response[pb.DeleteSchemaResponse], error) {
	tuples, err := s.prolog.GetTuples(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if len(tuples) > 0 {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("please remove all tuples before deleting schema"))
	}

	if err := s.prolog.DeleteSchema(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteSchemaResponse{}), nil
}

func (s *server) WriteTuples(ctx context.Context, req *connect.Request[pb.WriteTuplesRequest]) (*connect.Response[pb.WriteTuplesResponse], error) {
	if err := s.prolog.WriteTuples(ctx, req.Msg.GetTuples()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.WriteTuplesResponse{}), nil
}

func (s *server) GetTuples(ctx context.Context, req *connect.Request[pb.GetTuplesRequest]) (*connect.Response[pb.GetTuplesResponse], error) {
	tuples, err := s.prolog.GetTuples(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetTuplesResponse{
		Tuples: tuples,
	}), nil
}

func (s *server) DeleteTuples(ctx context.Context, req *connect.Request[pb.DeleteTuplesRequest]) (*connect.Response[pb.DeleteTuplesResponse], error) {
	if err := s.prolog.DeleteTuples(ctx, req.Msg.GetTuples()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteTuplesResponse{}), nil
}

func (s *server) Check(ctx context.Context, req *connect.Request[pb.CheckRequest]) (*connect.Response[pb.CheckResponse], error) {
	msg := req.Msg
	allowed, err := s.prolog.Check(ctx, &pb.Tuple{
		Namespace: msg.GetNamespace(),
		Id:        msg.GetId(),
		Relation:  msg.GetRelation(),
		User:      msg.GetUser(),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.CheckResponse{
		Allowed: allowed,
	}), nil
}

func (s *server) ListObjects(ctx context.Context, req *connect.Request[pb.ListObjectsRequest]) (*connect.Response[pb.ListObjectsResponse], error) {
	msg := req.Msg
	ids, err := s.prolog.ListObjects(ctx, msg.GetNamespace(), msg.GetRelation(), msg.GetUser())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.ListObjectsResponse{
		Ids: ids,
	}), nil
}
