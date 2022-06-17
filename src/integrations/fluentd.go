package integrations

import (
	"log"

	"github.com/fluent/fluent-logger-golang/fluent"
)

func SendLogToFluentd(Log string, config fluent.Config) {
	logger, err := fluent.New(config)
	if err != nil {
		log.Println(err)
	}
	defer logger.Close()
	tag := "datasage.access"
	err = logger.Post(tag, Log)
	if err != nil {
		log.Println("err: ", err)
	}
}
