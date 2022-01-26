package server

import (
	"Ali-DDNS/app/server/service/internal/conf"
	"Ali-DDNS/app/server/service/internal/service"
	"context"
	terrors "github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func NewCronServer(service *service.DomainTaskService, logger *zap.Logger) (*cron.Cron, error) {
	// 新建一个定时器
	cr := cron.New()

	// 新建一个定时器任务，定时触发 ip 变更检查
	if _, err := cr.AddFunc("@every "+conf.Option().TTL()+"s", func() {
		service.CheckAll(context.TODO())
	}); err != nil {
		return nil, terrors.Wrap(err, "init delay check job failed")
	}

	if _, err := cr.AddFunc(conf.Option().DelayCheckCron(), func() {
		service.CheckAll(context.TODO())
	}); err != nil {
		return nil, terrors.Wrap(err, "init cron check job failed")
	}

	return cr, nil
}
