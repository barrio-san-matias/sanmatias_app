package handler

import (
	"log"
	"net/http"

	"github.com/joeshaw/envdecode"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	var cfg struct {
		TwilioToken string `env:"TWILIO_TOKEN,required"`
		TwilioSID   string `env:"TWILIO_SID,required"`
	}
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	// Set your Twilio Account SID and Auth Token
	accountSid := cfg.TwilioSID
	authToken := cfg.TwilioToken

	// Create a new Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Get the country from the request body
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to parse form data:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	country := r.FormValue("Body")

	// Create a new TwiML response
	response := twilioApi.NewMessageResponse()

	response.SetBody("HI!!!!!")

	// Write the response
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.String()))
}
