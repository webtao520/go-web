package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var ViperConfig Configuration

func init (){
	runtimeViper := viper.New()
    runtimeViper.AddConfigPath(".")
    runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	err := runtimeViper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	runtimeViper.Unmarshal(&ViperConfig)
	

}
