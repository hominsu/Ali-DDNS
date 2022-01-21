package conf

type OptionConf struct {
	ttl            string
	delayCheckCron string
	showEachGetIp  string
}

func (o *OptionConf) TTL() string {
	return o.ttl
}

func (o *OptionConf) DelayCheckCron() string {
	return o.delayCheckCron
}

func (o *OptionConf) ShowEachGetIp() string {
	return o.showEachGetIp
}
