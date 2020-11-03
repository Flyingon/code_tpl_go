package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/xdg/scram"
	"hash"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var SHA256 scram.HashGeneratorFcn = func() hash.Hash { return sha256.New() }
var SHA512 scram.HashGeneratorFcn = func() hash.Hash { return sha512.New() }

var (
	user        string
	password    string
	brokers     string
	group       string
	topics      string
	sasl        string
	initialFlag int64
)

func KafkaConsumer() {
	user := user
	password := password
	brokers := strings.Split(brokers, ",")
	group := group
	topics := strings.Split(topics, ",")
	config := cluster.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.User = user
	config.Net.SASL.Password = password
	config.Net.SASL.Handshake = true
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	config.Net.SASL.Mechanism = sarama.SASLMechanism(sasl)
	config.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
	}
	config.Net.TLS.Enable = false

	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = time.Second
	config.ClientID = "test-client"
	config.Consumer.Offsets.Initial = initialFlag
	config.Version = sarama.V2_1_0_0
	consumer, err := cluster.NewConsumer(brokers, group, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	// Create signal channel
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	// Consume all channels, wait for signal to exit
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				fmt.Printf("Headers: %d\n", len(msg.Headers))
				for _, v := range msg.Headers {
					fmt.Printf("Headers: %s:%s\n", v.Key, v.Value)
				}
				fmt.Printf("Message: %s/%d/%d\t%s: %s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "")
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				fmt.Printf("Rebalanced: %+v\n", ntf)
			}
		case err, more := <-consumer.Errors():
			if more {
				fmt.Printf("Error: %s\n", err.Error())
			}
		case <-sigchan:
			return
		}
	}
}

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}

func main() {
	flag.StringVar(&user, "user", "story1001", "user")
	flag.StringVar(&password, "password", "123456", "password")
	flag.StringVar(&brokers, "brokers", "9.134.110.176:9092", "brokers")
	flag.StringVar(&group, "group", "test-consumer-group", "group")
	flag.StringVar(&topics, "topics", "tendata-story1001-431-wuji-test", "topics")
	flag.StringVar(&sasl, "sasl", "PLAIN", "SASL mechanism: PLAIN,SCRAM-SHA-512,GSSAPI,SCRAM-SHA-256")
	flag.Int64Var(&initialFlag, "init", -1, "initial flag: -1, -2")
	flag.Parse()
	fmt.Printf("----- user: %v, password: %v, brokers: %v, group: %v, topics: %v, sasl: %v, initialFlag: %v-----\n",
		user, password, brokers, group, topics, sasl, initialFlag)

	KafkaConsumer()
}
