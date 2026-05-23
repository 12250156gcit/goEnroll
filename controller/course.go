package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var c model.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.Create(); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"ststus" : "Course Added"})
}



func GetCourse (w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	c := model.Course{Cid: cid}
	if err := c.Read(); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, c)
}



func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCid := mux.Vars(r)["cid"]
	var c model.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.Update(oldCid); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, c)
}

func DeleteCourse (w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	c := model.Course{Cid: cid}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"status":"Course deleted"})
}

func GetAllCourses (w http.ResponseWriter, r *http.Request) {
	courses, err := model.GetAllCourse()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, courses)
}
