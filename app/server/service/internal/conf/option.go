package conf

type OptionConf struct {
	ttl            string
	delayCheckCron string
}

func (o *OptionConf) TTL() string {
	return o.ttl
}

func (o *OptionConf) DelayCheckCron() string {
	return o.delayCheckCron
}
