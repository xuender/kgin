package kgin_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kgin"
)

func TestString(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	eng := kgin.Default()

	eng.GET("/test", func(ctx *gin.Context) {
		kgin.String(ctx, http.StatusOK, "ok")
	})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(nil))

	eng.ServeHTTP(recorder, req)
	res := recorder.Result()
	ass.Equal(http.StatusOK, res.StatusCode)

	tag := res.Header.Get(kgin.Etag)
	ass.NotEqual("", tag)

	recorder = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(nil))
	req.Header.Set(kgin.IfNoneMatch, tag)

	eng.ServeHTTP(recorder, req)
	ass.Equal(http.StatusNotModified, recorder.Result().StatusCode)
}
