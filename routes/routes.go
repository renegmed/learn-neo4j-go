package routes

import (
	"learn-neo4j-go/config"
	"learn-neo4j-go/repository"
)

type Handler struct {
	Repository    repository.DbRepository
	Configuration *config.Neo4jConfiguration
}

func NewHandler(repo repository.DbRepository, config *config.Neo4jConfiguration) *Handler {
	return &Handler{
		Repository:    repo,
		Configuration: config,
	}
}
