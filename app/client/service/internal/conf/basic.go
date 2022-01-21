package conf

type BasicConf struct {
	endpoint   string // 阿里云服务地址
	domainName string // 监听变更的域名名称
	rr         string // 监听变更的域名的主机记录
	getIpUrl   string // 获取 ip 信息的 URL
	rpcUrl     string // rpc 服务地址
	rpcPort    string // rpc 服务端口
}

func (b *BasicConf) EndPoint() string {
	return b.endpoint
}

func (b *BasicConf) DomainName() string {
	return b.domainName
}

func (b *BasicConf) RR() string {
	return b.rr
}

func (b *BasicConf) GetIpUrl() string {
	return b.getIpUrl
}

func (b *BasicConf) RpcUrl() string {
	return b.rpcUrl
}

func (b *BasicConf) RpcPort() string {
	return b.rpcPort
}