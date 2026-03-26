package namespace

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	logs.Info(nil, "删除逻辑")
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
	}
	// 禁止删除保护的namespace
	if config.ProtectNameSpace[basicInfo.NameSpace] == true {
		returnData.Status = 400
		returnData.Message = basicInfo.NameSpace + " 禁止删除"
		c.JSON(200, returnData)
		return
	}
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.NameSpace, metav1.DeleteOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "删除 namespace: " + basicInfo.NameSpace + " 失败:" + err.Error()
		logs.Error(map[string]any{"Error": err}, "删除 namespace: "+basicInfo.NameSpace+" 失败:")
	} else {
		returnData.Status = 200
		returnData.Message = "删除 namespace: " + basicInfo.NameSpace + " 成功:"
	}
	c.JSON(200, returnData)
}
