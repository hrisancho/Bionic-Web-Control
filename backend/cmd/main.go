package main

import (
	"Bionic-Web-Control/internal/config"
	app_logger "Bionic-Web-Control/internal/logger"
	"Bionic-Web-Control/internal/mqtt"
	"Bionic-Web-Control/internal/server"
	app_validator "Bionic-Web-Control/internal/validator"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer ctxCancel()

	//TODO полностью переписать docker-compose файл
	//TODO Добавить swagger сервер для проверки http api

	mainValidator, err := app_validator.NewValidator()

	mainConfig, err := config.LoadConfig(mainValidator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mainConfig)

	logger, err := app_logger.NewLogger(config.DebugLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("config loaded", zap.Any("config", mainConfig))

	clientMQTT, err := mqtt.NewClientMQTT(ctx,
		logger,
		mainConfig,
	)

	if err != nil {
		log.Fatal(err)
	}
	// TODO автоматически узнавать какой порт сейчас свободен
	mainServer, err := server.NewServer(mainConfig, logger, mainValidator, clientMQTT)
	if err != nil {
		logger.Fatal("error while starting server", zap.Error(err))
	}

	mainServer.Run(ctx)
}
