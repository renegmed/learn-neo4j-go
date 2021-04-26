package config

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jConfiguration struct {
	URL          string
	Username     string
	Password     string
	Databasename string
}

func NewNeo4jConfiguration() *Neo4jConfiguration {
	neo4jUri, found := os.LookupEnv("NEO4J_URI")
	if !found {
		panic("NEO4J_URI not set")
	}
	neo4jUsername, found := os.LookupEnv("NEO4J_USERNAME")
	if !found {
		panic("NEO4J_USERNAME not set")
	}
	neo4jPassword, found := os.LookupEnv("NEO4J_PASSWORD")
	if !found {
		panic("NEO4J_PASSWORD not set")
	}
	neo4jDatabasename, found := os.LookupEnv("NEO4J_DATABASE_NAME")
	if !found {
		panic("NEO4J_DATABASE not set")
	}

	return &Neo4jConfiguration{
		URL:          neo4jUri,      // "neo4j://neo4j:7687",
		Username:     neo4jUsername, // "neo4j",
		Password:     neo4jPassword, // "testing",
		Databasename: neo4jDatabasename,
	}
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.URL, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func CloseDriver(driver neo4j.Driver) error {
	return driver.Close()
}
