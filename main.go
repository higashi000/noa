package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Text string `json:"text"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.POST("/send", func(c *gin.Context) {
		var tmp Msg
		c.BindJSON(&tmp)

		m.Broadcast([]byte(tmp.Name + ": " + tmp.Text))
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":5000")
}
