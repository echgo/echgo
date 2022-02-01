package tests

import (
	"github.com/echgo/echgo/telegram"
	"log"
	"testing"
)

// Test the smtp function
func TestTelegram(t *testing.T) {

	r := telegram.Request{
		ApiToken: "5054233711:AAGV_sK-5dzEqqWONF-rufeuUDJfeZ4NpfQ",
	}

	message, err := telegram.CreateMessage("New message!", "1998631548", "Markdown", r)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(message)
	}

}
