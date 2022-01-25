package service

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is domain task service providers.
var ProviderSet = wire.NewSet(NewDomainTaskService, NewDDNSInterfaceService)

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

type DDNSInterfaceService struct {
	v1.UnimplementedDDNSInterfaceServer

	domainUserUsecase *biz.DomainUserUsecase
}

func NewDDNSInterfaceService(domainUserUsecase *biz.DomainUserUsecase) *DDNSInterfaceService {
	return &DDNSInterfaceService{
		domainUserUsecase: domainUserUsecase,
	}
}
