package cache

type localCache struct{}

func newLocalCache(opt *Options) *localCache {
	return &localCache{}
}

func (r *localCache) MSet(keys map[string]interface{}) {

}

func (r *localCache) MGet(key []string) map[string]interface{} {
	return nil
}

func (r *localCache) MDel(key []string) {

}
