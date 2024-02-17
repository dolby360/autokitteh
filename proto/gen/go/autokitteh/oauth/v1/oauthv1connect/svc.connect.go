// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: autokitteh/oauth/v1/svc.proto

package oauthv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/oauth/v1"
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
	// OAuthServiceName is the fully-qualified name of the OAuthService service.
	OAuthServiceName = "autokitteh.oauth.v1.OAuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OAuthServiceRegisterProcedure is the fully-qualified name of the OAuthService's Register RPC.
	OAuthServiceRegisterProcedure = "/autokitteh.oauth.v1.OAuthService/Register"
	// OAuthServiceGetProcedure is the fully-qualified name of the OAuthService's Get RPC.
	OAuthServiceGetProcedure = "/autokitteh.oauth.v1.OAuthService/Get"
	// OAuthServiceStartFlowProcedure is the fully-qualified name of the OAuthService's StartFlow RPC.
	OAuthServiceStartFlowProcedure = "/autokitteh.oauth.v1.OAuthService/StartFlow"
	// OAuthServiceExchangeProcedure is the fully-qualified name of the OAuthService's Exchange RPC.
	OAuthServiceExchangeProcedure = "/autokitteh.oauth.v1.OAuthService/Exchange"
)

// OAuthServiceClient is a client for the autokitteh.oauth.v1.OAuthService service.
type OAuthServiceClient interface {
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	StartFlow(context.Context, *connect.Request[v1.StartFlowRequest]) (*connect.Response[v1.StartFlowResponse], error)
	Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)
}

// NewOAuthServiceClient constructs a client for the autokitteh.oauth.v1.OAuthService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OAuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &oAuthServiceClient{
		register: connect.NewClient[v1.RegisterRequest, v1.RegisterResponse](
			httpClient,
			baseURL+OAuthServiceRegisterProcedure,
			opts...,
		),
		get: connect.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+OAuthServiceGetProcedure,
			opts...,
		),
		startFlow: connect.NewClient[v1.StartFlowRequest, v1.StartFlowResponse](
			httpClient,
			baseURL+OAuthServiceStartFlowProcedure,
			opts...,
		),
		exchange: connect.NewClient[v1.ExchangeRequest, v1.ExchangeResponse](
			httpClient,
			baseURL+OAuthServiceExchangeProcedure,
			opts...,
		),
	}
}

// oAuthServiceClient implements OAuthServiceClient.
type oAuthServiceClient struct {
	register  *connect.Client[v1.RegisterRequest, v1.RegisterResponse]
	get       *connect.Client[v1.GetRequest, v1.GetResponse]
	startFlow *connect.Client[v1.StartFlowRequest, v1.StartFlowResponse]
	exchange  *connect.Client[v1.ExchangeRequest, v1.ExchangeResponse]
}

// Register calls autokitteh.oauth.v1.OAuthService.Register.
func (c *oAuthServiceClient) Register(ctx context.Context, req *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Get calls autokitteh.oauth.v1.OAuthService.Get.
func (c *oAuthServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// StartFlow calls autokitteh.oauth.v1.OAuthService.StartFlow.
func (c *oAuthServiceClient) StartFlow(ctx context.Context, req *connect.Request[v1.StartFlowRequest]) (*connect.Response[v1.StartFlowResponse], error) {
	return c.startFlow.CallUnary(ctx, req)
}

// Exchange calls autokitteh.oauth.v1.OAuthService.Exchange.
func (c *oAuthServiceClient) Exchange(ctx context.Context, req *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	return c.exchange.CallUnary(ctx, req)
}

// OAuthServiceHandler is an implementation of the autokitteh.oauth.v1.OAuthService service.
type OAuthServiceHandler interface {
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	StartFlow(context.Context, *connect.Request[v1.StartFlowRequest]) (*connect.Response[v1.StartFlowResponse], error)
	Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)
}

// NewOAuthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOAuthServiceHandler(svc OAuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	oAuthServiceRegisterHandler := connect.NewUnaryHandler(
		OAuthServiceRegisterProcedure,
		svc.Register,
		opts...,
	)
	oAuthServiceGetHandler := connect.NewUnaryHandler(
		OAuthServiceGetProcedure,
		svc.Get,
		opts...,
	)
	oAuthServiceStartFlowHandler := connect.NewUnaryHandler(
		OAuthServiceStartFlowProcedure,
		svc.StartFlow,
		opts...,
	)
	oAuthServiceExchangeHandler := connect.NewUnaryHandler(
		OAuthServiceExchangeProcedure,
		svc.Exchange,
		opts...,
	)
	return "/autokitteh.oauth.v1.OAuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OAuthServiceRegisterProcedure:
			oAuthServiceRegisterHandler.ServeHTTP(w, r)
		case OAuthServiceGetProcedure:
			oAuthServiceGetHandler.ServeHTTP(w, r)
		case OAuthServiceStartFlowProcedure:
			oAuthServiceStartFlowHandler.ServeHTTP(w, r)
		case OAuthServiceExchangeProcedure:
			oAuthServiceExchangeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOAuthServiceHandler struct{}

func (UnimplementedOAuthServiceHandler) Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.oauth.v1.OAuthService.Register is not implemented"))
}

func (UnimplementedOAuthServiceHandler) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.oauth.v1.OAuthService.Get is not implemented"))
}

func (UnimplementedOAuthServiceHandler) StartFlow(context.Context, *connect.Request[v1.StartFlowRequest]) (*connect.Response[v1.StartFlowResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.oauth.v1.OAuthService.StartFlow is not implemented"))
}

func (UnimplementedOAuthServiceHandler) Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.oauth.v1.OAuthService.Exchange is not implemented"))
}
