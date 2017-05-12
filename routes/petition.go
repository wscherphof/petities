package routes

import (
	"github.com/wscherphof/env"
	"github.com/wscherphof/essix/ratelimit"
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/routes/petition"
)

var test = (env.Get("GO_ENV", "") == "test")

func init() {
	router.GET("/petition", petition.Petition)

	router.GET("/signature", petition.SignatureForm)
	router.POST("/signature", ratelimit.Handle(petition.Signature))
	router.GET("/signature/post", petition.Signature)

	router.GET("/signature/confirm", petition.ConfirmForm)
	router.PUT("/signature/confirm", petition.Confirm)
	router.GET("/signature/confirm/put", petition.Confirm)

	router.GET("/synchronise", template.Handle("petition", "SynchroniseForm", ""))
	router.POST("/synchronise", petition.Synchronise)
	router.GET("/synchronise/post", petition.Synchronise)

	if test {
		router.GET("/provision", template.Handle("petition", "ProvisionForm", ""))
		router.POST("/provision", petition.Provision)
		router.GET("/provision/post", petition.Provision)
	}
}
