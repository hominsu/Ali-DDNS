package conf

type AccessKeyConf struct {
	id     string
	secret string
}

func (ak *AccessKeyConf) ID() string {
	return ak.id
}

func (ak *AccessKeyConf) Secret() string {
	return ak.secret
}
