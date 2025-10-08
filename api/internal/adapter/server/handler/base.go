package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService   port.StateService
	syncService    port.SyncService
	searchService  port.SearchService
	personService  port.PersonService
	rankingService port.RankingService
	authService    port.AuthService
}

func NewServerHandler(
	stateService port.StateService,
	syncService port.SyncService,
	searchService port.SearchService,
	personService port.PersonService,
	rankingService port.RankingService,
	authService port.AuthService,
) *ServerHandler {
	return &ServerHandler{
		stateService:   stateService,
		syncService:    syncService,
		searchService:  searchService,
		personService:  personService,
		rankingService: rankingService,
		authService:    authService,
	}
}
