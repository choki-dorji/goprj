package contolller

import (
	"encoding/json"
	"myapp/myapp/model"
	"myapp/myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json Body")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Admin Added"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json Body")
		return
	}
	defer r.Body.Close()
	getErr := admin.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}

	cookie := http.Cookie{
		Name:    "my-name",
		Value:   "my-value",
		Expires: time.Now().Add(24 * time.Hour),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)

	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "success"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-name",
		Expires: time.Now(),
	})
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}
