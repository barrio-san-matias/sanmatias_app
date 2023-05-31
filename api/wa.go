package handler

import (
	"log"
	"net/http"

	"github.com/joeshaw/envdecode"

	"github.com/twilio/twilio-go/twiml"
)

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	var cfg struct {
		TwilioToken string `env:"TWILIO_TOKEN,required"`
		TwilioSID   string `env:"TWILIO_SID,required"`
	}
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	message := &twiml.MessagingMessage{
		Body: "The Robots are coming! Head for the hills!",
	}

	twimlResult, err := twiml.Messages([]twiml.Element{message})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// Write the response
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(twimlResult))
}
