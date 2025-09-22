package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService  port.StateService
	syncService   port.SyncService
	searchService port.SearchService
}

func NewServerHandler(
	stateService port.StateService,
	syncService port.SyncService,
	searchService port.SearchService,
) *ServerHandler {
	return &ServerHandler{
		stateService:  stateService,
		syncService:   syncService,
		searchService: searchService,
	}
}
