package suggestions

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"destinations-suggester/internal/domain/models/suggestions"
	"destinations-suggester/internal/infra/http/server/shared"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ListQueryParams struct {
	Lat    float64   `query:"lat"`
	Lon    float64   `query:"lon"`
	UserID uuid.UUID `query:"userId"`
}

type ListResponse struct {
	Suggestions []shared.Suggestion `json:"suggestions"`
}

type suggestionsLister interface {
	List(ctx context.Context, userID uuid.UUID, userLocation places.Coordinates) ([]suggestions.Suggestion, error)
}

func List(suggestionsLister suggestionsLister) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var params ListQueryParams
		if err := c.Bind(&params); err != nil {
			return fmt.Errorf("cannot bind query params: %w", err)
		}

		suggestionsSlice, err := suggestionsLister.List(ctx, params.UserID, places.Coordinates{
			Lat: params.Lat,
			Lon: params.Lon,
		})
		if err != nil {
			return fmt.Errorf("cannot list suggestions: %w", err)
		}

		converted := make([]shared.Suggestion, 0, len(suggestionsSlice))
		for _, suggestion := range suggestionsSlice {
			converted = append(converted, shared.Suggestion{}.FromModel(&suggestion))
		}

		return c.JSON(http.StatusOK, ListResponse{
			Suggestions: converted,
		})
	}
}
