package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/go_book_api/pkg/common/models"
)

type UpdateBookBodyType struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *gin.Context) {
	var body UpdateBookBodyType

	err := c.BindJSON(&body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")

	var book models.Book

	result := h.DB.First(&book, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)

}
