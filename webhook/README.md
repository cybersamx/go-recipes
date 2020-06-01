# Webhook Example in Go

An example of a simple webhook system written in Go.

The client will push a random event message to the webhook on the sever every 5 seconds.

## Setup

1. Run the server

   ```bash
   $ go run model.go server.go
   ```

1. On another shell, run the client:

   ```bash
   $ go run model.go client.go
   ```

   alternatively, you can also run:

   ```bash
   $ curl -X POST \
       http://localhost:8000/webhook \
       -H 'Content-Type: application/json; charset=UTF-8' \
       -d '{"event": "create"}'
   ```
