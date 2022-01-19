package openapi

import (
	"Ali-DDNS/internal/openapi/defs"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func AddDomainRecord(domainName, rr, _type, value string) (*defs.Resp, error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}

	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(domainName),
		RR:         tea.String(rr),
		Type:       tea.String(_type),
		Value:      tea.String(value),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.AddDomainRecord(addDomainRecordRequest)
	if _err != nil {
		return nil, _err
	}

	return defs.ResponseBody{AddDomainRecordResponseBody: result.Body}.Format(), nil
}
