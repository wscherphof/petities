package model

import (
	"github.com/wscherphof/entity"
	"github.com/wscherphof/msg"
)

type Petition struct {
	*entity.Base
	Caption msg.MessageType
}

func init() {
	entity.Register(&Petition{})
	groningen := InitPetition("groningen")
	groningen.Caption = make(msg.MessageType, 2)
	groningen.Caption.Set("nl", "Laat Groningen niet zakken")
	groningen.Caption.Set("en", "Don’t let Groningen down")
	if err := groningen.Update(groningen); err != nil {
		panic(err)
	}
}

func InitPetition(id string) (petition *Petition) {
	petition = &Petition{Base: &entity.Base{ID: id}}
	return
}
