package adaptors

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	ConecStr string
	Connc    *sql.DB
}

//BUG: solve Postgres client connection string error
func NewPostgresClient(username, password, host string) (Postgres, error) {

	dsn := fmt.Sprintf("host=%s port=5432 dbname='postgres' user=%s password=%s sslmode=disable",
		host, username, password)

	connc, err := sql.Open("postgres", dsn)

	return Postgres{ConecStr: dsn, Connc: connc}, err
}

func (pg Postgres) Close() error {
	return pg.Connc.Close()
}
func (pg Postgres) Scan() (DbMetaInfo, error) {
	var dbMetaInfo DbMetaInfo

	schemas, err := pg.GetSchemaDetails()
	if err != nil {
		return DbMetaInfo{}, err
	}

	schemaData := []Schema{}
	for _, sc := range schemas {
		tables, err := pg.GetTableNamesFromDB(sc)
		if err != nil {
			log.Println(err.Error())
		}
		tbData := []Table{}
		for _, tb := range tables {
			tbcol, err := pg.GetTableDetails(sc, tb)
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
func (pg Postgres) GetSchemaDetails() ([]string, error) {

	var schemaNames []string
	rows, Qerr := pg.Connc.Query("SELECT datname FROM pg_database")
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

func (pg Postgres) GetTableNamesFromDB(dbname string) ([]string, error) {

	var tableNames []string
	rows, Qerr := pg.Connc.Query(`SELECT table_name FROM information_schema.tables 
								  WHERE table_schema='public' AND table_type = 'BASE TABLE' AND table_catalog=$1`, dbname)
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

func (pg Postgres) GetTableDetails(dbname string, tableName string) (Table, error) {
	var tableInfo Table

	tableInfo.Name = tableName

	q := fmt.Sprintf(`
		SELECT  column_name,
			    data_type as column_type,
			    col_description('public.%s'::regclass, ordinal_position)
				from information_schema.columns where table_name=$1
		`, tableName)

	rows, Qerr := pg.Connc.Query(q, tableName)
	if Qerr != nil {
		return tableInfo, Qerr
	}

	for rows.Next() {
		tempCol := Column{}
		tempComment := sql.NullString{}
		if err := rows.Scan(
			&tempCol.ColumnName, &tempCol.ColumnType,
			&tempComment); err != nil {
			Qerr = err
		} else {
			if tempComment.Valid {
				tempCol.ColumnComment = tempComment.String
			}
			tableInfo.Cols = append(tableInfo.Cols, tempCol)
		}
	}
	return tableInfo, Qerr
}
