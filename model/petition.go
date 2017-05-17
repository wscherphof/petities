package model

import (
	"errors"
	"github.com/wscherphof/entity"
	essix "github.com/wscherphof/essix/model"
	"github.com/wscherphof/essix/util"
	"github.com/wscherphof/msg"
	"log"
	"time"
)

var (
	ErrAlreadyConfirmed = errors.New("ErrAlreadyConfirmed")
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
	Name      string
	City      string
	Visible   bool
	Token     string
	Confirmed bool
}

func InitSignature(petition, email string) *Signature {
	return &Signature{
		Base: &entity.Base{
			Table: petition,
			ID:    email,
		},
		Token: util.NewToken(),
	}
}

func (s *Signature) Register() {
	entity.Register(s).
		Index("Created").
		Index("Confirmed")
}

var newSignatures = make(map[string]int, 100)

func (s *Signature) Confirm(token string) (err error, conflict bool) {
	if s.Confirmed {
		return ErrAlreadyConfirmed, true
	}
	if s.Token != token {
		return essix.ErrInvalidCredentials, true
	}
	s.Token, s.Confirmed = "", true
	if err = s.Update(s); err == nil {
		newSignatures[s.Table]++
	}
	return
}

func (p *Petition) Synchronise() (err error) {
	signature := InitSignature(p.ID, "")
	index := signature.Index(signature, "Confirmed")
	if err = index.Count(&(p.NumSignatures), true); err == nil {
		err = p.Update(p)
	}
	return
}

func init() {
	entity.Register(InitPetition())

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
					log.Printf("ERROR: newSignatures %v", err)
				} else {
					num := newSignatures[k]
					petition.NumSignatures += num
					if err := petition.Update(petition); err == nil {
						newSignatures[k] -= num
					} else {
						petition.NumSignatures -= num
						log.Printf("ERROR: newSignatures %v", err)
					}
				}
			}
		}
	}()
}
