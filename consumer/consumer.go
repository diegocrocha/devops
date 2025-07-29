package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

// Config defines the structure of the consumer configuration file.
type Config struct {
	Topic     string `json:"topic"`
	TargetURL string `json:"target_url"`
}

func main() {
	configPath := flag.String("config", "config.json", "path to configuration file")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("cannot parse config: %v", err)
	}

	brokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   cfg.Topic,
		GroupID: "consumer-group",
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		resp, err := http.Post(cfg.TargetURL, "application/json", bytes.NewReader(m.Value))
		if err != nil {
			log.Printf("error posting message: %v", err)
			continue
		}
		resp.Body.Close()
		log.Printf("forwarded message offset %d", m.Offset)
	}
}

// getEnv returns the value of the environment variable if set, otherwise the default value.
func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
