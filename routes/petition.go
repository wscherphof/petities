package routes

import (
	"github.com/wscherphof/essix/template"
	"github.com/wscherphof/petities/routes/petition"
)

func init() {
	router.GET("/petition", petition.Petition)

	router.POST("/signature", petition.Signature)
	router.GET("/signature/post", petition.Signature)

	router.GET("/signature/ack", petition.AcknowledgeForm)
	router.PUT("/signature/ack", petition.Acknowledge)
	router.GET("/signature/ack/put", petition.Acknowledge)

	router.GET("/provision", template.Handle("petition", "ProvisionForm", ""))
	router.POST("/provision", petition.Provision)
	router.GET("/provision/post", petition.Provision)

	router.GET("/synchronise", template.Handle("petition", "SynchroniseForm", ""))
	router.POST("/synchronise", petition.Synchronise)
	router.GET("/synchronise/post", petition.Synchronise)
}
