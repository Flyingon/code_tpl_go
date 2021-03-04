package main

import (
	"net/url"

	"github.com/astaxie/beego"

	"golang.org/x/net/websocket"
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

func (this *Client) SendMessage(body []byte) error {
	u := url.URL{Scheme: "ws", Host: this.Host, Path: this.Path}

	ws, err := websocket.Dial(u.String(), "", "http://"+this.Host+"/")
	defer ws.Close() //关闭连接
	if err != nil {
		beego.Error(err)
		return err
	}

	_, err = ws.Write(body)
	if err != nil {
		beego.Error(err)
		return err
	}

	return nil
}

func main() {
	clt := Client{
		Host: "9.134.110.176:8081",
		Path: "",
	}
	clt.SendMessage([]byte("我艹"))
}