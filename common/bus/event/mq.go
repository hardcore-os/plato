package event

type Manager struct {
	channelsConf map[Channel]*Options
	channels     map[Channel]mq
}

type mq interface {
	Send(event interface{})
	Receive() <-chan interface{}
}

func NewManager(opts map[Channel]*Options) *Manager {
	return &Manager{channelsConf: opts}
}

func (m *Manager) Send(c Channel, data interface{}) {

}

func (m *Manager) Receive(c Channel) interface{} {
	return nil
}
