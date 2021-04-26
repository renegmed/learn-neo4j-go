package routes

import (
	"learn-neo4j-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdatePerson(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error update person",
			"message": err,
		})
	}
	p, err := h.Repository.UpdatePerson(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error update person",
			"message": err,
		})
	}

	log.Println("....Person updated\n", p)
	c.JSON(http.StatusOK, p)
}
