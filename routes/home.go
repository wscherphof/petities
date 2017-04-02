package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/petition/groningen", http.StatusSeeOther)
	})
}
