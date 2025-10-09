package storage

const (
	userDomainCacheKey = "userDomainCacheKey:%d"
)

// 保持聚合的一致性
type StorageManager struct {
}

func NewStorageManager() *StorageManager {
	sm := &StorageManager{}
	return sm
}
