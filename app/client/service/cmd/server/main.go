package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
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

type cronService struct {
	cr    *cron.Cron
	serve func(stop chan struct{}, cr *cron.Cron) error
}

// App contain grpc serve and gin http serve
type App struct {
	cronService *cronService
}

func newApp(cronConf *cron.Cron) *App {
	cronServe := func(stop chan struct{}, cr *cron.Cron) error {
		go func(stop chan struct{}, cr *cron.Cron) {
			<-stop
			cr.Stop()
		}(stop, cr)

		cr.Run()
		return nil
	}

	return &App{
		cronService: &cronService{
			cr:    cronConf,
			serve: cronServe,
		},
	}
}

func (a *App) start(stop chan struct{}, errors chan error) {
	// cron server
	go func() {
		errors <- a.cronService.serve(stop, a.cronService.cr)
		log.Println("cron service stop...")
	}()
}

func main() {
	log.Printf("service.id: %v, service.name: %v, service.version: %v, git sha1: %v, build stamp: %v", id, Name, Version, GitSHA1, BuildStamp)

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
	errors := make(chan error, 1)

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
