package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	// create a variable of type student to store student information send by client
	var stud model.Student
	//extract the data from the request body send by client
	jsonObj := json.NewDecoder(r.Body)
	//store json data in the stud variable, converting json data to struct, Go object
	err := jsonObj.Decode(&stud)
	if err != nil {
		//sending the response back to client
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	//no error
	//call model and pass student info
	dbErr := stud.Create()
	if dbErr != nil {
		//sending the response back to client
		httpResp.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}
	//no err
	httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"Status": "Student added successfully"})
}

// helper function to convert string to int
func getUserId(userId string) (int64, error) {
	intID, err := strconv.ParseInt(userId, 10, 64)
	return intID, err
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	myMap := mux.Vars(r)
	stdid := myMap["sid"]
	stdID, idErr := getUserId(stdid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// no error,
	var sDetail model.Student
	sDetail = model.Student{StdId: stdID}
	// pass student data to model
	getErr := sDetail.Read()

	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return

	}
	httpResp.RespondWithJson(w, http.StatusOK, sDetail)
}

// handler function to update student details.
func UpdateStud(w http.ResponseWriter, r *http.Request) {
	old_sid := mux.Vars(r)["sid"]
	old_stdID, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	jsonObj := json.NewDecoder(r.Body)
	//store json data in the stud variable, converting json data to struct, Go object
	err := jsonObj.Decode(&stud)
	if err != nil {
		//sending the response back to client
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	updateErr := stud.Update(old_stdID)

	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, stud)
	}
}

// HANDLER FUNCTION TO DELETE A SINGLE STUDENT
func DeleteStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdID, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdID}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadGateway, err.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"ststus": "deleted"})

}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, students)
}


// test not passed need to check later
