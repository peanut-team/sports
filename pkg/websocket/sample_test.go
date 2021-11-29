package websocket

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sports/pkg/api/coach"
	"testing"
	"time"
)

func Test_server(t *testing.T) {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func Test_client(t *testing.T) {
	flag.Parse()
	log.SetFlags(0)

	// 用来接收命令行的终止信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 和服务端建立连接
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/v1/coach/training"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			// 从接收服务端message
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			// 向服务端发送message
			data := []*coach.AthleteTraining{
				{
					AthleteID: 1,
				},
				{
					AthleteID: 2,
				},
				{
					AthleteID: 3,
				},
			}
			msg, err := json.Marshal(data)
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// 收到命令行终止信号，通过发送close message关闭连接。
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			// 收到接收协程完成的信号或者超时，退出
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}


func Test_client_2(t *testing.T) {
	flag.Parse()
	log.SetFlags(0)

	// 用来接收命令行的终止信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 和服务端建立连接
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/v1/coach/training"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			// 从接收服务端message
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			// 向服务端发送message
			data := []*coach.AthleteTraining{
				{
					AthleteID: 1,
				},
				{
					AthleteID: 2,
				},
				{
					AthleteID: 3,
				},
			}
			msg, err := json.Marshal(data)
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// 收到命令行终止信号，通过发送close message关闭连接。
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			// 收到接收协程完成的信号或者超时，退出
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
