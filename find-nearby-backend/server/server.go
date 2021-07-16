package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"find-nearby-backend/logger"
	"find-nearby-backend/repository"
	"find-nearby-backend/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// Server represents the HTTP Server. Echo is used as the implementation.
type Server struct {
	address     string
	apiServer   *echo.Echo
	db          *sqlx.DB
	log         logger.Logger
	serverReady chan bool
}

// Start starts HTTP Server
func (s *Server) Start() {
	locationsRepo := repository.NewPostgresLocationRepository(s.db)
	locationsUsecase := usecase.NewLocationUsecase(locationsRepo)
	handler := NewHandler(s.log, locationsUsecase)
	s.apiServer.GET("/ping", handler.Ping)
	s.apiServer.GET("/locations/find", handler.FindLocations)
	go s.waitForShutdown(s.apiServer)
	go s.listenServer(s.apiServer)
	s.serverReady <- true
}

// ServerReady is a channel that signals whether a server is ready to serve the requests
func (s *Server) ServerReady() chan bool {
	return s.serverReady
}

func (s *Server) listenServer(apiServer *echo.Echo) {
	err := apiServer.Start(s.address)
	if err != http.ErrServerClosed {
		s.log.Fatalf(err.Error())
	}
}

func (s *Server) waitForShutdown(apiServer *echo.Echo) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	<-sig
	s.log.Infof("API server shutting down")

	if err := apiServer.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		s.log.Errorf(err.Error())
	}
	s.log.Infof("API server shutdown complete")
}

// NewServer is a constructor for a Server
func NewServer(addr string, db *sqlx.DB, logger logger.Logger) *Server {
	srv := Server{
		address:     addr,
		apiServer:   echo.New(),
		db:          db,
		log:         logger,
		serverReady: make(chan bool),
	}
	return &srv
}
