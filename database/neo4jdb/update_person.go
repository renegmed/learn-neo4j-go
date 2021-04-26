package neo4jdb

import (
	"learn-neo4j-go/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) UpdatePerson(person *models.Person) (*models.Person, error) {
	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return updatePersonFn(tx, person)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Person), nil
}

func updatePersonFn(tx neo4j.Transaction, person *models.Person) (interface{}, error) {
	records, err := tx.Run(
		`
		MATCH (p:Person)
		WHERE p.name = $name 
		SET p +={born: $born}
		RETURN p.name, p.born
		`,
		map[string]interface{}{
			"name": person.Name,
			"born": person.Born,
		})
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return &models.Person{
		Name: record.Values[0].(string),
		Born: record.Values[1].(int64),
	}, nil
}
