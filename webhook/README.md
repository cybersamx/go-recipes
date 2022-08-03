# Webhook Example in Go

The client will push a random event message to a webhook endpoint on the sever every 5 seconds.

## Setup

1. Run the server

   ```bash
   $ go run model.go server.go
   ```

1. On another shell, run the client:

   ```bash
   $ go run model.go client.go
   ```

   Or use curl to call the webhook.

   ```bash
   $ curl -X POST \
       http://localhost:8000/webhook \
       -H 'Content-Type: application/json; charset=UTF-8' \
       -d '{"event": "create"}'
   ```

1. Alternatively, you can run everything with just 1 command.

   ```bash
   $ make run
   ```
