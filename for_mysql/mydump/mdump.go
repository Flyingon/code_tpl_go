package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hpcloud/tail"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
	"time"
)

var process = Process{}

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
	DBName     string `yaml:"db_name"`
	OnceNumber int64  `yaml:"oncenumber"`
	FileSep    string `yaml:"file_sep"`
	LineSep    string `yaml:"line_sep"`
}

var cfg = Config{}

const (
	configFile   = "./config.yaml"
	tableFileTml = "./%s_tables.json"
	dataDirTml   = "./%s"
	processFile  = "./process.json"
)

type Process struct {
	TableIndex int `json:"table_index"`
}

type TableInfo struct {
	Name  string `json:"name"`
	MinID int64  `json:"min_id"`
	MaxID int64  `json:"max_id"`
}

func loadCfg() error {
	data, err := os.ReadFile(configFile)
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

func reloadTables(db *sql.DB, dbname string) ([]TableInfo, error) {
	tables, err := getTables(db)
	if err != nil {
		return nil, fmt.Errorf("getTables, err: %s", err)
	}
	//fmt.Println(tables)

	tableInfos := make([]TableInfo, 0, len(tables))
	tableFile := fmt.Sprintf(tableFileTml, dbname)
	if Exists(tableFile) {
		data, e := os.ReadFile(tableFile)
		if e != nil {
			return nil, fmt.Errorf("open table file failed: %s", e)
		}
		e = jsoniter.Unmarshal(data, &tableInfos)
		if e != nil {
			return nil, fmt.Errorf(" parse table file failed: %s", e)
		}
		return tableInfos, nil
	}

	for _, tableName := range tables {
		if tableName == "" {
			continue
		}

		minID, maxID, e := getMinMaxIds(db, tableName)
		if e != nil {
			fmt.Printf("[WARN] getMinMaxIds failed: %s, table: %s\n", e.Error(), tableName)
			continue
		}
		tableInfos = append(tableInfos, TableInfo{
			Name:  tableName,
			MinID: minID,
			MaxID: maxID,
		})
	}
	tableInfosJson, _ := jsoniter.Marshal(tableInfos)
	err = writeToFile(tableFile, tableInfosJson)
	if err != nil {
		return nil, fmt.Errorf("write to table file failed: %s", err.Error())
	}

	return tableInfos, nil
}

func doDump() (err error) {
	err = loadCfg()
	if err != nil {
		return err
	}

	dbname := cfg.Export.DBName
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port, dbname))
	if err != nil {
		return fmt.Errorf("opening.database: %v", err)
	}
	defer db.Close()

	tableInfos, err := reloadTables(db, dbname) // 更新要导出的表信息
	if err != nil {
		return err
	}
	//fmt.Println("tableInfos: ", tableInfos)

	if Exists(processFile) {
		data, e := os.ReadFile(processFile)
		if e != nil {
			return fmt.Errorf("read process file failed: %s", e.Error())
		}
		e = jsoniter.Unmarshal(data, &process)
		if e != nil {
			return fmt.Errorf("parse process file failed: %s", e.Error())
		}
	}
	defer func() {
		processJson, e := jsoniter.Marshal(&process)
		if e != nil {
			err = fmt.Errorf("marshal process file failed: %s", e.Error())
			return
		}
		e = writeToFile(processFile, processJson)
		if e != nil {
			err = fmt.Errorf("write to process file failed: %s", e.Error())
			return
		}
	}()

	for index, info := range tableInfos {
		if index < process.TableIndex {
			continue
		}
		process.TableIndex = index // 更新 index
		processJson, e := jsoniter.Marshal(&process)
		if e != nil {
			err = fmt.Errorf("marshal process file failed: %s", e.Error())
			return
		}
		e = writeToFile(processFile, processJson)
		if e != nil {
			err = fmt.Errorf("write to process file failed: %s", e.Error())
			return
		}

		exportErr := dumpOneTable(db, info)
		if exportErr != nil {
			return exportErr
		}
		//break
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

func dumpOneTable(db *sql.DB, tableInfo TableInfo) error {
	dirPath := fmt.Sprintf(dataDirTml, tableInfo.Name)
	if !Exists(dirPath) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	createSql, err := createTableSQL(db, tableInfo.Name)
	if err != nil {
		return err
	}
	//fmt.Println(createSql)
	createSqlFile := fmt.Sprintf("%s/create.sql", dirPath)
	err = writeToFile(createSqlFile, []byte(createSql))
	if err != nil {
		return err
	}

	dataFile := fmt.Sprintf("%s/data.csv", dirPath) // 数据文件存储
	pos := tableInfo.MinID
	if !Exists(dataFile) { // 第一次，存储表头
		stmt, err := sqlparser.Parse(createSql)
		if err != nil {
			return err
		}
		createTableStruct, _ := stmt.(*sqlparser.CreateTable)
		columnNames := make([]string, len(createTableStruct.Columns))
		for index, col := range createTableStruct.Columns {
			columnNames[index] = col.Name
		}
		err = writeToFile(dataFile, []byte(strings.Join(columnNames, ",")+cfg.Export.LineSep))
		if err != nil {
			return err
		}
	} else {
		pos, err = GetPosID(dataFile)
		if err != nil {
			return err
		}
		pos = pos + 1 // 当前这条数据已经有了
	}

	fmt.Printf("begin export table %s, from: %d\n", tableInfo.Name, pos)
	for pos < tableInfo.MaxID {
		start_ts := time.Now()
		value, e := createTableValues(db, tableInfo.Name, pos, pos+cfg.Export.OnceNumber, cfg.Export.FileSep)
		if e != nil {
			return e
		}
		//fmt.Println("VALUES: ", values)
		err = appendToFile(dataFile, []byte(value+cfg.Export.LineSep))
		if err != nil {
			return err
		}
		cost := time.Since(start_ts).Seconds()
		fmt.Printf("%0.5f, cost: %0.2f s\n", float32(pos)/float32(tableInfo.MaxID), cost)
		pos = pos + cfg.Export.OnceNumber
	}

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

func createTableValues(db *sql.DB, name string, minID, maxID int64, sep string) (string, error) {
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
	dataText := make([]string, 0)
	for rows.Next() {
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
				dataStrings[key], _ = jsoniter.MarshalToString(value.String)
			}
		}
		dataText = append(dataText, strings.Join(dataStrings, sep))

	}
	return strings.Join(dataText, cfg.Export.LineSep), rows.Err()
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func writeToFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func appendToFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func GetPosID(filePath string) (id int64, err error) {
	t, err := tail.TailFile(filePath, tail.Config{Follow: false})
	var lastLine = ""
	for line := range t.Lines {
		lastLine = line.Text
	}
	var maxIDStr string
	err = jsoniter.UnmarshalFromString(strings.Split(lastLine, cfg.Export.FileSep)[0], &maxIDStr)
	if err != nil {
		return 0, err
	}
	maxID, err := strconv.ParseInt(maxIDStr, 10, 64)
	if err != nil {
		return 0, err
	}
	//fmt.Println("AAAA ", maxIDStr, maxID)
	return maxID, nil
}
