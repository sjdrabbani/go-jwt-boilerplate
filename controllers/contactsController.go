package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gocontactmanager/models"
	u "github.com/gocontactmanager/utils"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}
