package main

import (
	"github.com/gin-gonic/gin"
	"github.com/higashi000/noa/initclient"
	"github.com/higashi000/noa/recvmsg"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	recvmsg.RecvMsg(r, m)
	initclient.InitClient(r)

	r.Run(":5000")
}
