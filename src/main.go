package main

import (
	"view"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := &view.ViewEngine{
		Engine: gin.Default(),
	}

	engine.InitRouter()
	engine.Run()
}
