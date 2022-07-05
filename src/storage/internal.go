package storage

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type InternalStorage struct {
	SqliteConnc *sql.DB
	Path        string
}

func (into InternalStorage) InsertDefaultData(tags []*Tag, classes []*Class) error {
	_, err := into.SqliteConnc.Exec(initSchema)
	if err != nil {
		return err
	}

	tagInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO tag ("id","tag_name","rule","description")
	 VALUES (?,?,?,?);
	`)

	for _, t := range tags {
		_, err = tagInsert.Exec(t.Id, t.TagName, t.Rule, t.Description)
	}
	tagInsert.Close()

	classInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO class ("id","description","rule","class")
	 VALUES (?,?,?,?);
	`)
	for _, c := range classes {
		_, err = classInsert.Exec(c.Id, c.Description, c.Rule, c.Class)
	}
	classInsert.Close()
	return err
}

func (into InternalStorage) SetSchemaData(dpDbDatabase DpDbDatabase) error {

	dbInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO dp_databases ("name","type")
	VALUES (?,?);
	`)
	if err != nil {
		log.Println("error1", err.Error())
		return err
	}
	dpDb, err := dbInsert.Exec(dpDbDatabase.Name, dpDbDatabase.Type)
	if err != nil {
		log.Println("error2", err.Error())
		return err
	}
	dpDbId, err := dpDb.LastInsertId()

	dbInsert.Close()

	for _, table := range dpDbDatabase.DpDbTables {

		tableInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_tables ("name","dp_db_id")
		 VALUES (?,?);
		`)
		if err != nil {
			log.Println("error3", err.Error())
			continue
		}

		//fmt.Println("INSERT INTO dp_db_tables :", table.Name, dpDbId)
		dbDbTable, err := tableInsert.Exec(table.Name, dpDbId)
		if err != nil {
			log.Println("error3b", err.Error())
			continue
		}

		dbDbTableId, err := dbDbTable.LastInsertId()
		tableInsert.Close()

		columnInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_columns ("dp_db_id","dp_db_table_id","column_name","column_type","column_comment","Tags","Classes")
		 VALUES (?,?,?,?,?,?,?);
		`)
		if err != nil {
			log.Println("error4", err.Error())
			continue
		}

		for _, column := range table.DpDbColumns {

			//fmt.Println("INSERT INTO dp_db_columns:", dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment)
			_, err = columnInsert.Exec(dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment, column.Tags, column.Classes)
			if err != nil {
				log.Println("error5", err.Error())
				//	continue
			}

		}
		columnInsert.Close()

	}
	return err
}

func NewInternalStorage(dsn string) (InternalStorage, error) {
	log.Println("NewInternalStorage enter")
	var isnew bool
	_, err := os.Stat(dsn)
	if os.IsNotExist(err) {
		log.Println("Creating sqlite database ", dsn)
		_, err := os.OpenFile(dsn, os.O_CREATE|os.O_WRONLY, 0660)
		isnew = true
		if err != nil {
			return InternalStorage{}, err
		}
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return InternalStorage{}, err
	}
	insto := InternalStorage{Path: dsn, SqliteConnc: db}
	if isnew {
		log.Println("Inserting Default data")
		tags, classes, err := GetAllDefaultClassAndTags()
		if err != nil {
			return InternalStorage{}, err
		}
		err = insto.InsertDefaultData(tags, classes)

	}
	log.Println("NewInternalStorage exit")
	return insto, err

}

func (insto InternalStorage) GetTags() ([]Tag, error) {
	tags := []Tag{}

	rows, err := insto.SqliteConnc.Query("SELECT id,tag_name,rule,description FROM tag")
	if err != nil {
		return tags, err
	}
	for rows.Next() {
		tempTag := Tag{}
		err = rows.Scan(&tempTag.Id, &tempTag.TagName,
			&tempTag.Rule, &tempTag.Description)

		tags = append(tags, tempTag)
	}
	return tags, err
}

func (insto InternalStorage) GetClasses() ([]Class, error) {
	classes := []Class{}

	rows, err := insto.SqliteConnc.Query("SELECT id,description,rule,class FROM class")
	if err != nil {
		return classes, err
	}
	for rows.Next() {
		tempClass := Class{}
		err = rows.Scan(&tempClass.Id, &tempClass.Description,
			&tempClass.Rule, &tempClass.Class)

		classes = append(classes, tempClass)
	}

	return classes, err
}

func (insto InternalStorage) GetAssociatedTags(class string) ([]Tag, error) {

	tags := []Tag{}
	rows, err := insto.SqliteConnc.Query("SELECT id,tag_name,rule,description FROM tag Where rule = ?", class)
	if err != nil {
		return tags, err
	}
	for rows.Next() {
		tempTag := Tag{}
		err = rows.Scan(&tempTag.Id, &tempTag.TagName,
			&tempTag.Rule, &tempTag.Description)
		tags = append(tags, tempTag)
	}

	return tags, err
}

func (insto InternalStorage) GetAssociatedClasses(rule string) ([]Class, error) {

	classes := []Class{}
	rows, err := insto.SqliteConnc.Query("SELECT id,description,rule,class FROM class Where rule = ?", rule)
	if err != nil {
		return classes, err
	}
	for rows.Next() {
		tempClass := Class{}
		err = rows.Scan(&tempClass.Id, &tempClass.Description,
			&tempClass.Rule, &tempClass.Class)
		classes = append(classes, tempClass)
	}

	return classes, err
}
