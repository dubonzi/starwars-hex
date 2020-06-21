package main

import (
	"net/http"
	"starwars-hex/pkg/planets"

	"github.com/labstack/echo/v4"
)

type planetHandler struct {
	svc planets.Service
}

// AddPlanetHandler setup handlers for the /planets endpoint.
func AddPlanetHandler(e *echo.Echo, svc planets.Service) {
	hand := planetHandler{svc}
	grp := e.Group("/planets")

	grp.GET("", hand.list)

}

func (h planetHandler) list(c echo.Context) error {
	planets, err := h.svc.List()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, planets)
	return nil
}
