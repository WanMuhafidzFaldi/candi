package app

import (
	"context"
	"log"

	"github.com/agungdwiprasetyo/backend-microservices/config"
	"github.com/agungdwiprasetyo/backend-microservices/internal/factory"
	"github.com/agungdwiprasetyo/backend-microservices/internal/factory/base"
	"github.com/agungdwiprasetyo/backend-microservices/internal/services"
	"github.com/agungdwiprasetyo/backend-microservices/pkg/helper"
	"github.com/agungdwiprasetyo/backend-microservices/pkg/middleware"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

// App user service
type App struct {
	config     *config.Config
	modules    []factory.ModuleFactory
	httpServer *echo.Echo
	grpcServer *grpc.Server
}

// New user service app
func New(cfg *config.Config) *App {

	mw := middleware.NewMiddleware(cfg)
	params := &base.ModuleParam{
		Config:     cfg,
		Middleware: mw,
	}

	selectedService := services.InitService(config.GlobalEnv.Service, params)
	modules := selectedService.Modules()

	// init http server
	echoServer := echo.New()

	// init grpc server
	grpcServer := grpc.NewServer(
		grpc.MaxSendMsgSize(200*int(helper.MByte)), grpc.MaxRecvMsgSize(200*int(helper.MByte)),
		grpc.UnaryInterceptor(mw.GRPCAuth),
		grpc.StreamInterceptor(mw.GRPCAuthStream),
	)

	return &App{
		config:     cfg,
		modules:    modules,
		httpServer: echoServer,
		grpcServer: grpcServer,
	}
}

// Shutdown graceful shutdown all server, panic if there is still a process running when the request exceed given timeout in context
func (a *App) Shutdown(ctx context.Context) {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}

	log.Println("Stopping GRPC server...")
	a.grpcServer.GracefulStop()
}
