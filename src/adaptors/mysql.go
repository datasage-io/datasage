package adaptors

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//
type Mysql struct {
	ConecStr string
	Connc    *sql.DB
}

func NewMysqlClient(username, password, host string) (Mysql, error) {
	//Create data source name with given information
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, host)

	//open connection with database
	connc, err := sql.Open("mysql", dsn)

	return Mysql{ConecStr: dsn, Connc: connc}, err
}

//close conncection
func (my *Mysql) Close() {
	my.Connc.Close()
}

func (my Mysql) Scan() (DbMetaInfo, error) {
	var dbMetaInfo DbMetaInfo

	schemas, err := my.GetSchemaDetails()
	if err != nil {
		return DbMetaInfo{}, err
	}

	schemaData := []Schema{}
	for _, sc := range schemas {
		tables, err := my.GetTableNamesFromDB(sc)
		if err != nil {
			log.Println(err.Error())
		}
		tbData := []Table{}
		for _, tb := range tables {
			tbcol, err := my.GetTableDetails(sc, tb)
			if err != nil {
				log.Println(err.Error())
			}
			tbData = append(tbData, tbcol)
		}
		schemaData = append(schemaData, Schema{Name: sc, Tables: tbData})
	}

	dbMetaInfo.Schemas = schemaData
	return dbMetaInfo, nil
}

func (mysql Mysql) GetSchemaDetails() ([]string, error) {

	var schemaNames []string
	rows, Qerr := mysql.Connc.Query("select SCHEMA_NAME from information_schema.schemata order by schema_name")
	if Qerr != nil {
		return schemaNames, Qerr
	}
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			Qerr = err
		} else {
			schemaNames = append(schemaNames, tmp)
		}
	}
	return schemaNames, Qerr
}

func (mysql Mysql) GetTableNamesFromDB(dbname string) ([]string, error) {

	var tableNames []string
	rows, Qerr := mysql.Connc.Query(`SELECT TABLE_NAME from information_schema.tables 
									 WHERE table_type = 'BASE TABLE' and table_schema = ?`, dbname)
	if Qerr != nil {
		return tableNames, Qerr
	}
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			Qerr = err
		} else {
			tableNames = append(tableNames, tmp)
		}
	}
	return tableNames, Qerr
}

func (mysql Mysql) GetTableDetails(dbname string, tableName string) (Table, error) {
	var tableInfo Table

	tableInfo.Name = tableName

	rows, Qerr := mysql.Connc.Query(`SELECT COLUMN_NAME as column_name, 
											COLUMN_TYPE as column_type, 
											COLUMN_COMMENT as column_comment 
											from information_schema.columns 
											where table_schema = ? and table_name = ?`, dbname, tableName)
	if Qerr != nil {
		return tableInfo, Qerr
	}

	for rows.Next() {
		tempCol := Column{}
		if err := rows.Scan(
			&tempCol.ColumnName, &tempCol.ColumnType,
			&tempCol.ColumnComment); err != nil {
			Qerr = err
		} else {
			tableInfo.Cols = append(tableInfo.Cols, tempCol)
		}
	}
	return tableInfo, Qerr
}
