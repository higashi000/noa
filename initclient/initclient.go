package initclient

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeClientUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}

	return u.String()
}

func InitClient(r *gin.Engine) {
	r.GET("/init", func(c *gin.Context) {
		uuid := MakeClientUUID()

		c.String(http.StatusOK, uuid)
	})
}
