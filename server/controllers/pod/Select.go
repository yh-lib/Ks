package pod

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
	var pod corev1.Pod
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c, &pod)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	_, err = clientSet.CoreV1().Pods(basicInfo.NameSpace).Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "创建 pod: " + pod.Name + " 失败:" + err.Error()
		logs.Error(map[string]any{"Error": err}, "创建 pod: "+pod.Name+" 失败:")
	} else {
		returnData.Status = 200
		returnData.Message = "创建 pod: " + pod.Name + " 成功:"
	}
	c.JSON(200, returnData)
}

func List(c *gin.Context) {
	logs.Info(nil, "列表逻辑")
	var pod corev1.Pod

	returnData := config.ReturnData{}
	returnData.Data = map[string]any{}
	clientSet, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	podList, err := clientSet.CoreV1().Pods(basicInfo.NameSpace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "创建 pod: "+pod.Name+" 失败:")
		returnData.Status = 400
		returnData.Message = "查询 pod 列表失败:" + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "查询 pod 列表成功"
		returnData.Data["items"] = podList.Items
	}
	c.JSON(200, returnData)
}
