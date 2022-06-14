package adaptors

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//
type Mysql struct {
	ConecStr string
	Connc    *sql.DB
}

func NewMysqlClient(username, password, address, dbname string) (Mysql, error) {
	//Create data source name with given information
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, dbname)

	//open connection with database
	connc, err := sql.Open("mysql", dsn)

	return Mysql{ConecStr: dsn, Connc: connc}, err
}

//close conncection
func (my *Mysql) Close() {
	my.Connc.Close()
}
