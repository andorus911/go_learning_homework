package main

import (
	chat "chatGrpc/chatpb"
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat.RegisterChatExampleServer(grpcServer, ServerExample{})
	_ = grpcServer.Serve(lis)
}

type ServerExample struct {
}

var i int64

func (s ServerExample) SendMessage(context context.Context, msg *chat.ChatMessage) (*chat.ChatMessage, error) {
	defer func() { i++ }()
	if msg.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "No empty txt allowed")
	}
	return &chat.ChatMessage{Text: "Pong: " + msg.Text, Id: i, Created: ptypes.TimestampNow()}, nil
}
