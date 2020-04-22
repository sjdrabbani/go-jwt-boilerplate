package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gocontactmanager/models"
	u "github.com/gocontactmanager/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResonseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w.u.Message(false, "Invalid Request"))
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
