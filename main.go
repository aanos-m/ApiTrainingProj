package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type animal struct {
	ID     string  `json:"ID"`
	Name   string  `json:"name"`
	DOB    string  `json:"dob"`
	Owner  string  `json:"owner"`
	Type   string  `json:"type"`
	Height string  `json:"height"`
	Weight float32 `json:"weight"`
	FavToy string  `json:"toy"`
}

var animals = []animal{
	{Name: "Goku", DOB: "1/1/10", Type: "Cat", Owner: "Tom", Height: "2ft", Weight: 50.5, FavToy: "strings"},
}

func main() {

	router := gin.Default()
	router.GET("/animals", getAnimals)
	router.POST("/animals", postAnimals)

	router.Run("localhost:8080")
}

func getAnimals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animals)
}

func postAnimals(c *gin.Context) {
	var newAnimal animal
	if err := c.BindJSON(&newAnimal); err != nil {
		return
	}

	animals = append(animals, newAnimal)
	c.IndentedJSON(http.StatusCreated, newAnimal)
}
