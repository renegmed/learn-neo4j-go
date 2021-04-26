package neo4jdb

import (
	"learn-neo4j-go/models"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) DeleteJob(person *models.Job) (*models.Job, error) {

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

func deleteJobFn(tx neo4j.Transaction, job *models.Job) (interface{}, error) {

	log.Println("...Job:\n", job)

	records, err := tx.Run(
		`
		MATCH (m:Movie)
		WHERE m.title = $title
		MATCH (p:Person)
		WHERE p.name = $name AND m.title = $title AND 
			exists((p) - [:`+job.Job+`]->(m)) 
		DELETE (p) - [:`+job.Job+`]->(m)
		RETURN p.name, m.title	
		`,
		map[string]interface{}{
			"title": job.Movie,
			"name":  job.Name,
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
		Movie: record.Values[1].(string),
		Job:   job.Job,
	}, nil
}
