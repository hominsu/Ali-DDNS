package conf

type OptionConf struct {
	ttl            string
	delayCheckCron string
	showEachGetIp  string
}

// TTL .
func (o *OptionConf) TTL() string {
	return o.ttl
}

// DelayCheckCron .
func (o *OptionConf) DelayCheckCron() string {
	return o.delayCheckCron
}

// ShowEachGetIp .
func (o *OptionConf) ShowEachGetIp() string {
	return o.showEachGetIp
}
