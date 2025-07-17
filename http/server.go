package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(apiKey string) *Server {
	r := mux.NewRouter()

	// Simple Auth Handler
	r.Use(APIKeyMiddleware(apiKey))

	// Middleware Logger and JSON Header
	r.Use(LoggingMiddleware)
	r.Use(JSONMiddleware)

	return &Server{
		router: r,
	}
}

func (s *Server) RegisterRoutes(handler interface{}) {
	if h, ok := handler.(interface{ RegisterRoutes(*mux.Router) }); ok {
		h.RegisterRoutes(s.router)
	}
}

func (s *Server) Start(addr string) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
