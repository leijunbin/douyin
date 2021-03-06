package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/comment"

	"douyin/pkg/snowflake"

	"douyin/cmd/comment/repository"
)

var snowflakeNode *snowflake.Node

func InitSnowflakeNode() {
	tmpNode, err := snowflake.NewNode(1)
	if err != nil {
		zapcomment.Logger.Panic("snowflake error: " + err.Error())
	}
	snowflakeNode = tmpNode
	zapcomment.Logger.Info("snowflake initialization succeeded")
}

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {
	commentId := snowflakeNode.Generate().Int64()

	user, err := rpc.GetUserInfo(s.ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	commentModel := &mysqldb.Comment{
		CommentID: commentId,
		VideoID:   req.VideoId,
		UserID:    req.UserId,
		State:     true,
		Content:   req.Content,
	}

	dbReq := repository.NewRepositoryCom(1).WithComment(commentModel).WithUser(user)

	if err := repository.ProducerComment(s.ctx, dbReq); err != nil {
		return nil, err
	}

	return pack.ChangeComment(commentModel, user), nil
}
