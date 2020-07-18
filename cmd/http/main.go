package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"starwars-hex/pkg/conf"
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/logger"
	"starwars-hex/pkg/planets"
	"starwars-hex/pkg/repository/mongodb"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	conf.Load()
	db := mongodb.Open(conf.MongoDBURI(), conf.MongoDBDatabaseName())

	router := createRouter(db)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("- Starting REST Api -")
		logger.Fatal("http.main", "router.Start", router.Start(conf.APIPort()))
	}()

	<-stop
	log.Println("- Stopping REST Api -")
	router.Shutdown(context.Background())
	mongodb.Close()
}

func httpErrorHandler(err error, c echo.Context) {
	var herr *errs.HTTPError
	switch err.(type) {
	case *echo.HTTPError:
		herr = errs.NewHTTPErrorFromEcho(err.(*echo.HTTPError))
	case *errs.HTTPError:
		herr = err.(*errs.HTTPError)
	default:
		logger.Error("[ECHO] main.httpErrorHandler", "", err)
		herr = errs.Unexpected
	}

	jerr := c.JSON(herr.Status, herr)
	if jerr != nil {
		logger.Error("main.httpErrorHandler", "c.JSON", jerr)
		c.Response().Header().Set("Content-type", "text/plain")
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte("an unexpected error occurred"))
	}
}

func createRouter(db *mongo.Database) *echo.Echo {
	router := echo.New()
	router.HTTPErrorHandler = httpErrorHandler
	router.HideBanner = true
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token", "Cache-Control"},
	}))

	planetDB := mongodb.NewPlanetDB(db.Collection("planets"))
	planetSvc := planets.NewService(planetDB, planets.NewSwapiClient())
	AddPlanetHandler(router, planetSvc)
	return router
}
