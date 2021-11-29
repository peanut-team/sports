package notifier

import (
	"sort"
	"sports/pkg/api/coach"
	"sports/pkg/logger"
	"sync"
)

// Registry contains all web-socket connections
type Registry struct {
	sync.RWMutex
	clients  map[int]*UserConnection
	Consumer *Consumer
}

// NewRegistry creates new Registry
func NewRegistry() *Registry {
	consumer := NewConsumer()
	go consumer.Run()
	return &Registry{
		clients:  make(map[int]*UserConnection),
		Consumer: consumer,
	}
}

// ListenAndSendMessages gets all messages from RabbitMQ and sends them to recipients
func (r *Registry) ListenAndSendMessages() {
	for message := range r.GetMessages() {
		r.Lock()
		userConnections := r.GetConnections(message.AthleteID)

		for _, uc := range userConnections {
			uc.trainingDate[message.AthleteID] = &message
			td := make([]*coach.AthleteTraining, 0)
			for _, t := range uc.trainingDate {
				// TODO 结束训练标志需要过滤
				td = append(td, t)
			}
			sort.SliceStable(td, func(i, j int) bool {
				return td[i].AthleteID < td[j].AthleteID
			})
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
func (r *Registry) Register(uc *UserConnection) {
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
