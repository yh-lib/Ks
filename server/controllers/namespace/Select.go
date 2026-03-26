package namespace

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(c *gin.Context) {
	logs.Info(nil, "详情逻辑")
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
	}
	var newNameSpace corev1.Namespace
	newNameSpace.Name = basicInfo.NameSpace
	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &newNameSpace, metav1.CreateOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "创建 namespace: " + newNameSpace.Name + " 失败:" + err.Error()
		logs.Error(map[string]any{"Error": err}, "创建 namespace: "+newNameSpace.Name+" 失败:")
	} else {
		returnData.Status = 200
		returnData.Message = "创建 namespace: " + newNameSpace.Name + " 成功:"
	}
	c.JSON(200, returnData)
}

func List(c *gin.Context) {
	logs.Info(nil, "列表逻辑")
	returnData := config.ReturnData{}
	returnData.Data = map[string]any{}
	clientSet, _, err := controllers.BasicInit(c)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
	}
	listNs, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "获取 namespace 列表失败： " + err.Error()
		logs.Error(map[string]any{"Error": err}, "获取 namespace 列表失败")
	} else {
		returnData.Status = 200
		returnData.Message = "获取 namespace 列表成功"
		returnData.Data["items"] = listNs.Items
	}
	c.JSON(200, returnData)
}
