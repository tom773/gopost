package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	// Routes
	e.GET("/api/health", hello)
	e.POST("/api/return", returnd)
	e.GET("/api/users/:id", getUser)
	e.GET("/api/users", createUser)
	e.POST("/api/posts", updateUser)
	e.POST("/api/users", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8070"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

type Request struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func returnd(c echo.Context) error {
	var request Request
	if c.Bind(&request) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	jsonr, err := json.Marshal(request.Message)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	return c.JSON(http.StatusOK, jsonr)
}

// Fake in-memory database
var users = map[string]*User{}
var request = Request{}

// Handlers
func getUser(c echo.Context) error {
	id := c.Param("id")
	user, ok := users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	users[user.ID] = user
	return c.JSON(http.StatusCreated, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	user, ok := users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	users[id] = user
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	_, ok := users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
