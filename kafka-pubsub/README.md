# Kafka Pubsub using Watermill Library

An example of implementing an event-driven application using pubsub pattern with Kafka as the message broker and [Watermill framework](https://github.com/ThreeDotsLabs/watermill).

## Overview

This sample code is basically based on the [Kafka example](https://github.com/ThreeDotsLabs/watermill/tree/master/_examples/pubsubs/kafka) found in [the Watermill project](https://github.com/ThreeDotsLabs/watermill). Other than the following differences, the sample code is similar to the Watermill Kafka example:

* Uses bitnami [Kafka](https://hub.docker.com/r/bitnami/kafka) and [Zookeeper](https://hub.docker.com/r/bitnami/zookeeper) docker images instead of [Confluence](https://hub.docker.com/u/confluentinc) docker images.
* The go program runs outside of Docker network, hence the broker URL points to localhost.
* Run the processors for sending and receiving Kafka messages as goroutines.

## Setup

1. Run Kafka and Zookeeper.

   ```bash
   $ docker-compose up
   ```

1. Run the go program.

   ```bash
   $ go run main.go
   ```

1. Terminate the programs.

   ```bash
   $ # Press CTRL-C to terminate the go program
   $ # Next shutdown kafka and zookeeper
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just one command.

   ```bash
   $ make
   $ make teardown    # Run this to remove containers
   ```

## Reference and Credits

* [Bitnami Kafka Docker Image](https://hub.docker.com/r/bitnami/kafka)
* [Bitnami Zookeeper Docker Image](https://hub.docker.com/r/bitnami/zookeeper)
* [Confluence Kafka Docker Image](https://hub.docker.com/r/confluentinc/cp-kafka)
* [Confluence Zookeeper Docker Image](https://hub.docker.com/r/confluentinc/cp-zookeeper)
* [Watermill Homepage](https://watermill.io/)
* [Watermill Github Project](https://github.com/ThreeDotsLabs/watermill)
