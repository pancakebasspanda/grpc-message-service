package server

import (
	"context"
	"io"
	"pancakebasspanda/grpc-message-service/protos"

	log "github.com/sirupsen/logrus"
)

// rpc Connection(stream ConnectionRequest) returns (stream ConnectionRequest); // database
// rpc SendMessage(SendMessageRequest) returns (SendMessageResponse);

func (s *MessageServiceServer) SendMessage(ctx context.Context, req *protos.SendMessageRequest) (*protos.SendMessageResponse, error) {
	log.Println(req)
	if s.conn[req.TargetId] != nil {
		s.conn[req.TargetId].Send(&protos.ConnectionResponse{
			Message:  req.Message,
			SenderId: req.SenderId,
		})
	}

	return &protos.SendMessageResponse{
		Sent:     true,
		SenderId: req.SenderId,
		TargetId: req.TargetId,
	}, nil
}

func (s *MessageServiceServer) Connection(stream protos.MessageService_ConnectionServer) error {
	//ctx := stream.Context()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//log.Print("no more data")
		}
		if req != nil {
			s.conn[req.ConnectionId] = stream
		}

	}

	return nil
}
