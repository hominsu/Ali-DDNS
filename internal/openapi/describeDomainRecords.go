package openapi

import (
	"Ali-DDNS/internal/openapi/defs/DescribeDomainRecords"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func DescribeDRecords(ds string) (*DescribeDomainRecords.Resp, error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}

	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(ds),
	}

	_result, _err := client.DescribeDomainRecords(describeDomainRecordsRequest)
	if _err != nil {
		return nil, _err
	}

	resp := DescribeDomainRecords.DescribeRespBody{DescribeDomainRecordsResponseBody: _result.Body}.Format()
	return resp, nil
}
