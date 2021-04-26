package neo4jdb

import (
	"learn-neo4j-go/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) DeleteMovie(movie *models.Movie) (*models.Movie, error) {
	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return deleteMovieFn(tx, movie)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Movie), nil
}

func deleteMovieFn(tx neo4j.Transaction, movie *models.Movie) (interface{}, error) {
	_, err := tx.Run(
		`
	MATCH (m:Movie {title: $title})
	DETACH DELETE m		
	`,
		map[string]interface{}{
			"title": movie.Title,
		})

	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return &models.Movie{
		Title: movie.Title,
	}, err
}
