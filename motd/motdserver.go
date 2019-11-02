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
	TopicsMap map[string]string
}

//Motd returns a random message for the requested topic
func (s *Server) Motd(ctx context.Context, request *MotdRequest) (*MotdResponse, error) {
	log.Printf("received request for topic: %s", request.Topic)
	message, err := s.getMessage(request.Topic)
	if err != nil {
		log.Printf("encountered err: %s", err)
	}
	return &MotdResponse{Motd: message}, nil
}

func (s *Server) getMessage(topic string) (string, error) {
	motds, ok := s.TopicsMap[topic]
	if !ok {
		return "error", fmt.Errorf("No topic %s in server", topic)
	}
	split := strings.Split(motds, "\n")
	items := len(split)
	return split[rand.Intn(items)], nil
}

func (s *Server) getTopics() []string {
	keys := make([]string, len(s.TopicsMap))
	i := 0
	for k := range s.TopicsMap {
		keys[i] = k
		i++
	}
	return keys
}

// Topics returns a list of all topics the motd server has
func (s *Server) Topics(ctx context.Context, request *TopicEnumRequest) (*TopicEnumResponse, error) {
	log.Print("Got topic request")
	return &TopicEnumResponse{Topics: s.getTopics()}, nil
}
