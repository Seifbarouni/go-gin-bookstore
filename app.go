package main

import (
	"projects/GinFramework/gin-bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("assets/*")
	r.GET("/books/show",handlers.Index)

	r.GET("/books/add",handlers.AddForm)
	r.POST("/books/add",handlers.ProcessAddForm)

	r.GET("/books/delete",handlers.DeleteForm)
	r.POST("/books/delete",handlers.ProcessDeleteForm)

	r.GET("/books/update",handlers.UpdateForm)
	r.POST("/books/update",handlers.ProcessUpdateForm)
	
	r.Run()

}