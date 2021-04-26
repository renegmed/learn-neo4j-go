package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteJob(c *gin.Context) {
	var job models.Job
	err := c.BindJSON(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error delete job",
			"message": err,
		})
	}

	log.Println("...Job to delete:\n", job)

	j, err := h.Repository.InsertJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error on delete job",
			"message": err,
		})
	}

	log.Println("...Job deleted\n", j)
	c.JSON(http.StatusOK, j)
}
