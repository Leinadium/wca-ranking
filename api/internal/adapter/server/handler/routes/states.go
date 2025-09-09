package states

import (
	"github.com/gin-gonic/gin"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
	"leinadium.dev/wca-ranking/internal/core/service"
)

type GetStates struct {
	svc *service.StateService
}

func (s *GetStates) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		states, err := s.svc.States(c.Request.Context())
		if err != nil {
			handler.Failure(c, 500, &handler.ErrDefault)
		} else {
			handler.Success(c, states)
		}
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
