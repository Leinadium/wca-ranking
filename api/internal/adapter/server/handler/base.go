package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService port.StateService
	syncService  port.SyncService
}

func NewServerHandler(
	stateService port.StateService,
	syncService port.SyncService,
) *ServerHandler {
	return &ServerHandler{
		stateService: stateService,
		syncService:  syncService,
	}
}
