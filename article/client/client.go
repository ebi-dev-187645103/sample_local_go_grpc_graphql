package client

import (
	"context"
	"fmt"
	"log"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/common"
	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"google.golang.org/grpc"
)

type Client struct{
	Conn   *grpc.ClientConn
	Client pb.ArticleServiceClient
}

func (c *Client)Create() {
	common.PrintStart("")

	articleInfo := &pb.ArticleInput{
		Author:  "fujito",
		Title:   "my hero academia",
		Content: "so nice Mange and Animetion",
	}

	req := &pb.CreateArticleRequest{
		ArticleInput: articleInfo,
	}
	res, err := c.Client.CreateArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to CreateArticle: %v\n",err)
	}
	fmt.Printf("CreateArticle Response: %v\n",res.GetArticle())
	common.PrintEnd("")
}

func (c *Client)Read() {
	var id int64 = 2
	res,err := c.Client.ReadArticle(
		context.Background(),
		&pb.ReadArticleRequest{Id: id},
	)
	if err != nil{
		log.Fatalf("Failde to ReadArticle: %v\n",err)
	}
	fmt.Printf("ReadArticle Response: %v\n",res)
}
