package storage

import (
	"database/sql"
	"os"
	"strconv"
	"sync"

	logger "github.com/datasage-io/datasage/src/logger"
	"github.com/rs/zerolog"
	_ "modernc.org/sqlite"
)

var log *zerolog.Logger = logger.GetInstance()

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
	INSERT INTO tag ("tag_name","rule","description")
	 VALUES (?,?,?);
	`)
	defer tagInsert.Close()
	for _, t := range tags {
		_, err = tagInsert.Exec(t.TagName, t.Rule, t.Description)
	}

	classInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO class ("description","rule","class")
	 VALUES (?,?,?);
	`)
	defer classInsert.Close()
	for _, c := range classes {
		_, err = classInsert.Exec(c.Description, c.Rule, c.Class)
	}

	return err
}

func (into InternalStorage) AddDataSource(dpDataSource DpDataSource) error {

	dbInsert, errP := into.SqliteConnc.Prepare(`
	INSERT INTO DpDataSource ("Datadomain","Dsname","Dsdecription","Dstype","DsKey","Dsversion","Host","Port","User","Password" )
	VALUES (?,?,?,?,?,?,?,?,?,?);
	`)
	if errP != nil {
		//log.Println("Prepare statement error :", errP.Error())
		return errP
	}
	defer dbInsert.Close()
	if _, err := dbInsert.Exec(dpDataSource.Datadomain,
		dpDataSource.Dsname,
		dpDataSource.Dsdecription,
		dpDataSource.Dstype,
		dpDataSource.DsKey,
		dpDataSource.Dsversion,
		dpDataSource.Host,
		dpDataSource.Port,
		dpDataSource.User,
		dpDataSource.Password); err != nil {
		log.Error().Err(err).Msg("AddDataSource")
		return err
	}
	return nil
}

func (into InternalStorage) GetDataSources() ([]DpDataSource, error) {
	log.Trace().Msgf("InternalStorage GetDpDataSources")
	dataSources := []DpDataSource{}

	rows, err := into.SqliteConnc.Query("SELECT * FROM DpDataSource")
	if err != nil {
		return dataSources, err
	}
	defer rows.Close()
	for rows.Next() {
		tempDataSource := DpDataSource{}
		err = rows.Scan(&tempDataSource.ID,
			&tempDataSource.Datadomain,
			&tempDataSource.Dsname,
			&tempDataSource.Dsdecription,
			&tempDataSource.Dstype,
			&tempDataSource.DsKey,
			&tempDataSource.Dsversion,
			&tempDataSource.Host,
			&tempDataSource.Port,
			&tempDataSource.User,
			&tempDataSource.Password)

		dataSources = append(dataSources, tempDataSource)
	}
	return dataSources, err
}

func (into InternalStorage) DeleteDataSources(ids []int64) (bool, error) {
	log.Trace().Msgf("InternalStorage DeleteDataSources %v", ids)
	for _, dsid := range ids {
		if res, err := into.SqliteConnc.Exec("DELETE FROM DpDataSource where ID=$1", dsid); err == nil {
			count, err := res.RowsAffected()
			log.Debug().Msgf("DeleteDpDataSources count is %v", count)
			if err == nil && count > 0 {
				return true, nil
			}
		} else {
			log.Error().Err(err).Msg("DeleteDataSources")
		}
	}
	return false, nil
}

func (into InternalStorage) SetSchemaData(dpDbDatabase DpDbDatabase) error {
	log.Trace().Msgf("InternalStorage SetSchemaData %v", dpDbDatabase)

	dbInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO dp_databases ("name","type","dskey")
	VALUES (?,?,?);
	`)
	/*
		if err != nil {
			log.Error().Err(err).Msg("SetSchemaData")
			return err
		}
	*/
	dpDb, err := dbInsert.Exec(dpDbDatabase.Name, dpDbDatabase.Type, dpDbDatabase.DsKey)
	var dpDbId int64
	if err != nil {
		//log.Error().Err(err).Msg("dbInsert")
		//ignore this error. may be the Schema already got scanned
		//find the id for this
		stmnt := "SELECT id FROM dp_databases WHERE name =" + "\"" + dpDbDatabase.Name +
			"\" AND type =" + "\"" + dpDbDatabase.Type + "\" AND dskey =" + "\"" + dpDbDatabase.DsKey + "\""
		log.Debug().Msgf("stmnt is: %v", stmnt)

		rows, err := into.SqliteConnc.Query(stmnt)
		if err != nil {
			log.Error().Err(err).Msg("select exec")
			return err
		}
		var id int64
		for rows.Next() {
			err = rows.Scan(&id)
			if err != nil {
				log.Error().Err(err).Msg("select scan")
				return err
			}
		}
		rows.Close()
		dpDbId = id
	} else {
		dpDbId, _ = dpDb.LastInsertId()
		dbInsert.Close()
	}

	log.Debug().Msgf("dpDbId: %v", dpDbId)

	var dbDbTableId int64
	for _, table := range dpDbDatabase.DpDbTables {

		tableInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_tables ("name","dp_db_id")
		 VALUES (?,?);
		`)
		if err != nil {
			log.Error().Err(err).Msg("insert prep")
			return err
		}

		//fmt.Println("INSERT INTO dp_db_tables :", table.Name, dpDbId)
		dbDbTable, err := tableInsert.Exec(table.Name, dpDbId)
		if err != nil {
			//log.Error().Err(err).Msg("insert exec")
			stmnt := "SELECT id FROM dp_db_tables WHERE name =" + "\"" + table.Name +
				"\" AND dp_db_id = " + strconv.FormatInt(dpDbId, 10)
			log.Debug().Msgf("stmnt is: %v", stmnt)

			rows, err := into.SqliteConnc.Query(stmnt)
			if err != nil {
				log.Error().Err(err).Msg("select exec")
				return err
			}
			var id int64
			for rows.Next() {
				err = rows.Scan(&id)
				if err != nil {
					log.Error().Err(err).Msg("select scan")
					return err
				}
			}
			dbDbTableId = id
			rows.Close()
		} else {

			dbDbTableId, _ = dbDbTable.LastInsertId()
			tableInsert.Close()
		}
		log.Debug().Msgf("TableId: %v", dbDbTableId)

		columnInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_columns ("dp_db_id","dp_db_table_id","column_name","column_type","column_comment","Tags","Classes")
		 VALUES (?,?,?,?,?,?,?);
		`)
		if err != nil {
			//log.Error().Err(err).Msg("insert prep")
			continue
		}

		for _, column := range table.DpDbColumns {
			log.Debug().Msgf("column: %v", column)

			//fmt.Println("INSERT INTO dp_db_columns:", dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment)
			_, err = columnInsert.Exec(dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment, column.Tags, column.Classes)
			if err != nil {
				//log.Error().Err(err).Msg("insert exec")
				//	continue
				//insert failed. We try update because column may exists already
				stmnt := "UPDATE dp_db_columns SET column_type=\"" + column.ColumnType + "\"" +
					" , column_comment=\"" + column.ColumnComment + "\"" +
					" , Tags=\"" + column.Tags + "\"" +
					" , Classes=\"" + column.Classes + "\"" +
					"  WHERE dp_db_id =" + strconv.FormatInt(dpDbId, 10) + " AND" +
					" dp_db_table_id =" + strconv.FormatInt(dbDbTableId, 10) + " AND" +
					" column_name = \"" + column.ColumnName + "\""

				log.Debug().Msgf("stmnt is: %v", stmnt)

				rows, err := into.SqliteConnc.Query(stmnt)
				if err != nil {
					log.Error().Err(err).Msg("update")
					return err
				}
				rows.Close()
			}
		}
		columnInsert.Close()
	}

	log.Trace().Msgf("InternalStorage SetSchemaData execution completed")
	return err
}

/*
var once sync.Once

func CreateInternalStorage(dsn string) (InternalStorage, error) {
	once.Do(func() { NewInternalStorage(dsn) })
}
*/

var (
	instance InternalStorage
	once     sync.Once
)

func getInternalStorageInstance(dsn string) (InternalStorage, error) {
	once.Do(func() {
		instance, _ = NewInternalStorage(dsn)
	})
	return instance, nil
}
func NewInternalStorage(dsn string) (InternalStorage, error) {
	log.Trace().Msgf("NewInternalStorage enter %v", dsn)
	var isnew bool
	_, err := os.Stat(dsn)
	if os.IsNotExist(err) {
		log.Debug().Msgf("Creating sqlite database : %v", dsn)
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
		log.Debug().Msgf("Inserting Default data")
		tags, classes, err := GetAllDefaultClassAndTags()
		if err != nil {
			return InternalStorage{}, err
		}
		err = insto.InsertDefaultData(tags, classes)
	}
	log.Trace().Msgf("NewInternalStorage exit %v", dsn)

	return insto, nil

}

func (insto InternalStorage) AddTag(name string, description string, rules []string) error {
	log.Trace().Msgf("AddTag name:  %v des: %v  rules: %v  ", name, description, rules)

	tagInsert, err := insto.SqliteConnc.Prepare(`
	INSERT INTO tag ("tag_name","description","rule")
	 VALUES (?,?,?);
	`)
	if err != nil {
		return err
	}
	//defer tagInsert.Close()
	for _, rule := range rules {
		log.Debug().Msgf("AddTag rule:: %v", rule)
		_, err = tagInsert.Exec(name, description, rule)
		if err != nil {
			return err
		}
	}
	tagInsert.Close()
	return nil
}

func (insto InternalStorage) AddClass(description string, rule string, class string) error {
	log.Trace().Msgf("AddClass description:  %v rule: %v  class: %v  ", description, rule, class)

	classInsert, err := insto.SqliteConnc.Prepare(`
	INSERT INTO class ("description","rule","class")
	 VALUES (?,?,?);
	`)
	_, err = classInsert.Exec(description, rule, class)

	classInsert.Close()
	return err

}

func (insto InternalStorage) DeleteTag(id int64) error {
	return nil
}

func (insto InternalStorage) GetTags() ([]Tag, error) {
	tags := []Tag{}

	rows, err := insto.SqliteConnc.Query("SELECT id,tag_name,rule,description FROM tag")
	if err != nil {
		return tags, err
	}
	//defer rows.Close()
	for rows.Next() {
		tempTag := Tag{}
		err = rows.Scan(&tempTag.Id, &tempTag.TagName,
			&tempTag.Rule, &tempTag.Description)

		tags = append(tags, tempTag)
	}
	rows.Close()

	return tags, err
}

func (insto InternalStorage) GetClasses() ([]Class, error) {
	classes := []Class{}

	rows, err := insto.SqliteConnc.Query("SELECT id,description,rule,class FROM class")
	if err != nil {
		return classes, err
	}
	//defer rows.Close()
	for rows.Next() {
		tempClass := Class{}
		err = rows.Scan(&tempClass.Id, &tempClass.Description,
			&tempClass.Rule, &tempClass.Class)

		classes = append(classes, tempClass)
	}
	rows.Close()
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
	rows.Close()
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
	rows.Close()
	return classes, err
}
