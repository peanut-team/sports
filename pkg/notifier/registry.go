package notifier

import (
	"encoding/json"
	"sort"
	"sports/pkg/api/coach"
	"sports/pkg/logger"
	"sync"
)

var bufferTime = int64(8)

// Registry contains all web-socket connections
type Registry struct {
	sync.RWMutex
	clients  map[int]*UserConnection
	Consumer *Consumer
	Pusher   *Pusher
}

// NewRegistry creates new Registry
func NewRegistry() *Registry {
	consumer := NewConsumer()
	go consumer.Run()
	pusher := NewPusher()
	go pusher.Run()
	return &Registry{
		clients:  make(map[int]*UserConnection),
		Consumer: consumer,
		Pusher:   pusher,
	}
}

// ListenAndSendMessages gets all messages from RabbitMQ and sends them to recipients
func (r *Registry) ListenAndSendMessages() {
	for message := range r.GetMessages() {
		r.Lock()
		userConnections := r.GetConnections(message.AthleteID)

		for _, uc := range userConnections {
			copyMsg := message
			// 删除离线用户
			if copyMsg.Status == coach.MatchType_Offline {
				logger.Infof("remove athlete[%d] data: %v", copyMsg.AthleteID, copyMsg)
				delete(uc.trainingDate, copyMsg.AthleteID)
			} else {
				logger.Infof("update athlete[%d] data: %v", copyMsg.AthleteID, copyMsg)
				uc.trainingDate[copyMsg.AthleteID] = &copyMsg
			}
			td := make([]coach.AthleteTraining, 0)
			for _, t := range uc.trainingDate {
				td = append(td, *t)
			}
			sort.SliceStable(td, func(i, j int) bool {
				return td[i].AthleteID < td[j].AthleteID
			})
			logger.Infof("send AthleteTraining data to user[%s]", uc.UID)
			uc.Send(td)
		}

		r.Unlock()
	}
}

// GetMessages returns Messages channel
func (r *Registry) GetMessages() chan coach.AthleteTraining {
	return r.Consumer.Messages
}

// GetConnection returns UserConnection for uid
func (r *Registry) GetConnections(aid int) map[int]*UserConnection {
	return r.clients
}

// Register add user connection to Registry
func (r *Registry) Register(uc *UserConnection, msg []byte) {
	r.Lock()
	defer r.Unlock()
	if uc == nil {
		return
	}
	// add cliet
	if _, ok := r.clients[uc.UID]; !ok {
		r.clients[uc.UID] = uc
		logger.Infof("User %d registered", uc.UID)
		logger.Infof("Connections %d", len(r.clients))
		// parse msg is start time, need to publish
		if msg != nil {
			start := &StartTopic{}
			err := json.Unmarshal(msg, start)
			if err == nil {
				start.NotifyTime = start.NotifyTime + bufferTime
				registry.Pusher.publishMessages(start)
			} else {
				logger.Errorf("json unmarshal websocket msg failed: %v", err)
			}
		}
	} else {
		// parse msg is start time, need to publish
		if msg != nil {
			logger.Warnf("only pub one times, invalid msg: %v", msg)
		}
	}
}

// Unregister remove user ws connection from Registry, when user leaves
func (r *Registry) Unregister(uc *UserConnection) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.clients[uc.UID]; ok {
		delete(r.clients, uc.UID)
		logger.Infof("User %d leave", uc.UID)
	}
}
