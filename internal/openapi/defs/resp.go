package defs

import (
	"github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Resp struct {
	RequestId string `json:"RequestId,omitempty"`
	RecordId  string `json:"RecordId,omitempty"`
}

type ResponseBody struct {
	*client.AddDomainRecordResponseBody
	*client.DeleteDomainRecordResponseBody
	*client.UpdateDomainRecordResponseBody
}

func (b ResponseBody) Format() *Resp {
	if b.AddDomainRecordResponseBody != nil {
		return &Resp{
			RequestId: tea.StringValue(b.AddDomainRecordResponseBody.RequestId),
			RecordId:  tea.StringValue(b.AddDomainRecordResponseBody.RecordId),
		}
	} else if b.DeleteDomainRecordResponseBody != nil {
		return &Resp{
			RequestId: tea.StringValue(b.DeleteDomainRecordResponseBody.RequestId),
			RecordId:  tea.StringValue(b.DeleteDomainRecordResponseBody.RecordId),
		}
	} else if b.UpdateDomainRecordResponseBody != nil {
		return &Resp{
			RequestId: tea.StringValue(b.UpdateDomainRecordResponseBody.RequestId),
			RecordId:  tea.StringValue(b.UpdateDomainRecordResponseBody.RecordId),
		}
	}
	return nil
}
