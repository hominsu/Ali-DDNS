//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Ali-DDNS/app/domain_check/internal/domain_check/data"
	"Ali-DDNS/app/domain_check/internal/domain_check/server"
	"Ali-DDNS/app/domain_check/internal/domain_check/service"
	"github.com/google/wire"
)

// initApp init ddns client application.
func initApp() (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
