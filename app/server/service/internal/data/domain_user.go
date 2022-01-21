package data

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"context"
	terrors "github.com/pkg/errors"
)

type domainUserRepo struct {
	data *Data
}

func NewDomainUserRepo(data *Data) biz.DomainUserRepo {
	return &domainUserRepo{
		data: data,
	}
}

func (r *domainUserRepo) AddUser(ctx context.Context, du *biz.DomainUser) (bool, error) {
	if du.Username == "" || du.Password == "" {
		return false, terrors.New("username, password should not be empty")
	}
	return r.data.db.HSet("users", du.Username, du.Password).Result()
}

func (r *domainUserRepo) IsUserExists(ctx context.Context, du *biz.DomainUser) (bool, error) {
	if du.Username == "" {
		return false, terrors.New("username should not be empty")
	}
	return r.data.db.HExists("user", du.Username).Result()
}

func (r *domainUserRepo) GetUserPassword(ctx context.Context, du *biz.DomainUser) (string, error) {
	if du.Username == "" {
		return "", terrors.New("username should not be empty")
	}
	return r.data.db.HGet("users", du.Username).Result()
}

func (r *domainUserRepo) AddDevice(ctx context.Context, du *biz.DomainUser) (bool, error) {
	if du.Username == "" || du.UUID == "" {
		return false, terrors.New("username, uuid should not be empty")
	}
	return r.data.db.HSet("devices", du.Username, du.UUID).Result()
}

func (r *domainUserRepo) GetDevice(ctx context.Context, du *biz.DomainUser) ([]string, error) {
	var ret []string

	if du.Username == "" {
		return ret, terrors.New("username should not be empty")
	}

	result, err := r.data.db.HGetAll("devices").Result()
	if err != nil {
		return ret, err
	}

	for f, v := range result {
		if v == du.Username {
			ret = append(ret, f)
		}
	}

	return ret, nil
}

func (r *domainUserRepo) DelDevice(ctx context.Context, du *biz.DomainUser) (int64, error) {
	if du.UUID == "" {
		return 0, terrors.New("uuid should not be empty")
	}

	return r.data.db.HDel("devices", du.UUID).Result()
}

func (r *domainUserRepo) GetAllDevice(ctx context.Context, du *biz.DomainUser) ([]string, error) {
	var ret []string

	result, err := r.data.db.HGetAll("devices").Result()
	if err != nil {
		return ret, err
	}

	for f := range result {
		ret = append(ret, f)
	}

	return ret, nil
}

func (r *domainUserRepo) AddDomainName(ctx context.Context, du *biz.DomainUser) (bool, error) {
	if du.Username == "" || du.DomainName == "" {
		return false, terrors.New("username, domain name should not be empty")
	}

	return r.data.db.HSet("domains", du.DomainName, du.Username).Result()
}

func (r *domainUserRepo) GetDomainName(ctx context.Context, du *biz.DomainUser) ([]string, error) {
	var ret []string

	if du.Username == "" {
		return ret, terrors.New("username should not be empty")
	}

	result, err := r.data.db.HGetAll("domains").Result()
	if err != nil {
		return ret, err
	}

	for f, v := range result {
		if v == du.Username {
			ret = append(ret, f)
		}
	}

	return ret, nil
}

func (r *domainUserRepo) DelDomainName(ctx context.Context, du *biz.DomainUser) (int64, error) {
	if du.Username == "" {
		return 0, terrors.New("username should not be empty")
	}

	return r.data.db.HDel("domains", du.DomainName).Result()
}

func (r *domainUserRepo) GetAllDomainName(ctx context.Context, du *biz.DomainUser) ([]string, error) {
	var ret []string

	result, err := r.data.db.HGetAll("domains").Result()
	if err != nil {
		return ret, err
	}

	for f := range result {
		ret = append(ret, f)
	}

	return ret, nil
}
