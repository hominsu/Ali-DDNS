package service

import (
	"Ali-DDNS/app/client/service/internal/data"
	"github.com/google/wire"
)

// ProviderSet is domain task service providers.
var ProviderSet = wire.NewSet(NewDomainCheckService)

type DomainCheckService struct {
	dataRepo *data.Data
}

func NewDomainCheckService(dataRepo *data.Data) *DomainCheckService {
	return &DomainCheckService{
		dataRepo: dataRepo,
	}
}
