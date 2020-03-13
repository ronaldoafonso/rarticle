package main

import (
	"context"
	pb "github.com/ronaldoafonso/rarticle/rarticlepb"
	"io/ioutil"
	"strings"
)

// RArticleServer ... gRPC server implementation of rarticle.
type RArticleServer struct{}

// GetRArticle ... Returns an rarticle. The content should be an HTML file.
func (*RArticleServer) GetRArticle(ctx context.Context, req *pb.RArticleRequest) (*pb.RArticleResponse, error) {
	name := req.GetName()
	lang := req.GetLang()

	file := "/html/" + lang + "/" + name + ".html"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &pb.RArticleResponse{Content: string(content)}, nil
}

// SearchRArticle ... Return a stream o links.
func (*RArticleServer) SearchRArticle(req *pb.SearchRequest, stream pb.RArticle_SearchRArticleServer) error {
	baseDir := "/html/" + req.GetLang()

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(baseDir + "/" + file.Name())
		if err != nil {
			return err
		}
		if strings.Contains(string(content), req.GetWord()) {
			link := strings.Split(file.Name(), ".")[0]
			stream.Send(&pb.SearchResponse{Link: link})
		}
	}

	return nil
}
