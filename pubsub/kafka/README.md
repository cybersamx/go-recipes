# Kafka Pubsub using Watermill Library

An example of implementing an event-driven application using pubsub pattern with kafka as the message broker. We leverage the [Watermill framework](https://github.com/ThreeDotsLabs/watermill) to set up our connection to kafka.

## Overview

The code in this example is based on the [Kafka example](https://github.com/ThreeDotsLabs/watermill/tree/master/_examples/pubsubs/kafka) found in [the Watermill project](https://github.com/ThreeDotsLabs/watermill) with the following differences:

* Our example uses bitnami [Kafka](https://hub.docker.com/r/bitnami/kafka) and [Zookeeper](https://hub.docker.com/r/bitnami/zookeeper) docker images instead of [Confluence](https://hub.docker.com/u/confluentinc) docker images.
* Our go program runs outside of Docker network, hence the broker URL points to localhost.
* The code for sending and receiving Kafka messages are implemented as goroutines.

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

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make run
   $ make teardown    # Run this to remove the containers
   ```

## Reference and Credits

* [Bitnami Kafka Docker Image](https://hub.docker.com/r/bitnami/kafka)
* [Bitnami Zookeeper Docker Image](https://hub.docker.com/r/bitnami/zookeeper)
* [Confluence Kafka Docker Image](https://hub.docker.com/r/confluentinc/cp-kafka)
* [Confluence Zookeeper Docker Image](https://hub.docker.com/r/confluentinc/cp-zookeeper)
* [Watermill Homepage](https://watermill.io/)
* [Watermill Github Project](https://github.com/ThreeDotsLabs/watermill)
