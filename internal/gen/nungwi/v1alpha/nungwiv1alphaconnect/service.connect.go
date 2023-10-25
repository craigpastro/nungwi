// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: nungwi/v1alpha/service.proto

package nungwiv1alphaconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// NungwiServiceName is the fully-qualified name of the NungwiService service.
	NungwiServiceName = "nungwi.v1alpha.NungwiService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NungwiServiceWriteSchemaProcedure is the fully-qualified name of the NungwiService's WriteSchema
	// RPC.
	NungwiServiceWriteSchemaProcedure = "/nungwi.v1alpha.NungwiService/WriteSchema"
	// NungwiServiceGetSchemaProcedure is the fully-qualified name of the NungwiService's GetSchema RPC.
	NungwiServiceGetSchemaProcedure = "/nungwi.v1alpha.NungwiService/GetSchema"
	// NungwiServiceDeleteSchemaProcedure is the fully-qualified name of the NungwiService's
	// DeleteSchema RPC.
	NungwiServiceDeleteSchemaProcedure = "/nungwi.v1alpha.NungwiService/DeleteSchema"
	// NungwiServiceWriteTuplesProcedure is the fully-qualified name of the NungwiService's WriteTuples
	// RPC.
	NungwiServiceWriteTuplesProcedure = "/nungwi.v1alpha.NungwiService/WriteTuples"
	// NungwiServiceGetTuplesProcedure is the fully-qualified name of the NungwiService's GetTuples RPC.
	NungwiServiceGetTuplesProcedure = "/nungwi.v1alpha.NungwiService/GetTuples"
	// NungwiServiceDeleteTuplesProcedure is the fully-qualified name of the NungwiService's
	// DeleteTuples RPC.
	NungwiServiceDeleteTuplesProcedure = "/nungwi.v1alpha.NungwiService/DeleteTuples"
	// NungwiServiceCheckProcedure is the fully-qualified name of the NungwiService's Check RPC.
	NungwiServiceCheckProcedure = "/nungwi.v1alpha.NungwiService/Check"
	// NungwiServiceListObjectsProcedure is the fully-qualified name of the NungwiService's ListObjects
	// RPC.
	NungwiServiceListObjectsProcedure = "/nungwi.v1alpha.NungwiService/ListObjects"
)

// NungwiServiceClient is a client for the nungwi.v1alpha.NungwiService service.
type NungwiServiceClient interface {
	WriteSchema(context.Context, *connect.Request[v1alpha.WriteSchemaRequest]) (*connect.Response[v1alpha.WriteSchemaResponse], error)
	GetSchema(context.Context, *connect.Request[v1alpha.GetSchemaRequest]) (*connect.Response[v1alpha.GetSchemaResponse], error)
	DeleteSchema(context.Context, *connect.Request[v1alpha.DeleteSchemaRequest]) (*connect.Response[v1alpha.DeleteSchemaResponse], error)
	WriteTuples(context.Context, *connect.Request[v1alpha.WriteTuplesRequest]) (*connect.Response[v1alpha.WriteTuplesResponse], error)
	GetTuples(context.Context, *connect.Request[v1alpha.GetTuplesRequest]) (*connect.Response[v1alpha.GetTuplesResponse], error)
	DeleteTuples(context.Context, *connect.Request[v1alpha.DeleteTuplesRequest]) (*connect.Response[v1alpha.DeleteTuplesResponse], error)
	Check(context.Context, *connect.Request[v1alpha.CheckRequest]) (*connect.Response[v1alpha.CheckResponse], error)
	ListObjects(context.Context, *connect.Request[v1alpha.ListObjectsRequest]) (*connect.Response[v1alpha.ListObjectsResponse], error)
}

// NewNungwiServiceClient constructs a client for the nungwi.v1alpha.NungwiService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNungwiServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NungwiServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &nungwiServiceClient{
		writeSchema: connect.NewClient[v1alpha.WriteSchemaRequest, v1alpha.WriteSchemaResponse](
			httpClient,
			baseURL+NungwiServiceWriteSchemaProcedure,
			opts...,
		),
		getSchema: connect.NewClient[v1alpha.GetSchemaRequest, v1alpha.GetSchemaResponse](
			httpClient,
			baseURL+NungwiServiceGetSchemaProcedure,
			opts...,
		),
		deleteSchema: connect.NewClient[v1alpha.DeleteSchemaRequest, v1alpha.DeleteSchemaResponse](
			httpClient,
			baseURL+NungwiServiceDeleteSchemaProcedure,
			opts...,
		),
		writeTuples: connect.NewClient[v1alpha.WriteTuplesRequest, v1alpha.WriteTuplesResponse](
			httpClient,
			baseURL+NungwiServiceWriteTuplesProcedure,
			opts...,
		),
		getTuples: connect.NewClient[v1alpha.GetTuplesRequest, v1alpha.GetTuplesResponse](
			httpClient,
			baseURL+NungwiServiceGetTuplesProcedure,
			opts...,
		),
		deleteTuples: connect.NewClient[v1alpha.DeleteTuplesRequest, v1alpha.DeleteTuplesResponse](
			httpClient,
			baseURL+NungwiServiceDeleteTuplesProcedure,
			opts...,
		),
		check: connect.NewClient[v1alpha.CheckRequest, v1alpha.CheckResponse](
			httpClient,
			baseURL+NungwiServiceCheckProcedure,
			opts...,
		),
		listObjects: connect.NewClient[v1alpha.ListObjectsRequest, v1alpha.ListObjectsResponse](
			httpClient,
			baseURL+NungwiServiceListObjectsProcedure,
			opts...,
		),
	}
}

// nungwiServiceClient implements NungwiServiceClient.
type nungwiServiceClient struct {
	writeSchema  *connect.Client[v1alpha.WriteSchemaRequest, v1alpha.WriteSchemaResponse]
	getSchema    *connect.Client[v1alpha.GetSchemaRequest, v1alpha.GetSchemaResponse]
	deleteSchema *connect.Client[v1alpha.DeleteSchemaRequest, v1alpha.DeleteSchemaResponse]
	writeTuples  *connect.Client[v1alpha.WriteTuplesRequest, v1alpha.WriteTuplesResponse]
	getTuples    *connect.Client[v1alpha.GetTuplesRequest, v1alpha.GetTuplesResponse]
	deleteTuples *connect.Client[v1alpha.DeleteTuplesRequest, v1alpha.DeleteTuplesResponse]
	check        *connect.Client[v1alpha.CheckRequest, v1alpha.CheckResponse]
	listObjects  *connect.Client[v1alpha.ListObjectsRequest, v1alpha.ListObjectsResponse]
}

// WriteSchema calls nungwi.v1alpha.NungwiService.WriteSchema.
func (c *nungwiServiceClient) WriteSchema(ctx context.Context, req *connect.Request[v1alpha.WriteSchemaRequest]) (*connect.Response[v1alpha.WriteSchemaResponse], error) {
	return c.writeSchema.CallUnary(ctx, req)
}

// GetSchema calls nungwi.v1alpha.NungwiService.GetSchema.
func (c *nungwiServiceClient) GetSchema(ctx context.Context, req *connect.Request[v1alpha.GetSchemaRequest]) (*connect.Response[v1alpha.GetSchemaResponse], error) {
	return c.getSchema.CallUnary(ctx, req)
}

// DeleteSchema calls nungwi.v1alpha.NungwiService.DeleteSchema.
func (c *nungwiServiceClient) DeleteSchema(ctx context.Context, req *connect.Request[v1alpha.DeleteSchemaRequest]) (*connect.Response[v1alpha.DeleteSchemaResponse], error) {
	return c.deleteSchema.CallUnary(ctx, req)
}

// WriteTuples calls nungwi.v1alpha.NungwiService.WriteTuples.
func (c *nungwiServiceClient) WriteTuples(ctx context.Context, req *connect.Request[v1alpha.WriteTuplesRequest]) (*connect.Response[v1alpha.WriteTuplesResponse], error) {
	return c.writeTuples.CallUnary(ctx, req)
}

// GetTuples calls nungwi.v1alpha.NungwiService.GetTuples.
func (c *nungwiServiceClient) GetTuples(ctx context.Context, req *connect.Request[v1alpha.GetTuplesRequest]) (*connect.Response[v1alpha.GetTuplesResponse], error) {
	return c.getTuples.CallUnary(ctx, req)
}

// DeleteTuples calls nungwi.v1alpha.NungwiService.DeleteTuples.
func (c *nungwiServiceClient) DeleteTuples(ctx context.Context, req *connect.Request[v1alpha.DeleteTuplesRequest]) (*connect.Response[v1alpha.DeleteTuplesResponse], error) {
	return c.deleteTuples.CallUnary(ctx, req)
}

// Check calls nungwi.v1alpha.NungwiService.Check.
func (c *nungwiServiceClient) Check(ctx context.Context, req *connect.Request[v1alpha.CheckRequest]) (*connect.Response[v1alpha.CheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// ListObjects calls nungwi.v1alpha.NungwiService.ListObjects.
func (c *nungwiServiceClient) ListObjects(ctx context.Context, req *connect.Request[v1alpha.ListObjectsRequest]) (*connect.Response[v1alpha.ListObjectsResponse], error) {
	return c.listObjects.CallUnary(ctx, req)
}

// NungwiServiceHandler is an implementation of the nungwi.v1alpha.NungwiService service.
type NungwiServiceHandler interface {
	WriteSchema(context.Context, *connect.Request[v1alpha.WriteSchemaRequest]) (*connect.Response[v1alpha.WriteSchemaResponse], error)
	GetSchema(context.Context, *connect.Request[v1alpha.GetSchemaRequest]) (*connect.Response[v1alpha.GetSchemaResponse], error)
	DeleteSchema(context.Context, *connect.Request[v1alpha.DeleteSchemaRequest]) (*connect.Response[v1alpha.DeleteSchemaResponse], error)
	WriteTuples(context.Context, *connect.Request[v1alpha.WriteTuplesRequest]) (*connect.Response[v1alpha.WriteTuplesResponse], error)
	GetTuples(context.Context, *connect.Request[v1alpha.GetTuplesRequest]) (*connect.Response[v1alpha.GetTuplesResponse], error)
	DeleteTuples(context.Context, *connect.Request[v1alpha.DeleteTuplesRequest]) (*connect.Response[v1alpha.DeleteTuplesResponse], error)
	Check(context.Context, *connect.Request[v1alpha.CheckRequest]) (*connect.Response[v1alpha.CheckResponse], error)
	ListObjects(context.Context, *connect.Request[v1alpha.ListObjectsRequest]) (*connect.Response[v1alpha.ListObjectsResponse], error)
}

// NewNungwiServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNungwiServiceHandler(svc NungwiServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	nungwiServiceWriteSchemaHandler := connect.NewUnaryHandler(
		NungwiServiceWriteSchemaProcedure,
		svc.WriteSchema,
		opts...,
	)
	nungwiServiceGetSchemaHandler := connect.NewUnaryHandler(
		NungwiServiceGetSchemaProcedure,
		svc.GetSchema,
		opts...,
	)
	nungwiServiceDeleteSchemaHandler := connect.NewUnaryHandler(
		NungwiServiceDeleteSchemaProcedure,
		svc.DeleteSchema,
		opts...,
	)
	nungwiServiceWriteTuplesHandler := connect.NewUnaryHandler(
		NungwiServiceWriteTuplesProcedure,
		svc.WriteTuples,
		opts...,
	)
	nungwiServiceGetTuplesHandler := connect.NewUnaryHandler(
		NungwiServiceGetTuplesProcedure,
		svc.GetTuples,
		opts...,
	)
	nungwiServiceDeleteTuplesHandler := connect.NewUnaryHandler(
		NungwiServiceDeleteTuplesProcedure,
		svc.DeleteTuples,
		opts...,
	)
	nungwiServiceCheckHandler := connect.NewUnaryHandler(
		NungwiServiceCheckProcedure,
		svc.Check,
		opts...,
	)
	nungwiServiceListObjectsHandler := connect.NewUnaryHandler(
		NungwiServiceListObjectsProcedure,
		svc.ListObjects,
		opts...,
	)
	return "/nungwi.v1alpha.NungwiService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NungwiServiceWriteSchemaProcedure:
			nungwiServiceWriteSchemaHandler.ServeHTTP(w, r)
		case NungwiServiceGetSchemaProcedure:
			nungwiServiceGetSchemaHandler.ServeHTTP(w, r)
		case NungwiServiceDeleteSchemaProcedure:
			nungwiServiceDeleteSchemaHandler.ServeHTTP(w, r)
		case NungwiServiceWriteTuplesProcedure:
			nungwiServiceWriteTuplesHandler.ServeHTTP(w, r)
		case NungwiServiceGetTuplesProcedure:
			nungwiServiceGetTuplesHandler.ServeHTTP(w, r)
		case NungwiServiceDeleteTuplesProcedure:
			nungwiServiceDeleteTuplesHandler.ServeHTTP(w, r)
		case NungwiServiceCheckProcedure:
			nungwiServiceCheckHandler.ServeHTTP(w, r)
		case NungwiServiceListObjectsProcedure:
			nungwiServiceListObjectsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNungwiServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNungwiServiceHandler struct{}

func (UnimplementedNungwiServiceHandler) WriteSchema(context.Context, *connect.Request[v1alpha.WriteSchemaRequest]) (*connect.Response[v1alpha.WriteSchemaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.WriteSchema is not implemented"))
}

func (UnimplementedNungwiServiceHandler) GetSchema(context.Context, *connect.Request[v1alpha.GetSchemaRequest]) (*connect.Response[v1alpha.GetSchemaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.GetSchema is not implemented"))
}

func (UnimplementedNungwiServiceHandler) DeleteSchema(context.Context, *connect.Request[v1alpha.DeleteSchemaRequest]) (*connect.Response[v1alpha.DeleteSchemaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.DeleteSchema is not implemented"))
}

func (UnimplementedNungwiServiceHandler) WriteTuples(context.Context, *connect.Request[v1alpha.WriteTuplesRequest]) (*connect.Response[v1alpha.WriteTuplesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.WriteTuples is not implemented"))
}

func (UnimplementedNungwiServiceHandler) GetTuples(context.Context, *connect.Request[v1alpha.GetTuplesRequest]) (*connect.Response[v1alpha.GetTuplesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.GetTuples is not implemented"))
}

func (UnimplementedNungwiServiceHandler) DeleteTuples(context.Context, *connect.Request[v1alpha.DeleteTuplesRequest]) (*connect.Response[v1alpha.DeleteTuplesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.DeleteTuples is not implemented"))
}

func (UnimplementedNungwiServiceHandler) Check(context.Context, *connect.Request[v1alpha.CheckRequest]) (*connect.Response[v1alpha.CheckResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.Check is not implemented"))
}

func (UnimplementedNungwiServiceHandler) ListObjects(context.Context, *connect.Request[v1alpha.ListObjectsRequest]) (*connect.Response[v1alpha.ListObjectsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nungwi.v1alpha.NungwiService.ListObjects is not implemented"))
}
