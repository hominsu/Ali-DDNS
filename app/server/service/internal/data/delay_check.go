package data

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"context"
	terrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

type delayCheckRepo struct {
	data   *Data
	logger *zap.SugaredLogger
}

func NewDelayCheckRepo(data *Data, logger *zap.Logger) biz.DelayCheckRepo {
	return &delayCheckRepo{
		data:   data,
		logger: logger.Sugar(),
	}
}

func (r *delayCheckRepo) SetDelayTask(ctx context.Context, dc *biz.DelayCheck) (int64, error) {
	if dc.DomainName == "" {
		return 0, terrors.New("domain name should not be empty")
	}
	return r.data.db.SAdd("delay_task", dc.DomainName).Result()
}

func (r *delayCheckRepo) DelDelayTask(ctx context.Context, dc *biz.DelayCheck) (int64, error) {
	if dc.DomainName == "" {
		return 0, terrors.New("domain name should not be empty")
	}
	return r.data.db.SRem("delay_task", dc.DomainName).Result()
}

func (r *delayCheckRepo) IsDelayTask(ctx context.Context, dc *biz.DelayCheck) (bool, error) {
	if dc.DomainName == "" {
		return false, terrors.New("domain name should not be empty")
	}
	return r.data.db.SIsMember("delay_task", dc.DomainName).Result()
}
