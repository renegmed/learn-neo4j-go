package neo4jdb

import (
	"learn-neo4j-go/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) DeletePerson(person *models.Person) (*models.Person, error) {
	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return deletePersonFn(tx, person)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Person), nil
}

func deletePersonFn(tx neo4j.Transaction, person *models.Person) (interface{}, error) {
	_, err := tx.Run(
		`
	MATCH (p:Person {name: $name})
	DETACH DELETE p		
	`,
		map[string]interface{}{
			"name": person.Name,
		})

	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return &models.Person{
		Name: person.Name,
	}, err
}
