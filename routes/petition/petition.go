package petition

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"net/http"
)

func Petition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.GET(w, r, "petition", "Petition")
	petition := model.InitPetition(r.FormValue("id"))
	if err, empty := petition.Read(petition); err != nil {
		template.Error(w, r, err, empty)
	} else {
		t.Set("petition", petition)
		t.Run()
	}
}
