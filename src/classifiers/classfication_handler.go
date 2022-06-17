package classifiers

import (
	"log"

	"github.com/datasage-io/datasage/src/adaptors"
	"github.com/datasage-io/datasage/src/storage"
)

func Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

	if err != nil {
		log.Println(err.Error())
	}

	st, err := storage.New(storage.StorageConfig{Type: "internal", Path: "datasage.db"})
	if err != nil {
		log.Fatal(err.Error())
	}
	//get all classes
	classes, err := st.GetClasses()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("no of classes:= ", len(classes))

	//get all tags
	tags, err := st.GetTags()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("no of tags:= ", len(tags))

	phonetag, err := st.GetAssociatedTags("Phone Number")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(phonetag)
}
