package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// Service作成
	service,err := service.NewService()
	if err != nil{
		panic(err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// gRPCサーバーにGreetingServiceを登録
	// pb.RegisterArticleServiceServer(s, NewMyServer())
	pb.RegisterArticleServiceServer(s, service)

	// サーバーリフレクションの設定
	reflection.Register(s)

	// 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

type myServer struct{
	pb.UnimplementedArticleServiceServer
}

// 自作サービス構造体のコンストラクタを定義
func NewMyServer()*myServer{
	return &myServer{}
}
// func (s *myServer) CreateArticle(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {

func (s *myServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error){

	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &pb.CreateArticleResponse{
		Article: fmt.Sprintf("Hello, %s!", req.GetArticleInput()),
	}, nil
}
