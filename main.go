package main

import (
	"context"
	"keruen-geo/controller"
	"keruen-geo/models"
	"keruen-geo/service"
	"keruen-geo/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("started keruen-geo")
	storage := make(map[int64]models.Location)

	store := store.NewStore(storage)
	service := service.NewStoreService(store)
	controller := controller.NewController(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer", "Authority"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://geo.qkeruen.kz"
		},
	}))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	r.POST("/", controller.GetDrivers)
	r.POST("/getAll", controller.GetAllDrivers)
	r.POST("/getNearby", controller.GetNearby)

	// for drivers
	r.POST("/:id", controller.Create)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3002",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, os.Interrupt, syscall.SIGTERM)
	<-quite

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Println("Server stopped")
}
