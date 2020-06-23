package main

import (
	"net/http"
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/logger"
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
	grp.GET("/:name", hand.byName)
	grp.POST("", hand.add)
	grp.DELETE("/:name", hand.delete)

}

func (h planetHandler) list(c echo.Context) error {
	planets, err := h.svc.List()
	if err != nil {
		return err
	}
	if err = c.JSON(http.StatusOK, planets); err != nil {
		logger.Error("planetHandler.list", "c.JSON", err, planets)
		return errs.Unexpected
	}
	return nil
}

func (h planetHandler) byName(c echo.Context) error {
	name := c.Param("name")
	planet, err := h.svc.FindByName(name)
	if err != nil {
		return err
	}
	if err = c.JSON(http.StatusOK, planet); err != nil {
		logger.Error("planetHandler.byName", "c.JSON", err, planet)
		return errs.Unexpected
	}
	return nil
}

func (h planetHandler) add(c echo.Context) error {
	var planet planets.Planet
	err := c.Bind(&planet)
	if err != nil {
		return errs.BadRequest
	}
	planet, err = h.svc.Add(planet)
	if err != nil {
		return err
	}
	if err = c.JSON(http.StatusCreated, planet); err != nil {
		logger.Error("planetHandler.add", "c.JSON", err, planet)
		return errs.Unexpected
	}
	return nil
}

func (h planetHandler) delete(c echo.Context) error {
	name := c.Param("name")
	if err := h.svc.Delete(name); err != nil {
		return err
	}
	return nil
}
