# Long Polling Example in Go

An example of client-server long polling system written in Go.

An event will surface anytime between 1 to 10 seconds and is then sent to the client. If time for event to surface > 5 seconds, timeout and return a response with no content.

Both the server and client will run concurrently in a main function.

## Setup

1. Run the program to launch the client and server. 

   ```bash
   $ go run main.go
   ```
