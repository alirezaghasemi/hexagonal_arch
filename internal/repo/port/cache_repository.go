package port

type CacheRepository interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
