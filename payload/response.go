package payload

import (
	"encoding/json"
	"net/http"
)

type ResponsePayload struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

func _json(w http.ResponseWriter, code int, b []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func JSON(w http.ResponseWriter, code int, i interface{}) (err error) {
	rp := &ResponsePayload{
		Data: i,
		Code: code,
	}
	b, err := json.Marshal(rp)
	if err != nil {
		return err
	}
	_json(w, code, b)
	return
}

type TopWordResponse struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}
