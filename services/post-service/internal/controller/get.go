package controller

import "github.com/gin-gonic/gin"

func get(c *gin.Context) {
	c.JSON(200, "yunxinxinshidabendan")
}
