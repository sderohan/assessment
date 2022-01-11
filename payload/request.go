package payload

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator"
)

type TopWordRequest struct {
	Text string `json:"text"`
}

func DecodeReq(r *http.Request, to interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(to); err != nil {
		if err != io.EOF {
			return err
		}
	}
	validate := validator.New()
	return validate.Struct(to)
}
