package openapi

import (
	"Ali-DDNS/internal/openapi/defs"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func DeleteDomainRecord(recordId string) (*defs.Resp, error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}

	deleteDomainRecordRequest := &alidns20150109.DeleteDomainRecordRequest{
		RecordId: tea.String(recordId),
	}

	result, _err := client.DeleteDomainRecord(deleteDomainRecordRequest)
	if _err != nil {
		return nil, _err
	}

	return defs.ResponseBody{DeleteDomainRecordResponseBody: result.Body}.Format(), nil
}
