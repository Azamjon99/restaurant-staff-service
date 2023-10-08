package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcserver "github.com/Azamjon99/restaurant-staff-service/src/application/grpc"
	pb "github.com/Azamjon99/restaurant-staff-service/src/application/protos/restaurant_staff"
	appsvc "github.com/Azamjon99/restaurant-staff-service/src/application/services"
	managersvc "github.com/Azamjon99/restaurant-staff-service/src/domain/manager/services"
	"github.com/Azamjon99/restaurant-staff-service/src/infrastructure/config"
	"github.com/Azamjon99/restaurant-staff-service/src/infrastructure/jwt"
	managerrepo "github.com/Azamjon99/restaurant-staff-service/src/infrastructure/repositories/manager"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger, err := config.NewLogger()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDatabase,
		60,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	managerRepo := managerrepo.NewMangerRepository(db)

	managerSvc := managersvc.NewManagerService(managerRepo)
	tokenSvc := jwt.NewService(config.JWTSecret, config.JWTExpiresInSec)

	managerApp := appsvc.NewManagerAppService(managerSvc, tokenSvc)

	root := gin.Default()

	root.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	osSignals := make(chan os.Signal, 1)

	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(osSignals)

	// start http server

	var httpServer *http.Server

	g.Go(func() error {
		httpServer = &http.Server{
			Addr:    config.HttpPort,
			Handler: root,
		}

		logger.Debug("main: started http server", zap.String("port", config.HttpPort))
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// grpc server
	var grpcServer *grpc.Server

	g.Go(func() error {
		server := grpcserver.NewServer(
			managerApp,
		)
		grpcServer = grpc.NewServer()
		pb.RegisterStaffServiceServer(grpcServer, server)

		lis, err := net.Listen("tcp", config.GrpcPort)
		if err != nil {
			logger.Fatal("main: net.Listen", zap.Error(err))
		}

		defer lis.Close()

		logger.Debug("main: started grps server", zap.String("port", config.GrpcPort))

		return grpcServer.Serve(lis)
	})

	select {
	case <-osSignals:
		logger.Info("main: received os signal, shutting down")
		break
	case <-ctx.Done():
		break
	}

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Shutdown(shutdownCtx)
	}

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if err := g.Wait(); err != nil {
		logger.Error("main: server returned an error", zap.Error(err))
	}
}
