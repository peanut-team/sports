package redis

import (
	"context"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"sports/pkg/logger"
	"time"
)

type Config struct {
	Address  string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

var rdb *redis.Client
var redisConfig *Config

func initConfig(cnf *Config) {
	if cnf == nil {
		cnf = &Config{}
	}
	if cnf.Address == "" {
		cnf.Address = "localhost:6379"
	}
	if cnf.Password == "" {
		cnf.Password = "12345678"
	}
	if cnf.PoolSize == 0 {
		cnf.PoolSize = 100
	}
	redisConfig = cnf
}

// 初始化连接
func initClient(cnf *Config) (err error) {
	initConfig(cnf)

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func PubHandler(ctx context.Context, channerl string, msg interface{}) {
	n, err := rdb.Publish(ctx, channerl, msg).Result()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("%d clients received the message\n", n)
}

func SubHandler(ctx context.Context, handler func(msg *redis.Message), channerls ...string) {
	sub := rdb.Subscribe(ctx, channerls...)
	defer sub.Close()
	for msg := range sub.Channel() {
		handler(msg)
	}
}

func PubSubHandler(ctx context.Context, handler func(msg *redis.Message), channerls ...string) {
	pubSub := rdb.PSubscribe(ctx, channerls...)
	defer pubSub.Close()
	for msg := range pubSub.Channel() {
		handler(msg)
	}
}
