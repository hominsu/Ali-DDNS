package service

import (
	v1 "Ali-DDNS/api/domain_record/v1"
	"Ali-DDNS/app/domain_service/internal/domain_task/biz"
	"github.com/google/wire"
)

// ProviderSet is domain task service providers.
var ProviderSet = wire.NewSet(NewDomainTaskService)

type DomainTaskService struct {
	v1.UnimplementedDomainServiceServer

	delayCheckUsecase   *biz.DelayCheckUsecase
	domainRecordUsecase *biz.DomainRecordUsecase
	domainUserUsecase   *biz.DomainUserUsecase
}

func NewDomainTaskService(delayCheckUsecase *biz.DelayCheckUsecase,
	domainRecordUsecase *biz.DomainRecordUsecase,
	domainUserUsecase *biz.DomainUserUsecase) *DomainTaskService {
	return &DomainTaskService{
		delayCheckUsecase:   delayCheckUsecase,
		domainRecordUsecase: domainRecordUsecase,
		domainUserUsecase:   domainUserUsecase,
	}
}
