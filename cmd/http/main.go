package main

import (
	"log"
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
	herr, ok := err.(*errs.HTTPError)
	if !ok {
		herr = errs.Unexpected
	}
	c.JSON(herr.Status, herr)
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
