package service

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/biz"
	"Ali-DDNS/internal/openapi"
	"Ali-DDNS/internal/openapi/defs/DescribeDomainRecords"
	"context"
	"encoding/json"
)

// GetDomainRecord get the domain record from data repo
func (s *DomainTaskService) GetDomainRecord(ctx context.Context, in *v1.DRRequest) (*v1.DRResponse, error) {
	var ret = &v1.DRResponse{DomainRecords: ""}

	allDomainRecord, err := s.domainRecordUsecase.GetAllDomainRecord(ctx, &biz.DomainRecord{
		DomainName: in.GetDomainName(),
	})
	if err != nil {
		s.logger.Errorf("get all domain record from redis failed, err: %v", err)
		return ret, err
	}

	var records []*DescribeDomainRecords.DRecord
	for _, v := range allDomainRecord {
		var record = &DescribeDomainRecords.DRecord{}
		err := json.Unmarshal([]byte(v), record)
		if err != nil {
			return ret, err
		}
		records = append(records, record)
	}
	bytes, err := json.Marshal(DescribeDomainRecords.DRecords{Records: records})
	if err != nil {
		return ret, err
	}

	return &v1.DRResponse{
		DomainRecords: string(bytes),
	}, nil
}

// UpdateDomainRecord update the domain record in data repo
func (s *DomainTaskService) UpdateDomainRecord(ctx context.Context, in *v1.UpdateDomainRequest) (*v1.UpdateDomainResponse, error) {
	s.logger.Infof("Update Doamin Record: [Domain Name: %s, RecordId: %s, RR: %s, Type: %s, Value: %s]", in.DomainName, in.RecordId, in.Rr, in.Type, in.Value)
	resp, err := openapi.UpdateDomainRecord(in.RecordId, in.Rr, in.Type, in.Value)
	if err != nil {
		return nil, err
	}

	err = s.AddTask(ctx, in.DomainName)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDomainResponse{
		RequestId: resp.RequestId,
		RecordId:  resp.RecordId,
	}, nil
}
