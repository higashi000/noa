package recvmsg

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Text   string `json:"text"`
	Line   int    `json:"line"`
	Uuid   string `json:"uuid"`
	RoomID string `json:"roomid"`
}

func RecvMsg(r *gin.Engine, m *melody.Melody) {
	var recv Msg

	r.POST("/send", func(c *gin.Context) {
		c.BindJSON(&recv)

		strLine := strconv.Itoa(recv.Line)

		returnData := `{"line": ` + strLine + `, "text":"` + recv.Text + `", "uuid":"` + recv.Uuid + `"}`

		fmt.Println(returnData)

		m.Broadcast([]byte(returnData))

		m.BroadcastFilter([]byte(returnData), func(q *melody.Session) bool {
			return q.Request.URL.Path == recv.RoomID
		})
		c.JSON(http.StatusOK, "ok")
	})
}
