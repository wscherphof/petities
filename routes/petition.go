package routes

import (
	"github.com/wscherphof/petities/routes/petition"
)

func init() {
	router.GET("/petition", petition.Petition)

	router.POST("/signature", petition.Signature)
	router.GET("/signature/post", petition.Signature)

	router.POST("/provision", petition.Provision)
}
