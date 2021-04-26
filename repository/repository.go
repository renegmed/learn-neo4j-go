package repository

import (
	"learn-neo4j-go/models"
)

type DbRepository interface {
	InsertMovie(movie *models.Movie) (*models.Movie, error)
	UpdateMovie(movie *models.Movie) (*models.Movie, error)
	DeleteMovie(movie *models.Movie) (*models.Movie, error)
	InsertPerson(person *models.Person) (*models.Person, error)
	UpdatePerson(person *models.Person) (*models.Person, error)
	DeletePerson(person *models.Person) (*models.Person, error)
	InsertJob(job *models.Job) (*models.Job, error)
	DeleteJob(job *models.Job) (*models.Job, error)
}
