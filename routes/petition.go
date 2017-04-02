package routes

import (
	"github.com/wscherphof/petities/routes/petition"
)

func init() {
	router.GET("/petition/:id", petition.SignatureForm)
}
