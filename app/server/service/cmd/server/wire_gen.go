// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"Ali-DDNS/app/server/service/internal/data"
	"Ali-DDNS/app/server/service/internal/server"
	"Ali-DDNS/app/server/service/internal/service"
)

// Injectors from wire.go:

// initApp init ddns server application.
func initApp() (*App, func(), error) {
	client := data.NewRedisClient()
	dataData, cleanup, err := data.NewData(client)
	if err != nil {
		return nil, nil, err
	}
	delayCheckRepo := data.NewDelayCheckRepo(dataData)
	delayCheckUsecase := biz.NewDelayCheckUsecase(delayCheckRepo)
	domainRecordRepo := data.NewDomainRecordRepo(dataData)
	domainRecordUsecase := biz.NewDomainRecordUsecase(domainRecordRepo)
	domainUserRepo := data.NewDomainUserRepo(dataData)
	domainUserUsecase := biz.NewDomainUserUsecase(domainUserRepo)
	domainTaskService := service.NewDomainTaskService(delayCheckUsecase, domainRecordUsecase, domainUserUsecase)
	grpcServer, err := server.NewGRPCServer(domainTaskService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	engine := server.NewGinServer(domainTaskService)
	cron, err := server.NewCronServer(domainTaskService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app := newApp(grpcServer, engine, cron)
	return app, func() {
		cleanup()
	}, nil
}
