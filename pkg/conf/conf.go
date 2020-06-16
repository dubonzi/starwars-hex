package conf

import (
	"fmt"
	"os"
)

var (
	mongoDBURI          = "mongodb://localhost:27017/"
	mongoDBDatabaseName = "starwars"
	apiPort             = "9080"
	swapiURL            = "https://swapi.dev/api"
)

// Load config variables from the environment.
func Load() {
	if os.Getenv("MONGODB_URI") != "" {
		mongoDBURI = os.Getenv("MONGODB_URI")
	}
	if os.Getenv("MONGODB_DATABASE_NAME") != "" {
		mongoDBDatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	}
	if os.Getenv("API_PORT") != "" {
		apiPort = os.Getenv("API_PORT")
	}
	if os.Getenv("SWAPI_URL") != "" {
		swapiURL = os.Getenv("SWAPI_URL")
	}

}

// MongoDBURI URI used to connect to the MongoDB instance.
//	Default: "mongodb://localhost:27017/"
func MongoDBURI() string {
	return mongoDBURI
}

// MongoDBDatabaseName name of the database.
//	Default: "starwars"
func MongoDBDatabaseName() string {
	return mongoDBDatabaseName
}

// APIPort returns the port for the http server.
//	Default: :9080
func APIPort() string {
	return fmt.Sprintf(":" + apiPort)
}

// SwapiURL base url for the swapi api.
//	Default: https://swapi.dev/api
func SwapiURL() string {
	return swapiURL
}
