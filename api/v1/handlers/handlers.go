package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/berke581/go-contact-form/email"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/contact", contactHandler)

	return r
}

type FormData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	data := &FormData{}
	json.NewDecoder(r.Body).Decode(data)

	userEmail := os.Getenv("USER_EMAIL")
	userPw := os.Getenv("USER_PASSWORD")
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")

	sender := email.NewEmailSender(userEmail, userPw)
	var formattedName string
	if len(data.Name) > 0 {
		formattedName = data.Name
	} else {
		formattedName = "--"
	}
	formattedBody := fmt.Sprintf("Name: %s, Email: %s\n\n%s", formattedName, data.Email, data.Body)
	sender.SendEmail(recipientEmail, data.Title, formattedBody)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(data)
}
