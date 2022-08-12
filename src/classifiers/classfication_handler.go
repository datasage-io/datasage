package classifiers

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/datasage-io/datasage/src/adaptors"
	"github.com/datasage-io/datasage/src/storage"
	"github.com/datasage-io/datasage/src/utils/constants"

	logger "github.com/datasage-io/datasage/src/logger"
	"github.com/rs/zerolog"

	"github.com/spf13/viper"
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

var log *zerolog.Logger = logger.GetInstance()

func Run() {

	scanInterval := viper.GetInt("classifiers.dbschema-scan-interval")
	ticker := time.NewTicker(time.Duration(scanInterval) * time.Minute)
	for ; ; <-ticker.C {
		log.Debug().Msg("periodic scan started")
		st, err := storage.GetStorageInstance()
		if err != nil {
			log.Error().Err(err).Msg("GetStorageInstance Internal Error")
			continue
		}

		datasources, err := st.GetDataSources()
		if err != nil {
			log.Error().Err(err).Msg("Datasources not found")
		}
		for _, datasource := range datasources {
			ScanDataSource(datasource)
		}
		log.Debug().Msg(" periodic scan completed")

	}

}

func ScanDataSource(datasource storage.DpDataSource) error {
	log.Debug().Msgf("ScanDataSource started %v", datasource)

	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
		return err
	}
	adaptor, err := adaptors.New(adaptors.AdaptorConfig{
		Type:     datasource.Dstype,
		Username: datasource.User,
		Password: datasource.Password,
		Host:     datasource.Host,
	})

	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
		return err
	}

	//get all classes
	classes, err := st.GetClasses()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")

	}
	log.Info().Msgf("no of classes:=%v ", len(classes))

	//get all tags
	tags, err := st.GetTags()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	log.Info().Msgf("no of tags:=%v ", len(tags))
	columsCount := 0
	tableCount := 0
	dBCount := 0
	skipDBs := []string{"mysql", "performance_schema", "datadefender"}
	info, err := adaptor.Scan()
	for _, sc := range info.Schemas {
		dBCount = dBCount + 1
		log.Info().Msgf("DB name:= %v", sc.Name)
		skip := false
		for _, skipDB := range skipDBs {

			if skipDB == sc.Name {
				log.Info().Msgf("skip DB name:= %v", sc.Name)
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
				colName, err := removeSpecialChars(cols.ColumnName)
				if err != nil {
					log.Error().Err(err).Msg("Internal Error")
					continue
				}
				relatedtags, _ := st.GetAssociatedTags(colName)
				relatedclasses, _ := st.GetAssociatedClasses(colName)

				//if err != nil {
				//	log.Println(err.Error())
				//}
				tags := []string{}
				classes := []string{}
				if len(relatedclasses) > 0 {
					for _, relatedclass := range relatedclasses {
						log.Info().Msgf("Class:= %v", relatedclass.Class)
						classes = append(classes, relatedclass.Class)
					}
				}
				if len(relatedtags) > 0 {
					for _, relatedtag := range relatedtags {
						log.Info().Msgf("TagName:= %v", relatedtag.TagName)
						//tags = tags + ";" + relatedtag.TagName
						tags = append(tags, relatedtag.TagName)
					}
				}
				col := storage.DpDbColumn{
					ColumnName:    colName,
					ColumnType:    cols.ColumnType,
					ColumnComment: cols.ColumnComment,
					Tags:          strings.Join(tags, ";"),
					Classes:       strings.Join(classes, ";"),
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
			DsKey:      datasource.DsKey,
			Name:       sc.Name,
			Type:       datasource.Dstype,
			DpDbTables: dpDbTables,
		}
		err := st.SetSchemaData(schema)
		if err != nil {
			fmt.Println(err)
		}

	}
	log.Trace().Msgf("scan completed for datasource: %v", datasource)

	st.UpdateDSStatus(int64(datasource.ID), constants.DataSourceInitialScanCompleted)
	return nil
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
