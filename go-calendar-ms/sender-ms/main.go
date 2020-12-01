package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go_learning_homework/go-calendar-ms/api-ms/logger"
	"log"
)

type config struct {
	LogFile    string `mapstructure:"log_file"`
	LogLevel   string `mapstructure:"log_level"`

	RabbitUser string `mapstructure:"rabbit_user"`
	RabbitPassword string `mapstructure:"rabbit_pass"`
	RabbitHost string `mapstructure:"rabbit_host"`
	RabbitPort string `mapstructure:"rabbit_port"`
}

func main() {
	// config
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

	// init of a logger
	lg := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	defer lg.Sync()

	// running of a sender
	RunReminder(context.Background(), lg, cfg.RabbitUser, cfg.RabbitPassword, cfg.RabbitHost, cfg.RabbitPort)
}
