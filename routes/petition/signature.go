package petition

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"net/http"
)

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
				emailMessage := template.Email(r, "petition", "Acknowledge-email", "lang")
				link := "https://" + r.Host + "/signature/ack"
				link += "?ack=" + signature.AcknowledgeToken
				link += "&petition=" + petition.ID
				link += "&email=" + email
				emailMessage.Set("link", link)
				emailMessage.Run(email, "Please acknowledge")
				t.Set("email", email)
				t.Run()
			}
		}
	}
}

func AcknowledgeForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.GET(w, r, "petition", "AcknowledgeForm")
	petition := model.InitPetition(r.FormValue("petition"))
	if err, _ := petition.Read(petition); err != nil {
		template.Error(w, r, err, false)
	} else {
		t.Set("petition", petition.ID)
		t.Set("caption", petition.Caption)
		t.Set("email", r.FormValue("email"))
		t.Set("ack", r.FormValue("ack"))
		t.Run()
	}
}

func Acknowledge(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if t := template.PRG(w, r, "petition", "Acknowledge"); t == nil {
		return
	} else {
		petition, email, ack := r.FormValue("petition"), r.FormValue("email"), r.FormValue("ack")
		signature := model.InitSignature(petition, email)
		if err, empty := signature.Read(signature); err != nil {
			template.Error(w, r, err, empty)
		} else if err, conflict := signature.Acknowledge(ack); err != nil {
			template.Error(w, r, err, conflict)
		} else {
			t.Run()
		}
	}
}
