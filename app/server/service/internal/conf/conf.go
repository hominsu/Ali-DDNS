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

	session = &SessionConf{}

	option = &OptionConf{}

	env = &envDef{
		basicEndpoint:        "ALIDDNSSERVER_BASIC_ENDPOINT",
		basicWebPort:         "ALIDDNSSERVER_BASIC_WEB_PORT",
		basicRpcNetwork:      "ALIDDNSSERVER_BASIC_RPC_NETWORK",
		basicRpcPort:         "ALIDDNSSERVER_BASIC_RPC_PORT",
		redisAddr:            "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_ADDR",
		redisPort:            "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PORT",
		redisPassword:        "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_PASSWORD",
		redisDb:              "ALIDDNSSERVER_DOMAIN_RECORD_REDIS_DB",
		sessionSize:          "ALIDDNSSERVER_SESSION_SIZE",
		sessionNetwork:       "ALIDDNSSERVER_SESSION_REDIS_NETWORK",
		sessionAddress:       "ALIDDNSSERVER_SESSION_REDIS_ADDRESS",
		sessionPort:          "ALIDDNSSERVER_SESSION_REDIS_PORT",
		sessionPassword:      "ALIDDNSSERVER_SESSION_REDIS_PASSWORD",
		sessionSecret:        "ALIDDNSSERVER_SESSION_SECRET",
		optionTtl:            "ALIDDNSCLIENT_OPTION_TTL",
		optionDelayCheckCron: "ALIDDNSCLIENT_OPTION_DELAY_CHECK_CRON",
	}
)

type envDef struct {
	basicEndpoint        string
	basicWebPort         string
	basicRpcNetwork      string
	basicRpcPort         string
	redisAddr            string
	redisPort            string
	redisPassword        string
	redisDb              string
	sessionSize          string
	sessionNetwork       string
	sessionAddress       string
	sessionPort          string
	sessionPassword      string
	sessionSecret        string
	optionTtl            string
	optionDelayCheckCron string
}

func Basic() *BasicConf {
	return basic
}

func Redis() *RedisConf {
	return redis
}

func Session() *SessionConf {
	return session
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
	case "basicWebPort":
		Basic().webPort = value
	case "basicRpcNetwork":
		Basic().rpcNetwork = value
	case "basicRpcPort":
		Basic().rpcPort = value
	case "redisAddr":
		Redis().addr = value
	case "redisPort":
		Redis().port = value
	case "redisPassword":
		Redis().password = value
	case "redisDb":
		Redis().db = strToInt(value)
	case "sessionSize":
		Session().size = strToInt(value)
	case "sessionNetwork":
		Session().network = value
	case "sessionAddress":
		Session().address = value
	case "sessionPassword":
		Session().password = value
	case "sessionPort":
		Session().port = value
	case "sessionSecret":
		Session().secret = []byte(value)
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
