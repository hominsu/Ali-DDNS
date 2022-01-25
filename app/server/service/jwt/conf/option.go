package conf

type OptionConf struct {
	jwtToken string
}

func (o *OptionConf) JwtToken() string {
	return o.jwtToken
}
