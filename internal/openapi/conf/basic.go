package conf

type BasicConf struct {
	endpoint string
}

func (b BasicConf) EndPoint() string {
	return b.endpoint
}
