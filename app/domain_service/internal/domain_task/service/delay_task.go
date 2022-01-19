package service

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/biz"
	"Ali-DDNS/internal/openapi"
	"context"
	"encoding/json"
	"log"
	"time"
)

// AddTask add a domain update task into redis
func (s *DomainTaskService) AddTask(ctx context.Context, domainName string) error {
	if isDelay, err := s.delayCheckUsecase.IsDelayTask(ctx, &biz.DelayCheck{
		DomainName: domainName},
	); err != nil {
		return err
	} else {
		// if the domain name to be updated is already waiting, the timer is not set
		if isDelay {
			return nil
		} else {
			// add the domain name into Redis which will be updated
			if _, err := s.delayCheckUsecase.SetDelayTask(ctx, &biz.DelayCheck{
				DomainName: domainName,
			}); err != nil {
				return err
			}
			// the domain name is updated after 10 seconds
			time.AfterFunc(time.Second*10, func() {
				// delete the delay check domain name from Redis, if error is nil then check
				_, err := s.delayCheckUsecase.DelDelayTask(ctx, &biz.DelayCheck{
					DomainName: domainName,
				})
				if err != nil {
					log.Println(err.Error())
					return
				}
				s.Check(ctx, domainName)
			})
			return nil
		}
	}
}

// Check will get the domain record from openapi the check whether the domain record is expired
func (s *DomainTaskService) Check(ctx context.Context, domainName string) bool {
	// delete all key/value pairs from the domain name hashset
	if err := s.domainRecordUsecase.DelAllDomainRecord(ctx, &biz.DomainRecord{
		DomainName: domainName,
	}); err != nil {
		return true
	}

	// obtain the corresponding domain name records through ali openapi
	records, err := openapi.DescribeDRecords(domainName)
	if err != nil {
		log.Println(err.Error())
		return true
	}

	// iterate over the obtained record, marshal it and add it to the domain hashset as a key-value pair of: host-value/return-value
	for _, domainRecord := range records.DomainRecords.Records {
		bytes, err := json.Marshal(domainRecord)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if _, err = s.domainRecordUsecase.SetDomainRecord(ctx, &biz.DomainRecord{
			DomainName: domainRecord.DomainName,
			RR:         domainRecord.RR,
			Value:      string(bytes),
		}); err != nil {
			log.Println(err.Error())
			continue
		}
	}
	return false
}

func (s *DomainTaskService) CheckAll(ctx context.Context) {
	// get all domain names from Redis
	domainNames, err := s.domainUserUsecase.GetAllDomainName(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if domainNames == nil {
		return
	} else {
		// iterate over each domain name
		for _, domainName := range domainNames {
			if s.Check(ctx, domainName) {
				continue
			}
		}
	}
}
