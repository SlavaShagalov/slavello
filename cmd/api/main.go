package main

import (
	"context"
	boardsRepository "github.com/SlavaShagalov/slavello/internal/boards/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	authDel "github.com/SlavaShagalov/slavello/internal/auth/delivery/http"
	authUsecase "github.com/SlavaShagalov/slavello/internal/auth/usecase"
	"github.com/SlavaShagalov/slavello/internal/boards"
	boardsUsecase "github.com/SlavaShagalov/slavello/internal/boards/usecase"
	mw "github.com/SlavaShagalov/slavello/internal/middleware"
	"github.com/SlavaShagalov/slavello/internal/pkg/config"
	pHasher "github.com/SlavaShagalov/slavello/internal/pkg/hasher/bcrypt"
	pLog "github.com/SlavaShagalov/slavello/internal/pkg/log/zap"
	pStorages "github.com/SlavaShagalov/slavello/internal/pkg/storages"
	"github.com/SlavaShagalov/slavello/internal/pkg/storages/postgres"
	sessionsRepository "github.com/SlavaShagalov/slavello/internal/sessions/repository/redis"
	"github.com/SlavaShagalov/slavello/internal/users"
	usersDel "github.com/SlavaShagalov/slavello/internal/users/delivery/http"
	usersRepository "github.com/SlavaShagalov/slavello/internal/users/repository/postgres"
	usersUsecase "github.com/SlavaShagalov/slavello/internal/users/usecase"
	"github.com/SlavaShagalov/slavello/internal/workspaces"
	workspacesDel "github.com/SlavaShagalov/slavello/internal/workspaces/delivery/http"
	workspacesRepository "github.com/SlavaShagalov/slavello/internal/workspaces/repository/postgres"
	workspacesUsecase "github.com/SlavaShagalov/slavello/internal/workspaces/usecase"
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

	// ===== Hasher =====
	hasher := pHasher.New()

	// ===== Repositories =====
	var usersRepo users.Repository
	var workspacesRepo workspaces.Repository
	var boardsRepo boards.Repository
	usersRepo = usersRepository.New(db, logger)
	workspacesRepo = workspacesRepository.New(db, logger)
	boardsRepo = boardsRepository.New(db, logger)

	sessionsRepo := sessionsRepository.New(redisClient, context.Background(), logger)

	serverType := viper.GetString(config.ServerType)

	// ===== Usecases =====
	authUC := authUsecase.New(usersRepo, sessionsRepo, hasher, logger)
	usersUC := usersUsecase.New(usersRepo)
	workspacesUC := workspacesUsecase.New(workspacesRepo)
	boardsUC := boardsUsecase.New(boardsRepo)

	router := mux.NewRouter()

	// ===== Middleware =====
	checkAuth := mw.NewCheckAuth(authUC, logger)
	cors := mw.NewCors()
	accessLog := mw.NewAccessLog(serverType, logger)

	// ===== Delivery =====
	authDel.RegisterHandlers(router, authUC, usersUC, logger, checkAuth)
	usersDel.RegisterHandlers(router, usersUC, logger, checkAuth)
	workspacesDel.RegisterHandlers(router, workspacesUC, boardsUC, logger, checkAuth)

	server := http.Server{
		Addr:    ":" + viper.GetString(config.ServerPort),
		Handler: accessLog(cors(router)),
	}

	// ===== Start =====
	logger.Info("API service started", zap.String("port", viper.GetString(config.ServerPort)))
	if err = server.ListenAndServe(); err != nil {
		logger.Error("API server stopped", zap.Error(err))
	}
}
