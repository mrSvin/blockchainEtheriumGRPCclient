package service

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetWallet(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server gRPC!"}, nil
}
