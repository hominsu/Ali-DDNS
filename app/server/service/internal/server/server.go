package server

import "github.com/google/wire"

// ProviderSet is grpc server and http server providers.
var ProviderSet = wire.NewSet(NewDomainGRPCServer, NewInterfaceGRPCServer, NewInterfaceHTTPServer, NewCronServer)
