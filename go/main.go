package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type Person struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var person Person = Person{"1", "nhi", 25, "123 Rosehill"}
var db *sql.DB

func getPerson(c *gin.Context) {
	c.JSON(http.StatusOK, person)
}

func main() {
	// db, err := sql.Open("postgres", "postgres://nhi_nguyen:nhi_nguyen@localhost:5432/matches_db?sslmode=disable")

	// if err != nil {
	// 	// handle error
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// err = db.Ping()

	// if err != nil {
	// 	panic(err)
	// }

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	fmt.Println("Successfully connected!")

	// albumID, err := addAlbum(Album{Title: "Hahah", Artist: "Nhi", Price: 99.99}, db)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albumID)

	router := gin.Default()

	router.GET("/person", getPerson)
	router.POST("/albums", addNewAlbum)
	router.GET("/albums/:id", getAlbum)

	router.Run("localhost:8080")
}

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	var album Album

	if err != nil {
		panic(err)
	}
	if err := db.QueryRow("SELECT * FROM album WHERE id = (?)", i).Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, sql.ErrNoRows)
		}
	}
	c.JSON(http.StatusOK, album)

}

func addNewAlbum(c *gin.Context) {
	var album Album
	if err := c.BindJSON(&album); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	albumID, err := addAlbum(album, db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, albumID)
}

func addAlbum(album Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", album.Title, album.Artist, album.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
