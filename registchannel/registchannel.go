package registchannel

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Channel struct {
	RoomId   string   `json:"roomid" bson:"roomid"`
	Password string   `json:"password" bson:"password"`
	Admin    string   `json:"admin" bson:"admin"`
	FileType string   `json:"filetype" bson:"filetype"`
	Text     []string `json:"text" bson:"text"`
}

func RegistChannel(r *gin.Engine, channelColle *mongo.Collection) error {

	findOptions := options.Find()
	cur, err := channelColle.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return errors.Wrap(err, "failed find collection")
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

		res, err := channelColle.InsertOne(context.Background(), recvChannelData)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, `{"status": "true", "message":"Correct"}`)
		log.Println(res)
	})

	return nil
}
