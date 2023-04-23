package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/go_book_api/pkg/common/models"
)

func (h handler) GetBookById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	result := h.DB.First(&book, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)

}
