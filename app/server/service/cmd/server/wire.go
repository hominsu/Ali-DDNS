//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"Ali-DDNS/app/server/service/internal/data"
	"Ali-DDNS/app/server/service/internal/server"
	"Ali-DDNS/app/server/service/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// initApp init ddns server application.
func initApp(logger *zap.Logger) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
