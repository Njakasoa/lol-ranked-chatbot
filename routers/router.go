package routers

import (
	"github.com/Njakasoa/lol-ranked-chatbot/controllers"

	"github.com/gorilla/mux"
)

// InitializeRouter for the router
func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/webhook", controllers.WebhookVerificationHandler).Methods("GET")
	router.HandleFunc("/webhook", controllers.WebhookMessagesHandler).Methods("POST")
	router.HandleFunc("/", controllers.ServerHealthHandler)
	return router
}
