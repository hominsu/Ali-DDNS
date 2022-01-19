package conf

type BasicConf struct {
	endpoint   string
	webPort    string
	rpcNetwork string
	rpcPort    string
}

func (b BasicConf) EndPoint() string {
	return b.endpoint
}

func (b BasicConf) WebPort() string {
	return b.webPort
}

func (b BasicConf) RpcNetwork() string {
	return b.rpcNetwork
}

func (b BasicConf) RpcPort() string {
	return b.rpcPort
}
