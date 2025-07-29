# devops

## Kafka Consumer

This repository includes a simple Kafka consumer implemented in Go. The consumer reads a configuration file and forwards each message received from Kafka to a target HTTP endpoint.

### Building

```
cd consumer
go build
```

### Running

Create a configuration JSON file (an example is provided at `consumer/sample-config.json`) with the following format:

```
{
  "topic": "<kafka-topic>",
  "target_url": "<http-endpoint>"
}
```

Then run the consumer:

```
./consumer -config /path/to/config.json
```

The consumer expects Kafka to be available at `localhost:9092`. This can be overridden by setting the environment variable `KAFKA_BROKERS` to a comma separated list of broker addresses.
