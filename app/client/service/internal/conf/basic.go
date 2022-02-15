package conf

type BasicConf struct {
	endpoint   string // 阿里云服务地址
	domainName string // 监听变更的域名名称
	rr         string // 监听变更的域名的主机记录
	rpcUrl     string // rpc 服务地址
	rpcPort    string // rpc 服务端口
}

// EndPoint .
func (b *BasicConf) EndPoint() string {
	return b.endpoint
}

// DomainName .
func (b *BasicConf) DomainName() string {
	return b.domainName
}

// RR .
func (b *BasicConf) RR() string {
	return b.rr
}

// RpcUrl .
func (b *BasicConf) RpcUrl() string {
	return b.rpcUrl
}

// RpcPort .
func (b *BasicConf) RpcPort() string {
	return b.rpcPort
}
