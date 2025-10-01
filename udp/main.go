package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	info "github.com/nginx/kic-test-containers"
)

var version string

func main() {
	commitHash, commitTime, dirtyBuild := info.GetBuildInfo()
	log.Printf("version: %v, commit: %v, date: %v, dirty: %v\n", version, commitHash, commitTime, dirtyBuild)

	ip := os.Getenv("POD_IP")
	log.Printf("ip: %v\n", ip)
	if ip == "" {
		log.Fatalf("missing required env var: POD_IP")
	}
	port := flag.String("port", "3334", "The port the server listens to")
	flag.Parse()
	listener, err := (&net.ListenConfig{}).ListenPacket(context.Background(), "udp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()
	log.Printf("listening to udp connections at: :%v", *port)
	buffer := make([]byte, 1024)
	for {
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			log.Panicln(err)
		}

		request := string(buffer[:n])

		log.Printf("packet-received: request=%q bytes=%d from=%s", request, n, addr.String())

		response := fmt.Sprintf("%v:%v", ip, *port)
		if request == "health" {
			response = "healthy"
		}

		log.Printf("write data to connection: %q", response)
		n, err = listener.WriteTo([]byte(response), addr)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("packet-written: bytes=%d to=%s", n, addr.String())
	}
}
