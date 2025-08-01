package routes

import (
	"btx/bts"
	"btx/core"

	"github.com/gin-gonic/gin"
)

var c = &bts.CellTower{}

func GetCellTower() gin.IRoutes {

	return core.R.GET("/", func(ctx *gin.Context) {

		towers, err := c.GetAll()

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		ctx.JSON(200, gin.H{"data": towers})
	})

}
