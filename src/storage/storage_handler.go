package storage

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

var initSchema = `

CREATE TABLE IF NOT EXISTS "class" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"description"	TEXT,
	"rule"	TEXT,
	"class"	TEXT
);
CREATE TABLE IF NOT EXISTS "tag" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"tag_name"	TEXT,
	"rule"	TEXT,
	"description"	TEXT
);
CREATE TABLE IF NOT EXISTS "dp_databases" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"dskey" TEXT NOT NULL
  );
  CREATE TABLE IF NOT EXISTS  "dp_db_tables" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" TEXT NOT NULL,
	"dp_db_id" INTEGER NOT NULL
  );
  CREATE TABLE IF NOT EXISTS  "dp_db_columns" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"dp_db_id" INTEGER NOT NULL,
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
  CREATE UNIQUE INDEX index_databases ON dp_databases(name,type,dskey);
  CREATE UNIQUE INDEX index_tables ON dp_db_tables(dp_db_id,name);
  CREATE UNIQUE INDEX index_columns ON dp_db_columns(dp_db_id,dp_db_table_id,column_name);
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
	AddClass(string, string, string) error
	GetTags() ([]Tag, error)
	AddTag(string, string, []string) error

	GetAssociatedTags(string) ([]Tag, error)
	GetAssociatedClasses(string) ([]Class, error)
	SetSchemaData(DpDbDatabase) error

	AddDataSource(DpDataSource) error
	GetDataSources() ([]DpDataSource, error)
	DeleteDataSources(ids []int64) (bool, error)
}

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
	DsKey      string      `json:"DbKey"`
	Name       string      `json:"Name"`
	Type       string      `json:"Type"`
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

func GetAllDefaultClassAndTags() ([]*Tag, []*Class, error) {

	defaultTagsFile := viper.GetString("storage.default_tags")
	defaultClassesFile := viper.GetString("storage.default_classes")
	log.Debug().Msgf("default Tags File:  %v", defaultTagsFile)
	log.Debug().Msgf("default Classes File:  %v", defaultClassesFile)

	tagr, err := os.Open(defaultTagsFile)
	if err != nil {
		return []*Tag{}, []*Class{}, err
	}

	classr, err := os.Open(defaultClassesFile)
	if err != nil {
		return []*Tag{}, []*Class{}, err
	}

	tags := []*Tag{}
	classes := []*Class{}

	if err := gocsv.Unmarshal(tagr, &tags); err != nil {
		log.Error().Err(err).Msg("Unmarshal tag failed")
		return []*Tag{}, []*Class{}, err

	}
	if err := gocsv.UnmarshalFile(classr, &classes); err != nil {
		log.Error().Err(err).Msg("Unmarshal class failed")
		return []*Tag{}, []*Class{}, err

	}
	return tags, classes, nil
}
func GetStorageInstance() (Storage, error) {
	stype := viper.GetString("storage.type")
	spath := viper.GetString("storage.path")

	return New(StorageConfig{Type: stype, Path: spath})

}

func New(config StorageConfig) (Storage, error) {

	switch config.Type {
	case "internal":
		//return NewInternalStorage(config.Path)
		return getInternalStorageInstance(config.Path)

	default:
		return nil, fmt.Errorf("%s storage not found", config.Type)
	}
}
