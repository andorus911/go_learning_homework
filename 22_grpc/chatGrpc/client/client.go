package main

import (
	"bufio"
	chat "chatGrpc/chatpb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

func writeRoutine(end chan interface{}, ctx context.Context, conn chat.ChatExampleClient) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			if str == "exit" {
				break OUTER
			}
			msg, err := conn.SendMessage(context.Background(), &chat.ChatMessage{Text: str,})
			if err != nil {
				fmt.Printf("error:%s\n", status.Convert(err).Message())
			}

			if msg != nil {
				created, _ := ptypes.Timestamp(msg.Created)
				fmt.Printf("[%s]id:%d msg:%s\n", created.Local(), msg.Id, msg.Text)
			}

		}

	}
	log.Printf("Finished writeRoutine")
	close(end)
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := chat.NewChatExampleClient(cc)
	end := make(chan interface{})
	go writeRoutine(end, ctx, c)

	<-end
}
