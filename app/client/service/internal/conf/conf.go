package conf

import (
	"log"
	"os"
	"reflect"
)

var (
	basic = &BasicConf{}

	option = &OptionConf{}

	env = &envDef{
		basicEndpoint:        "ALIDDNSCLIENT_BASIC_ENDPOINT",
		basicDomainName:      "ALIDDNSCLIENT_BASIC_DOMAIN_NAME",
		basicRR:              "ALIDDNSCLIENT_BASIC_RR",
		basicRpcUrl:          "ALIDDNSCLIENT_BASIC_RPC_URL",
		basicRpcPort:         "ALIDDNSCLIENT_BASIC_RPC_PORT",
		optionTtl:            "ALIDDNSCLIENT_OPTION_TTL",
		optionDelayCheckCron: "ALIDDNSCLIENT_OPTION_DELAY_CHECK_CRON",
		optionShowEachGetIp:  "ALIDDNSCLIENT_OPTION_SHOW_EACH_GET_IP",
	}
)

type envDef struct {
	basicEndpoint        string
	basicDomainName      string
	basicRR              string
	basicRpcUrl          string
	basicRpcPort         string
	optionTtl            string
	optionDelayCheckCron string
	optionShowEachGetIp  string
}

// Basic .
func Basic() *BasicConf {
	return basic
}

// Option .
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
		basic.endpoint = value
	case "basicDomainName":
		basic.domainName = value
	case "basicRR":
		basic.rr = value
	case "basicRpcUrl":
		basic.rpcUrl = value
	case "basicRpcPort":
		basic.rpcPort = value
	case "optionTtl":
		option.ttl = value
	case "optionShowEachGetIp":
		option.showEachGetIp = value
	case "optionDelayCheckCron":
		option.delayCheckCron = value
	default:
		log.Println("Unknown Env Name")
		os.Exit(-1)
	}
}
