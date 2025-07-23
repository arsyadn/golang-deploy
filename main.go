package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID int   `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Arsyad", Email: "arsyadnurzaky@gmail.com", Password: "password123"},
	{ID: 2, Name: "Alyssa", Email: "alyssa@gmail.com", Password: "password456"},
}

func main() {
	_ = godotenv.Load()

	r := gin.Default()


	r.GET("/v1/users", func(c *gin.Context) {
		var userReponse []UserResponse
		for _, user := range users {
			userReponse = append(userReponse, UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			})
		}
		c.JSON(200, gin.H{
			"users": userReponse,
		})
	})

	r.POST("/v1/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		newUser.ID = len(users) + 1 
		users = append(users, newUser)

		c.JSON(201, gin.H{
			"message": "User created successfully",
			"user": UserResponse{
				ID:    newUser.ID,
				Name:  newUser.Name,
				Email: newUser.Email,
			},
		})
	})

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping Arsyad!",
		})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Arsyad!",
		})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adalah hands-on deployment VPS pada bootcamp dibimbing",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	r.Run(":" + port)
}
