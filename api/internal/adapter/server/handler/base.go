package handler

import "leinadium.dev/wca-ranking/internal/core/port"

type ServerHandler struct {
	stateService   port.StateService
	syncService    port.SyncService
	searchService  port.SearchService
	personService  port.PersonService
	rankingService port.RankingService
	authService    port.AuthService
	userService    port.UserService
}

func NewServerHandler(
	stateService port.StateService,
	syncService port.SyncService,
	searchService port.SearchService,
	personService port.PersonService,
	rankingService port.RankingService,
	authService port.AuthService,
	userService port.UserService,
) *ServerHandler {
	return &ServerHandler{
		stateService:   stateService,
		syncService:    syncService,
		searchService:  searchService,
		personService:  personService,
		rankingService: rankingService,
		authService:    authService,
		userService:    userService,
	}
}
