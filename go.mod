module code_tpl_go

go 1.14

//replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
//replace github.com/pingcap/tidb => github.com/pingcap/tidb v1.1.0-beta.0.20200403123430-6d02bc72d9c0

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/Shopify/sarama v1.28.0
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/buger/jsonparser v1.1.1
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.4
	github.com/google/go-cmp v0.5.4
	github.com/jmoiron/sqlx v1.3.1 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/olivere/elastic v6.2.35+incompatible // indirect
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.5 // indirect
	github.com/pingcap/check v0.0.0-20200212061837-5e12011dc712 // indirect
	github.com/pingcap/errors v0.11.5-0.20190809092503-95897b64e011 // indirect
	github.com/pingcap/log v0.0.0-20200117041106-d28c14d3b1cd // indirect
	github.com/pingcap/parser v0.0.0-20200331080149-8dce7a46a199 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/siddontang/go-log v0.0.0-20190221022429-1e957dd83bed
	github.com/siddontang/go-mysql v1.1.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/protobuf v1.25.0
	gopkg.in/olivere/elastic.v6 v6.2.35
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
