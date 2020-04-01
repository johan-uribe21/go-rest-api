package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johan-uribe21/go-rest-api/controllers"
	"github.com/johan-uribe21/go-rest-api/models"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
