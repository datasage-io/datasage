package integrations

import "log"

func (intg Integrations) StreamLogToAll(Log string) error {
	err := StreamLogToGRPC(Log, intg.Rpc)
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	err = StreamLogToKafka(Log, intg.Kafka)
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	return nil
}
