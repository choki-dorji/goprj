package contolller

import (
	"encoding/json"
	"fmt"
	"myapp/myapp/model"
	"myapp/myapp/utils/date"
	"myapp/myapp/utils/httpResp"
	"strconv"

	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	e.Date_Enrolled = date.GetDate()

	if err := e.EnrollStud(); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, "Duplicate keys")
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled"})
}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stdID, _ := strconv.ParseInt(vars["sid"], 10, 64)
	cid := vars["cid"]
	fmt.Println("stdID", stdID)
	fmt.Println("cid", cid)

	e, err := model.GetEnroll(stdID, cid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Enrollment not found")
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, e)
}

func GetAllEnrolls(w http.ResponseWriter, r *http.Request) {
	enrolls, err := model.GetAllEnrolls()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, enrolls)
}

func DeleteEnroll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stdID, _ := strconv.ParseInt(vars["sid"], 10, 64)
	cid := vars["cid"]

	if err := model.DeleteEnroll(stdID, cid); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Delete failed")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func UpdateEnroll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oldStdID, _ := strconv.ParseInt(vars["sid"], 10, 64)

	var e model.Enroll
	json.NewDecoder(r.Body).Decode(&e)
	defer r.Body.Close()

	e.Date_Enrolled = date.GetDate()

	if err := e.UpdateEnroll(oldStdID); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, e)
}
