package openapi

import (
	"Ali-DDNS/internal/openapi/conf"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(conf.AK().ID()),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(conf.AK().Secret()),
	}

	// 访问的域名
	config.Endpoint = tea.String(conf.Basic().EndPoint())
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}
