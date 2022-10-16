package main

import (
	"restful-api/config"
	"restful-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/persons/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/persons", inDB.CreatePerson)
	router.PUT("/persons/:id", inDB.UpdatePerson)
	router.DELETE("/persons/:id", inDB.DeletePerson)

	router.Run(":3000")
}
