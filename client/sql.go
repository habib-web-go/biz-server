package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/habib-web-go/biz-server/gen/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultMessageId = 2
	defaultAuthKey   = ""
)

var (
	addr      = flag.String("addr", "localhost:5062", "the address to connect to")
	messageId = flag.Uint64("messageId", defaultMessageId, "next message id")
	authKey   = flag.String("authKey", defaultAuthKey, "auth key")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSQLServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// without userId
	r, err := c.GetUsers(ctx, &pb.GetUsersRequest{AuthKey: *authKey, MessageId: *messageId, UserId: nil})
	if err != nil {
		log.Fatalf("failed: %v\n", err)
	}
	log.Printf("Result: %d\n", len(r.Users))

	// with userId
	var userId uint64 = 5
	r, err = c.GetUsers(ctx, &pb.GetUsersRequest{AuthKey: *authKey, MessageId: *messageId, UserId: &userId})
	if err != nil {
		log.Fatalf("failed: %v\n", err)
	}
	log.Printf("Result: %s\n", r.String())

}
