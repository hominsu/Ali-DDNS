package openapi

import (
	"Ali-DDNS/internal/openapi/defs"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func UpdateDomainRecord(recordId, rr, _type, value string) (*defs.Resp, error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}

	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(recordId),
		RR:       tea.String(rr),
		Type:     tea.String(_type),
		Value:    tea.String(value),
	}

	result, _err := client.UpdateDomainRecord(updateDomainRecordRequest)
	if _err != nil {
		return nil, _err
	}

	return defs.ResponseBody{UpdateDomainRecordResponseBody: result.Body}.Format(), nil
}
