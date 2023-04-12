package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"os"
)

pgUser := os.Getenv('DB_USER');
pgPassword := os.Getenv('DB_PASSWORD');
dbName := os.Getenv('DB_NAME');

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// Handle error, e.g. return a JSON response with error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Establish a database connection
	db, err := sql.Open("postgres", "user=%s password=%s dbname=%s sslmode=disable", pgUser, pgPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL database!")
	
	
	insrt, err := db.Prepare(`INSERT INTO users (username, email)
	VALUES($1, $2) returning *`);
	if err != nil {
		log.Fatal(err)
	}
	defer insrt.Close();
	
	_, err = insrt.Exec(user.Name, user.Email);
	if err != nil {
		log.Fatal(err)
	}
	// Return a JSON response with the parsed user object
	c.JSON(http.StatusOK, user)
}