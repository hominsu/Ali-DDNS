package server

import (
	"Ali-DDNS/app/client/service/internal/conf"
	"Ali-DDNS/app/client/service/internal/service"
	"context"
	terrors "github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"log"
)

// NewCronServer new a cron server
func NewCronServer(service *service.DomainCheckService) (*cron.Cron, error) {
	// 新建一个定时器
	cr := cron.New()

	check := func() {
		if ip, err := service.Check(context.TODO()); err != nil {
			log.Println(err)
		} else if ip != "" {
			log.Println("Change the Domain Value to: ", ip)
		}
	}

	// 新建一个定时器任务，定时触发 ip 变更检查
	if _, err := cr.AddFunc("@every "+conf.Option().TTL()+"s", check); err != nil {
		return nil, terrors.Wrap(err, "init delay check job failed")
	}

	if _, err := cr.AddFunc(conf.Option().DelayCheckCron(), check); err != nil {
		return nil, terrors.Wrap(err, "init cron check job failed")
	}

	return cr, nil
}
