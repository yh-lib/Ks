package namespace

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	logs.Info(nil, "更新逻辑")
}
