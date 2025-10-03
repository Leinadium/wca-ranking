package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService  port.StateService
	syncService   port.SyncService
	searchService port.SearchService
	personService port.PersonService
}

func NewServerHandler(
	stateService port.StateService,
	syncService port.SyncService,
	searchService port.SearchService,
	personService port.PersonService,
) *ServerHandler {
	return &ServerHandler{
		stateService:  stateService,
		syncService:   syncService,
		searchService: searchService,
		personService: personService,
	}
}
