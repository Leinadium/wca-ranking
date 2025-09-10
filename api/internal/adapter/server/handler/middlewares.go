package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors.Errors() {
			go fmt.Printf("Error: %v", err)
		}
	}
}
