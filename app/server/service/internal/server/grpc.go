package server

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/service"
	"Ali-DDNS/app/server/service/pkg/interceptor"
	"Ali-DDNS/pkg"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DomainServer struct {
	*grpc.Server
}

// NewDomainGRPCServer new a gRPC server.
func NewDomainGRPCServer(service *service.DomainTaskService, logger *zap.Logger) (*DomainServer, error) {
	var opts []grpc.ServerOption

	// get the creds then append into the grpc options
	creds, err := pkg.GetServerCreds()
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.Creds(creds))

	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_zap.StreamServerInterceptor(
			interceptor.ZapInterceptor(logger),
			grpc_zap.WithLevels(interceptor.CodeLevel),
		),
		grpc_recovery.StreamServerInterceptor(
			interceptor.RecoveryInterceptor(),
		),
	)))

	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(
			interceptor.ZapInterceptor(logger),
			grpc_zap.WithLevels(interceptor.CodeLevel),
		),
		grpc_recovery.UnaryServerInterceptor(
			interceptor.RecoveryInterceptor(),
		),
	)))

	srv := grpc.NewServer(opts...)

	v1.RegisterDomainServiceServer(srv, service)

	return &DomainServer{srv}, nil
}

type InterfaceServer struct {
	*grpc.Server
}

func NewInterfaceGRPCServer(service *service.DDNSInterfaceService, logger *zap.Logger) (*InterfaceServer, error) {
	var opts []grpc.ServerOption

	// get the creds then append into the grpc options
	creds, err := pkg.GetServerCreds()
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.Creds(creds))

	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_zap.StreamServerInterceptor(
			interceptor.ZapInterceptor(logger),
			grpc_zap.WithLevels(interceptor.CodeLevel),
		),
		grpc_recovery.StreamServerInterceptor(
			interceptor.RecoveryInterceptor(),
		),
	)))

	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(
			interceptor.ZapInterceptor(logger),
			grpc_zap.WithLevels(interceptor.CodeLevel),
		),
		grpc_recovery.UnaryServerInterceptor(
			interceptor.RecoveryInterceptor(),
		),
	)))

	srv := grpc.NewServer(opts...)

	v1.RegisterDDNSInterfaceServer(srv, service)

	return &InterfaceServer{srv}, nil
}
