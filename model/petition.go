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
	return &Signature{
		Base:     &entity.Base{ID: petition + "|" + email},
		Petition: petition,
		Email:    email,
	}
}

var newSignatures = make(map[string]int, 100)

func (p *Petition) Sign(name, email, city string) (err error, conflict bool) {
	signature := InitSignature(p.ID, email)
	signature.Name = name
	signature.City = city
	if err, conflict = signature.Create(signature); err == nil {
		newSignatures[p.ID]++
	}
	return
}

func (p *Petition) newSignatures(num int) (err error) {
	var spare int
	spare, p.NumSignatures, newSignatures[p.ID] = newSignatures[p.ID], num, 0
	if err = p.Update(p); err != nil {
		newSignatures[p.ID] += spare
	}
	return
}

func (p *Petition) Synchronise() (err error) {
	var count int
	signature := InitSignature("", "")
	if err = signature.Index(signature, "Petition").Count(p.ID, &count); err == nil {
		err = p.newSignatures(count)
	}
	return
}

func init() {
	entity.Register(&Petition{})
	entity.Register(&Signature{}).Index("Petition").Index("Created")

	// Periodically update NumSignatures for each petition.
	go func() {
		for {
			time.Sleep(30 * time.Second)
			for k, v := range newSignatures {
				if v < 1 {
					continue
				}
				petition := InitPetition(k)
				if err, _ := petition.Read(petition); err != nil {
					log.Println("ERROR:", err)
				} else if err := petition.newSignatures(petition.NumSignatures + newSignatures[k]); err != nil {
					log.Println("ERROR:", err)
				}
			}
		}
	}()
}
