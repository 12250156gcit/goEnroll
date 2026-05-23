package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"myapp/utils/httpResp/date"
	"net/http"
	"strings"
)

func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Println(e)
	current_date := date.Getdate()
	e.Date_Enrolled = current_date

	// pass e to model
	saveErr := e.EnrollStud()
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, saveErr.Error())
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"ststus": "student enrolled success"})
	}
}
