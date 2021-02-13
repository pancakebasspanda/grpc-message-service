package client

import (
	"context"
	"flag"
	"pancakebasspanda/grpc-message-service/protos"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

func New(ctx context.Context) (protos.MessageServiceClient, error) {

	if !flag.Parsed() {
		flag.Parse()
	}

	logrus.WithField("address", address).Info("connecting to message-service")

	retryOptions := []grpc_retry.CallOption{
		grpc_retry.WithMax(maxRetries),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(time.Duration(backoff))),
	}

	conn, err := grpc.DialContext(
		ctx, address,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOptions...)),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOptions...)),
	)
	if err != nil {
		logrus.WithError(err).Error("failed connecting to message-service")
		return nil, err
	}

	defer conn.Close()

	logrus.Info("successfully connected to data service")

	return protos.NewMessageServiceClient(conn), nil

}
