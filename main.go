package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sollambert/gin-test/routes"
	"github.com/sollambert/gin-test/auth"
	"net/http"
)

func main() {
	// Create new router
	r := gin.Default()

	// Configure Authboss
	ab := configureAuthboss()

	// Register Authboss routes and handlers
	ab.RegisterAuthRoutes(r.Group("/auth"))

	r.POST("/user", routes.AddUser)
	r.GET("/", func(c *gin.Context) {
		// Check if user is authenticated
		if ab.CurrentUser(c.Request) == nil {
			// User is not authenticated, redirect to login page
			c.Redirect(http.StatusFound, "/auth/login")
			return
		}

		// User is authenticated, show welcome page
		c.JSON(200, gin.H{"message": "Welcome, " + ab.CurrentUser(c.Request).GetPID()})
	})

	r.GET("/logout", func(c *gin.Context) {
		// Log out user
		err := ab.Logout(c.Writer, c.Request)
		if err != nil {
			// Handle error
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Redirect to login page after logout
		c.Redirect(http.StatusFound, "/auth/login")
	})

	r.Run(":8080")
}
