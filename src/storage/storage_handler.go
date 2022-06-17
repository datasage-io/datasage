package storage

import (
	"fmt"
	"log"
	"os"

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

type Storage interface {
	GetClasses() ([]Class, error)
	GetTags() ([]Tag, error)
	GetAssociatedTags(string) ([]Tag, error)
}

type StorageConfig struct {
	Type string
	Path string
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

		return NewInternalStorage(config.Path)

	default:
		return nil, fmt.Errorf("%s storage not found", config.Type)
	}
}
