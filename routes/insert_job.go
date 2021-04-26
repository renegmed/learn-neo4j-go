package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertJob(c *gin.Context) {
	var job models.Job
	err := c.BindJSON(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error insert job",
			"message": err,
		})
	}

	log.Println("...Job received:\n", job)

	j, err := h.Repository.InsertJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error on insert job",
			"message": err,
		})
	}

	log.Println("....Movie Job of a person\n", j)
	c.JSON(http.StatusOK, j)
}
