package main

import (
	"log"
	"os"
	"os/signal"
	"starwars-hex/pkg/conf"
	"starwars-hex/pkg/logger"
	"starwars-hex/pkg/planets"
	"starwars-hex/pkg/repository/mongodb"

	"github.com/labstack/echo/v4"
)

func main() {
	conf.Load()
	db := mongodb.Open(conf.MongoDBURI(), conf.MongoDBDatabaseName())

	router := echo.New()
	router.HideBanner = true
	// TODO: Refactor error handling https://echo.labstack.com/guide/error-handling
	// TODO: Refactor main to keep it small

	planetDB := mongodb.NewPlanetDB(db.Collection("planets"))
	planetSvc := planets.NewService(planetDB, planets.NewSwapiClient())
	AddPlanetHandler(router, planetSvc)

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
