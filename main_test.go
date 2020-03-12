package main

import (
	"context"
	pb "github.com/ronaldoafonso/rarticle/rarticlepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

func initGRPCServer() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRArticleServer(grpcServer, &RArticleServer{})
	reflection.Register(grpcServer)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func TestServerGetRArticle(t *testing.T) {
	initGRPCServer()
	conn, err := grpc.Dial("localhost:80", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("GRPC error. Did not connect: %v.", err)
	}
	defer conn.Close()

	grpcClient := pb.NewRArticleClient(conn)

	req := pb.RArticleRequest{
		Name: "wrongname",
		Lang: "wronglang",
	}
	res, err := grpcClient.GetRArticle(context.Background(), &req)
	if err == nil {
		t.Errorf("GetRArticle should return an error but returned nil.")
	}
	if res != nil {
		t.Errorf("GetRArticle should return nil not %v.", res)
	}

	tests := []struct {
		name string
		lang string
		want string
	}{
		{"zero", "en", "<p>I'm a zero.</p>"},
		{"zero", "br", "<p>Eu sou um zero.</p>"},
	}

	for _, test := range tests {
		req := pb.RArticleRequest{
			Name: test.name,
			Lang: test.lang,
		}
		res, err := grpcClient.GetRArticle(context.Background(), &req)
		if err != nil {
			t.Errorf("GetRArticle error: %v.", err)
		}
		got := res.GetContent()
		if got != test.want {
			t.Errorf("GetRArticle error. Want: %v, got: %v.", test.want, got)
		}
	}
}
