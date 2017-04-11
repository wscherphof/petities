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

type Counter struct {
	*entity.Base
	Num int
}

func InitCounter(opt_petition ...string) (counter *Counter) {
	counter = &Counter{Base: &entity.Base{}}
	if len(opt_petition) == 1 {
		counter.ID = opt_petition[0]
	}
	return
}

func (p *Petition) Sign(name, email, city string) (err error, conflict bool) {
	signature := InitSignature(p.ID, email)
	signature.Name = name
	signature.City = city
	if err, conflict = signature.Create(signature); err == nil {
		counter := InitCounter(p.ID)
		if e, empty := counter.Read(counter); e != nil && !empty {
			err = e
		} else {
			counter.Num++
			err = counter.Update(counter)
		}
	}
	return
}

func (p *Petition) Synchronise() {
	counter := InitCounter(p.ID)
	signature := InitSignature("", "")
	counter.Read(counter)
	signature.Index(signature, "Petition").Count(p.ID,
		&counter.Num,
	)
	p.NumSignatures = counter.Num
	if err := counter.Update(counter); err != nil {
		log.Println("ERROR:", err)
	} else if err := p.Update(p); err != nil {
		log.Println("ERROR:", err)
	}
	return
}

func init() {
	entity.Register(&Petition{})
	entity.Register(&Signature{}).Index("Petition").Index("Created")
	entity.Register(&Counter{})

	// Periodically update NumSignatures for each petition.
	go func() {
		for {
			counter := InitCounter()
			if cursor, err := counter.ReadAll(counter); err != nil {
				log.Println("ERROR:", err)
			} else {
				defer cursor.Close()
				for cursor.Next(counter) {
					petition := InitPetition(counter.ID)
					if err, _ := petition.Read(petition); err != nil {
						log.Println("ERROR:", err)
					} else {
						petition.NumSignatures = counter.Num
						if err := petition.Update(petition); err != nil {
							log.Println("ERROR:", err)
						}
					}
				}
				if cursor.Err() != nil {
					log.Println("ERROR:", cursor.Err())
				}
			}
			time.Sleep(60 * time.Second)
		}
	}()
}
