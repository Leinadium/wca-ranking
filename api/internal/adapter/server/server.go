package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

func NewServer(
	config *config.Server,
	handlers []handler.Handler,
	groups []*handler.HandlerGroup,
) *Server {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Host, "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	for _, group := range groups {
		g := engine.Group(group.Pattern, group.Middlewares...)
		for _, handler := range group.Handlers {
			g.Handle(string(handler.Metadata().Method), handler.Metadata().Pattern, handler.Handle())
		}
	}
	for _, handler := range handlers {
		engine.Handle(string(handler.Metadata().Method), handler.Metadata().Pattern, handler.Handle())
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: engine.Handler(),
	}

	return &Server{engine: engine, server: server}
}

func (s *Server) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return err
	}

	go s.server.Serve(ln)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
