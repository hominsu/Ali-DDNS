//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/biz"
	"Ali-DDNS/app/domain_service/internal/domain_task/data"
	"Ali-DDNS/app/domain_service/internal/domain_task/server"
	"Ali-DDNS/app/domain_service/internal/domain_task/service"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp() (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
