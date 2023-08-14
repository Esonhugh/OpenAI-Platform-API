package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var C *GlobalConfig

type ConfigData struct {
	Openai struct {
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
	} `json:"openai" yaml:"openai"`
	AccessToken     string `json:"access_token" yaml:"access_token"`
	SecretKeyPrefix string `json:"secret_key_prefix" yaml:"secret_key_prefix"`
}

type GlobalConfig struct {
	Config ConfigData
	*viper.Viper
}

func Init() {
	C = &GlobalConfig{
		ConfigData{},
		viper.New(),
	}
	C.SetConfigName("openai-cli")
	C.SetConfigType("yaml")
	C.AddConfigPath("$HOME/.config")
	C.AutomaticEnv()
	err := C.ReadInConfig()
	if err != nil {
		log.Errorln("fail to read in config $HOME/.config/openai-cli.yaml", err)
		os.Exit(-1)
	}
	err = C.Unmarshal(&C.Config)
	if err != nil {
		log.Errorln("Bad Content in config $HOME/.config/openai-cli.yaml", err)
		os.Exit(-2)
	}
	C.Config.AccessToken = C.GetString("access_token")
	C.Config.SecretKeyPrefix = C.GetString("secret_key_prefix")
	log.Infoln("Got Config successfully in $HOME/.config/openai-cli.yaml")
}
