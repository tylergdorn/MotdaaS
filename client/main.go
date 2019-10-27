package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tylergdorn/MotdaaS/motd"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Need two arguments: address, topic")
	}

	conn, err := grpc.Dial(args[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()
	c := motd.NewMotdClient(conn)

	response, err := c.Motd(context.Background(), &motd.MotdRequest{
		Topic: args[1],
	})
	if err != nil {
		log.Fatalf("Error when calling motd: %s", err)
	}
	fmt.Println(response.Motd)
}
