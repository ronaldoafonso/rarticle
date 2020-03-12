package main

import (
	"context"
	pb "github.com/ronaldoafonso/rarticle/rarticlepb"
	"io/ioutil"
)

// RArticleServer ... gRPC server implementation of rarticle.
type RArticleServer struct{}

// GetRArticle ... Returns an rarticle. The content should be an HTML file.
func (*RArticleServer) GetRArticle(ctx context.Context, req *pb.RArticleRequest) (*pb.RArticleResponse, error) {
	HTMLFile := "/html/" + req.GetName() + "-" + req.GetLang() + ".html"
	content, err := ioutil.ReadFile(HTMLFile)
	if err != nil {
		return nil, err
	}
	res := pb.RArticleResponse{
		Content: string(content),
	}
	return &res, nil
}
