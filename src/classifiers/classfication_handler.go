package classifiers

import (
	"log"

	"github.com/datasage-io/datasage/src/adaptors"
	"github.com/datasage-io/datasage/src/storage"
)

func Run() {

	//Fetch MetaData
	adpt, err := adaptors.New(adaptors.AdaptorConfig{
		Type:     "mysql",
		Username: "appuser",
		Password: "appuserpassword",
		Host:     "localhost"})

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

	//Create new internal storage
	st, err := storage.NewInternalStorage("datasage.db")
	if err != nil {
		log.Println(err.Error())
	}

	//get all classes
	classes, err := st.GetAllClasses()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(len(classes))

	//get all tags
	tags, err := st.GetAllTags()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(len(tags))
}
