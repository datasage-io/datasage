package storage

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

var initSchema = `

CREATE TABLE IF NOT EXISTS "class" (
	"id"	INTEGER,
	"description"	TEXT,
	"rule"	TEXT,
	"class"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "tag" (
	"id"	INTEGER,
	"tag_name"	TEXT,
	"rule"	TEXT,
	"description"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "dp_databases" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" TEXT,
	"type" TEXT
  );
  CREATE TABLE IF NOT EXISTS  "dp_db_tables" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" TEXT DEFAULT NULL,
	"dp_db_id" INTEGER DEFAULT NULL
  );
  CREATE TABLE IF NOT EXISTS  "dp_db_columns" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"dp_db_id" INTEGER DEFAULT NULL,
	"dp_db_table_id" INTEGER NOT NULL,
	"column_name" TEXT NOT NULL,
	"column_type" TEXT NOT NULL,
	"column_comment" TEXT NOT NULL,
	"Tags" TEXT NOT NULL,
	"Classes" TEXT NOT NULL
  ) ;
  CREATE TABLE IF NOT EXISTS  "DpDataSource" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"Datadomain" TEXT NOT NULL,
	"Dsname" TEXT NOT NULL,
	"Dsdecription" TEXT NOT NULL,
	"Dstype" TEXT NOT NULL,
	"DsKey" TEXT NOT NULL,
	"Dsversion" TEXT NOT NULL,
	"Host" TEXT NOT NULL,
	"Port" TEXT NOT NULL,
	"User" TEXT NOT NULL,
	"Password" TEXT NOT NULL
  ) ;
`

type Class struct {
	Id          int    `csv:"id"`
	Description string `csv:"description"`
	Rule        string `csv:"rule"`
	Class       string `csv:"class"`
}

type Tag struct {
	Id          int    `csv:"id"`
	TagName     string `csv:"tag_name"`
	Rule        string `csv:"rule"`
	Description string `csv:"description"`
}
type StorageConfig struct {
	Type string
	Path string
}
type Storage interface {
	GetClasses() ([]Class, error)
	GetTags() ([]Tag, error)
	GetAssociatedTags(string) ([]Tag, error)
	GetAssociatedClasses(string) ([]Class, error)
	SetSchemaData(DpDbDatabase) error
	SetDpDataSourceData(DpDataSource) error
	GetDpDataSources() ([]DpDataSource, error)
	DeleteDpDataSources(id int64) (bool, error)
}

/*
type DatabaseScanDto struct {
	ID                      int         `json:"id"`
	DbKey                   string      `json:"db_key"`
	Name                    string      `json:"name"`
	Type                    string      `json:"type"`
	ModifiedDeletedAtSource int         `json:"modified_deleted_at_source"`
	Deleted                 time.Time   `json:"deleted"`
	Key                     string      `json:"key"`
	LastScanID              int         `json:"last_scan_id"`
	DpDbTable               []DpDbTable `json:"dp_db_table"`
}

type DpDbTable struct {
	ID                      int          `json:"id"`
	Name                    string       `json:"name"`
	Tags                    string       `json:"tags"`
	ModifiedDeletedAtSource int          `json:"modified_deleted_at_source"`
	DeletedAt               time.Time    `json:"deleted_at"`
	DpDbID                  int          `json:"dp_db_id"`
	DpDbColumn              []DpDbColumn `json:"dp_db_column"`
}

type DpDbColumn struct {
	ID                      int       `json:"id"`
	ColumnName              string    `json:"column_name"`
	ColumnType              string    `json:"column_type"`
	ColumnComment           string    `json:"column_comment"`
	ModifiedDeletedAtSource int       `json:"modified_deleted_at_source"`
	DeletedAt               time.Time `json:"deleted_at"`
	DpDbTableID             int       `json:"dp_db_table_id"`
}
*/
type DpDataSource struct {
	ID           int    `json:"id"`
	Datadomain   string `json:"Datadomain"`
	Dsname       string `json:"Dsname"`
	Dsdecription string `json:"Dsdecription"`
	Dstype       string `json:"Dstype"`
	DsKey        string `json:"DsKey"`
	Dsversion    string `json:"Dsversion"`
	Host         string `json:"Host"`
	Port         string `json:"Port"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	CreatedAt    string `json:"CreatedAt"`
}

type DpDbDatabase struct {
	DbKey      string      `json:"DbKey"`
	Name       string      `json:"Name"`
	Type       string      `json:"Type"`
	Key        string      `json:"Key"`
	DpDbTables []DpDbTable `json:"DpDbTable"`
}

type DpDbTable struct {
	Name        string       `json:"Name"`
	Tags        string       `json:"Tags"`
	DpDbColumns []DpDbColumn `json:"DpDbColumns"`
}

type DpDbColumn struct {
	ColumnName    string `json:"column_name"`
	ColumnType    string `json:"ColumnType"`
	ColumnComment string `json:"Column_Comment"`
	Tags          string `json:"Tags"`
	Classes       string `json:"Classes"`
}

type SensitiveElementTag struct {
	SensitiveElementTags []SensitiveElementTags `json:"sensitive_data_element"`
}

type SensitiveElementTags struct {
	SensitivityClass string `json:"sensitivity_class"`
	Tags             string `json:"tags"`
}

type SensitiveDataElement struct {
	ID                      int                         `json:"id" gorm:"primary_key"`
	WorkspaceID             int                         `json:"workspace_id"`
	DpDbColumn              string                      `json:"dp_db_column"`
	ElementID               int                         `json:"element_id"`
	DeletedAt               time.Time                   `json:"deleted_at"`
	SensitivityClass        string                      `json:"sensitivity_class"`
	SensitiveDataElementTag []DpSensitiveDataElementTag `gorm:"foreignKey:SensitiveDataElementID;references:ID" json:"sensitive_data_element_tag"`
}

type DpSensitiveDataElementTag struct {
	ID                     int    `json:"id" gorm:"primary_key"`
	SensitiveDataElementID int    `json:"sensitive_data_element_id"`
	WorkspaceID            int    `json:"workspace_id"`
	Tag                    string `json:"tag"`
}

type ScanDto struct {
	DatabaseId int `json:"db_id"`
}

/*TODO: code to read .csv files form repository goes
here.
for now reading .csv files from local filesystem use
use go embed in final development
*/
func GetAllDefaultClassAndTags() ([]*Tag, []*Class, error) {
	tagr, err := os.Open("./storage/default/tags.csv")
	if err != nil {
		return []*Tag{}, []*Class{}, err
	}

	classr, err := os.Open("./storage/default/class.csv")
	if err != nil {
		return []*Tag{}, []*Class{}, err
	}

	tags := []*Tag{}
	classes := []*Class{}

	if err := gocsv.Unmarshal(tagr, &tags); err != nil {
		log.Println(err.Error())
		return []*Tag{}, []*Class{}, err

	}
	if err := gocsv.UnmarshalFile(classr, &classes); err != nil {
		log.Println(err.Error())
		return []*Tag{}, []*Class{}, err

	}
	return tags, classes, nil
}
func New(config StorageConfig) (Storage, error) {

	switch config.Type {
	case "internal":
		log.Println("New StorageConfig")
		//return NewInternalStorage(config.Path)
		return getInternalStorageInstance(config.Path)

	default:
		return nil, fmt.Errorf("%s storage not found", config.Type)
	}
}
