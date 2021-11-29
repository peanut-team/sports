package redis
//
//import (
//	"context"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestSubClient(t *testing.T) {
//	cnf := &Config{}
//	err := initClient(cnf)
//	assert.NoError(t, err)
//	ctx := context.Background()
//
//	handler := func(msg *redis.Message) {
//		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
//	}
//	SubHandler(ctx, handler, "chat")
//}
//
//func TestPubSubClient(t *testing.T) {
//	cnf := &Config{}
//	err := initClient(cnf)
//	assert.NoError(t, err)
//	ctx := context.Background()
//
//	handler := func(msg *redis.Message) {
//		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
//	}
//	PubSubHandler(ctx, handler, "chat")
//
//	ctx.Deadline()
//}
//
//func TestPubClient(t *testing.T) {
//	cnf := &Config{}
//	err := initClient(cnf)
//	assert.NoError(t, err)
//	ctx := context.Background()
//
//	PubHandler(ctx, "chat", "hello")
//}
//
