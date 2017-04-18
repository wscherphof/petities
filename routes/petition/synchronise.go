package petition

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/model"
	"log"
	"net/http"
)

func Synchronise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if t := template.PRG(w, r, "petition", "Synchronise"); t == nil {
		return
	} else {
		petition := model.InitPetition()
		status := ""
		if cursor, err := petition.ReadAll(petition); err == nil {
			defer cursor.Close()
			for cursor.Next(petition) {
				go func() {
					if err := petition.Synchronise(); err != nil {
						log.Println("ERROR: Synchronise - petition.Synchronise", err)
					}
				}()
				status = fmt.Sprintf("%s%s: %d, ", status, petition.ID, petition.NumSignatures)
			}
			if cursor.Err() != nil {
				log.Println("ERROR: Synchronise - cursor.Err()", cursor.Err())
			}
		}
		t.Set("status", status)
		t.Run()
	}
}
