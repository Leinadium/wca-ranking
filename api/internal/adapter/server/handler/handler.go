package handler

import "github.com/gin-gonic/gin"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type RouteMetadata struct {
	Method  Method
	Pattern string
}

type Handler interface {
	Handle() gin.HandlerFunc
	Metadata() *RouteMetadata
}

type HandlerGroup struct {
	Handlers    []Handler
	Middlewares []gin.HandlerFunc
	Pattern     string
}

func NewHandlerGroup(pattern string, handlers []Handler) *HandlerGroup {
	return &HandlerGroup{
		Pattern:  pattern,
		Handlers: handlers,
	}
}
