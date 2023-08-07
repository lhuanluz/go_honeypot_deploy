package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
)

var webhookURL = os.Getenv("WEBHOOK_URL")

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    webhookURL = os.Getenv("WEBHOOK_URL")
    if webhookURL == "" {
        log.Fatal("WEBHOOK_URL not set. Ensure .env file is correctly set up.")
    }
}



func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: honeypot <ssh/http>")
		return
	}

	switch os.Args[1] {
	case "ssh":
		deploySSHHoneypot()
	case "http":
		deployHTTPHoneypot()
	default:
		fmt.Println("Unknown option. Use either 'ssh' or 'http'.")
	}
}

func sendToWebhook(content string) {
	payload := map[string]string{
		"content": content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to encode payload to JSON: %s", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send POST request: %s", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK response from webhook: %d", resp.StatusCode)
	}
}

func deploySSHHoneypot() {
	ssh.Handle(func(s ssh.Session) {
		log.Printf("SSH connection from: %s", s.RemoteAddr())
	})

	server := &ssh.Server{
		Addr: ":2222",
		PasswordHandler: func(ctx ssh.Context, pass string) bool {
			message := fmt.Sprintf("SSH login attempt: user=%s pass=%s from IP: %s", ctx.User(), pass, ctx.RemoteAddr())
			log.Println(message)
			sendToWebhook(message)
			return false // Always deny the login
		},
	}

	log.Printf("SSH honeypot running on :2222")
	log.Fatal(server.ListenAndServe())
}

func deployHTTPHoneypot() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("HTTP request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
		log.Println(message)
		sendToWebhook(message)
	})

	log.Printf("HTTP honeypot running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
