package petition

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"net/http"
	"strconv"
)

func SignatureForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	t := template.GET(w, r, "petition", "SignatureForm")
	petition := model.InitPetition(id)
	if err, empty := petition.Read(petition); err != nil && !empty {
		template.Error(w, r, err, false)
	} else {
		t.Set("petition", petition)
		t.Run()
	}
}

func Signature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if t := template.PRG(w, r, "petition", "Signature"); t == nil {
		return
	} else {
		name, email, city := r.FormValue("name"), r.FormValue("email"), r.FormValue("city")
		petition := model.InitPetition(r.FormValue("petition"))
		if err, _ := petition.Read(petition); err != nil {
			template.Error(w, r, err, false)
		} else if err, conflict := petition.Sign(name, email, city); err != nil {
			template.Error(w, r, err, conflict)
		} else {
			t.Set("petition", petition.ID)
			t.Set("num", strconv.Itoa(petition.NumSignatures()))
			t.Set("name", name)
			t.Set("email", email)
			t.Set("city", city)
			t.Run()
		}
	}
}
