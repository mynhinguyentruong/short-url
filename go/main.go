package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"
)

type Person struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var person Person = Person{"1", "nhi", 25, "123 Rosehill"}

func getPerson(c *gin.Context) {
	c.JSON(http.StatusOK, person)
}

func main() {
	db, err := sql.Open("postgres", "postgres://nhi_nguyen:nhi_nguyen@localhost:5432/matches_db?sslmode=disable")

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	router := gin.Default()

	router.GET("/person", getPerson)

	router.Run("localhost:8080")
}
