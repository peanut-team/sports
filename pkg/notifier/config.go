package notifier

import (
	"flag"
)

var (
	rabbit   = flag.String("rabbit", "amqp://root:maxwit2021@121.199.27.6:5672/", "AMQP URI")
	exchange = flag.String("exchange", "canoe_exchange", "Durable, non-auto-deleted AMQP exchange name")
	registry *Registry
)

type Server struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type ConfigValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Usage string `json:"usage"`
}

type RabbitConfig struct {
	Server   Server      `json:"server"`
	Exchange ConfigValue `json:"exchange"`
	Routing  ConfigValue `json:"routing"`
	Queue    ConfigValue `json:"queue"`
}

func InitNotifier() {
	registry = NewRegistry()
	go registry.ListenAndSendMessages()
}
