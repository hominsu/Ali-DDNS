package data

import (
	v1 "Ali-DDNS/api/domain_record/v1"
	"Ali-DDNS/app/domain_check/internal/domain_check/conf"
	"Ali-DDNS/pkg"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	v1.DomainServiceClient
}

func NewData() (*Data, func(), error) {
	conn, err := grpc.Dial(conf.Basic().RpcUrl()+":"+conf.Basic().RpcPort(), grpc.WithTransportCredentials(pkg.GetClientCreds()))
	if err != nil {
		log.Panicln(err.Error())
	}

	cleanup := func() {
		log.Println("closing the grpc connect")
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}

	return &Data{
		v1.NewDomainServiceClient(conn),
	}, cleanup, err
}
