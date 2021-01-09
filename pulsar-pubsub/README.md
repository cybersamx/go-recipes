# Pulsar Pubsub

This recipe implements an event-driven application using pubsub pattern with Apache Pulsar, a high-performance event streaming platform.

## Overview

This sample code is based on the original [Kafka recipe](../kafka-pubsub). The program runs the processors for sending and receiving Kafka messages as goroutines.

## Pulsar Setup

1. Run Apache Pulsar. Note this will run Pulsar-Manager - see next section for details.

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
   $ # Next shutdown pulsar
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just one command.

   ```bash
   $ make
   $ docker-compose down    # Run this to remove containers
   ```

## Pulsar Manager Setup

Pulsar Manager is a visual dashboard, a nice-to-have to quick inspect what Pulsar is doing (it's not required to get the code running). There's a bit of extra setup that you need to do before you can start using the tool.

1. Run docker-compose.

   ```bash
   $ docker-compose up
   ```

1. Generate an admin account to access the tool.

   ```bash
   $ ./gen-admin.sh
   ```

1. Launch the web browser and go to <http://localhost:9527>. Use username `admin` and password `password` to log into the tool.

1. Add an environment by connecting to the Pulsar host. Use the following parameters when creating a new environment.

   | Config Name      | Config Value        |
   |------------------|---------------------|
   | Environment Name | standalone          |
   | Service URL      | http://pulsar:8080/ |

## Reference and Credits

* [Set up a standalone Pulsar in Docker](https://pulsar.apache.org/docs/en/standalone-docker/)
* [Github: Pulsar Manager](https://github.com/apache/pulsar-manager) - Use Pulsar Manager instead of Pulsar Dashboard, which has been deprecated.
