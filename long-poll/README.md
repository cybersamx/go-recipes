# Long Polling Example in Go

An example of client-server long polling system written in Go.

An event will surface anytime between 1 to 10 seconds and is then sent to the client. If time for event to surface > 5 seconds, timeout and return a response with no content.

## Setup

1. Run the server

   ```bash
   $ go run server.go
   ```
   
2. On another shell, run the client

   ```bash
   $ go run client.go
   ```
