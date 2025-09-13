//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ../../../openapi/config.yaml ../../../openapi/api.yaml

package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

func NewServer(
	config *config.Server,
	handler *handler.ServerHandler,
) *Server {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Host, "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	swagger, _ := schema.GetSwagger()
	engine.Use(ginmiddleware.OapiRequestValidator(swagger))

	strictHandler := schema.NewStrictHandler(handler, []schema.StrictMiddlewareFunc{})
	schema.RegisterHandlers(engine, strictHandler)

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
