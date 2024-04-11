package server

import (
	"Bionic-Web-Control/internal/mqtt"
	"context"
	"log"
	"sync"

	"Bionic-Web-Control/internal/config"
	main_logger "Bionic-Web-Control/internal/logger"
	//"Bionic-Web-Control/internal/mqtt"
	"Bionic-Web-Control/internal/validator"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
)

type Server struct {
	App        *fiber.App
	Config     config.Config
	Logger     *main_logger.Logger
	Validator  *validator.AppValidator
	clientMQTT *mqtt.ClientMQTT
}

func NewServer(config config.Config, logger *main_logger.Logger, validator *validator.AppValidator, clientMQTT *mqtt.ClientMQTT) (server *Server, err error) {
	server = &Server{
		App:        fiber.New(GetFiberConfig()),
		Config:     config,
		Logger:     logger,
		Validator:  validator,
		clientMQTT: clientMQTT,
	}

	server.SetupRoutes()

	return
}

func (server Server) Run(ctx context.Context) {
	eventServerStopped := sync.WaitGroup{}
	eventServerStopped.Add(1)
	go func() {
		<-ctx.Done()
		if err := server.App.Shutdown(); err != nil {
			server.Logger.Error("HTTP server shutdown error", zap.Error(err))
		}

		eventServerStopped.Done()
	}()

	err := server.App.ListenTLS(server.Config.ServerAddr, "./config/http-ssl/server.crt", "./config/http-ssl/server.key")
	if err != nil {
		server.Logger.Info("SSL keys not found, using HTTP")
		err = server.App.Listen(server.Config.ServerAddr)
		if err != nil {
			log.Println(err)
		}
	}

	server.Logger.Info("Server stopping...")
	eventServerStopped.Wait()
	server.Logger.Info("Server stopped successfully")
}
