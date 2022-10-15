# Retry

A simple code that retries connecting a client to a http server.

> **Note**
>
> This is a very basic retry algorithm, which works for a simple or small-scale application. For a more robust retry algorithm, especially in large-scale, distributed context, we need to implement an algorithm with some backoff and even jitter features. We will implement something like this later.

## Setup

1. Run the application.

   ```bash
   $ make run
   ```
