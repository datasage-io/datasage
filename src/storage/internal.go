package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

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

func (into InternalStorage) AddDataSource(dpDataSource DpDataSource) (int64, error) {

	dbInsert, errP := into.SqliteConnc.Prepare(`
	INSERT INTO DpDataSource ("Datadomain","Dsname","Dsdecription","Dstype","DsKey","Dsversion","Host","Port","User","Password" )
	VALUES (?,?,?,?,?,?,?,?,?,?);
	`)
	if errP != nil {
		//log.Println("Prepare statement error :", errP.Error())
		return -1, errP
	}
	defer dbInsert.Close()
	dpDb, err := dbInsert.Exec(dpDataSource.Datadomain,
		dpDataSource.Dsname,
		dpDataSource.Dsdecription,
		dpDataSource.Dstype,
		dpDataSource.DsKey,
		dpDataSource.Dsversion,
		dpDataSource.Host,
		dpDataSource.Port,
		dpDataSource.User,
		dpDataSource.Password)
	if err != nil {
		log.Error().Err(err).Msg("AddDataSource")
		return -1, err
	}
	var dpDbId int64
	dpDbId, _ = dpDb.LastInsertId()

	return dpDbId, nil
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
	fmt.Println("Before sleep the time is:", time.Now().Unix()) // Before sleep the time is: 1257894000
	time.Sleep(10 * time.Second)                                // pauses execution for 2 seconds
	fmt.Println("After sleep the time is:", time.Now().Unix())
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

func (insto InternalStorage) GetScanLog(ds string, db string, table string, columns []string) ([]CScan, error) {
	log.Trace().Msgf("InternalStorage GetScanLog : %v : %v :%v :%v", ds, db, table, columns)

	cscan := []CScan{}

	rows, err := insto.SqliteConnc.Query("SELECT DpDatasource.Dsname , dp_databases.name ,  dp_db_tables.name , dp_db_columns.column_name , dp_db_columns.Tags,dp_db_columns.Classes  FROM (((dp_db_columns INNER JOIN dp_db_tables ON  dp_db_columns.dp_db_table_id = dp_db_tables.id ) INNER JOIN dp_databases  ON  dp_db_tables.dp_db_id = dp_databases.id) INNER JOIN DpDatasource  ON  dp_databases.dskey = DpDatasource.DsKey)")
	if err != nil {
		return cscan, err
	}
	//defer rows.Close()
	for rows.Next() {
		temp := CScan{}
		err = rows.Scan(&temp.DsName, &temp.DbName,
			&temp.TableName, &temp.Columns, &temp.Classes, &temp.Tags)

		cscan = append(cscan, temp)
	}
	rows.Close()
	return cscan, nil

}

func (into InternalStorage) Scan(dsname string) error {
	log.Trace().Msgf("InternalStorage Scan :%v ", dsname)

	return nil

}
func (into InternalStorage) GetStatus(dsname string) (string, error) {
	log.Trace().Msgf("InternalStorage GetStatus :%v ", dsname)
	return "", nil
}
func (into InternalStorage) GetRecommendedPolicy(dsname string) ([]RecommendedPolicy, error) {
	log.Trace().Msgf("InternalStorage GetRecommendedPolicy :%v ", dsname)
	var recommedPolicy = []string{"GDPR Audit", "PII Audit", "SOC 2 Audit", "HIPAA Audit", "UDI Audit"}
	var rPolicies []RecommendedPolicy
	for i := range recommedPolicy {
		name := recommedPolicy[i]
		rPolicies = append(rPolicies, RecommendedPolicy{PolicyId: int32(i), PolicyName: name})
	}
	return rPolicies, nil
}

func (into InternalStorage) GetDataSource(dsname string) (DpDataSource, error) {
	log.Trace().Msgf("InternalStorage GetDpDataSources")
	dataSource := DpDataSource{}
	qryString := "SELECT * FROM DpDataSource WHERE Dsname=" + dsname

	rows, err := into.SqliteConnc.Query(qryString)
	if err != nil {
		return dataSource, err
	}
	defer rows.Close()
	err = rows.Scan(&dataSource.ID,
		&dataSource.Datadomain,
		&dataSource.Dsname,
		&dataSource.Dsdecription,
		&dataSource.Dstype,
		&dataSource.DsKey,
		&dataSource.Dsversion,
		&dataSource.Host,
		&dataSource.Port,
		&dataSource.User,
		&dataSource.Password)
	return dataSource, err
}

func (into InternalStorage) ApplyPolicy(dsname string, ids []int64) error {
	log.Trace().Msgf("InternalStorage ApplyPolicy :%v :%v", dsname, ids)
	return nil
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
	instance, _ = NewInternalStorage(dsn)
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
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(10)

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

func (into InternalStorage) DeleteClasses(ids []int64) (bool, error) {
	log.Trace().Msgf("InternalStorage DeleteClasses %v", ids)
	distictids := removeDuplicateValues(ids)
	fmt.Println(distictids)

	for _, classid := range ids {
		if res, err := into.SqliteConnc.Exec("DELETE FROM class where ID=$1", classid); err == nil {
			fmt.Println(res)
			if err != nil {
				return false, err
			}
			count, err := res.RowsAffected()
			log.Debug().Msgf("Deleteclasses count is %v", count)
			if count == 0 {
				return false, err
			}
		} else {
			log.Error().Err(err).Msg("Deleteclasses")
			return false, err
		}
	}
	return true, nil
}

func (into InternalStorage) DeleteTags(ids []int64) (bool, error) {
	log.Trace().Msgf("InternalStorage DeleteTags %v", ids)
	distictids := removeDuplicateValues(ids)
	fmt.Println(distictids)
	for _, tagsid := range distictids {
		if res, err := into.SqliteConnc.Exec("DELETE FROM Tag where ID=$1", tagsid); err == nil {
			fmt.Println(res)
			if err != nil {
				return false, err
			}

			count, err := res.RowsAffected()
			log.Debug().Msgf("DeleteTags count is %v", count)
			if count == 0 {
				fmt.Println("delete failed ", tagsid)
				return false, err
			}
		} else {
			log.Error().Err(err).Msg("DeleteTags")
			return false, err
		}
	}
	return true, nil

}

//Remove Duplicate Keys
func removeDuplicateValues(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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

func (insto InternalStorage) UpdateDSStatus(dsid int64, statusid int64) error {
	insstmnt := "INSERT INTO ds_scanstatus (\"ds_id\",\"status_id\") VALUES (" + strconv.FormatInt(dsid, 10) + "," + strconv.FormatInt(statusid, 10) + ")"
	updatestmnt := "UPDATE ds_scanstatus SET status_id=" + strconv.FormatInt(statusid, 10) + "  WHERE ds_id =" + strconv.FormatInt(dsid, 10)
	urows, err := insto.SqliteConnc.Query(insstmnt)
	log.Debug().Msgf("insstmnt:: %v", insstmnt)

	if err != nil {
		log.Error().Err(err).Msg("update")
		//return err
		rows, err1 := insto.SqliteConnc.Query(updatestmnt)
		log.Debug().Msgf("updatestmnt:: %v", updatestmnt)

		if err1 != nil {
			return err1
		}
		rows.Close()
	} else {
		urows.Close()
	}
	return nil
}
