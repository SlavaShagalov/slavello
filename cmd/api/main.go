package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/SlavaShagalov/slavello/internal/pkg/config"
	pLog "github.com/SlavaShagalov/slavello/internal/pkg/log/zap"
	pStorages "github.com/SlavaShagalov/slavello/internal/pkg/storages"
	"github.com/SlavaShagalov/slavello/internal/pkg/storages/postgres"
)

func main() {
	// ===== Logger =====
	logger := pLog.NewDev()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Println(err)
		}
	}()
	logger.Info("API server starting...")

	// ===== Configuration =====
	viper.SetConfigName("api")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Failed to read configuration", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Configuration read successfully")

	// ===== Database =====
	db, err := postgres.NewStd(logger)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			logger.Error("Failed to close Postgres connection", zap.Error(err))
		}
		logger.Info("Postgres connection closed")
	}()

	// ===== Sessions Storage =====
	ctx := context.Background()
	redisClient, err := pStorages.NewRedis(logger, ctx)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		err = redisClient.Close()
		if err != nil {
			logger.Error("Failed to close Redis client", zap.Error(err))
		}
		logger.Info("Redis client closed")
	}()

	router := mux.NewRouter()

	server := http.Server{
		Addr:    ":" + viper.GetString(config.ServerPort),
		Handler: router,
	}

	// ===== Start =====
	logger.Info("API service started", zap.String("port", viper.GetString(config.ServerPort)))
	if err = server.ListenAndServe(); err != nil {
		logger.Error("API server stopped", zap.Error(err))
	}
}
