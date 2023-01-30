// Code generated by Kitex v0.3.2. DO NOT EDIT.

package commentservice

import (
	"context"
	"dousheng/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateComment(ctx context.Context, createCommentReq *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error)
	DeleteComment(ctx context.Context, deleteCommentReq *comment.DeleteCommentReq, callOptions ...callopt.Option) (r *comment.DeleteCommentResp, err error)
	GetCommentsByVideoId(ctx context.Context, getCommentsByVideoIdReq *comment.GetCommentsByVideoIdReq, callOptions ...callopt.Option) (r *comment.GetCommentsByVideoIdResp, err error)
	CountComment(ctx context.Context, countCommentReq *comment.CountCommentReq, callOptions ...callopt.Option) (r *comment.CountCommentResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CreateComment(ctx context.Context, createCommentReq *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, createCommentReq)
}

func (p *kCommentServiceClient) DeleteComment(ctx context.Context, deleteCommentReq *comment.DeleteCommentReq, callOptions ...callopt.Option) (r *comment.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, deleteCommentReq)
}

func (p *kCommentServiceClient) GetCommentsByVideoId(ctx context.Context, getCommentsByVideoIdReq *comment.GetCommentsByVideoIdReq, callOptions ...callopt.Option) (r *comment.GetCommentsByVideoIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentsByVideoId(ctx, getCommentsByVideoIdReq)
}

func (p *kCommentServiceClient) CountComment(ctx context.Context, countCommentReq *comment.CountCommentReq, callOptions ...callopt.Option) (r *comment.CountCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountComment(ctx, countCommentReq)
}