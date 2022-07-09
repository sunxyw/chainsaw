package redis

type RedisConf struct {
	Enable   bool
	Host     string
	Port     string
	Password string
	Database int
}
