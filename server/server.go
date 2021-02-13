package server

import (
	"pancakebasspanda/grpc-message-service/protos"
)

type MessageServiceServer struct {
	conn map[string]protos.MessageService_ConnectionServer
}

func NewServer() *MessageServiceServer {
	return &MessageServiceServer{conn: make(map[string]protos.MessageService_ConnectionServer)}
}
