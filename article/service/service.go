package service

import (
	"context"
	"strconv"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/common"
	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/repository"
)

// 自作サービスのインターフェース
type Service struct{
	pb.UnimplementedArticleServiceServer
	repository repository.Repository
}

// 自作サービスの構造体のコンストラクタ
func NewService()(*Service,error){
	repo,err := repository.NewSqliteRepo()
	if err != nil{
		return nil,err
	}
	return &Service{
		repository: repo,
	},nil
}

//
func (s *Service)CreateArticle(ctx context.Context, req *pb.CreateArticleRequest)(*pb.CreateArticleResponse,error){
	common.PrintStart("")

	// 記事をDBにINSERTし、INSERTした記事のIDを返す
	// id,err := s.repository.InsertArticle(ctx,input)
	id,err := s.repository.InsertArticle(ctx)
	if err != nil{
		return nil,err
	}
	// return &pb.CreateArticleResponse{
	// 	Article: &pb.Article{
	// 		Id:      id,
	// 		Author:  input.Author,
	// 		Title:   input.Title,
	// 		Content: input.Content,
	// 	},
	// },nil

	common.PrintEnd("")
	return &pb.CreateArticleResponse{
		Article: strconv.FormatInt(id, 10),
	},nil
}
