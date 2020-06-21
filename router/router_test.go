package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInitClient(t *testing.T) {
	r := NewRouter()

	req := httptest.NewRequest("GET", "/init", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestRecvMsg(t *testing.T) {

}
