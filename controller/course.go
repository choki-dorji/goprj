package contolller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"

	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var c model.Course
	json.NewDecoder(r.Body).Decode(&c)
	defer r.Body.Close()
	if err := c.Add(); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, c)
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := model.GetAllCourses()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]
	c, err := model.GetCourse(cid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]

	var c model.Course
	json.NewDecoder(r.Body).Decode(&c)
	defer r.Body.Close()
	c.Cid = cid
	if err := c.Update(); err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]
	if err := model.DeleteCourse(cid); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
