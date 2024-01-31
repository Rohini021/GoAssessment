package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	StringField string  `json:"stringField"`
	IntField    int     `json:"intField"`
	BoolField   bool    `json:"boolField"`
	PtrField    *string `json:"ptrField"`
	RuneField   string  `json:"runeField"`
}

func main() {
	router := gin.Default()
	router.POST("/api/bind", func(c *gin.Context) {
		var dataVar data

		if err := c.ShouldBindJSON(&dataVar); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("Received: %+v\n", dataVar)
		c.JSON(http.StatusOK, gin.H{"message": "Request successfully processed"})
	})
	router.Run(":8080")
}

// use command to test: curl -X POST -H "Content-Type: application/json" -d '{"stringField":"Hello","intField":42,"boolField":true,"ptrField":"Pointer","runeField":"A"}' http://localhost:8080/api/bind
