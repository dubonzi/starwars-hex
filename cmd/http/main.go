package main

import (
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

	planetDB := mongodb.NewPlanetDB(db.Collection("planets"))
	planetSvc := planets.NewService(planetDB, planets.NewSwapiClient())
	AddPlanetHandler(router, planetSvc)
	return router
}
