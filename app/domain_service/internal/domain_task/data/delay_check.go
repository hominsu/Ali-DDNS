package data

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/biz"
	"context"
	terrors "github.com/pkg/errors"
)

type delayCheckRepo struct {
	data *Data
}

func NewDelayCheckRepo(data *Data) biz.DelayCheckRepo {
	return &delayCheckRepo{
		data: data,
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
