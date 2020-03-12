/* File: main.go */
/* Description: Handle all rarticles requests. */

package main

import (
	pb "github.com/ronaldoafonso/rarticle/rarticlepb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Fail to listen: %v.", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterRArticleServer(gRPCServer, &RArticleServer{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v.", err)
	}
}
