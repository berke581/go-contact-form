package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	"github.com/berke581/go-contact-form/email"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/contact", contactHandler)

	return r
}

type FormData struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email"`
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	data := &FormData{}
	json.NewDecoder(r.Body).Decode(data)

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		json.NewEncoder(rw).Encode(err.(validator.ValidationErrors).Error())
		return
	}

	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")

	sender := email.NewEmailSender(senderEmail, senderPassword)
	sender.SendEmail(receiverEmail, data.Name, data.Title, data.Email, data.Body)

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(data)
}
