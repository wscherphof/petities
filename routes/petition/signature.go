package petition

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/env"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"net/http"
)

var test = (env.Get("GO_ENV", "") == "test")

func SignatureForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.GET(w, r, "petition", "SignatureForm", "SignatureForm-form")
	petition := model.InitPetition(r.FormValue("petition"))
	if err, _ := petition.Read(petition); err != nil {
		template.Error(w, r, err, false)
	} else {
		t.Set("petitionID", petition.ID)
		t.Run()
	}
}

func Signature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if t := template.PRG(w, r, "petition", "Signature", "lang"); t == nil {
		return
	} else {
		name, email, city, visible := r.FormValue("name"), r.FormValue("email"), r.FormValue("city"), r.FormValue("visible")
		petition := model.InitPetition(r.FormValue("petition"))
		if err, _ := petition.Read(petition); err != nil {
			template.Error(w, r, err, false)
		} else {
			signature := model.InitSignature(petition.ID, email)
			signature.Name = name
			signature.City = city
			signature.Visible = (visible == "on")
			if err, conflict := signature.Create(signature); err != nil {
				template.Error(w, r, err, conflict)
			} else {
				emailMessage := template.Email(r, "petition", "Confirm-email", "lang")
				link := "https://" + r.Host + "/signature/confirm"
				link += "?confirm=" + signature.Token
				link += "&petition=" + petition.ID
				link += "&email=" + email
				emailMessage.Set("link", link)
				emailMessage.Run(email, "Please acknowledge")
				t.Set("email", email)
				if test {
					// for load testing purposes, this REALLY should only be in the email
					t.Set("confirm", signature.Token)
					t.Set("email", signature.ID)
				} else {
					t.Set("confirm", "")
					t.Set("email", "")
				}
				t.Run()
			}
		}
	}
}

func ConfirmForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.GET(w, r, "petition", "ConfirmForm")
	petition := model.InitPetition(r.FormValue("petition"))
	if err, _ := petition.Read(petition); err != nil {
		template.Error(w, r, err, false)
	} else {
		t.Set("petition", petition.ID)
		t.Set("caption", petition.Caption)
		t.Set("email", r.FormValue("email"))
		t.Set("confirm", r.FormValue("confirm"))
		t.Run()
	}
}

func Confirm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if t := template.PRG(w, r, "petition", "Confirm"); t == nil {
		return
	} else {
		petition, email, confirm := r.FormValue("petition"), r.FormValue("email"), r.FormValue("confirm")
		signature := model.InitSignature(petition, email)
		if err, empty := signature.Read(signature); err != nil {
			template.Error(w, r, err, empty)
		} else if err, conflict := signature.Confirm(confirm); err != nil {
			template.Error(w, r, err, conflict)
		} else {
			t.Run()
		}
	}
}
