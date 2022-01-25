package conf

type BasicConf struct {
	endpoint             string
	interfacePort        string
	domainGrpcNetwork    string
	domainGrpcPort       string
	interfaceGrpcNetwork string
	interfaceGrpcPort    string
}

func (b *BasicConf) EndPoint() string {
	return b.endpoint
}

func (b *BasicConf) InterfacePort() string {
	return b.interfacePort
}

func (b *BasicConf) DomainGrpcNetwork() string {
	return b.domainGrpcNetwork
}

func (b *BasicConf) DomainGrpcPort() string {
	return b.domainGrpcPort
}

func (b *BasicConf) InterfaceGrpcNetwork() string {
	return b.interfaceGrpcNetwork
}

func (b *BasicConf) InterfaceGrpcPort() string {
	return b.interfaceGrpcPort
}
