package books

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/go_book_api/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book
	result := h.DB.Find(&books)

	fmt.Println(books)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}
