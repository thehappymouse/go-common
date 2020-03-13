package rabbitmq

// 基本的一个绑定信息
type BindInfo struct  {
	Exchange string
	QueueName string
	RouterKey string
	Qos int
	//ExchangeName() string // 获取接收者绑定的交换机
	//QueueName() string     // 获取接收者需要监听的队列
	//RouterKey() string     // 这个队列绑定的路由
	//Qos() int              //获取设置的 Qos 数量
}
