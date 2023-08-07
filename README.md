# Honeypot in Go

This is a simple honeypot script written in Go. It can deploy either an SSH honeypot or an HTTP honeypot. Captured requests are sent to a Discord webhook.

## Prerequisites

-   Go installed on your system.
-   `github.com/gliderlabs/ssh` and `github.com/joho/godotenv` Go libraries.

## Setup

1.  Clone the repository (or download the Go script).
    
2.  Inside the project directory, create a `.env` file with your Discord webhook URL:

`WEBHOOK_URL=https://discord.com/api/webhooks/your_webhook_id/your_webhook_token` 
    
    Replace the URL with your actual webhook URL.
    
3.  Install the required Go libraries:
    
    `go get github.com/gliderlabs/ssh
    go get github.com/joho/godotenv` 
   
## Running the Honeypot

To deploy an SSH honeypot:

`go run your_script_name.go ssh` 

To deploy an HTTP honeypot:

`go run your_script_name.go http` 

Once deployed, any interaction with the honeypot will be logged and sent to the Discord webhook.

## How It Works

-   The SSH honeypot listens on port `2222`.
-   The HTTP honeypot listens on port `8080`.

Any attempts to connect or authenticate are logged and sent to the specified Discord webhook. For the SSH honeypot, this includes any attempted usernames and passwords.

## Important Note

This honeypot is a basic tool for educational and research purposes. Ensure you understand the implications and risks before deploying in a production environment.
