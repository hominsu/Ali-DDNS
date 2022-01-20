package main

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/conf"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

	id, _ = os.Hostname()
)

type grpcService struct {
	gs    *grpc.Server
	lis   *net.Listener
	serve func(stop chan struct{}, gs *grpc.Server, lis *net.Listener) error
}

type ginService struct {
	hs    *http.Server
	serve func(stop chan struct{}, hs *http.Server) error
}

type cronService struct {
	cr    *cron.Cron
	serve func(stop chan struct{}, cr *cron.Cron) error
}

// App contain grpc serve and gin http serve
type App struct {
	grpcService *grpcService
	ginService  *ginService
	cronService *cronService
}

func newApp(grpcServer *grpc.Server, ginEngine *gin.Engine, cronConf *cron.Cron) *App {
	grpcListener, err := net.Listen(conf.Basic().RpcNetwork(), ":"+conf.Basic().RpcPort())
	if err != nil {
		log.Fatal(err)
	}

	// grpc serve func, use stop channel to smooth exit
	grpcServe := func(stop chan struct{}, gs *grpc.Server, lis *net.Listener) error {
		go func(stop chan struct{}) {
			<-stop
			gs.Stop()
		}(stop)

		return gs.Serve(*lis)
	}

	ginServer := http.Server{
		Addr:    ":" + conf.Basic().WebPort(),
		Handler: ginEngine,
	}

	// gin serve func, use stop channel to smooth exit
	ginServe := func(stop chan struct{}, hs *http.Server) error {
		go func(stop chan struct{}, hs *http.Server) {
			<-stop
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if err := hs.Shutdown(ctx); err != nil {
				log.Println(err)
			}
		}(stop, hs)

		return hs.ListenAndServe()
	}

	cronServe := func(stop chan struct{}, cr *cron.Cron) error {
		go func(stop chan struct{}, cr *cron.Cron) {
			<-stop
			cr.Stop()
		}(stop, cr)

		cr.Run()
		return nil
	}

	return &App{
		grpcService: &grpcService{
			gs:    grpcServer,
			lis:   &grpcListener,
			serve: grpcServe,
		},
		ginService: &ginService{
			hs:    &ginServer,
			serve: ginServe,
		},
		cronService: &cronService{
			cr:    cronConf,
			serve: cronServe,
		},
	}
}

func (a *App) start(stop chan struct{}, errors chan error) {
	// set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// cron server
	go func() {
		errors <- a.cronService.serve(stop, a.cronService.cr)
	}()

	// grpc server
	go func() {
		errors <- a.grpcService.serve(stop, a.grpcService.gs, a.grpcService.lis)
	}()

	// web server
	go func() {
		errors <- a.ginService.serve(stop, a.ginService.hs)
	}()
}

func main() {
	log.Printf("service.id: %v, service.name: %v, service.version: %v", id, Name, Version)

	done := make(chan bool)
	ch := make(chan os.Signal)

	// monitor the exit signal
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func(done chan bool) {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Printf("Signal: %v, Program Exit...", s)
				done <- true
			default:
				log.Println("Other Signal: ", s)
			}
		}
	}(done)

	// stop channel use to control all server's status, error channel use to receive the server error
	stop := make(chan struct{})
	errors := make(chan error, 3)

	app, dataRepoCleanup, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	defer dataRepoCleanup()

	app.start(stop, errors)

	// blocking waiting to exit
	<-done

	// smooth exit
	close(stop)
	for i := 0; i < cap(errors); i++ {
		if err := <-errors; err != nil {
			log.Println(err)
		}
	}
}
