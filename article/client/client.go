package client

import (
	"context"
	"fmt"
	"log"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func NewClient(port string)(*Client,error) {
func NewClient(port string)(error) {
	fmt.Println("client.NewClient: start")
	// 2. gRPCサーバーとのコネクションを確立
	address := fmt.Sprintf("localhost:%s",port)
	fmt.Println("address",address)
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		fmt.Println("NewClient: end")
		// return nil,err
		return err
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client :=&Client{conn,pb.NewArticleServiceClient(conn)}


	client.Hello()
	fmt.Println("client.NewClient: end")

	return nil
}

type Client struct{
	conn   *grpc.ClientConn
	client pb.ArticleServiceClient
}

func (c *Client)Hello() {
	fmt.Println("client.Hello: start")
	name := "fujito"

	req := &pb.CreateArticleRequest{
		ArticleInput: name,
	}
	res, err := c.client.CreateArticle(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetArticle())
	}
	fmt.Println("client.Hello: end")
}
