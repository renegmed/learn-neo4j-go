package neo4jdb

import (
	"learn-neo4j-go/repository"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AppDbRepository struct {
	Driver neo4j.Driver
}

func NewDbRepository(driver neo4j.Driver) repository.DbRepository {
	appRepository := AppDbRepository{
		Driver: driver,
	}
	return appRepository
}
