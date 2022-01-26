package biz

import (
	"context"
	"go.uber.org/zap"
)

type DelayCheck struct {
	DomainName string
}

type DelayCheckRepo interface {
	SetDelayTask(ctx context.Context, dc *DelayCheck) (int64, error)
	DelDelayTask(ctx context.Context, dc *DelayCheck) (int64, error)
	IsDelayTask(ctx context.Context, dc *DelayCheck) (bool, error)
}

type DelayCheckUsecase struct {
	repo   DelayCheckRepo
	logger *zap.SugaredLogger
}

func NewDelayCheckUsecase(repo DelayCheckRepo, logger *zap.Logger) *DelayCheckUsecase {
	return &DelayCheckUsecase{
		repo:   repo,
		logger: logger.Sugar(),
	}
}

func (uc *DelayCheckUsecase) SetDelayTask(ctx context.Context, dc *DelayCheck) (int64, error) {
	return uc.repo.SetDelayTask(ctx, dc)
}

func (uc *DelayCheckUsecase) DelDelayTask(ctx context.Context, dc *DelayCheck) (int64, error) {
	return uc.repo.DelDelayTask(ctx, dc)
}

func (uc *DelayCheckUsecase) IsDelayTask(ctx context.Context, dc *DelayCheck) (bool, error) {
	return uc.repo.IsDelayTask(ctx, dc)
}
