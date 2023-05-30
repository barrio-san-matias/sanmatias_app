package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	webhookSecret = "verify" // Replace with your actual secret token
)

type WebhookMessage struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("incoming request (wa)")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	/***
	var cfg struct {
		TelegramToken string `env:"WHATSAPP_TOKEN,required"`
	}
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}
	***/

	// Verify the webhook secret
	if !verifyWebhookSecret(r) {
		log.Println("Webhook secret verification failed")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var webhookData WebhookMessage
	err = json.Unmarshal(body, &webhookData)
	if err != nil {
		log.Println("Failed to parse webhook data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process the incoming message
	log.Printf("Received message from %s: %s\n", webhookData.From, webhookData.Message)

	// Send a response (optional)
	response := "Received your message: " + webhookData.Message
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, response)
}

func verifyWebhookSecret(r *http.Request) bool {
	// Retrieve the X-Hub-Signature header from the request
	signature := r.Header.Get("X-Hub-Signature")

	// Retrieve the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		return false
	}

	// Calculate the expected signature
	expectedSignature := calculateSignature(body)

	// Compare the signatures
	return subtle.ConstantTimeCompare([]byte(signature), []byte(expectedSignature)) == 1
}

func calculateSignature(data []byte) string {
	secret := []byte(webhookSecret)
	hash := hmacSHA256(data, secret)
	return fmt.Sprintf("sha256=%x", hash)
}

func hmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
