package cache

type mode = string

const (
	Local  mode = "local"
	Remote mode = "remote"
)

type Options struct {
	Mode mode
}

type cache interface {
	MSet(keys map[string]interface{})
	MGet(key []string) map[string]interface{}
	MDel(key []string)
}

type Manager struct {
	opts      []*Options
	cacheList []cache
}

func NewManager(opts []*Options) *Manager {
	m := &Manager{opts: opts, cacheList: make([]cache, 0)}
	for _, opt := range opts {
		switch opt.Mode {
		case Local:
			m.cacheList = append(m.cacheList, newLocalCache(opt))
		case Remote:
			m.cacheList = append(m.cacheList, newRedisCache(opt))
		}
	}
	return m
}

func (m *Manager) MSet(keys map[string]interface{}) {

}

func (m *Manager) MGet(key []string) map[string]interface{} {
	return nil
}

func (m *Manager) MDel(key []string) {
}
