package main

import (
	essix "github.com/wscherphof/essix/server"
	"github.com/wscherphof/petities/messages"
	"github.com/wscherphof/petities/routes"
)

func init() {
	messages.Init()
	routes.Init()
}

func main() {
	essix.Run()
}
