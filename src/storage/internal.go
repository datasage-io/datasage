package storage

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "modernc.org/sqlite"
)

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
		log.Println("Prepare statement error :", errP.Error())
		return errP
	}
	defer dbInsert.Close()
	_, errE := dbInsert.Exec(dpDataSource.Datadomain, dpDataSource.Dsname, dpDataSource.Dsdecription, dpDataSource.Dstype, dpDataSource.DsKey, dpDataSource.Dsversion, dpDataSource.Host, dpDataSource.Port, dpDataSource.User, dpDataSource.Password)
	if errE != nil {
		log.Println("Exec statement error", errE.Error())
		return errE
	}
	return nil
}

func (into InternalStorage) GetDataSources() ([]DpDataSource, error) {
	log.Println("InternalStorage GetDpDataSources")
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
	log.Println("InternalStorage DeleteDpDataSources", ids)
	for _, dsid := range ids {
		res, err := into.SqliteConnc.Exec("DELETE FROM DpDataSource where ID=$1", dsid)
		if err == nil {
			count, err := res.RowsAffected()
			log.Println("InternalStorage DeleteDpDataSources count is ", count)
			if err == nil && count > 0 {

				return true, nil
			}

		} else {
			log.Println(err)
		}
	}
	log.Println("InternalStorage DeleteDpDataSources count exit ")

	return false, nil
}

func (into InternalStorage) SetSchemaData(dpDbDatabase DpDbDatabase) error {

	log.Println("InternalStorage SetSchemaData", dpDbDatabase.Name)

	dbInsert, err := into.SqliteConnc.Prepare(`
	INSERT INTO dp_databases ("name","type","dskey")
	VALUES (?,?,?);
	`)
	if err != nil {
		log.Println("error1", err.Error())
		return err
	}
	dpDb, err := dbInsert.Exec(dpDbDatabase.Name, dpDbDatabase.Type, dpDbDatabase.DsKey)
	if err != nil {
		log.Println("error2", err.Error())
		return err
	}
	dpDbId, err := dpDb.LastInsertId()

	dbInsert.Close()

	for _, table := range dpDbDatabase.DpDbTables {

		tableInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_tables ("name","dp_db_id")
		 VALUES (?,?);
		`)
		if err != nil {
			log.Println("error3", err.Error())
			return err
		}

		//fmt.Println("INSERT INTO dp_db_tables :", table.Name, dpDbId)
		dbDbTable, err := tableInsert.Exec(table.Name, dpDbId)
		if err != nil {
			log.Println("error3b", err.Error())
			return err
		}

		dbDbTableId, err := dbDbTable.LastInsertId()
		tableInsert.Close()

		columnInsert, err := into.SqliteConnc.Prepare(`
		INSERT INTO dp_db_columns ("dp_db_id","dp_db_table_id","column_name","column_type","column_comment","Tags","Classes")
		 VALUES (?,?,?,?,?,?,?);
		`)
		if err != nil {
			log.Println("error4", err.Error())
			continue
		}

		for _, column := range table.DpDbColumns {

			//fmt.Println("INSERT INTO dp_db_columns:", dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment)
			_, err = columnInsert.Exec(dpDbId, dbDbTableId, column.ColumnName, column.ColumnType, column.ColumnComment, column.Tags, column.Classes)
			if err != nil {
				log.Println("error5", err.Error())
				//	continue
			}

		}
		columnInsert.Close()

	}
	log.Println("InternalStorage SetSchemaData done")
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

	log.Println("NewInternalStorage enter")
	var isnew bool
	_, err := os.Stat(dsn)
	if os.IsNotExist(err) {
		log.Println("Creating sqlite database ", dsn)
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
		log.Println("Inserting Default data")
		tags, classes, err := GetAllDefaultClassAndTags()
		if err != nil {
			return InternalStorage{}, err
		}
		err = insto.InsertDefaultData(tags, classes)
	}
	log.Println("NewInternalStorage exit")

	return insto, nil

}

func (insto InternalStorage) AddTag(name string, description string, rules []string) error {
	log.Println("AddTag ")
	tagInsert, err := insto.SqliteConnc.Prepare(`
	INSERT INTO tag ("tag_name","description","rule")
	 VALUES (?,?,?);
	`)
	if err != nil {
		return err
	}
	//defer tagInsert.Close()
	for _, rule := range rules {
		log.Println("AddTag rule: %v ", rule)
		_, err = tagInsert.Exec(name, description, rule)
		if err != nil {
			return err
		}
	}
	tagInsert.Close()
	log.Println("AddTag done")
	return nil
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
