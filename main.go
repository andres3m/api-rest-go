package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type driver struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Team   string `json:"team"`
	Points int    `json:"points"`
}

var drivers = []driver{
	{ID: "1", Name: "Max Verstapen", Team: "Red Bull Racing Honda RBPT", Points: 277},
	{ID: "2", Name: "Lando Norris", Team: "McLaren Mercedes", Points: 199},
	{ID: "3", Name: "Charles Leclerc", Team: "Ferrari", Points: 177},
	{ID: "4", Name: "Oscar Piastri", Team: "McLaren Mercedes", Points: 167},
	{ID: "5", Name: "Carlos Sainz", Team: "Ferrari", Points: 162},
}

func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drivers)
}

func main() {
	router := gin.Default()
	router.GET("/drivers", getDrivers)

	router.Run("Localhost:8080")
}
