package service

import (
	"context"
	"fmt"
	"runtime"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
)

// 自作サービスのインターフェース
type Service struct{
	pb.UnimplementedArticleServiceServer
}

// 自作サービスの構造体のコンストラクタ
func NewService()*Service{
	return &Service{}
}

//
func (s *Service)CreateArticle(ctx context.Context, req *pb.CreateArticleRequest)(*pb.CreateArticleResponse,error){
	programCounter,pwd,line,ok := runtime.Caller(0)
	fn := runtime.FuncForPC(programCounter)
	fmt.Println("Start: ",fn.Name())
	fmt.Println(programCounter,pwd,line,ok)

	fmt.Println("Serviceやで！")
	id := "100"

	fmt.Println("End: ",fn.Name())
	return &pb.CreateArticleResponse{
		Article: id,
	},nil
}
