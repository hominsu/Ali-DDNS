package conf

type RedisConf struct {
	addr     string
	port     string
	password string
	db       int
}

func (r *RedisConf) Addr() string {
	return r.addr + ":" + r.port
}

func (r *RedisConf) Password() string {
	return r.password
}

func (r *RedisConf) DB() int {
	return r.db
}
