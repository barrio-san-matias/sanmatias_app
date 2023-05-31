package handler

import (
	"fmt"
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

	r.ParseForm() // Parses the request body
	from := r.Form.Get("From")
	log.Printf(">>>> FROM: %s\n", from)

	body := r.Form.Get("Body")
	log.Printf(">>>> BODY: %s\n", body)

	message := &twiml.MessagingMessage{
		Body: fmt.Sprintf("Hola %s 👋🏻 estoy en desarrollo. Por ahora no puedo responder ninguna pregunta.", from),
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
