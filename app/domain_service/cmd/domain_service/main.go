package main

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/conf"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

// App contain grpc serve and gin http serve
type App struct {
	gs *grpcService
	hs *ginService
}

func newApp(grpcServer *grpc.Server, ginEngine *gin.Engine) *App {
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

	return &App{
		gs: &grpcService{
			gs:    grpcServer,
			lis:   &grpcListener,
			serve: grpcServe,
		},
		hs: &ginService{
			hs:    &ginServer,
			serve: ginServe,
		},
	}
}

func (a *App) start(stop chan struct{}, done chan bool, errors chan error) {
	// set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// grpc server
	go func() {
		errors <- a.gs.serve(stop, a.gs.gs, a.gs.lis)
	}()

	// web server
	go func() {
		errors <- a.hs.serve(stop, a.hs.hs)
	}()

	// blocking waiting to exit
	<-done
}

func main() {
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
	errors := make(chan error, 2)

	app, redisCleanup, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	defer redisCleanup()

	app.start(stop, done, errors)

	// smooth exit
	close(stop)
	for i := 0; i < cap(errors); i++ {
		if err := <-errors; err != nil {
			log.Println(err)
		}
	}
}
