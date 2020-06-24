package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v6"
)

var esUrl, esUser, esSecret, esHost string
var EsClt *elastic.Client

func EsCltInit() {
	esHost = fmt.Sprintf("http://%s:%s@%s", esUser, esSecret, esUrl)
	EsClt = cltInit()
}

func cltInit() *elastic.Client {
	fmt.Printf("es host: %s", esUrl)

	// elastic.SetSniff(false)不设置嗅探器，嗅探器用于自动发现Elasticsearch的节点，并设置为restClient实例
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esUrl),
		elastic.SetBasicAuth(esUser, esSecret))
	if err != nil {
		msg := fmt.Sprintf("esClient init failed, err: %v", err)
		fmt.Printf(msg, err)
		return nil
	}

	return client
}

func main() {
	esUrl = "http://es.sh.sh-global2.db:9200"
	esUser = "txhdxs"
	esSecret = "txhd%5ex23s%5ego0"
	EsCltInit()
}
