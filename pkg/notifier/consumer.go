package notifier

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"sports/pkg/api/coach"
	"sports/pkg/logger"
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
	//go c.publishMessages()
}

// GetMessages listens RabbitMQ and send new messages into Messages channel
func (c *Consumer) GetMessages() {
	channel := c.GetChannel()
	defer channel.Close()

	err := channel.ExchangeDeclare(*exchange, "topic", false, false, false, false, nil)
	checkError(err)
	for message := range c.GetDeliveries(*queue, *routing, channel) {
		parsedMessage, err := ParseMessage(message)
		if err != nil {
			logger.Error(err)
			continue
		}
		c.Messages <- *parsedMessage
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
		logger.Infof("Check your RabbitMQ connection!")
	}
}

// GetDeliveries returns channel with messages from RabbitMQ, for particular queue
func (c *Consumer) GetDeliveries(qname, routing string, channel *amqp.Channel) <-chan amqp.Delivery {
	queue, err := channel.QueueDeclare(qname, false, false, false, false, nil)
	checkError(err)

	err = channel.QueueBind(queue.Name, routing, *exchange, false, nil)
	checkError(err)

	deliveries, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	checkError(err)
	return deliveries
}

// ParseMessage parses amqp.Delivery and returns Message instance
func ParseMessage(message amqp.Delivery) (*coach.AthleteTraining, error) {
	message.Ack(false)
	matches := re.FindStringSubmatch(message.RoutingKey)
	if len(matches) == 0 {
		return nil, errors.New("Unknown routing key ")
	}
	result := &coach.AthleteTraining{}
	err := json.Unmarshal(message.Body, result)
	return result, err
}
