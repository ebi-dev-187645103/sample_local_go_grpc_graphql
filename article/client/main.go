package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  pb.ArticleServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = pb.NewArticleServiceClient(conn)

	// 4. 実行
	Hello()
}

func Hello() {
	name := "fujito"

	req := &pb.CreateArticleRequest{
		ArticleInput: name,
	}
	res, err := client.CreateArticle(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetArticle())
	}
}
