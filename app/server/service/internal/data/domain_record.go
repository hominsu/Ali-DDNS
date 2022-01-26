package data

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"context"
	terrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

type domainRecordRepo struct {
	data   *Data
	logger *zap.SugaredLogger
}

func NewDomainRecordRepo(data *Data, logger *zap.Logger) biz.DomainRecordRepo {
	return &domainRecordRepo{
		data:   data,
		logger: logger.Sugar(),
	}
}

func (r *domainRecordRepo) SetDomainRecord(ctx context.Context, dr *biz.DomainRecord) (bool, error) {
	if dr.DomainName == "" || dr.RR == "" || dr.Value == "" {
		return false, terrors.New("domain name, rr, value should not be empty")
	}
	return r.data.db.HSet(dr.DomainName, dr.RR, dr.Value).Result()
}

func (r *domainRecordRepo) GetDomainRecord(ctx context.Context, dr *biz.DomainRecord) (string, error) {
	if dr.DomainName == "" || dr.RR == "" {
		return "", terrors.New("domain name, rr should not be empty")
	}
	return r.data.db.HGet(dr.DomainName, dr.RR).Result()
}

func (r *domainRecordRepo) DelDomainRecord(ctx context.Context, dr *biz.DomainRecord) (int64, error) {
	if dr.DomainName == "" || dr.RR == "" {
		return 0, terrors.New("domain name, rr should not be empty")
	}
	return r.data.db.HDel(dr.DomainName, dr.RR).Result()
}

func (r *domainRecordRepo) GetAllDomainRecord(ctx context.Context, dr *biz.DomainRecord) (map[string]string, error) {
	if dr.DomainName == "" {
		return nil, terrors.New("domain name should not be empty")
	}
	return r.data.db.HGetAll(dr.DomainName).Result()
}

func (r *domainRecordRepo) DelAllDomainRecord(ctx context.Context, dr *biz.DomainRecord) error {
	if dr.DomainName == "" {
		return terrors.New("domain name should not be empty")
	}
	rrs, err := r.data.db.HVals(dr.DomainName).Result()
	if err != nil {
		return terrors.Wrap(err, "get all domain name failed")
	}
	if rrs == nil {
		return nil
	} else {
		for _, rr := range rrs {
			if _, err := r.DelDomainRecord(ctx, &biz.DomainRecord{
				DomainName: dr.DomainName,
				RR:         rr,
			}); err != nil {
				return terrors.Wrap(err, "del domain record failed")
			}
		}
	}
	return nil
}
