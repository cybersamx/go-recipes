# Long Polling Recipe in Go and Go Deployment as a Docker Container

This is a 2 part example:

* A client and a server communicate with each other using long polling.
* Encapsulate the Go server in a Docker file and running it as a container.

## Long Polling

50% of the time, the server will respond to incoming requests immediately with an event. The other 50% of the time, the server will respond after a delay of 1 to 10 seconds. Independently, the server will time out in 5 seconds after receiving a request. So if the event is only emitted after 5 seconds, the existing connection will be close and the client will have to send the request again.

One pattern that we can use to address this behavior is long-polling.
## Setup

### Running the App

1. Run the program to launch the client and server.

   ```bash
   $ go run main.go
   ```

1. Call the server

   ```bash
   $ curl -i http://localhost:8000/
   ```

1. You can also run both the client and server at the same time.

   ```bash
   $ ENABLE_CLIENT=true go run main.go
   ```

   Alternatively, you can also run the above.

   ```bash
   $ make run
   ```
