package conf

import (
	"log"
	"os"
	"reflect"
)

var (
	ak = &AccessKeyConf{}

	basic = &BasicConf{}

	env = &envDef{
		akId:            "ALIDDNSSERVER_ACCESSKEY_ID",
		akSecret:        "ALIDDNSSERVER_ACCESSKEY_SECRET",
		basicEndpoint:   "ALIDDNSSERVER_BASIC_ENDPOINT",
		basicWebPort:    "ALIDDNSSERVER_BASIC_WEB_PORT",
		basicRpcNetwork: "ALIDDNSSERVER_BASIC_RPC_NETWORK",
		basicRpcPort:    "ALIDDNSSERVER_BASIC_RPC_PORT",
	}
)

type envDef struct {
	akId            string
	akSecret        string
	basicEndpoint   string
	basicWebPort    string
	basicRpcNetwork string
	basicRpcPort    string
}

func AK() *AccessKeyConf {
	return ak
}

func Basic() *BasicConf {
	return basic
}

func init() {
	typeOfEnv := reflect.TypeOf(*env)
	valueOfEnv := reflect.ValueOf(*env)

	for i := 0; i < typeOfEnv.NumField(); i++ {
		if value := os.Getenv(valueOfEnv.Field(i).String()); value != "" {
			setEnv(typeOfEnv.Field(i).Name, value)
		} else {
			envError("Env should not be empty")
		}
	}
}

func setEnv(key, value string) {
	switch key {
	case "akId":
		AK().id = value
	case "akSecret":
		AK().secret = value
	case "basicEndpoint":
		Basic().endpoint = value
	case "basicWebPort":
		Basic().webPort = value
	case "basicRpcNetwork":
		Basic().rpcNetwork = value
	case "basicRpcPort":
		Basic().rpcPort = value
	default:
		envError("Unknown Env Name")
	}
}

func envError(err string) {
	log.Println(err)
	os.Exit(-1)
}
