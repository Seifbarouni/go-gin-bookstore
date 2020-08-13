package handlers

import (
	"net/http"
	"projects/GinFramework/gin-bookstore/data"
	"strconv"

	"github.com/gin-gonic/gin"
)


func Index(c *gin.Context){
	books:=data.GetBooks()
	c.HTML(200,"books.html",books)
}

func AddForm(c *gin.Context){
	c.HTML(200,"AddForm.html",nil)
}

func ProcessAddForm(c *gin.Context){
	title:=c.PostForm("title")
	author:=c.PostForm("author")
	priceString:=c.PostForm("price")

	price,err:=strconv.ParseFloat(priceString,64)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Please write a valid price.",
		})
	}else{
		book:=data.NewBook()
		book.Title=title
		book.Author=author
		book.Price=price
		data.AddBook(book)
		c.Redirect(http.StatusFound,"/books/show")
	}

}
func DeleteForm(c *gin.Context){
	books:=data.GetBooks()
	c.HTML(200,"DeleteBook.html",books)
}
func ProcessDeleteForm(c *gin.Context){
	
	idstr:=c.PostForm("ID")
	id,err:=strconv.ParseInt(idstr,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
		"message":"Please write a valid ID",
	})
	return
	}else if valid:=data.IsValidID(id); !valid{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Non existant ID",
		})
		return
	}else{
		data.DeleteBook(id)
		c.Redirect(http.StatusFound,"/books/show")
	}
}

func UpdateForm(c *gin.Context){
	c.HTML(200,"UpdateBook.html",data.GetBooks())
}

func ProcessUpdateForm(c *gin.Context){
	id,err:=strconv.ParseInt(c.PostForm("ID"),10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
		"message":"Please write a valid ID",
	})
return
}

if valid:=data.IsValidID(id); !valid {
	c.JSON(http.StatusBadRequest,gin.H{
		"message":"Non existant ID",
	})
	return
}
price,err:=strconv.ParseFloat(c.PostForm("price"),64)
if err != nil {
	c.JSON(http.StatusBadRequest,gin.H{
		"message":"Please write a valid price.",
	})
	return
}
	title:=c.PostForm("title")
	author:=c.PostForm("author")
	data.UpdateBook(id,price,title,author)
	c.Redirect(http.StatusFound,"/books/show")
}

