package server

import "github.com/google/wire"

// ProviderSet is grpc server and http server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewGinServer, NewCronServer)
