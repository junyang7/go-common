package _render

import (
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_response"
	"net/http"
)

func JSON(w http.ResponseWriter, res *_response.Response) {
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(_json.Encode(res))
}
