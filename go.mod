module code_tpl_go

go 1.18

//replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
//replace github.com/pingcap/tidb => github.com/pingcap/tidb v1.1.0-beta.0.20200403123430-6d02bc72d9c0
replace github.com/siddontang/go-mysql => github.com/go-mysql-org/go-mysql v1.1.2

require (
	cloud.google.com/go/pubsub v1.3.1
	github.com/BurntSushi/toml v1.2.0
	github.com/Flyingon/go-common v0.1.12
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/Shopify/sarama v1.36.0
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/astaxie/beego v1.12.3
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/buger/jsonparser v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-cmp v0.5.8
	github.com/google/wire v0.5.0
	github.com/googollee/go-socket.io v1.6.2
	github.com/gorilla/websocket v1.5.0
	github.com/json-iterator/go v1.1.12
	github.com/kcorlidy/dangerous v0.0.0-20200211105345-70577de0e5c4
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.3.1
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/siddontang/go-log v0.0.0-20190221022429-1e957dd83bed
	github.com/siddontang/go-mysql v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.36
	github.com/valyala/fasthttp v1.39.0
	github.com/xdg/scram v1.0.5
	go.uber.org/automaxprocs v1.5.1
	golang.org/x/net v0.7.0
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9
	google.golang.org/protobuf v1.28.1
	gopkg.in/olivere/elastic.v6 v6.2.37
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cloud.google.com/go v0.102.1 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	cloud.google.com/go/kms v1.4.0 // indirect
	github.com/andreburgaud/crypt2go v0.13.0 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-mysql-org/go-mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.3 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/olivere/elastic v6.2.37+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pingcap/errors v0.11.5-0.20201126102027-b0a155152ca3 // indirect
	github.com/pingcap/log v0.0.0-20210317133921-96f4fcab92a4 // indirect
	github.com/pingcap/parser v0.0.0-20210415081931-48e7f467fd74 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/oauth2 v0.0.0-20220622183110-fd043fe589d2 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/api v0.93.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220812140447-cec7f5303424 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
