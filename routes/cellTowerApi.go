package routes

import (
	"btx/core"

	"github.com/gin-gonic/gin"
)

func GetCellTower() gin.IRoutes {

	return core.R.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"message": "Hello, world!",
		})
	})

}
