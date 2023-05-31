package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"fmt"
	"log"
	"net/http"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

const (
	token_secret = "verify"
)

//type WebhookMessage map[string]interface{}

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	// Set your Twilio Account SID and Auth Token
	accountSid := "YOUR_TWILIO_ACCOUNT_SID"
	authToken := "YOUR_TWILIO_AUTH_TOKEN"

	// Create a new Twilio client
	client := twilio.NewRestClient(accountSid, authToken)

	// Get the country from the request body
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to parse form data:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	country := r.FormValue("Body")

	// Retrieve data from the API
	resp, err := client.Get(fmt.Sprintf("https://restcountries.eu/rest/v2/name/%s?fullText=true", country), nil)
	if err != nil {
		log.Println("Failed to retrieve data from API:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Create a new TwiML response
	response := api.NewMessageResponse()

	// Check the response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Failed to retrieve data for the following country - %s. Here is a more verbose reason: %s\n", country, resp.Status)
		response.SetBody("Sorry we could not process your request. Please try again or check a different country")
	} else {
		// Parse the response JSON
		var data []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			log.Println("Failed to decode JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if len(data) > 0 {
			countryData := data[0]
			nativeName := countryData["nativeName"].(string)
			capital := countryData["capital"].(string)
			people := countryData["demonym"].(string)
			region := countryData["region"].(string)

			response.SetBody(fmt.Sprintf("%s is a country in %s. Its capital city is %s, while its native name is %s. A person from %s is called a %s.", country, region, capital, nativeName, country, people))
		} else {
			response.SetBody("No data found for the specified country")
		}
	}

	// Write the response
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.String()))
}
