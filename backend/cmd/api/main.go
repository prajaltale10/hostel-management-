package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hostel-saas/config"
	"hostel-saas/models"
	"hostel-saas/pkg/db"
	"hostel-saas/pkg/logger"
	"hostel-saas/pkg/redis"
	"hostel-saas/router"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}

	logger.InitLogger(cfg.Env)
	defer logger.Log.Sync()

	if err := db.InitPostgres(cfg); err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
	}

	if err := models.AutoMigrate(db.DB); err != nil {
		logger.Log.Fatal("Failed to auto migrate database", zap.Error(err))
	}

	if err := redis.InitRedis(cfg); err != nil {
		logger.Log.Fatal("Redis connection failed", zap.Error(err))
	}

	r := router.SetupRouter(cfg)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		logger.Log.Info("Starting server on port " + cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Fatal("Server startup failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Log.Info("Server exiting")
}
