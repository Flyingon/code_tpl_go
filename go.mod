module code_tpl_go

go 1.14

//replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
//replace github.com/pingcap/tidb => github.com/pingcap/tidb v1.1.0-beta.0.20200403123430-6d02bc72d9c0
replace github.com/siddontang/go-mysql => github.com/go-mysql-org/go-mysql v1.1.2

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/Shopify/sarama v1.29.0
	github.com/astaxie/beego v1.12.3
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/buger/jsonparser v1.1.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-mysql-org/go-mysql v1.2.0 // indirect
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-cmp v0.5.5
	github.com/google/wire v0.5.0
	github.com/googollee/go-socket.io v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/json-iterator/go v1.1.11
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/olivere/elastic v6.2.35+incompatible // indirect
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/siddontang/go-log v0.0.0-20190221022429-1e957dd83bed
	github.com/siddontang/go-mysql v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.173
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims v0.0.0-20210609005011-d5c8b6048211
	github.com/tencentyun/cos-go-sdk-v5 v0.7.25
	github.com/tencentyun/qcloud-cos-sts-sdk v0.0.0-20210601063555-b0ef43159af8
	github.com/xdg/scram v1.0.3
	go.uber.org/automaxprocs v1.4.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/protobuf v1.26.0
	gopkg.in/olivere/elastic.v6 v6.2.35
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)
