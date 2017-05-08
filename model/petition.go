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
	petition  string
	Name      string
	City      string
	Visible   bool
	Token     string `gorethink:",omitempty"`
	Confirmed bool
}

func InitSignature(petition, email string) *Signature {
	return &Signature{
		Base: &entity.Base{
			Table: petition + "_Signature",
			ID:    email,
		},
		Token:    util.NewToken(),
		petition: petition,
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
		newSignatures[s.petition]++
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
					log.Println("ERROR:", err)
				} else if err := petition.newSignatures(petition.NumSignatures + newSignatures[k]); err != nil {
					log.Println("ERROR:", err)
				}
			}
		}
	}()
}
