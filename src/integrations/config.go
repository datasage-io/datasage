package integrations

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Integrations Integrations
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

func ReadLogConfig(path string) Config {
	config := Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("err: %s", err)
	}
	if yaml.Unmarshal(data, &config) != nil {
		log.Printf("err: %s", err)
	}
	return config
}
