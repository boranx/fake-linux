package main

import (
	"fmt"
	"log"

	"github.com/abhayptp/go-chatgpt"
)

const (
	// BearerToken is the bearer token from browser developer tools.
	TOKEN = "<session-token>"
)

func main() {

	// Initialize. Copy bearer-token and session-token from browser developer tools.
	c := chatgpt.NewChatGpt(chatgpt.NewClient(&chatgpt.Credentials{
		BearerToken:  "Bearer <bearer-token>",
		SessionToken: TOKEN,
	}))

	// Send init message
	_, err := c.SendMessage(`
I want you to act as a linux terminal until I tell you to stop.
I will send you commands and you will send me the output of those commands.
`)
	if err != nil {
		log.Fatal(err)
	}

	var param string

	for {
		fmt.Print("$ ")
		// Send Message Continuously
		_, err := fmt.Scanln(&param)
		if err != nil {
			log.Fatal(err)
		}

		// Send Message to ChatGPT
		res, err := c.SendMessage(param)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
	}
}
