package conf

type SessionConf struct {
	size     int
	network  string
	address  string
	port     string
	password string
	secret   []byte
}

func (s SessionConf) Size() int {
	return s.size
}

func (s SessionConf) Network() string {
	return s.network
}

func (s SessionConf) Address() string {
	return s.address
}

func (s SessionConf) Port() string {
	return s.port
}

func (s SessionConf) Password() string {
	return s.password
}

func (s SessionConf) Secret() []byte {
	return s.secret
}
