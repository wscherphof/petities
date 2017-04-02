package petition

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"net/http"
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
