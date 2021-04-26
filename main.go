package main

import (
	"learn-neo4j-go/config"
	"learn-neo4j-go/database/neo4jdb"
	"learn-neo4j-go/routes"
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {

	configuration := config.NewNeo4jConfiguration()
	driver, err := configuration.NewDriver()
	if err != nil {
		log.Fatal(err)
	}

	defer config.CloseDriver(driver)

	repo := neo4jdb.NewDbRepository(driver)
	handler := routes.NewHandler(repo, configuration)

	router.POST("/movies", handler.InsertMovie)
	router.PUT("/movies", handler.UpdateMovie)
	router.DELETE("/movies", handler.DeleteMovie)
	router.POST("/movies/person", handler.InsertPerson)
	router.PUT("/movies/person", handler.UpdatePerson)
	router.DELETE("/movies/person", handler.DeletePerson)
	router.POST("/movies/job", handler.InsertJob)
	router.DELETE("/movies/job", handler.DeleteJob)

	log.Println("...Started server port :8080")

	log.Fatal(router.Run(":8080"))
}
