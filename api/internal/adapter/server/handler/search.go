package handler

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/server/schema"
	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/pkg/utils"
)

func (s *ServerHandler) GetSearch(
	ctx context.Context,
	request schema.GetSearchRequestObject,
) (schema.GetSearchResponseObject, error) {
	search, err := s.searchService.Search(ctx, request.Params.S)
	if err != nil {
		return schema.GetSearch500JSONResponse(ErrDefault), nil
	} else {
		return schema.GetSearch200JSONResponse(
			utils.MapNotNull(search, func(s *domain.SearchResult) schema.SearchItem {
				return schema.SearchItem{
					WcaId:   (*string)(&s.WCAID),
					WcaName: &s.Name,
					StateId: &s.StateID.String,
				}
			}),
		), nil
	}
}
