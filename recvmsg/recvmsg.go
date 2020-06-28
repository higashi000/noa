package recvmsg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Text   []string `json:"text"`
	Line   int      `json:"line"`
	Uuid   string   `json:"uuid"`
	RoomID string   `json:"roomid"`
}

func RecvMsg(r *gin.Engine, m *melody.Melody) {
	var recv Msg

	r.POST("/send", func(c *gin.Context) {
		c.BindJSON(&recv)

		//strLine := strconv.Itoa(recv.Line)

		//returnData := `{"line": ` + strLine + `, "text":[` + recv.Text + `], "uuid":"` + recv.Uuid + `"}`

		sendJSON, _ := json.Marshal(recv)

		//		m.Broadcast([]byte(sendJSON))

		fmt.Println(string(sendJSON))
		m.BroadcastFilter([]byte(sendJSON), func(q *melody.Session) bool {
			fmt.Println(q.Request.URL.Path)
			return q.Request.URL.Path == "/channel/"+recv.RoomID+"/ws"
		})
		c.JSON(http.StatusOK, "ok")
	})
}
