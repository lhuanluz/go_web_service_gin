package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type person struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

var persons = []person{
	{ID: 1, Name: "Jack", Age: 30, Height: 1.76},
	{ID: 2, Name: "Ane", Age: 25, Height: 1.63},
	{ID: 3, Name: "Jacob", Age: 17, Height: 1.82},
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPerson)
	router.GET("/persons/:id", getPersonByID)
	router.POST("/persons", postPerson)

	router.Run("localhost:8080")
}

func getPerson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func postPerson(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	persons = append(persons, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range persons {
		if strconv.Itoa(a.ID) == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}
