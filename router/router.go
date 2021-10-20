package router

import (
	h "gin/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static/")
	r.GET("/", h.Login)
	r.POST("/login", h.LoginCheck)
	r.POST("/sign", h.SignCheck)
	listGroup := r.Group("/list", h.AuthMiddle())
	{
		listGroup.GET("/", h.GetListHTML)
		listGroup.POST("/", h.AddNewList)
		listGroup.DELETE("/:title", h.DeleteList)
		listGroup.PUT("/", h.UpdateList)
		listGroup.GET("/all", h.GetAllList)
	}
	return r
}
