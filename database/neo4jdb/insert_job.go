package neo4jdb

import (
	"learn-neo4j-go/models"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) InsertJob(person *models.Job) (*models.Job, error) {

	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return createJobFn(tx, person)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Job), nil
}

func createJobFn(tx neo4j.Transaction, job *models.Job) (interface{}, error) {

	log.Println("...Job:\n", job)

	records, err := tx.Run(
		`
		MATCH (m:Movie)
		WHERE m.title = $title
		MATCH (p:Person)
		WHERE p.name = $name
		MERGE (p)-[rel:`+job.Job+`]->(m)
		ON CREATE SET rel.roles = $roles
		RETURN p.name,rel,m.title	
		`,
		map[string]interface{}{
			"title": job.Movie,
			"name":  job.Name,
			//"job":   job.Job,
			"roles": job.Roles,
		})

	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	return &models.Job{
		Name:  record.Values[0].(string),
		Movie: record.Values[2].(string),
	}, nil
}
