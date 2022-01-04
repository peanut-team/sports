package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"sports/internal/account/service"
	"sports/pkg/api/coach"
	"sports/pkg/logger"
	"strconv"
)

var (
	cRouting = "game_user_data_upload_route"
	cQueue   = "game_user_data_upload_queue"
)

// Consumer is listening RabbitMQ and send messages to Registry
type Consumer struct {
	Messages chan coach.AthleteTraining

	connection *amqp.Connection
}

// NewConsumer is constructor for Consumer
func NewConsumer() *Consumer {
	return &Consumer{Messages: make(chan coach.AthleteTraining)}
}

// Run starts two goroutine
// first is listening RabbitMQ for new messages
// second - publishes messages for offline users
func (c *Consumer) Run() {
	conn, err := amqp.Dial(*rabbit)
	if err != nil {
		panic(err)
	}
	c.connection = conn

	go c.GetMessages()
}

// GetMessages listens RabbitMQ and send new messages into Messages channel
func (c *Consumer) GetMessages() {
	channel := c.GetChannel()
	defer channel.Close()

	for message := range c.GetDeliveries(channel) {
		parsedMessage, err := ParseMessage(message)
		if err != nil {
			logger.Error(err)
			continue
		}
		if parsedMessage != nil {
			logger.Infof("sub mq data: %v", parsedMessage)
			c.Messages <- *parsedMessage
		}
	}
}

// GetChannel returns new AMQP channel
func (c *Consumer) GetChannel() *amqp.Channel {
	channel, err := c.connection.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}

func checkError(err error) {
	if err != nil {
		logger.Errorf("Check your RabbitMQ connection! err: %#v", err)
	}
}

// GetDeliveries returns channel with messages from RabbitMQ, for particular queue
func (c *Consumer) GetDeliveries(channel *amqp.Channel) <-chan amqp.Delivery {
	deliveries, err := channel.Consume(cQueue, "", false, false, false, false, nil)
	checkError(err)
	return deliveries
}

// ParseMessage parses amqp.Delivery and returns Message instance
func ParseMessage(message amqp.Delivery) (*coach.AthleteTraining, error) {
	message.Ack(false)
	source := &UploadData{}
	err := json.Unmarshal(message.Body, source)
	if err != nil {
		return nil, err
	}

	if source.UserId == "" {
		return nil, nil
	}
	// trans model
	result, err := NewAthleteTraining(source)

	return result, err
}

type UploadData struct {
	UserId       string                `json:"userId"`       // 用户ID
	Status       coach.SportsmanStatus `json:"status"`       // 离线：0，在线：1，训练中：2
	Distance     float64               `json:"distance"`     // 学员训练距离，单位：m
	PotSpeed     float64               `json:"spotSpeed"`    // 加速度，单位：m/s2（米每二次方秒）
	AvguserSpeed float64               `json:"avguserSpeed"` // 平均速度，单位：m/s
	NotifyTime   int64                 `json:"notifyTime"`   // 开始时间
}

func NewAthleteTraining(data *UploadData) (*coach.AthleteTraining, error) {
	// userId
	id, err := strconv.ParseInt(data.UserId, 10, 64)
	if err != nil {
		logger.Warnf("invalid userId: %s, err %v", data.UserId, err)
		return nil, err
	}

	// query user info
	user, err := service.GetUser(int(id))
	if err != nil {
		return nil, fmt.Errorf("can not fetch user[%d], err: %v", data.UserId, err)
	}
	return &coach.AthleteTraining{
		StartTime: data.NotifyTime,
		// TODO change
		SportImg:           "https://pic1.zhimg.com/80/v2-6c5ff4ef0bb78991ed03ab720f1b2447_720w.jpg?source=1940ef5c",
		AthleteID:          int(user.ID),
		AthleteName:        user.Username,
		Status:             data.Status,
		Distance:           data.Distance,
		InstantaneousSpeed: data.PotSpeed,
		AverageSpeed:       data.AvguserSpeed,
	}, nil
}
