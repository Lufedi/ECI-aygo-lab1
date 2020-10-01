package main

import (
	"aygolabone/client"
	"aygolabone/model"
	"aygolabone/repository/mongo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	//Initialize the database
	client.Connect();

	// Repositories

	textRepostitory := mongo.NewTextRepositoryMongo()

	//Register the controllers
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//controllers
	r.POST("/api/text", func(context *gin.Context) {
		var text model.Text
		context.BindJSON(&text)
		textRepostitory.Save(&text)

		texts, error := textRepostitory.GetRecent()
		if error != nil {
			context.String(http.StatusInternalServerError, fmt.Sprintf("Error getting the texts"));
		}
		context.JSON(http.StatusOK, texts);

	})

	r.GET("/api/text", func(context *gin.Context) {
		texts, error := textRepostitory.GetRecent()
		if error != nil {
			context.String(http.StatusInternalServerError, fmt.Sprintf("Error getting the texts"));
		}
		context.JSON(http.StatusOK, texts);
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run("0.0.0.0:4000")
}
