package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	timeout := flag.String("timeout", "10s", "timeout seconds")
	flag.Parse()

	toDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTimeout("tcp", flag.Args()[0] + ":" + flag.Args()[1], toDuration)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), toDuration)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(ctx, cancel, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn)
		wg.Done()
	}()

	wg.Wait()
}

func readRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				cancel()
				break
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)

			os.Stdout.Write([]byte(fmt.Sprintf("%s\n", text)))
		}
	}
	log.Println("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				break
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Println("Finished writeRoutine")
}