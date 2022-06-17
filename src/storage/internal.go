package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type InternalStorage struct {
	SqliteConnc *sql.DB
	Path        string
}

func NewInternalStorage(dsn string) (InternalStorage, error) {

	db, err := sql.Open("sqlite", dsn)
	if err := db.Ping(); err != nil {
		log.Println(err.Error())
	}
	return InternalStorage{SqliteConnc: db, Path: dsn}, err

}

func (insto InternalStorage) GetAllTags() ([]Tag, error) {
	tags := []Tag{}

	rows, err := insto.SqliteConnc.Query("SELECT id,tag_name,rule,description FROM dp_tag_rule_system")
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

func (insto InternalStorage) GetAllClasses() ([]Class, error) {
	classes := []Class{}

	rows, err := insto.SqliteConnc.Query("SELECT id,description,rule,class FROM dp_sensitivity_class_rule_system;")
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
