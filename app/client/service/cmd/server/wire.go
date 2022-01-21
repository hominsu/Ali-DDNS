//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Ali-DDNS/app/client/service/internal/data"
	"Ali-DDNS/app/client/service/internal/server"
	"Ali-DDNS/app/client/service/internal/service"
	"github.com/google/wire"
)

// initApp init ddns client application.
func initApp() (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
