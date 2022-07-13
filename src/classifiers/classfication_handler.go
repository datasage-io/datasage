package classifiers

import (
	"fmt"
	"log"
	"regexp"

	"github.com/datasage-io/datasage/src/adaptors"
	"github.com/datasage-io/datasage/src/storage"
)

type DpDataSource struct {
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
}

func Run(dpDataSource DpDataSource) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Classification Handler Run")
	//Fetch MetaData
	/*
		adpt, err := adaptors.New(adaptors.AdaptorConfig{
			Type:     "mysql",
			Username: "user1",
			Password: "Accu0104#",
			Host:     "localhost"})
	*/

	adpt, err := adaptors.New(adaptors.AdaptorConfig{
		Type:     dpDataSource.Dstype,
		Username: dpDataSource.User,
		Password: dpDataSource.Password,
		Host:     dpDataSource.Host})

	if err != nil {
		log.Fatal(err.Error())
	}
	info, err := adpt.Scan()

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err != nil {
		log.Println(err.Error())
	}

	st, err := storage.New(storage.StorageConfig{Type: "internal", Path: "datasageD.db"})
	if err != nil {
		log.Fatal(err.Error())
	}
	//get all classes
	classes, err := st.GetClasses()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("no of classes:= ", len(classes))

	//get all tags
	tags, err := st.GetTags()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("no of tags:= ", len(tags))

	storageDpDataSourceObj := storage.DpDataSource{
		ID:           -1,
		Datadomain:   dpDataSource.Datadomain,
		Dsname:       dpDataSource.Dsname,
		Dsdecription: dpDataSource.Dsdecription,
		Dstype:       dpDataSource.Dstype,
		DsKey:        dpDataSource.DsKey,
		Dsversion:    dpDataSource.Dsversion,
		Host:         dpDataSource.Host,
		Port:         dpDataSource.Port,
		User:         dpDataSource.User,
		Password:     dpDataSource.Password,
	}

	err1 := st.SetDpDataSourceData(storageDpDataSourceObj)
	if err1 != nil {
		fmt.Println("SetDpDataSourceData error: ", err)
		return
	}

	/*
		phonetag, err := st.GetAssociatedTags("Phone Number")
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(phonetag)
	*/
	log.Println(" identifying class and tags start")
	columsCount := 0
	tableCount := 0
	dBCount := 0
	skipDBs := []string{"mysql", "performance_schema", "datadefender"}

	for _, sc := range info.Schemas {
		dBCount = dBCount + 1
		log.Println("DB name:= ", sc.Name)
		skip := false
		for _, skipDB := range skipDBs {

			if skipDB == sc.Name {
				log.Println("skip DB name:= ", sc.Name)
				skip = true
			}
		}
		if skip == true {
			continue
		}

		dpDbTables := []storage.DpDbTable{}
		for _, tb := range sc.Tables {
			tableCount = tableCount + 1
			//dpDbColumn := storage.DpDbColumn

			dpDbColumns := []storage.DpDbColumn{}

			for _, cols := range tb.Cols {
				columsCount = columsCount + 1
				colName, err1 := removeSpecialChars(cols.ColumnName)
				if err1 != nil {
					log.Println(err.Error())
					continue
				}
				relatedtags, _ := st.GetAssociatedTags(colName)
				relatedclasses, _ := st.GetAssociatedClasses(colName)

				//if err != nil {
				//	log.Println(err.Error())
				//}
				tags := ""
				classes := ""
				if len(relatedclasses) > 0 {
					for _, relatedclass := range relatedclasses {
						log.Println("Class:= ", relatedclass.Class)
						classes = classes + ";" + relatedclass.Class
					}
				}
				if len(relatedtags) > 0 {
					for _, relatedtag := range relatedtags {
						log.Println("TagName:", relatedtag.TagName)
						tags = tags + ";" + relatedtag.TagName
					}
				}
				/*
					else {
						continue
					}
				*/

				col := storage.DpDbColumn{
					ColumnName:    colName,
					ColumnType:    cols.ColumnType,
					ColumnComment: cols.ColumnComment,
					Tags:          tags,
					Classes:       classes,
				}
				dpDbColumns = append(dpDbColumns, col)
			}
			dpDbTable := storage.DpDbTable{

				Name:        tb.Name,
				Tags:        "",
				DpDbColumns: dpDbColumns,
			}
			dpDbTables = append(dpDbTables, dpDbTable)

		}
		schema := storage.DpDbDatabase{
			DbKey:      "todo",
			Name:       sc.Name,
			Type:       "mysql",
			Key:        "987654321",
			DpDbTables: dpDbTables,
		}
		err := st.SetSchemaData(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

	log.Println(" identifying class and tags completed")
	log.Println(" columsCount", columsCount)

}

func ListDatasources() ([]storage.DpDataSource, error) {
	log.Println("ListDatasources")
	st, err := storage.New(storage.StorageConfig{Type: "internal", Path: "datasageD.db"})
	if err != nil {
		log.Fatal(err.Error())
	}
	dataSources, err := st.GetDpDataSources()
	log.Println(dataSources)
	return dataSources, err

}

func DeleteDatasource(ids []int64) bool {
	overallStatus := true
	log.Println("DeleteDatasources ")
	st, err := storage.New(storage.StorageConfig{Type: "internal", Path: "datasageD.db"})
	if err == nil {
		for i := range ids {
			status, error := st.DeleteDpDataSources(ids[i])
			if error != nil || status == false {
				log.Println(error)
				overallStatus = false
			}
		}
	}
	return overallStatus
}

//removeSpecialChars - Removes the special characters from the string.
func removeSpecialChars(char string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(char, " ")
	return processedString, nil
}
