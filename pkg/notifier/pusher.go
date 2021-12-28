package notifier

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"sports/pkg/logger"
)

var (
	pubRouting = "game_status_notify_by_app_route"
	pubQueue   = "game_status_notify_by_app_queue"
)

type StartTopic struct {
	NotifyTime int64 `json:"notifyTime"` // 运动员开始运动的时间戳
	Status     int   `json:"status"`     // 比赛状态，比赛开始 1， 比赛结束 0
}

// Pusher through RabbitMQ to send messages
type Pusher struct {
	connection *amqp.Connection
}

// NewPusher is constructor for Pusher
func NewPusher() *Pusher {
	return &Pusher{}
}

// Run starts two goroutine
// first is listening RabbitMQ for new messages
// second - publishes messages for offline users
func (c *Pusher) Run() {
	conn, err := amqp.Dial(*rabbit)
	if err != nil {
		panic(err)
	}
	c.connection = conn
}

// publishMessages publishes messages into queue undelivered.user.<UID>
func (c *Pusher) publishMessages(msg *StartTopic) {
	channel := c.GetChannel()
	defer channel.Close()
	err := channel.QueueBind(pubQueue, pubRouting, *exchange, false, nil)
	checkError(err)

	pMsg, err := json.Marshal(msg)
	if err != nil {
		logger.Errorf("pub failed, object [%v], err: %v", msg, err)
	}
	channel.Publish(*exchange, pubRouting, false, false, amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Body:            pMsg,
		DeliveryMode:    amqp.Transient,
		Priority:        0,
	})
	logger.Infof("Pub Start Msg: %v", msg)
}

// GetChannel returns new AMQP channel
func (c *Pusher) GetChannel() *amqp.Channel {
	channel, err := c.connection.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}
