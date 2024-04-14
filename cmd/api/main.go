package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/SlavaShagalov/slavello/internal/pkg/config"
	pLog "github.com/SlavaShagalov/slavello/internal/pkg/log/zap"
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
