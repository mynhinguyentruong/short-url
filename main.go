package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var person Person = Person{"nhi", 25, "123 Rosehill"}

func getPerson(c *gin.Context) {
	c.JSON(http.StatusOK, person)
}

func main() {
	router := gin.Default()

	router.GET("/person", getPerson)

	router.Run("localhost:8080")
}
