package main

import (
	"flag"
	"fmt"
	"sync"

	dataclassifier "github.com/datasage-io/datasage/src/classifiers"
	integrations "github.com/datasage-io/datasage/src/integrations"

	logger "github.com/datasage-io/datasage/src/logger"
	grpcServer "github.com/datasage-io/datasage/src/server"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup
var configFilePath *string
var log = logger.GetInstance()

func main() {
	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()

	log.Debug().Msgf("config path is  %v \n", *configFilePath)

	loadConfig()
	wg.Add(1)
	//start integration component dependent servers
	go integrations.RunServer()
	//Run a gRPC Server for CLI command processing
	go grpcServer.RunServer()
	go dataclassifier.Run()
	wg.Wait()

}
func loadConfig() {
	viper.SetConfigName("datasage")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			var log *zerolog.Logger = logger.GetInstance()
			log.Panic().Msgf("No config file found at %v\n", *configFilePath)
		} else {
			var log *zerolog.Logger = logger.GetInstance()
			log.Panic().Msgf("Error reading config file: %s\n", readErr)
		}
	}

	fmt.Println("logging level:", viper.GetString("logging.level"))

	logger.SetLogLevel(viper.GetString("logging.level"))

}
