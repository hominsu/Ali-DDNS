package data

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/client/service/internal/conf"
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
	creds, err := pkg.GetClientCreds()
	if err != nil {
		return nil, nil, err
	}

	conn, err := grpc.Dial(conf.Basic().RpcUrl()+":"+conf.Basic().RpcPort(), grpc.WithTransportCredentials(creds))
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
