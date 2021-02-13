package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"pancakebasspanda/grpc-message-service/protos"
	"pancakebasspanda/grpc-message-service/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const appName = "message-service"

var (
	level   string
	address string
)

func init() {
	flag.StringVar(&level, "level", "debug", "log level defaults to Info")
	flag.StringVar(&address, "address", ":8080", "address to run the server on")
}

func main() {
	flag.Parse()
	logging()

	grpcServer := grpc.NewServer()

	messageServer := server.NewServer()

	protos.RegisterMessageServiceServer(grpcServer, messageServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			grpcServer.GracefulStop()
			os.Exit(1)
		}
	}()

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}

func logging() {
	if level, err := logrus.ParseLevel(level); err != nil {
		logrus.WithError(err).Fatal("parsing log level")
	} else {
		logrus.SetLevel(level)
	}
}
