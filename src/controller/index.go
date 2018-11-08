package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func Regist(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
