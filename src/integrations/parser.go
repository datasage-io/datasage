package integrations

import (
	"encoding/json"
	"time"
)

type Log struct {
	DataDomainID     string    `json:"DataDomainID"`
	Database         string    `json:"Database"`
	Operation        string    `json:"Operation"`
	OperationDetails string    `json:"OperationDetails"`
	Timestamp        time.Time `json:"Timestamp"`
	User             string    `json:"User"`
}

func ParseLog(logdata string) Log {
	log := Log{}
	json.Unmarshal([]byte(logdata), &log)
	return log
}
