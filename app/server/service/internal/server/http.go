package server

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/conf"
	"Ali-DDNS/pkg"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type HttpServer struct {
	Mux        *runtime.ServeMux
	GRPCCancel context.CancelFunc
}

func NewInterfaceHTTPServer() (*HttpServer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	var opts []grpc.DialOption

	// get the creds then append into the grpc options
	creds, err := pkg.GetClientCreds()
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	mux := runtime.NewServeMux()

	if err := v1.RegisterDDNSInterfaceHandlerFromEndpoint(ctx, mux, "localhost:"+conf.Basic().InterfaceGrpcPort(), opts); err != nil {
		return nil, err
	}

	return &HttpServer{
		Mux:        mux,
		GRPCCancel: cancel,
	}, err
}
