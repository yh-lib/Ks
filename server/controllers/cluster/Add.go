package cluster

import (
	"server/config"

	"github.com/gin-gonic/gin"
)

var returnData = config.NewRetrunData()

func Add(c *gin.Context) {
	returnData.Message = "集群添加成功"
	c.JSON(200, returnData)
}
