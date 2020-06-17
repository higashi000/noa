package main

import (
	"github.com/gin-gonic/gin"
	"github.com/higashi000/noa/initclient"
	"github.com/higashi000/noa/recvmsg"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Text string `json:"text"`
	Line int    `json:"line"`
	Uuid string `json:"uuid"`
}

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	recvmsg.RecvMsg(r, m)
	initclient.InitClient(r)

	r.Run(":5000")
}
