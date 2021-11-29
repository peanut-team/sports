package notifier

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
)

func TestMessageParsing(t *testing.T) {
	headers := make(amqp.Table, 0)
	var ttl int32 = 13
	headers["ttl"] = ttl
	body := "test message"
	message := amqp.Delivery{
		RoutingKey: "user.18",
		Body:       []byte(body),
		Headers:    headers,
	}
	parsed, _ := ParseMessage(message)
	fmt.Sprintf("msg；%v", parsed)
}
