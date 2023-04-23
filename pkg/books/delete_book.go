package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandeymangg/go_book_api/pkg/common/models"
)

type Success struct {
	status  string
	message string
}

func (h handler) DeleteBook(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	result := h.DB.First(&book, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	success := Success{
		status:  "Success",
		message: "Deleted!",
	}

	c.JSON(http.StatusOK, success)
}
