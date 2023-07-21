package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Mysql  MysqlCfg  `yaml:"mysql"`
	Export ExportCfg `yaml:"export"`
}

type MysqlCfg struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
}

type ExportCfg struct {
	OnceNumber int64 `yaml:"oncenumber"`
}

var cfg = Config{}

const (
	configFile   = "./config.yaml"
	tableFileTml = "./%s_tables.json"
	dataDirTml   = "./%s_dir"
)

type TableInfo struct {
	//Name  string `json:"-"`
	MinID int64 `json:"min_id"`
	MaxID int64 `json:"max_id"`
}

func loadCfg() error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}
	return nil
}

func getTables(db *sql.DB) ([]string, error) {
	tables := make([]string, 0)

	// Get table list
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return tables, err
	}
	defer rows.Close()

	// Read result
	for rows.Next() {
		var table sql.NullString
		if err := rows.Scan(&table); err != nil {
			return tables, err
		}
		tables = append(tables, table.String)
	}
	return tables, rows.Err()
}

func getMinMaxIds(db *sql.DB, tableName string) (int64, int64, error) {

	// Get table list
	querySql := fmt.Sprintf("select min(id), max(id) from %s;", tableName)
	rows, err := db.Query(querySql)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	// Read result
	var minID, maxID int64
	for rows.Next() {
		if err := rows.Scan(&minID, &maxID); err != nil {
			return 0, 0, err
		}
	}
	return minID, maxID, nil
}

func reloadTables(db *sql.DB, dbname string) (map[string]TableInfo, error) {
	tables, err := getTables(db)
	if err != nil {
		return nil, fmt.Errorf("getTables, err: ", err)
	}
	//fmt.Println(tables)

	tableInfos := make(map[string]TableInfo)
	tableFile := fmt.Sprintf(tableFileTml, dbname)
	if exist, _ := exists(tableFile); exist {
		data, e := ioutil.ReadFile(tableFile)
		if e != nil {
			return nil, fmt.Errorf("open table file failed: ", e)
		}
		e = jsoniter.Unmarshal(data, &tableInfos)
		if e != nil {
			return nil, fmt.Errorf(" parse table file failed: ", e)
		}
		return tableInfos, nil
	}

	f, err := os.OpenFile(tableFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open table file failed:  ", err)
	}
	defer f.Close()

	for _, tableName := range tables {
		minID, maxID, e := getMinMaxIds(db, tableName)
		if e != nil {
			fmt.Printf("[WARN] getMinMaxIds failed: %s, table: %s\n", e.Error(), tableName)
			continue
		}
		tableInfos[tableName] = TableInfo{
			MinID: minID,
			MaxID: maxID,
		}
	}
	tableInfosJson, _ := jsoniter.MarshalToString(tableInfos)
	_, err = f.WriteString(tableInfosJson)
	if err != nil {
		return nil, fmt.Errorf("write to table file failed: %s", err.Error())
	}

	return tableInfos, nil
}

func doDump() error {
	err := loadCfg()
	if err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return errors.New("args is not enough")
	}
	dbname := os.Args[1]

	fmt.Println(cfg, dbname)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port, dbname))
	if err != nil {
		return fmt.Errorf("opening.database: ", err)
	}
	defer db.Close()

	tableInfos, err := reloadTables(db, dbname) // 更新要导出的表信息
	if err != nil {
		return err
	}
	//fmt.Println("tableInfos: ", tableInfos)
	for tableName, info := range tableInfos {
		e := dumpOneTable(db, tableName, info)
		if e != nil {
			return e
		}
		break
	}
	return nil
}

func main() {
	err := doDump()
	if err != nil {
		fmt.Println("[ERROR] dump failed, err: ", err)
		return
	}
	fmt.Println("[INFO] dump success")
}

func dumpOneTable(db *sql.DB, tableName string, tableInfo TableInfo) error {
	createSql, err := createTableSQL(db, tableName)
	if err != nil {
		return err
	}
	fmt.Println(createSql)

	pos := tableInfo.MinID

	values, err := createTableValues(db, tableName, pos, pos+cfg.Export.OnceNumber)
	if err != nil {
		return err
	}
	fmt.Println("VALUES: ", values)

	return nil
}

func createTableSQL(db *sql.DB, name string) (string, error) {
	// Get table creation SQL
	var table_return sql.NullString
	var table_sql sql.NullString
	err := db.QueryRow("SHOW CREATE TABLE "+name).Scan(&table_return, &table_sql)

	if err != nil {
		return "", err
	}
	if table_return.String != name {
		return "", errors.New("Returned table is not the same as requested table")
	}

	return table_sql.String, nil
}

func createTableValues(db *sql.DB, name string, minID, maxID int64) (string, error) {
	// Get Data
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE id >= %d AND id < %d", name, minID, maxID))
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	if len(columns) == 0 {
		return "", errors.New("No columns in table " + name + ".")
	}

	// Read data
	data_text := make([]string, 0)
	for rows.Next() {
		// Init temp data storage

		//ptrs := make([]interface{}, len(columns))
		//var ptrs []interface {} = make([]*sql.NullString, len(columns))

		data := make([]*sql.NullString, len(columns))
		ptrs := make([]interface{}, len(columns))
		for i, _ := range data {
			ptrs[i] = &data[i]
		}

		// Read data
		if err := rows.Scan(ptrs...); err != nil {
			return "", err
		}

		dataStrings := make([]string, len(columns))

		for key, value := range data {
			if value != nil && value.Valid {
				dataStrings[key] = value.String
			}
		}

		data_text = append(data_text, "('"+strings.Join(dataStrings, "','")+"')")
	}

	return strings.Join(data_text, ","), rows.Err()
}

func exists(p string) (bool, os.FileInfo) {
	f, err := os.Open(p)
	if err != nil {
		return false, nil
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, nil
	}
	return true, fi
}
