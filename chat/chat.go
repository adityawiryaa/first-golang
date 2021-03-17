package chat

import (
	"context"
	"fmt"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	fmt.Println("message", message.Body)
	return &Message{Body: "Hello From Server!"}, nil
}
