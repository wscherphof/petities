package model

import (
	"github.com/wscherphof/entity"
	"github.com/wscherphof/msg"
	"log"
	"time"
)

type Petition struct {
	*entity.Base
	Address       msg.MessageType
	Desk          msg.MessageType
	Closed        time.Time
	Answered      time.Time
	Status        string
	NumSignatures int
	Websites      []*Website
	Caption       msg.MessageType
	Intro         msg.MessageType
	We            msg.MessageType
	Observations  []msg.MessageType
	Requests      []msg.MessageType
}

type Website struct {
	URL     string
	Caption msg.MessageType
}

func InitPetition(opt_id ...string) (petition *Petition) {
	petition = &Petition{Base: &entity.Base{}}
	if len(opt_id) == 1 {
		petition.ID = opt_id[0]
	}
	return
}

type Signature struct {
	*entity.Base
	Petition string
	Name     string
	Email    string
	City     string
}

func InitSignature(petition, email string) *Signature {
	return &Signature{Base: &entity.Base{ID: petition + "|" + email}}
}

func (p *Petition) Sign(name, email, city string) (err error, conflict bool) {
	signature := InitSignature(p.ID, email)
	signature.Petition = p.ID
	signature.Name = name
	signature.Email = email
	signature.City = city
	err, conflict = signature.Create(signature)
	return
}

func init() {
	entity.Register(&Petition{})
	entity.Register(&Signature{}).Index("Petition")

	// Periodically update NumSignatures for each petition.
	go func() {
		for {
			time.Sleep(15 * time.Second)
			if cursor, err := entity.ReadAll(&Petition{}); err == nil {
				defer cursor.Close()
				petition := InitPetition()
				for cursor.Next(petition) {
					entity.Index(&Signature{}, "Petition").Count(petition.ID,
						&petition.NumSignatures,
					)
					if err := petition.Update(petition); err != nil {
						log.Println("ERROR:", err)
					}
				}
				if cursor.Err() != nil {
					log.Println("ERROR:", cursor.Err())
				}
			}
		}
	}()
}
