package controller

import (
	"net/http"
	"strings"

	"github.com/sderohan/assessment.git/payload"
	"github.com/sderohan/assessment.git/service"
)

func GetTopTenWords(w http.ResponseWriter, r *http.Request) {
	req := new(payload.TopWordRequest)
	if err := payload.DecodeReq(r, req); err != nil {
		payload.JSON(w, 400, err)
		return
	}

	req.Text = strings.Trim(req.Text, " ")
	if req.Text == "" {
		payload.JSON(w, 400, "Request body is empty!")
	}

	data, err := service.GetTopTenWords(req.Text)
	if err != nil {
		payload.JSON(w, 400, err)
		return
	}
	payload.JSON(w, 200, data)
}
