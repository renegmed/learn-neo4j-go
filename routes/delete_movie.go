package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteMovie(c *gin.Context) {
	var movie models.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error delete movie",
			"message": err,
		})
	}

	log.Println("...Movie to delete:\n", movie)

	m, err := h.Repository.DeleteMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error on delete movie",
			"message": err,
		})
	}

	log.Println("....Movie deleted\n", m)
	c.JSON(http.StatusOK, m)
}
