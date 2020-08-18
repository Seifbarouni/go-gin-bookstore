package handlers

import (
	"net/http"
	"projects/GinFramework/gin-bookstore/data"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	books := data.GetBooks()
	c.HTML(200, "books.html", books)
}
func Redirect(c *gin.Context){
	c.Redirect(http.StatusFound,"/books/show")
}

func AddForm(c *gin.Context) {
	c.HTML(200, "AddForm.html", nil)
}

func UpdateForm(c *gin.Context){
	c.HTML(200,"UpdateBook.html",data.GetBooks())
}

func DeleteForm(c *gin.Context) {
	books := data.GetBooks()
	c.HTML(200, "DeleteBook.html", books)
}