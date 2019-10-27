package motd

import (
	"context"
	fmt "fmt"
	"log"
	"math/rand"
	"strings"
)

//Server a
type Server struct {
	Topics map[string]string
}

//Motd a
func (s *Server) Motd(ctx context.Context, request *MotdRequest) (*MotdResponse, error) {
	log.Printf("received request for topic: %s", request.Topic)
	message, err := s.getMessage(request.Topic)
	if err != nil {
		log.Printf("encountered err: %s", err)
	}
	return &MotdResponse{Motd: message}, nil
}

func (s *Server) getMessage(topic string) (string, error) {
	motds, ok := s.Topics[topic]
	if !ok {
		return "error", fmt.Errorf("No topic %s in server", topic)
	}
	split := strings.Split(motds, "\n")
	items := len(split)
	return split[rand.Intn(items)], nil
}
