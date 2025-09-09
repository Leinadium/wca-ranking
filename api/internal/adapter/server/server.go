package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	config *config.Server,
	handlers []handler.Handler,
	groups []handler.HandlerGroup,
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

	return &Server{engine: engine}
}

func (s *Server) Run() {
	return s.engine.Run()
}