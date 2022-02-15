package service

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/biz"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProviderSet is domain task service providers.
var ProviderSet = wire.NewSet(NewDomainTaskService, NewDDNSInterfaceService)

// DomainTaskService .
type DomainTaskService struct {
	v1.UnimplementedDomainServiceServer

	delayCheckUsecase   *biz.DelayCheckUsecase
	domainRecordUsecase *biz.DomainRecordUsecase
	domainUserUsecase   *biz.DomainUserUsecase
	logger              *zap.SugaredLogger
}

// NewDomainTaskService new a domain task service
func NewDomainTaskService(delayCheckUsecase *biz.DelayCheckUsecase,
	domainRecordUsecase *biz.DomainRecordUsecase,
	domainUserUsecase *biz.DomainUserUsecase,
	logger *zap.Logger) *DomainTaskService {
	return &DomainTaskService{
		delayCheckUsecase:   delayCheckUsecase,
		domainRecordUsecase: domainRecordUsecase,
		domainUserUsecase:   domainUserUsecase,
		logger:              logger.Sugar(),
	}
}

type DDNSInterfaceService struct {
	v1.UnimplementedDDNSInterfaceServer

	domainUserUsecase *biz.DomainUserUsecase
	logger            *zap.SugaredLogger
}

// NewDDNSInterfaceService new ddns interface
func NewDDNSInterfaceService(domainUserUsecase *biz.DomainUserUsecase, logger *zap.Logger) *DDNSInterfaceService {
	return &DDNSInterfaceService{
		domainUserUsecase: domainUserUsecase,
		logger:            logger.Sugar(),
	}
}
