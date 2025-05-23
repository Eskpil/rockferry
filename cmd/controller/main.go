package main

import (
	"log"
	"log/slog"
	"net"
	"sync"
	"time"
	"fmt"

	"github.com/eskpil/rockferry/controllerapi"
	"github.com/eskpil/rockferry/internal/controller"
	"github.com/eskpil/rockferry/internal/controller/api"
	"github.com/eskpil/rockferry/internal/controller/controllers/resource"
	"github.com/eskpil/rockferry/internal/controller/db"
	"github.com/eskpil/rockferry/internal/controller/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

func runDb(dir string) {
	cfg := embed.NewConfig()
	cfg.Dir = dir
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer e.Close()
	select {
	case <-e.Server.ReadyNotify():
		log.Printf("Database is running")
	case <-time.After(60 * time.Second):
		e.Server.Stop() // trigger a shutdown
		log.Printf("Server took too long to start!")
	}

	log.Fatal(<-e.Err())
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		// TODO: Add clustering support
		runDb("salmon_vm.etcd")
		wg.Done()
	}(wg)

	if err := controller.Initialize(); err != nil {
		panic(err)
	}

	// TODO: Enable some kind of config
	// TODO: Avoid multiple db connections
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	// TODO: Avoid panic
	if err != nil {
		panic(err)
	}

	r := runtime.New(cli)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		server := echo.New()

		server.Use(r.EchoMiddleware())

		//server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//	AllowOrigins:     []string{"http://10.100.0.186:5173", "*"},                                                                              // Replace with your client's origin
		//	AllowMethods:     []string{http.MethodHead, http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},    // Allow only GET for SSE
		//	AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "Host", "Accept-Encoding", "Accept-Language", "Origin", "Referer"}, // Allow necessary headers
		//	AllowCredentials: true,                                                                                                                   // Allow credentials if required
		//
		//}))

		server.Use(middleware.CORS())

		server.Use(db.Middleware())

		server.GET("/v1/resources/events", resource.Watch())
		server.GET("/v1/resources", resource.List())
		server.POST("v1/resources", resource.Create())
		server.DELETE("/v1/resources", resource.Delete())
		server.PATCH("/v1/resources", resource.Patch())

		if err := server.Start("0.0.0.0:8080"); err != nil {
			panic(err)
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		listener, err := net.Listen("tcp", "0.0.0.0:9090")
		if err != nil {
			panic(err)
		}

		api, err := api.New(r)
		if err != nil {
			panic(err)
		}

		server := grpc.NewServer()
		controllerapi.RegisterControllerApiServer(server, api)

		reflection.Register(server)

		fmt.Println("Serving gRPC api on port 9090")

		if err := server.Serve(listener); err != nil {
			slog.Error("could not serve requests", slog.Any("err", err))
		}

	}(wg)

	wg.Wait()
}
