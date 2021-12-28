package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sports/pkg/api/coach"
	"sports/pkg/logger"
)

// UserConnection represents user ws-connection and his UID
type UserConnection struct {
	UID          int
	ws           *websocket.Conn
	trainingDate map[int]*coach.AthleteTraining
}

// NewUserConnection constructor for UserConnection
func NewUserConnection(ws *websocket.Conn) (*UserConnection, error) {
	if ws == nil {
		return nil, fmt.Errorf("invalide ws(%v)", ws)
	}
	uid := int(uuid.New().ID())
	return &UserConnection{ws: ws, UID: uid}, nil
}

// Send method deliver message on websocket
func (u *UserConnection) Send(data []coach.AthleteTraining) {
	msg, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed Marshal AthleteTraining %v: %v", data, msg)
	} else {
		u.ws.WriteMessage(websocket.TextMessage, msg)
	}
}

// Listen method listens ws-connection and tries to stop someone training
func (u *UserConnection) Listen() {
	defer func() {
		u.ws.Close()
		registry.Unregister(u)
	}()

	for {
		_, wsMsg, err := u.ws.ReadMessage()
		if err != nil {
			logger.Errorf("user connection read err: %v ", err)
			break
		}
		if wsMsg != nil {
			logger.Infof("get ws msg: %s", wsMsg)
		}

		if u.trainingDate == nil {
			u.trainingDate = make(map[int]*coach.AthleteTraining)
		}
		registry.Register(u, wsMsg)
	}
}
