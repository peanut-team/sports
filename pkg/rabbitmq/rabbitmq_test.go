package rabbitmq

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"log"
	"sports/pkg/api/coach"
	"testing"
)

var username = "root"
var password = "maxwit2021"
var host = "121.199.27.6"
var port = "5672"

func TestPub(t *testing.T) {
	var exchange = flag.String("exchange", "notifications", "Durable, non-auto-deleted AMQP exchange name")
	var routing  = flag.String("routing key", "test.*", "Routing key for queue")
	var queue    = flag.String("queue", "notifications", "Queue name")

	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	connAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	conn, err := amqp.Dial(connAddr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()
	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return
	}
	err = ch.ExchangeDeclare(*exchange, "topic", false, false, false, false, nil)
	assert.NoError(t, err)


	defer ch.Close()
	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		*queue, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return
	}
	err = ch.QueueBind(q.Name, *routing, *exchange, false, nil)
	assert.NoError(t, err)

	// 向服务端发送message
	data := &coach.AthleteTraining{
		AthleteID: 2,
		Status:    coach.SportsmanStatus_Online,
	}
	body,_  := json.Marshal(data)
	// 4.将消息发布到声明的队列
	err = ch.Publish(
		*exchange,     // exchange
		"test.1", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return
	}
	log.Printf(" [x] Sent %s", body)
}

func TestSub(t *testing.T) {
	// 建立连接
	connAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	conn, err := amqp.Dial(connAddr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	// 获取channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	// 声明队列
	//请注意，我们也在这里声明队列。因为我们可能在发布者之前启动使用者，所以我们希望在尝试使用队列中的消息之前确保队列存在。
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return
	}
	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
