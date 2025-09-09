package handler

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Success bool `json:"success"`
	Payload any  `json:"payload"`
}

func Response(c *gin.Context, code int, r *response) {
	c.JSON(code, r)
}

func SuccessWithCode(c *gin.Context, code int, payload any) {
	Response(c, 200, &response{Success: true, Payload: payload})
}

func Success(c *gin.Context, payload any) {
	SuccessWithCode(c, 200, payload)
}

func Failure(c *gin.Context, code int, err *serverError) {
	Response(c, code, &response{Success: false, Payload: err})
	// TODO: log error
}
