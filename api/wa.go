package handler

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	token_secret = "verify"
)

type WebhookMessage struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func WhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("incoming request (wa)")

	switch r.Method {
	case "GET":
		verifyToken(w, r)
	case "POST":
		receiveMessage(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("hub.verify_token") == token_secret {
		w.Write([]byte(r.FormValue("hub.challenge")))
	} else {
		w.Write([]byte("Error; wrong verify token"))
	}
}

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Error reading empty response body", http.StatusBadRequest)

		return
	}

	defer r.Body.Close()

	message, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)

		return
	}

	// TODO: validate xhub signature

	var msg WebhookMessage
	err = json.Unmarshal(message, &msg)

	if err != nil {
		http.Error(w, "Error parsing response body format", http.StatusBadRequest)

		return
	}

	log.Printf(">>>>> msg: %+v", msg)
}

func verifySignature(signature, secret, message []byte) bool {
	mac := hmac.New(sha1.New, secret)
	mac.Write(message)

	expectedSignature := mac.Sum(nil)

	return hmac.Equal(expectedSignature, hexSignature(signature))
}

func hexSignature(signature []byte) []byte {
	s := make([]byte, 20)

	hex.Decode(s, signature)

	return s
}
