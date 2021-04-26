package neo4jdb

import (
	"learn-neo4j-go/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) UpdateMovie(movie *models.Movie) (*models.Movie, error) {
	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return updateMovieFn(tx, movie)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Movie), nil
}

func updateMovieFn(tx neo4j.Transaction, movie *models.Movie) (interface{}, error) {
	records, err := tx.Run(
		`
		MATCH (m:Movie)
		WHERE m.title = $title
		SET m +={released: $released, tagline: $tagline}
		RETURN m.title, m.released, m.tagline
		`,
		map[string]interface{}{
			"title":    movie.Title,
			"released": movie.Released,
			"tagline":  movie.Tagline,
		})

	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return &models.Movie{
		Title:    record.Values[0].(string),
		Released: record.Values[1].(int64),
		Tagline:  record.Values[2].(string),
	}, nil
}
