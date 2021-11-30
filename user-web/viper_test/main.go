package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
}

func main() {
	viper.SetConfigFile("user-web/viper_test/config.yaml")
	//viper.SetConfigName("user-web\\viper_test\\config")
	//viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	serverConfig := ServerConfig{}
	err = viper.Unmarshal(&serverConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	//fmt.Printf("name:%s\n", viper.Get("name"), viper.ConfigFileUsed())
	fmt.Println(serverConfig)
}
