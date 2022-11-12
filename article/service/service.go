package service

import (
	"context"

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

func (s *Service)CreateArticle(ctx context.Context, req *pb.CreateArticleRequest)(*pb.CreateArticleResponse,error){
	common.PrintStart("")
	// INSERTする記事のInputを取得
	input := req.ArticleInput

	// 記事をDBにINSERTし、INSERTした記事のIDを返す
	id,err := s.repository.InsertArticle(ctx,input)
	if err != nil{
		return nil,err
	}
	common.PrintEnd("")
	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	},nil
}

func (s *Service)ReadArticle(ctx context.Context, req *pb.ReadArticleRequest)(*pb.ReadArticleResponse,error){
	// INSERTする記事のInputを取得
	id := req.GetId()

	// DBから該当IDの記事を取得
	a,err := s.repository.SelectArticleByID(ctx,id)
	if err != nil{
		return nil,err
	}

	// 取得した記事をレスポンスとして返す
	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	},nil
}

func (s *Service)UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest)(*pb.UpdateArticleResponse,error){
	id    := req.GetId()
	input := req.GetArticleInput()

	// 該当記事をUPDATE
	if err := s.repository.UpdateArticle(ctx,id,input); err != nil{
		return nil,err
	}

	// 取得した記事をレスポンスとして返す
	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Title:   input.Title,
			Author:  input.Author,
			Content: input.Content,
		},
	},nil
}
