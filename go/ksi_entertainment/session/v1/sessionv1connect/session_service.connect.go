// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ksi_entertainment/session/v1/session_service.proto

package sessionv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/kodmm/ksi-entertainment-proto/go/ksi_entertainment/session/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SessionServiceName is the fully-qualified name of the SessionService service.
	SessionServiceName = "ksi_entertainment.session.v1.SessionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SessionServiceCreateSessionProcedure is the fully-qualified name of the SessionService's
	// CreateSession RPC.
	SessionServiceCreateSessionProcedure = "/ksi_entertainment.session.v1.SessionService/CreateSession"
)

// SessionServiceClient is a client for the ksi_entertainment.session.v1.SessionService service.
type SessionServiceClient interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
}

// NewSessionServiceClient constructs a client for the ksi_entertainment.session.v1.SessionService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSessionServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SessionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sessionServiceClient{
		createSession: connect_go.NewClient[v1.CreateSessionRequest, v1.CreateSessionResponse](
			httpClient,
			baseURL+SessionServiceCreateSessionProcedure,
			opts...,
		),
	}
}

// sessionServiceClient implements SessionServiceClient.
type sessionServiceClient struct {
	createSession *connect_go.Client[v1.CreateSessionRequest, v1.CreateSessionResponse]
}

// CreateSession calls ksi_entertainment.session.v1.SessionService.CreateSession.
func (c *sessionServiceClient) CreateSession(ctx context.Context, req *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return c.createSession.CallUnary(ctx, req)
}

// SessionServiceHandler is an implementation of the ksi_entertainment.session.v1.SessionService
// service.
type SessionServiceHandler interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
}

// NewSessionServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSessionServiceHandler(svc SessionServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(SessionServiceCreateSessionProcedure, connect_go.NewUnaryHandler(
		SessionServiceCreateSessionProcedure,
		svc.CreateSession,
		opts...,
	))
	return "/ksi_entertainment.session.v1.SessionService/", mux
}

// UnimplementedSessionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSessionServiceHandler struct{}

func (UnimplementedSessionServiceHandler) CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ksi_entertainment.session.v1.SessionService.CreateSession is not implemented"))
}
