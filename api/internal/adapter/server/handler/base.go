package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService port.StateService
}

func NewServerHandler(
	stateService port.StateService,
) *ServerHandler {
	return &ServerHandler{
		stateService: stateService,
	}
}
