package mqtt

var broker = "127.0.0.1"
var port = 1883

//func TestMqttClientPub(t *testing.T) {
//	opts := mqtt.NewClientOptions()
//	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
//	opts.SetClientID("go_mqtt_client_1")
//	opts.SetUsername("emqx")
//	opts.SetPassword("public")
//	opts.SetDefaultPublishHandler(messagePubHandler)
//	opts.OnConnect = connectHandler
//	opts.OnConnectionLost = connectLostHandler
//	client := mqtt.NewClient(opts)
//	if token := client.Connect(); token.Wait() && token.Error() != nil {
//		panic(token.Error())
//	}
//
//	sub(client)
//	publish(client)
//
//	client.Disconnect(1250)
//}