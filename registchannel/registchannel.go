package registchannel

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Channel struct {
	RoomId   string `json:"roomid" bson:"roomid"`
	Password string `json:"password" bson:"password"`
	Admin    string `json:"admin" bson:"admin"`
}

func RegistChannel(r *gin.Engine, db *mongo.Database) {
	channelColl := db.Collection("channel")

	findOptions := options.Find()
	cur, err := channelColl.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println(err)
	}

	r.POST("noa/registchannel", func(c *gin.Context) {
		var recvChannelData Channel
		c.BindJSON(&recvChannelData)

		for cur.Next(context.Background()) {
			var tmp Channel

			cur.Decode(&tmp)

			if recvChannelData.RoomId == tmp.RoomId {
				c.JSON(http.StatusOK, `{"status": "false", "message":"This ID is already used."}`)
				return
			}
		}

		res, err := channelColl.InsertOne(context.Background(), recvChannelData)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, `{"status": "true", "message":"Correct"}`)
		log.Println(res)
	})
}
