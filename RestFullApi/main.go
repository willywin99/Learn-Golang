package main

import (
	"RestFullApi/config"
	"RestFullApi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @title Person API
// @version 1.0
// @description This is a sample service for Rest API
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /persons
func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":3000")
}
