package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertMovie(c *gin.Context) {
	var movie models.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err,
		})
	}
	m, err := h.Repository.InsertMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err,
		})
	}

	log.Println("....Movie inserted", m)
	c.JSON(http.StatusOK, m)
}
