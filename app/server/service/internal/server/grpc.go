package server

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/service"
	"Ali-DDNS/pkg"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

type DomainServer struct {
	*grpc.Server
}

// NewDomainGRPCServer new a gRPC server.
func NewDomainGRPCServer(service *service.DomainTaskService) (*DomainServer, error) {
	var opts []grpc.ServerOption

	// get the creds then append into the grpc options
	creds, err := pkg.GetServerCreds()
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.Creds(creds))

	srv := grpc.NewServer(opts...)

	v1.RegisterDomainServiceServer(srv, service)

	return &DomainServer{srv}, nil
}

type InterfaceServer struct {
	*grpc.Server
}

func NewInterfaceGRPCServer(service *service.DDNSInterfaceService) (*InterfaceServer, error) {
	var opts []grpc.ServerOption

	// get the creds then append into the grpc options
	creds, err := pkg.GetServerCreds()
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.Creds(creds),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(service.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(service.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(service.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(service.RecoveryInterceptor()),
		)))

	srv := grpc.NewServer(opts...)

	v1.RegisterDDNSInterfaceServer(srv, service)

	return &InterfaceServer{srv}, nil
}
