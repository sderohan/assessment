package controller

import (
	"net/http"
	"strings"

	"github.com/sderohan/assessment.git/payload"
	"github.com/sderohan/assessment.git/service"
)

func GetTopTenWords(w http.ResponseWriter, r *http.Request) {
	req := new(payload.TopWordRequest)
	// decode the request and get the data
	if err := payload.DecodeReq(r, req); err != nil {
		payload.JSON(w, 400, err)
		return
	}

	// remove spaces and check for empty
	req.Text = strings.Trim(req.Text, " ")
	if req.Text == "" {
		payload.JSON(w, 400, "Request body is empty!")
	}

	// call the service for performing word count
	data, err := service.GetTopTenWords(req.Text)
	if err != nil {
		payload.JSON(w, 400, err)
		return
	}

	// return the word counts
	payload.JSON(w, 200, data)
}
