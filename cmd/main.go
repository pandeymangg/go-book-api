package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/go_book_api/pkg/books"
	"github.com/pandeymangg/go_book_api/pkg/common/db"
)

func main() {
	port := ":8080"
	dbUrl := "your db url"

	r := gin.Default()
	h := db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	books.RegisterRoutes(r, h)

	r.Run(port)
}
