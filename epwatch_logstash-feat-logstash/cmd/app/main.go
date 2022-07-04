package main

import (
	"flag"
	"fmt"

	"github.com/CMKL-PTTEP/epwatch_logstash/internal/app"
	"github.com/spf13/viper"
)

func main() {

	profile := flag.String("profile", "local_config", "[local_config | staging_config | production_config]")
	configPath := flag.String("config-path", "config/goapp", "path to config")
	flag.Parse()
	fmt.Println("Active profile:", *profile)
	fmt.Println("config path:", *configPath)
	// Read in config
	initConfig(*profile, *configPath)
	app.Start()
}

func initConfig(configName string, configPath string) {
	viper.AutomaticEnv()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
