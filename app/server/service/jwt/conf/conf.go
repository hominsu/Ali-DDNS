package conf

import (
	"log"
	"os"
	"reflect"
)

var (
	option = &OptionConf{}

	env = &envDef{
		optionJwtToken: "ALIDDNSCLIENT_OPTION_JWT_TOKEN",
	}
)

type envDef struct {
	optionJwtToken string
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
	case "optionJwtToken":
		Option().jwtToken = value
	default:
		log.Printf("unknown env name: %v \n", key)
		os.Exit(-1)
	}
}
