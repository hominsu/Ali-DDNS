package main

import (
	"Ali-DDNS/app/server/service/internal/conf"
	"Ali-DDNS/app/server/service/internal/server"
	"Ali-DDNS/pkg"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// GitSHA1 is the sha1 of the HEAD
	GitSHA1 string
	// BuildStamp is the date of the build
	BuildStamp string

	id, _ = os.Hostname()
)

type domainGrpcService struct {
	gs     *grpc.Server
	lis    *net.Listener
	logger *zap.Logger
	serve  func(stop chan struct{}, gs *grpc.Server, lis *net.Listener, logger *zap.Logger) error
}

type interfaceGrpcService struct {
	gs     *grpc.Server
	lis    *net.Listener
	logger *zap.Logger
	serve  func(stop chan struct{}, gs *grpc.Server, lis *net.Listener, logger *zap.Logger) error
}

type interfaceHttpService struct {
	hs     *http.Server
	rc     context.CancelFunc
	logger *zap.Logger
	serve  func(stop chan struct{}, hs *http.Server, rc context.CancelFunc, logger *zap.Logger) error
}

type cronService struct {
	cr     *cron.Cron
	logger *zap.Logger
	serve  func(stop chan struct{}, cr *cron.Cron, logger *zap.Logger) error
}

// App contain grpc serve and gin http serve
type App struct {
	domainGrpcService    *domainGrpcService
	interfaceGrpcService *interfaceGrpcService
	interfaceHttpService *interfaceHttpService
	cronService          *cronService
}

func newApp(ds *server.DomainServer, is *server.InterfaceServer, hs *server.HttpServer, cr *cron.Cron, logger *zap.Logger) *App {
	dsListener, err := net.Listen(conf.Basic().DomainGrpcNetwork(), ":"+conf.Basic().DomainGrpcPort())
	if err != nil {
		logger.Sugar().Fatalf("create listener failed, err: %v", err)
	}

	// grpc serve func, use stop channel to smooth exit
	dsServe := func(stop chan struct{}, gs *grpc.Server, lis *net.Listener, logger *zap.Logger) error {
		go func(stop chan struct{}) {
			<-stop
			gs.Stop()
		}(stop)

		return gs.Serve(*lis)
	}

	isListener, err := net.Listen(conf.Basic().InterfaceGrpcNetwork(), ":"+conf.Basic().InterfaceGrpcPort())
	if err != nil {
		logger.Sugar().Fatalf("create listener failed, err: %v", err)
	}

	// grpc serve func, use stop channel to smooth exit
	isServe := func(stop chan struct{}, gs *grpc.Server, lis *net.Listener, logger *zap.Logger) error {
		go func(stop chan struct{}) {
			<-stop
			gs.Stop()
		}(stop)

		return gs.Serve(*lis)
	}

	httpServer := &http.Server{
		Addr:    ":" + conf.Basic().InterfacePort(),
		Handler: hs.Mux,
	}

	// http serve func, use stop channel to smooth exit
	httpServe := func(stop chan struct{}, hs *http.Server, rc context.CancelFunc, logger *zap.Logger) error {
		go func(stop chan struct{}, hs *http.Server, rc context.CancelFunc, logger *zap.Logger) {
			defer rc()
			<-stop
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if err := hs.Shutdown(ctx); err != nil {
				logger.Sugar().Warnf("http server shutdown failed, error: %v", err)
			}
		}(stop, hs, rc, logger)

		return hs.ListenAndServe()
	}

	cronServe := func(stop chan struct{}, cr *cron.Cron, logger *zap.Logger) error {
		go func(stop chan struct{}, cr *cron.Cron) {
			<-stop
			cr.Stop()
		}(stop, cr)

		cr.Run()
		return nil
	}

	return &App{
		domainGrpcService: &domainGrpcService{
			gs:     ds.Server,
			lis:    &dsListener,
			serve:  dsServe,
			logger: logger,
		},
		interfaceGrpcService: &interfaceGrpcService{
			gs:     is.Server,
			lis:    &isListener,
			serve:  isServe,
			logger: logger,
		},
		interfaceHttpService: &interfaceHttpService{
			hs:     httpServer,
			rc:     hs.GRPCCancel,
			serve:  httpServe,
			logger: logger,
		},
		cronService: &cronService{
			cr:     cr,
			serve:  cronServe,
			logger: logger,
		},
	}
}

func (a *App) start(stop chan struct{}, errors chan error) {
	// cron server
	go func() {
		s := a.cronService
		errors <- s.serve(stop, s.cr, s.logger)
		s.logger.Info("cron service stop...")
	}()

	// domain grpc server
	go func() {
		s := a.domainGrpcService
		errors <- s.serve(stop, s.gs, s.lis, s.logger)
		s.logger.Info("domain grpc service stop...")
	}()

	// interface grpc server
	go func() {
		s := a.interfaceGrpcService
		errors <- s.serve(stop, s.gs, s.lis, s.logger)
		s.logger.Info("interface grpc service stop...")
	}()

	// interface http server
	go func() {
		s := a.interfaceHttpService
		errors <- s.serve(stop, s.hs, s.rc, s.logger)
		s.logger.Info("interface http service stop...")
	}()
}

func main() {
	log.Printf("service.id: %v, service.name: %v, service.version: %v, git sha1: %v, build stamp: %v", id, Name, Version, GitSHA1, BuildStamp)

	logger := pkg.NewProductionLogger()

	logger.Sugar().Infof("service.id: %v, service.name: %v, service.version: %v, git sha1: %v, build stamp: %v", id, Name, Version, GitSHA1, BuildStamp)

	done := make(chan bool)
	ch := make(chan os.Signal)

	// monitor the exit signal
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func(done chan bool, logger *zap.Logger) {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logger.Sugar().Infof("Signal: %v, Program Exit...", s)
				done <- true
			default:
				logger.Sugar().Infof("Other Signal: %v", s)
			}
		}
	}(done, logger)

	// stop channel use to control all server's status, error channel use to receive the server error
	stop := make(chan struct{})
	errors := make(chan error, 4)

	app, dataRepoCleanup, err := initApp(logger)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
	defer dataRepoCleanup()

	app.start(stop, errors)

	// blocking waiting to exit
	<-done

	// smooth exit
	close(stop)
	for i := 0; i < cap(errors); i++ {
		if err := <-errors; err != nil {
			logger.Error(err.Error())
		}
	}
}
