package biz

import (
	"context"
	"go.uber.org/zap"
)

type DomainRecord struct {
	DomainName string
	RR         string
	Value      string
}

type DomainRecordRepo interface {
	SetDomainRecord(ctx context.Context, dr *DomainRecord) (bool, error)
	GetDomainRecord(ctx context.Context, dr *DomainRecord) (string, error)
	DelDomainRecord(ctx context.Context, dr *DomainRecord) (int64, error)
	GetAllDomainRecord(ctx context.Context, dr *DomainRecord) (map[string]string, error)
	DelAllDomainRecord(ctx context.Context, dr *DomainRecord) error
}

type DomainRecordUsecase struct {
	repo   DomainRecordRepo
	logger *zap.SugaredLogger
}

func NewDomainRecordUsecase(repo DomainRecordRepo, logger *zap.Logger) *DomainRecordUsecase {
	return &DomainRecordUsecase{
		repo:   repo,
		logger: logger.Sugar(),
	}
}

func (uc *DomainRecordUsecase) SetDomainRecord(ctx context.Context, dr *DomainRecord) (bool, error) {
	return uc.repo.SetDomainRecord(ctx, dr)
}

func (uc *DomainRecordUsecase) GetDomainRecord(ctx context.Context, dr *DomainRecord) (string, error) {
	return uc.repo.GetDomainRecord(ctx, dr)
}

func (uc *DomainRecordUsecase) DelDomainRecord(ctx context.Context, dr *DomainRecord) (int64, error) {
	return uc.repo.DelDomainRecord(ctx, dr)
}

func (uc *DomainRecordUsecase) GetAllDomainRecord(ctx context.Context, dr *DomainRecord) (map[string]string, error) {
	return uc.repo.GetAllDomainRecord(ctx, dr)
}

func (uc *DomainRecordUsecase) DelAllDomainRecord(ctx context.Context, dr *DomainRecord) error {
	return uc.repo.DelAllDomainRecord(ctx, dr)
}
