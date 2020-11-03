package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func binLogWithCanal() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = ""
	cfg.Flavor = "mysql"
	cfg.User = "root"
	cfg.Password = ""
	// We only care table canal_test in test db
	cfg.Dump.TableDB = ""
	//cfg.Dump.Tables = []string{"canal_test"}

	c, err := canal.NewCanal(cfg)
	fmt.Printf("new canal: %v", err)

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.RunFrom(mysql.Position{
		Name: "mysql-bin.000001",
		Pos:  122701,
	})

}

func main() {
	binLogWithCanal()
}
