package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Njakasoa/lol-ranked-chatbot/models"
)

const (
	// FacebookAPI : URL to use for callback
	FacebookAPI = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
	// Image used for test
	Image = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
)

// WebhookVerificationHandler for facebook bot messenger
func WebhookVerificationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mode := query.Get("hub.mode")
	token := query.Get("hub.verify_token")
	challenge := query.Get("hub.challenge")

	if mode == "subscribe" && token == os.Getenv("LOUPGAROU_BOT_VERIFY_TOKEN") {
		fmt.Println("WEBHOOK_VERIFIED")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("SOMETHING WRONG HAPPENED")
	}
}

// WebhookMessagesHandler manage message from client
func WebhookMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var callback models.Callback
	json.NewDecoder(r.Body).Decode(&callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				ProcessMessage(event)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Message not supported"))
	}
}

// ProcessMessage from client
func ProcessMessage(event models.Messaging) {
	client := &http.Client{}
	response := models.Response{
		Recipient: models.User{
			ID: event.Sender.ID,
		},
		Message: models.Message{
			Attachment: &models.Attachment{
				Type: "image",
				Payload: models.Payload{
					URL: Image,
				},
			},
		},
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	url := fmt.Sprintf(FacebookAPI, os.Getenv("BOT_GG_PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
