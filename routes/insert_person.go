package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertPerson(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error insert person",
			"message": err,
		})
	}
	p, err := h.Repository.InsertPerson(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err,
		})
	}

	log.Println("....Person inserted\n", p)
	c.JSON(http.StatusOK, p)
}
