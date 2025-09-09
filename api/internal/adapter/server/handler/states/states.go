package states

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
	"leinadium.dev/wca-ranking/internal/core/service"
)

type GetStates struct {
	svc *service.StateService
}

func (s *GetStates) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		states, _ := s.svc.States(c.Request.Context())
		c.JSON(http.StatusOK, gin.H{
			"results": states,
		})
	}
}

func (s *GetStates) Metadata() *handler.RouteMetadata {
	return &handler.RouteMetadata{Method: handler.GET, Pattern: "/"}
}

func StatesGroup(svc *service.StateService) *handler.HandlerGroup {
	return handler.NewHandlerGroup("/states", []handler.Handler{
		&GetStates{svc: svc},
	})
}
