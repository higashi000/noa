package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitClient(t *testing.T) {
	r := NewRouter()

	req := httptest.NewRequest("GET", "/init", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestRecvMsg(t *testing.T) {
	router := NewRouter()

	msgData := `
   {"roomid":"testdata","text":["test1","test2","test3"],"password":"testdata","admin":""}
   `

	req := httptest.NewRequest("POST", "/send", strings.NewReader(msgData))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, msgData, rec.Body.String())
}
