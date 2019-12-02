package main

type (
    // Event represents a system event.
    Event struct {
        Event string `json:"event"`
    }

    // Message represents a system message.
    Message struct {
        Message string `json:"message"`
    }
)
