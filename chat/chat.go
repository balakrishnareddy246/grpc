package chat

import (
	"log"
	"golang.org/x/net/content"
	"github.com/balakrishnareddy246/grpc/chat"
)

type Server struct {

}
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	log.Printf(*Received message body from client: %s", err)

}

grpcServer := grpc.NewServer()

s:= chat.Server()
chat.RegisterchatServer(grpcServer, &s)
