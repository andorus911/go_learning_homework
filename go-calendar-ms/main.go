package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go_learning_homework/go-calendar-ms/api"
	"go_learning_homework/go-calendar-ms/logger"
	"log"
)

type config struct {
	HttpListen string `mapstructure:"http_listen"`
	LogFile    string `mapstructure:"log_file"`
	LogLevel   string `mapstructure:"log_level"`
}

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "path to the config")
	flag.Parse()

	if configFilePath == "" {
		log.Fatal("Config is not presented")
	}

	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	cfg := &config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
	}

	lg := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	defer lg.Sync()

	api.StartServer(cfg.HttpListen, lg)
}
