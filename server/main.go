package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gobuffalo/packr/v2"
	"github.com/tylergdorn/MotdaaS/motd"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	box := packr.New("topics", "./topics")

	m := loadTopics(box)
	log.Printf("Loaded %d topics", len(m))

	s := motd.Server{TopicsMap: m}
	grpcServer := grpc.NewServer()

	motd.RegisterMotdServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func loadTopics(box *packr.Box) map[string]string {
	var topicMap = make(map[string]string)

	for _, topic := range box.List() {
		data, _ := box.Find(topic)
		topicMap[topic] = string(data)
	}

	return topicMap
}
