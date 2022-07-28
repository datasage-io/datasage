package integrations

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Integrations
}

type Integrations struct {
	Rpc   []GRPCLogConfig
	Kafka []KafkaLogConfig
}

type GRPCLogConfig struct {
	Host string
	Port string
}

type KafkaLogConfig struct {
	Broker string
	Topic  string
	Port   string
}

func ReadLogConfig() (Config, error) {
	log.Printf("ReadLogConfig")
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("err: %s", err)
		log.Fatal(err)

	}
	return config, nil
}
