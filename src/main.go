package main

import (
	"log"

	"github.com/datasage-io/datasage/src/adaptors"
)

func main() {
	adpt, err := adaptors.New(adaptors.AdaptorConfig{
		Type:     "mysql",
		Username: "appuser",
		Password: "appuserpassword",
		Host:     "localhost:3306"})

	if err != nil {
		log.Fatal(err.Error())
	}
	info, err := adpt.Scan()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, sc := range info.Schemas {
		log.Println("DB name:= ", sc.Name)

		for _, tb := range sc.Tables {
			log.Println("Table Name:= ", tb.Name)

			for _, cols := range tb.Cols {
				log.Println("columns:= ", cols)
			}
		}
	}
}
