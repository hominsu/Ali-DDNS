package DescribeDomainRecords

import (
	"github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Resp struct {
	TotalCount    int64    `json:"TotalCount,omitempty"`
	PageSize      int64    `json:"PageSize,omitempty"`
	RequestId     string   `json:"RequestId,omitempty"`
	DomainRecords DRecords `json:"DomainRecords,omitempty"`
	PageNumber    int64    `json:"PageNumber,omitempty"`
}

type DescribeRespBody struct {
	*client.DescribeDomainRecordsResponseBody
}

func (b DescribeRespBody) Format() *Resp {
	return &Resp{
		TotalCount:    tea.Int64Value(b.TotalCount),
		PageSize:      tea.Int64Value(b.PageSize),
		RequestId:     tea.StringValue(b.RequestId),
		DomainRecords: DescribeDomainRecords{b.DomainRecords}.Format(),
		PageNumber:    tea.Int64Value(b.PageNumber),
	}
}
