package neo4jdb

import (
	"learn-neo4j-go/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (db AppDbRepository) InsertMovie(movie *models.Movie) (*models.Movie, error) {
	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.
	session := db.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return createMovieFn(tx, movie)
	})
	if err != nil {
		return nil, err
	}
	return result.(*models.Movie), nil
}

func createMovieFn(tx neo4j.Transaction, movie *models.Movie) (interface{}, error) {
	records, err := tx.Run("CREATE (m:Movie { title: $title, released: $released, tagline: $tagline }) RETURN m.title, m.released, m.tagline",
		map[string]interface{}{
			"title":    movie.Title,
			"released": movie.Released,
			"tagline":  movie.Tagline,
		})
	// In face of driver native errors, make sure to return them directly.
	// Depending on the error, the driver may try to execute the function again.
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
