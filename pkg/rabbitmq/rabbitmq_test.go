package rabbitmq

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sports/pkg/api/coach"
	"testing"
	"time"
)

var username = "root"
var password = "maxwit2021"
var host = "121.199.27.6"
var port = "5672"

var (
	cRouting = "game_user_data_upload_route"
	cQueue   = "game_user_data_upload_queue"
)


type UploadData struct {
	UserId       int                   `json:"userId"`       // 用户ID
	Status       coach.SportsmanStatus `json:"status"`       // 离线：0，在线：1，训练中：2
	Distance     float64               `json:"distance"`     // 学员训练距离，单位：m
	PotSpeed     float64               `json:"spotSpeed"`    // 加速度，单位：m/s2（米每二次方秒）
	AvguserSpeed float64               `json:"avguserSpeed"` // 平均速度，单位：m/s
	NotifyTime   int64                 `json:"notifyTime"`   // 开始时间
}

func TestPub(t *testing.T) {
	var exchange = flag.String("exchange", "canoe_exchange", "Durable, non-auto-deleted AMQP exchange name")

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

	defer ch.Close()

	assert.NoError(t, err)

	// 用来接收命令行的终止信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	now := time.Now().Unix()

	for {
		select {
		case _ = <-ticker.C:
			aa := []int{1}
			for _, a := range aa {
				// 向服务端发送message
				t := rand.Intn(100)
				data := &UploadData{
					UserId:       a,
					Status:       coach.MatchType_Offline,
					Distance:     10.1 + float64(t),
					PotSpeed:     44.1 + float64(t),
					AvguserSpeed: 22.1 + float64(t),
					NotifyTime:   now,
				}


				body,_  := json.Marshal(data)
				// 4.将消息发布到声明的队列
				err = ch.Publish(
					*exchange,     // exchange
					cRouting, // routing key
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


		case <-interrupt:
			log.Println("interrupt")

			select {
			case <-time.After(time.Second):
			}
			return
		}
	}
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
