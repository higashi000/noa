package recvmsg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/higashi000/noa/registchannel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Text   []string `json:"text"`
	Line   int      `json:"line"`
	Uuid   string   `json:"uuid"`
	RoomID string   `json:"roomid"`
}

func RecvMsg(r *gin.Engine, m *melody.Melody, channelColle *mongo.Collection) {
	var recv Msg

	findOptions := options.FindOne()

	r.POST("/send", func(c *gin.Context) {
		c.BindJSON(&recv)

		var doc registchannel.Channel

		err := channelColle.FindOne(context.Background(), bson.M{"roomid": recv.RoomID}, findOptions).Decode(&doc)
		if err != nil {
			log.Println(err)
		}

		if len(recv.Text) == len(doc.Text) {
			necessaryUpdate := false

			for i := 0; i < len(recv.Text); i++ {
				if recv.Text[i] != doc.Text[i] {
					necessaryUpdate = true
					break
				}
			}

			if !necessaryUpdate {
				c.JSON(http.StatusOK, doc)
				return
			}
		}

		update := bson.D{{"$set",
			bson.D{
				{"roomid", doc.RoomId},
				{"password", doc.Password},
				{"admin", doc.Admin},
				{"text", recv.Text},
			},
		}}
		res, err := channelColle.UpdateOne(context.Background(), doc, update)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res)

		sendJSON, _ := json.Marshal(recv)

		m.BroadcastFilter([]byte(sendJSON), func(q *melody.Session) bool {
			return q.Request.URL.Path == strings.Join([]string{"/channel", recv.RoomID, "ws"}, "/")
		})
		fmt.Println(doc.Text)
		c.JSON(http.StatusOK, doc)
	})
}
