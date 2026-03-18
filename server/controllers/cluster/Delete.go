package cluster

import "github.com/gin-gonic/gin"

func Delete(c *gin.Context) {
	returnData.Message = "集群删除成功"
	c.JSON(200, returnData)
}
