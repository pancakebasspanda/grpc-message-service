package client

import (
	"flag"
	"time"
)

var (
	address    string
	maxRetries uint
	backoff    int
	tracing    bool
)

func init() {
	flag.StringVar(&address, "message-service-address", "localhost:8080", "message-service grpc address")
	flag.UintVar(&maxRetries, "message-service-max-retries", 3, "max number of request retries")
	flag.IntVar(&backoff, "message-service-backoff", int(time.Second*1), "backoff between connection retries")
}
