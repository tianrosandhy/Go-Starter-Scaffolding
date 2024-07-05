package main

import (
	"os"
	"os/signal"
	"skeleton/bootstrap"
	"skeleton/src/routes"
	"sync"
	"syscall"
	"time"
)

var version = "dev"
var builddate = "realtime"

// @title Skeleton Service
// @version 1.0.0
// @description Skeleton Service
// @contact.name TianRosandhy
// @contact.email tianrosandhy@gmail.com
// @host localhost
// @schemes https
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := bootstrap.NewApplication()

	app.Config.Set("APP_VERSION", version)
	app.Config.Set("BUILDDATE", builddate)
	app.Log.Infof("Running on version %s (build date %s)", version, builddate)

	wg := sync.WaitGroup{}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		app.Log.Println("Initiate gracefully shutdown with exit signal")
		WaitTimeout(&wg, 10*time.Second)
		app.Log.Println("Gracefully shutting down...")
		_ = app.App.Close()
	}()

	routes.Handle(app)

	app.Log.Fatal(app.App.Start(":" + app.Config.GetString("PORT")))
}

// WaitTimeout to wait with timeout
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()
	select {
	case <-done:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
