package adaptors

import (
	"fmt"
)

type DbMetaInfo struct {
	Schemas []Schema
}
type Schema struct {
	Name   string
	Tables []Table
}
type Table struct {
	Name string
	Cols []Column
}
type Column struct {
	ColumnName    string
	ColumnType    string
	ColumnComment string
}

type Adaptor interface {
	Scan() (DbMetaInfo, error)
}
type AdaptorConfig struct {
	Type, Username, Password, Host string
}

func New(config AdaptorConfig) (Adaptor, error) {

	switch config.Type {

	case "mysql":
		return NewMysqlClient(config.Username, config.Password, config.Host)

	/*case "postgres":
	return NewPostgresClient(config.Username, config.Password, config.Host)*/

	default:
		return nil, fmt.Errorf("%s Database adoptor not found", config.Type)
	}

}
