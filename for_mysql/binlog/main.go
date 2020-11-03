package main

import (
	"code_tpl_go/for_mysql/mysql-binlog/river"
	"fmt"
	"time"
)

func main() {
	r, err := river.NewRiver(&river.Config{
		MyAddr:         "",
		MyUser:         "root",
		MyPassword:     "",
		MyCharset:      "utf8",
		ServerID:       1001,
		Flavor:         "mysql",
		DataDir:        "/Users/yuanzhaoyi/Develop/github.com/code_tpl_go/for_mysql/binlog/",
		DumpExec:       "",
		SkipMasterData: false,
		Sources:        nil,

		BulkSize:       1,
		FlushBulkTime:  river.TomlDuration{3 * time.Millisecond},
		SkipNoPkTable:  false,
		Rules:          nil,
	})
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	r.Run()
}
