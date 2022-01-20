package service

import (
	v1 "Ali-DDNS/api/domain_record/v1"
	"Ali-DDNS/app/domain_check/internal/domain_check/conf"
	"Ali-DDNS/internal/openapi/defs"
	"Ali-DDNS/internal/openapi/defs/DescribeDomainRecords"
	"context"
	"encoding/json"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

// GetPeerAddr get peer ip addr
func (s *DomainCheckService) GetPeerAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}

func (s *DomainCheckService) GetDomainRecord(ctx context.Context, domainName string) (*DescribeDomainRecords.DRecords, error) {
	domainRequest := &v1.DRRequest{DomainName: domainName}
	resp, err := s.dataRepo.GetDomainRecord(ctx, domainRequest)
	if err != nil {
		return nil, err
	}

	var records = &DescribeDomainRecords.DRecords{}
	err = json.Unmarshal([]byte(resp.GetDomainRecords()), records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *DomainCheckService) UpdateDomainRecord(ctx context.Context, domainName, recordId, rr, _type, value string) (*defs.Resp, error) {
	updateRequest := &v1.UpdateDomainRequest{
		DomainName: domainName,
		RecordId:   recordId,
		Rr:         rr,
		Type:       _type,
		Value:      value,
	}
	resp, err := s.dataRepo.UpdateDomainRecord(ctx, updateRequest)
	if err != nil {
		return nil, err
	}

	return &defs.Resp{
		RequestId: resp.RequestId,
		RecordId:  resp.RecordId,
	}, nil
}

func (s *DomainCheckService) Check(ctx context.Context) (value string, _error error) {
	// 查询 "haomingsu.cn" 的 DNS 解析信息

	//_records, err := openapi2.DescribeDRecords(config.Basic.DomainName())
	_records, err := s.GetDomainRecord(ctx, conf.Basic().DomainName())
	if err != nil {
		return "", err
	}
	//for _, domainRecord := range _records.DomainRecords.Records {
	for _, domainRecord := range _records.Records {
		// 如果记录的域名为 "haomingsu.cn" 并且主机记录为 "home"
		if domainRecord.DomainName == conf.Basic().DomainName() && domainRecord.RR == conf.Basic().RR() {
			ip := s.GetPeerAddr(context.TODO())
			if conf.Option().ShowEachGetIp() == "true" {
				log.Printf("Current Address: %s, Doamin Value: %s", ip, domainRecord.Value)
			}
			if ip != domainRecord.Value {
				//_, err := openapi2.UpdateDomainRecord(domainRecord.RecordId, domainRecord.RR, domainRecord.Type, ip)
				_, err := s.UpdateDomainRecord(ctx, domainRecord.DomainName, domainRecord.RecordId, domainRecord.RR, domainRecord.Type, ip)
				if err != nil {
					return "", err
				}
				return ip, nil
			}
		}
	}

	return "", nil
}
