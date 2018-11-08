package view

import (
	"controller"
	"view/middleware"

	"github.com/gin-gonic/gin"
)

type ViewEngine struct {
	Engine *gin.Engine
}

type ViewGroup struct {
	Group *gin.RouterGroup
}

func (this *ViewEngine) Run() {
	this.Run()
}

func (this *ViewEngine) InitRouter() {
	group := ViewGroup{
		Group: this.Engine.Group("/api/v1"),
	}

	group.Group.Use(middleware.JWTMiddleware)
	group.Group.POST("/login", controller.Login)
	group.Group.POST("/regist", controller.Regist)
}
