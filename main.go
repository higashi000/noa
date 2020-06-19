package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/higashi000/noa/initclient"
	"github.com/higashi000/noa/recvmsg"
	"github.com/higashi000/noa/registchannel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println(err)
	}

	db := client.Database("noa")

	r := gin.Default()
	m := melody.New()

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	registchannel.RegistChannel(r, db)
	recvmsg.RecvMsg(r, m)
	initclient.InitClient(r)

	r.Run(":5000")
}
