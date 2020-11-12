package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go_learning_homework/go-calendar-ms/api"
	"go_learning_homework/go-calendar-ms/internal/domain/services"
	"go_learning_homework/go-calendar-ms/logger"
	postgres "go_learning_homework/go-calendar-ms/tools/postgres"
	"log"
)

type config struct {
	HttpListen string `mapstructure:"http_listen"`
	LogFile    string `mapstructure:"log_file"`
	LogLevel   string `mapstructure:"log_level"`

	SqlUser string `mapstructure:"sql_user"`
	SqlPassword string `mapstructure:"sql_pass"`
	SqlHost string `mapstructure:"sql_host"`
	SqlPort string `mapstructure:"sql_port"`
	DbName string `mapstructure:"db_name"`
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

	db, err := postgres.InitDB(context.Background(), lg, cfg.SqlUser, cfg.SqlPassword, cfg.SqlHost, cfg.SqlPort, cfg.DbName)
	if err != nil {
		lg.Fatal(err.Error())
	}
	defer func() {
		if err := postgres.CloseDBCxn(); err != nil {
			lg.Error(err.Error())
		}
	}()

	eventService := services.EventService{EventStorage: db}

	//evt, _ := service.CreateEvent(context.Background(), "окулист", "окулист", 1, time.Now(), time.Now().Add(time.Hour * 2))
	//service.RemoveEventById(context.Background(), evt.Id)

	api.StartServer(cfg.HttpListen, *lg, &eventService)
}
