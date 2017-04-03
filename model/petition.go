package model

import (
	"github.com/wscherphof/entity"
	"github.com/wscherphof/msg"
	"time"
)

type Petition struct {
	*entity.Base
	Address       msg.MessageType
	Desk          msg.MessageType
	Closed        time.Time
	Answered      time.Time
	Status        string
	Websites      []*Website
	NumSignatures msg.MessageType
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

func init() {
	entity.Register(&Petition{})

	groningen := InitPetition("groningen")
	groningen.Address = msg.New().
		Set("nl", "Tweede Kamer").
		Set("en", "House of Representatives")
	groningen.Desk = msg.New().
		Set("nl", "Nederland").
		Set("en", "The Netherlands")
	const format = "2006-Jan-02"
	groningen.Closed, _ = time.Parse(format, "2017-May-08")
	groningen.Answered, _ = time.Parse(format, "2017-May-18")
	groningen.Status = "Signable"
	groningen.Websites = append(groningen.Websites,
		&Website{"http://www.laatgroningennietzakken.nl/",
			msg.New().
				Set("nl", "Campagnepagina 'Laat Groningen niet zakken'").
				Set("en", "Campain page 'Don't let Groningen down'"),
		},
		&Website{"https://www.facebook.com/permalink.php?story_fbid=1371451446239730&amp;id=225251577526395",
			msg.New().
				Set("nl", "Bericht op Facebook met de aankondiging van deze petitie").
				Set("en", "Facebook message announcing this petition"),
		},
		&Website{"https://twitter.com/petities/status/829063398361092096",
			msg.New().
				Set("nl", "Tweet met aankondiging van deze petitie").
				Set("en", "Tweet announcing this petition"),
		},
	)
	groningen.NumSignatures = msg.New().
		Set("nl", "179.538").
		Set("en", "179,538")
	groningen.Caption = msg.New().
		Set("nl", "Laat Groningen niet zakken").
		Set("en", "Don’t let Groningen down")
	groningen.Intro = msg.New().
		Set("nl", "Aardbevingen, veroorzaakt door gaswinning, hebben het leefklimaat in de provincie Groningen voor een groot deel verwoest. Schadeafhandeling stuit op spanning tussen rechtvaardig en rechtmatig. Woningen en gebouwen zijn onveilig. Gedupeerden voelen zich vaak ‘gevangen in een onveilige gevangenis’. - Freek de Jonge").
		Set("en", "Earthquakes caused by gas extraction, have destroyed a large part of the living environment in the province of Groningen. The process of claiming damages is causing tension between what is righteous and what is lawful. Homes and buildings are unsafe. And victims often feel ‘trapped in an unstable prison’. - Freek de Jonge")
	groningen.We = msg.New().
		Set("nl", "Nederlanders, en overige gebruikers van Gronings gas, solidair met de gedupeerde Groningers.").
		Set("en", "The people living in the Netherlands, and other users of gas coming from Groningen, stand by the victims of Groningen.")
	groningen.Observations = append(groningen.Observations,
		msg.New().
			Set("nl", "Dat de aardgaswinning in Groningen tot een catastrofe leidt").
			Set("en", "That gas extraction in Groningen is leading to a catastrophe"),
		msg.New().
			Set("nl", "Dat verwevenheid van belangen van overheid en bedrijfsleven en leveringsafspraken spotten met de veiligheid van de inwoners").
			Set("en", "That the interdependence of interests of government and industry and supply arrangements mock the safety of the local residents"),
		msg.New().
			Set("nl", "Dat zowel vaststelling als afhandeling van de aardbevingsschade door de NAM van de gedupeerde een machteloos, tot eindeloos afwachten gedwongen, slachtoffer maakt").
			Set("en", "That the manner in which the earthquake damage is determined and settled by the NAM leaves victims feeling powerless and forced to wait endlessly for a resolution"),
		msg.New().
			Set("nl", "Dat de consequenties van toekomstige bevingen door overheid en bestuurders onderschat worden").
			Set("en", "That the consequences of future earthquakes are grossly understated by government and directors"),
	)
	groningen.Requests = append(groningen.Requests,
		msg.New().
			Set("nl", "Afbouwplan versneld stoppen gaswinning").
			Set("en", "A speed up plan to reduce gas extraction"),
		msg.New().
			Set("nl", "'Generaal pardon' voor alle in behandeling zijnde schades").
			Set("en", "A ‘general pardon’ for all damage claims currently pending"),
		msg.New().
			Set("nl", "Daarna omkering bewijslast in een onafhankelijk schadeproces ").
			Set("en", "Then reverse burden of proof in an independent claims process"),
		msg.New().
			Set("nl", "Uitkoopregeling voor iedereen die het gebied wil verlaten ").
			Set("en", "A buy-out settlement for everyone wanting to leave the region"),
		msg.New().
			Set("nl", "Scheiding gas en staat").
			Set("en", "The separation of gas and state"),
		msg.New().
			Set("nl", "Herstel gemeenschappelijk overleg").
			Set("en", "Resume communal discussion"),
		msg.New().
			Set("nl", "Gasbaten investeren").
			Set("en", "Invest gas profits"),
		msg.New().
			Set("nl", "Deltaplan dat Groningen koploper maakt in transitie van economie en energievoorziening").
			Set("en", "Delta plans making Groningen a leader in economic and energy transition"),
		msg.New().
			Set("nl", "Internationaal kenniscentrum voor milieu in Groningen vestigen.").
			Set("en", "Set up international knowledge center for environment in Groningen"),
	)
	if err := groningen.Update(groningen); err != nil {
		panic(err)
	}
}

func InitPetition(id string) *Petition {
	return &Petition{Base: &entity.Base{ID: id}}
}
