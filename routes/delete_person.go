package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeletePerson(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error delete person",
			"message": err,
		})
	}
	p, err := h.Repository.DeletePerson(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error delete person",
			"message": err,
		})
	}

	log.Println("....Person deleted\n", p)
	c.JSON(http.StatusOK, p)
}
