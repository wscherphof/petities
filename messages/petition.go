package messages

import (
	msg "github.com/wscherphof/essix/messages"
)

func init() {
	msg.Key("signatures").
		Set("nl", "ondertekeningen").
		Set("en", "signatures")

	msg.Key("Petition").
		Set("nl", "Petitie").
		Set("en", "Petition")

	msg.Key("We").
		Set("nl", "Wij").
		Set("en", "We")

	msg.Key("observe").
		Set("nl", "constateren").
		Set("en", "observe")

	msg.Key("and request").
		Set("nl", "en verzoeken").
		Set("en", "and request")

	msg.Key("Sign this petition").
		Set("nl", "Onderteken deze petitie").
		Set("en", "Sign this petition")

	msg.Key("Name").
		Set("nl", "Naam").
		Set("en", "Name")

	msg.Key("Name placeholder").
		Set("nl", "Voorletter en achternaam minstens").
		Set("en", "Initial and surname at least")

	msg.Key("Email address").
		Set("nl", "E-mailadres").
		Set("en", "Email address")

	msg.Key("Email placeholder").
		Set("nl", "gebruikersnaam@iets.iets").
		Set("en", "username@something.something")

	msg.Key("City").
		Set("nl", "Woonplaats").
		Set("en", "Place of residence")

	msg.Key("City placeholder").
		Set("nl", "Waar u nu woont, op het tijdstip van ondertekenen").
		Set("en", "Where you live now, at the time of signing")

	msg.Key("Visible").
		Set("nl", "mijn naam en woonplaats mogen publiekelijk zichtbaar zijn onder de petitie.").
		Set("en", "Publish my name visibly below the petition.")

	msg.Key("Understand").
		Set("nl", "ik begrijp dat dit een showcase is en niet de echte petitie.").
		Set("en", "I understand this is a showcase, not the actual petition.")

	msg.Key("Submit signature").
		Set("nl", "Ja, ik onderteken deze petitie").
		Set("en", "Yes, I sign this petition")

	msg.Key("Please acknowledge").
		Set("nl", "Bevestig alstublieft").
		Set("en", "Please confirm")

	msg.Key("Confirm").
		Set("nl", "Mijn ondertekening bevestigen").
		Set("en", "Confirm my subscription")

	msg.Key("ErrAlreadyConfirmed").
		Set("nl", "Deze ondertekening was al bevestigd 👍").
		Set("en", "This signature has been confirmed already 👍")

	msg.Key("Details").
		Set("nl", "Details").
		Set("en", "Details")

	msg.Key("Addressed to").
		Set("nl", "Ontvanger").
		Set("en", "Addressed to")

	msg.Key("Petition desk").
		Set("nl", "Petitieloket").
		Set("en", "Petition desk")

	msg.Key("Closing date").
		Set("nl", "Einddatum").
		Set("en", "Closing date")

	msg.Key("Answer expected").
		Set("nl", "Antwoord verwacht").
		Set("en", "Answer expected")

	msg.Key("Date format").
		Set("nl", "2-1-2006").
		Set("en", "1/2/2006")

	msg.Key("Status").
		Set("nl", "Status").
		Set("en", "Status")

	msg.Key("Signable").
		Set("nl", "Ondertekenbaar").
		Set("en", "Signable")

	msg.Key("Websites").
		Set("nl", "Websites").
		Set("en", "Websites")

	msg.Key("History").
		Set("nl", "Geschiedenis").
		Set("en", "History")
}
