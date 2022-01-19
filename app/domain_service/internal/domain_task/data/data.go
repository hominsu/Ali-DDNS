package data

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/conf"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	terrors "github.com/pkg/errors"
	"log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedisClient, NewDelayCheckRepo, NewDomainRecordRepo, NewDomainUserRepo)

// Data .
type Data struct {
	db *redis.Client
}

// NewRedisClient .
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis().Addr(),
		Password: conf.Redis().Password(),
		DB:       conf.Redis().DB(),
	})
}

// NewData .
func NewData(redisClient *redis.Client) (*Data, func(), error) {
	d := &Data{
		db: redisClient,
	}

	if _, err := d.db.Ping().Result(); err != nil {
		return nil, nil, terrors.Wrap(err, "New redis connect failed")
	}

	cleanup := func() {
		log.Println("closing the redis connect")
		if err := d.db.Close(); err != nil {
			log.Println(err)
		}
	}

	return d, cleanup, nil
}
