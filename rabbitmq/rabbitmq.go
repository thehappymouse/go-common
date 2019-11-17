package rabbitmq

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

// RabbitMQ 用于管理和维护rabbitmq的对象
type RabbitMQ struct {
	wg            sync.WaitGroup
	channel       *amqp.Channel
	connectString string //连接字符串
	exchangeName  string // exchange的名称
	exchangeType  string // exchange的类型
	receivers     []Receiver
}

// New 创建一个新的操作RabbitMQ的对象
func New(exchange, exchangeType, connect string) *RabbitMQ {
	// 这里可以根据自己的需要去定义
	return &RabbitMQ{
		exchangeName:  exchange,
		exchangeType:  exchangeType,
		connectString: connect,
	}
	//amqp.ExchangeTopic
}

// prepareExchange 声明交换机
func (mq *RabbitMQ) prepareExchange() error {
	// 申明Exchange
	err := mq.channel.ExchangeDeclare(
		mq.exchangeName, // exchange
		mq.exchangeType, // type
		true,            // durable 持久化
		false,           // autoDelete 自动删除
		false,           // internal
		false,           // noWait  异步的
		nil,             // args
	)

	if nil != err {
		return err
	}
	return nil
}

// Start 启动Rabbitmq的客户端
func (mq *RabbitMQ) Start() {
	for {
		mq.run()
		// 一旦连接断开，那么需要隔一段时间去重连
		time.Sleep(5 * time.Second)
	}
}

// RegisterReceiver 注册一个用于接收指定队列指定路由的数据接收者
func (mq *RabbitMQ) RegisterReceiver(receiver Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

// run 开始获取连接并初始化相关操作
// todo 隐藏密码信息
func (mq *RabbitMQ) run() {
	log.Debug().Msgf("尝试连接:%s", mq.connectString)
	conn, err := amqp.Dial(mq.connectString)
	if err != nil {
		log.Error().Msgf("[%s]连接失败，将重连", mq.connectString)
		return
	}
	defer conn.Close()

	mq.channel, err = conn.Channel()
	if err != nil {
		log.Error().Msgf("[%s]获取通道失败，将重连", err)
		return
	}
	defer mq.channel.Close()

	// 初始化Exchange
	mq.prepareExchange()
	log.Info().Msgf("[%s]已连接", mq.connectString)

	for _, receiver := range mq.receivers {
		mq.wg.Add(1)
		go mq.listen(receiver) // 每个接收者单独启动一个goroutine用来初始化queue并接收消息
	}

	mq.wg.Wait()

	log.Error().Msg("所有处理队列的携程都意外退出了，即将重新开始")
}

// Listen 监听指定路由发来的消息
// 这里需要针对每一个接收者启动一个goroutine来执行listen
// 该方法负责从每一个接收者监听的队列中获取数据，并负责重试
func (mq *RabbitMQ) listen(receiver Receiver) {
	defer mq.wg.Done()

	// 这里获取每个接收者需要监听的队列和路由
	queueName := receiver.QueueName()
	routerKey := receiver.RouterKey()

	// 申明队列 todo 默认开始持久化了
	_, err := mq.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if nil != err {
		// 当队列初始化失败的时候，需要告诉这个接收者相应的错误
		receiver.OnError(fmt.Errorf("初始化队列 %s 失败: %s", queueName, err.Error()))
	}

	// 将Queue绑定到Exchange上去
	err = mq.channel.QueueBind(
		queueName,       // queue name
		routerKey,       // routing key
		mq.exchangeName, // exchange
		false,           // no-wait
		nil,
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("绑定队列 [%s - %s] 到交换机失败: %s", queueName, routerKey, err.Error()))
	}
	log.Info().Msgf("队列已绑定:[%s][%s][%s]", mq.exchangeName, queueName, routerKey)

	// 获取消费通道 todo 功效
	err = mq.channel.Qos(receiver.Qos(), 0, true) // 确保rabbitmq会一个一个发消息
	if err != nil {
		receiver.OnError(fmt.Errorf("设置 Qos 失败: %s", queueName, err.Error()))
	}
	// consumerTag 为空
	messages, err := mq.channel.Consume(queueName, "", false, false, false, false, nil)
	if nil != err {
		receiver.OnError(fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error()))
	}
	log.Warn().Msgf("[*][%s] Waiting for messages\n", queueName)
	// 使用callback消费数据
	for msg := range messages {
		//log.Debug().Msgf("[*] receiver new msg:%s", msg.Body)
		// 当接收者消息处理失败的时候，
		// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
		// 通过重试可以成功的操作，那么这个时候是需要重试的
		// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
		if !receiver.OnReceive(msg.Body) {
			log.Warn().Msg("receiver 数据处理失败，重启程序时将重试")
		} else {
			// 确认收到本条消息, multiple必须为false
			msg.Ack(false)
		}
	}
}
