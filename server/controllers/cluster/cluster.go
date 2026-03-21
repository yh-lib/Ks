package cluster

// 定义一个结构体，用于描述创建集群所用的配置信息
type ClusterInfo struct {
	Id       string `json:"id"`
	Alias    string `json:"alias"`
	Location string `json:"city"`
}

type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}
