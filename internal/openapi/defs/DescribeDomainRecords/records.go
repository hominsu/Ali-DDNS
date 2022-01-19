package DescribeDomainRecords

import (
	"github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type DRecord struct {
	Status     string `json:"Status,omitempty"`
	Type       string `json:"Type,omitempty"`
	Remark     string `json:"Remark,omitempty"`
	TTL        int64  `json:"TTL,omitempty"`
	RecordId   string `json:"RecordId,omitempty"`
	Priority   int64  `json:"Priority,omitempty"`
	RR         string `json:"RR,omitempty"`
	DomainName string `json:"DomainName,omitempty"`
	Weight     int32  `json:"Weight,omitempty"`
	Value      string `json:"Value,omitempty"`
	Line       string `json:"Line,omitempty"`
	Locked     bool   `json:"Locked,omitempty"`
}

type DRecords struct {
	Records []*DRecord `json:"Record"`
}

type DescribeDomainRecord struct {
	*client.DescribeDomainRecordsResponseBodyDomainRecordsRecord
}

func (b *DescribeDomainRecord) Format() *DRecord {
	return &DRecord{
		Status:     tea.StringValue(b.Status),
		Type:       tea.StringValue(b.Type),
		Remark:     tea.StringValue(b.Remark),
		TTL:        tea.Int64Value(b.TTL),
		RecordId:   tea.StringValue(b.RecordId),
		Priority:   tea.Int64Value(b.Priority),
		RR:         tea.StringValue(b.RR),
		DomainName: tea.StringValue(b.DomainName),
		Weight:     tea.Int32Value(b.Weight),
		Value:      tea.StringValue(b.Value),
		Line:       tea.StringValue(b.Line),
		Locked:     tea.BoolValue(b.Locked),
	}
}

type DescribeDomainRecords struct {
	*client.DescribeDomainRecordsResponseBodyDomainRecords
}

func (b DescribeDomainRecords) Format() DRecords {
	var records []*DRecord
	for _, r := range b.Record {
		dr := DescribeDomainRecord{r}
		records = append(records, dr.Format())
	}

	return DRecords{Records: records}
}
