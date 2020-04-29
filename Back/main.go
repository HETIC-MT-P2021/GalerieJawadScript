package main

import (
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/api"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/database"
	"github.com/gin-gonic/gin"
)

func main() {

	// initializes database
	db, _ := database.Initialize()

	port := "8000"
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)   // listen to given port
}