package integrations

func (intg Integrations) StreamLogToAll(Log string) {
	StreamLogToGRPC(Log, intg.Rpc)
	StreamLogToKafka(Log, intg.Kafka)
}
