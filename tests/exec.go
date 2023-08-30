package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/samber/lo"
)

func Exec(router http.Handler, method, path string, elems ...any) *http.Response {
	var data []byte

	switch len(elems) {
	case 0:
		data = nil
	case 1:
		data = lo.Must1(json.Marshal(elems[0]))
	default:
		data = lo.Must1(json.Marshal(elems))
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder.Result()
}
