package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateMovie(c *gin.Context) {
	var movie models.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error update movie",
			"message": err,
		})
	}
	m, err := h.Repository.UpdateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error update movie",
			"message": err,
		})
	}

	log.Println("....Movie updated", m)
	c.JSON(http.StatusOK, m)
}
