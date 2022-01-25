package conf

import (
	"log"
	"os"
	"reflect"
	"strconv"
)

var (
	basic = &BasicConf{}

	redis = &RedisConf{}

	option = &OptionConf{}

	env = &envDef{
		basicEndpoint:             "ALIDDNSSERVER_BASIC_ENDPOINT",
		basicInterfacePort:        "ALIDDNSSERVER_BASIC_INTERFACE_PORT",
		basicDomainGrpcNetwork:    "ALIDDNSSERVER_BASIC_DOMAIN_GRPC_NETWORK",
		basicDomainGrpcPort:       "ALIDDNSSERVER_BASIC_DOMAIN_GRPC_PORT",
		basicInterfaceGrpcNetwork: "ALIDDNSSERVER_BASIC_INTERFACE_GRPC_NETWORK",
		basicInterfaceGrpcPort:    "ALIDDNSSERVER_BASIC_INTERFACE_GRPC_PORT",
		redisAddr:                 "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_ADDR",
		redisPort:                 "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PORT",
		redisPassword:             "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PASSWORD",
		redisDb:                   "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_DB",
		optionTtl:                 "ALIDDNSCLIENT_OPTION_TTL",
		optionDelayCheckCron:      "ALIDDNSCLIENT_OPTION_DELAY_CHECK_CRON",
	}
)

type envDef struct {
	basicEndpoint             string
	basicInterfacePort        string
	basicDomainGrpcNetwork    string
	basicDomainGrpcPort       string
	basicInterfaceGrpcNetwork string
	basicInterfaceGrpcPort    string
	redisAddr                 string
	redisPort                 string
	redisPassword             string
	redisDb                   string
	optionTtl                 string
	optionDelayCheckCron      string
}

func Basic() *BasicConf {
	return basic
}

func Redis() *RedisConf {
	return redis
}

func Option() *OptionConf {
	return option
}

func init() {
	typeOfEnv := reflect.TypeOf(*env)
	valueOfEnv := reflect.ValueOf(*env)

	for i := 0; i < typeOfEnv.NumField(); i++ {
		if value := os.Getenv(valueOfEnv.Field(i).String()); value != "" {
			setEnv(typeOfEnv.Field(i).Name, value)
		} else {
			log.Printf("env %v should not be empty\n", valueOfEnv.Field(i).String())
			os.Exit(-1)
		}
	}
}

func setEnv(key, value string) {
	switch key {
	case "basicEndpoint":
		Basic().endpoint = value
	case "basicInterfacePort":
		Basic().interfacePort = value
	case "basicDomainGrpcNetwork":
		Basic().domainGrpcNetwork = value
	case "basicDomainGrpcPort":
		Basic().domainGrpcPort = value
	case "basicInterfaceGrpcNetwork":
		Basic().interfaceGrpcNetwork = value
	case "basicInterfaceGrpcPort":
		Basic().interfaceGrpcPort = value
	case "redisAddr":
		Redis().addr = value
	case "redisPort":
		Redis().port = value
	case "redisPassword":
		Redis().password = value
	case "redisDb":
		Redis().db = strToInt(value)
	case "optionTtl":
		Option().ttl = value
	case "optionDelayCheckCron":
		Option().delayCheckCron = value
	default:
		log.Printf("unknown env name: %v \n", key)
		os.Exit(-1)
	}
}

func strToInt(value string) int {
	ret, err := strconv.Atoi(value)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return ret
}
