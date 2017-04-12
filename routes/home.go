package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"net/http"
)

func init() {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		template.Handle("petition", "Home", "")(w, r, ps)
	})

	router.GET("/loaderio-:code", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/static/loaderio-"+ps.ByName("code")+".txt", http.StatusSeeOther)
	})
}
