package initclient

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/higashi000/noa/registchannel"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InitialData struct {
	Uuid     string   `json:"uuid"`
	FileType string   `json:"filetype"`
	Text     []string `json:"text"`
}

func MakeClientUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "Failed generate uuid")
	}

	return u.String(), nil
}

func InitClient(r *gin.Engine, channelColle *mongo.Collection) {
	findOptions := options.FindOne()

	r.GET("/init", func(c *gin.Context) {
		uuid, err := MakeClientUUID()

		if err != nil {
			fmt.Println(err)
			uuid = ""
		}

		var doc registchannel.Channel

		roomid := c.Query("roomid")

		err = channelColle.FindOne(context.Background(), bson.M{"roomid": roomid}, findOptions).Decode(&doc)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(roomid)

		initData := InitialData{
			Uuid:     uuid,
			FileType: doc.FileType,
			Text:     doc.Text,
		}

		c.JSON(http.StatusOK, initData)
	})
}
