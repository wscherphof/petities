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
	if err, empty := groningen.Read(groningen); err != nil && !empty {
		panic(err)
	}
	groningen.Caption.
		Set("nl", "Laat Groningen niet zakken").
		Set("en", "Don’t let Groningen down")
	if err := groningen.Update(groningen); err != nil {
		panic(err)
	}
}

func InitPetition(id string) *Petition {
	return &Petition{
		Base:    &entity.Base{ID: id},
		Caption: make(msg.MessageType, 2),
	}
}
