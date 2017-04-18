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
	ErrAlreadyAcknowledged = errors.New("ErrAlreadyAcknowledged")
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
	Petition         string
	Name             string
	Email            string
	City             string
	Visible          bool
	AcknowledgeToken string
}

func InitSignature(petition, email string) *Signature {
	return &Signature{
		Base:             &entity.Base{ID: petition + "|" + email},
		Petition:         petition,
		Email:            email,
		AcknowledgeToken: util.NewToken(),
	}
}

var newSignatures = make(map[string]int, 100)

func (s *Signature) Acknowledge(ack string) (err error, conflict bool) {
	if s.AcknowledgeToken == "" {
		return ErrAlreadyAcknowledged, true
	}
	if s.AcknowledgeToken != ack {
		return essix.ErrInvalidCredentials, true
	}
	s.AcknowledgeToken = ""
	if err = s.Update(s); err == nil {
		newSignatures[s.Petition]++
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
	index := signature.Index(signature, "Petition+AcknowledgeToken")
	if err = index.Count(&count, p.ID, ""); err == nil {
		err = p.newSignatures(count)
	}
	return
}

func init() {
	entity.Register(&Petition{})
	entity.Register(&Signature{}).
		Index("Created").
		Index("Petition+AcknowledgeToken", "Petition", "AcknowledgeToken")

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
