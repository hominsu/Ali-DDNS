package server

import (
	v1 "Ali-DDNS/api/domain_record/v1"
	"Ali-DDNS/app/domain_service/internal/domain_task/service"
	"Ali-DDNS/pkg"
	"google.golang.org/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(service *service.DomainTaskService) *grpc.Server {
	var opts []grpc.ServerOption
	// get the creds then append into the grpc options
	opts = append(opts, grpc.Creds(pkg.GetServerCreds()))

	srv := grpc.NewServer(opts...)

	v1.RegisterDomainServiceServer(srv, service)

	return srv
}
