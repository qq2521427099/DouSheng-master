// Code generated by Kitex v0.3.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"dousheng/kitex_gen/favorite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	IsFavorite(ctx context.Context, isFavoriteReq *favorite.IsFavoriteReq, callOptions ...callopt.Option) (r *favorite.IsFavoriteResp, err error)
	Like(ctx context.Context, likeReq *favorite.LikeReq, callOptions ...callopt.Option) (r *favorite.LikeResp, err error)
	UnLike(ctx context.Context, unLikeReq *favorite.UnLikeReq, callOptions ...callopt.Option) (r *favorite.UnLikeResp, err error)
	GetFavoritesByUserId(ctx context.Context, getFavoritesByUserIdReq *favorite.GetFavoritesByUserIdReq, callOptions ...callopt.Option) (r *favorite.GetFavoritesByUserIdResp, err error)
	CountFavorite(ctx context.Context, countFavoriteReq *favorite.CountFavoriteReq, callOptions ...callopt.Option) (r *favorite.CountFavoriteResp, err error)
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
	return &kFavoriteServiceClient{
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

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) IsFavorite(ctx context.Context, isFavoriteReq *favorite.IsFavoriteReq, callOptions ...callopt.Option) (r *favorite.IsFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFavorite(ctx, isFavoriteReq)
}

func (p *kFavoriteServiceClient) Like(ctx context.Context, likeReq *favorite.LikeReq, callOptions ...callopt.Option) (r *favorite.LikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Like(ctx, likeReq)
}

func (p *kFavoriteServiceClient) UnLike(ctx context.Context, unLikeReq *favorite.UnLikeReq, callOptions ...callopt.Option) (r *favorite.UnLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnLike(ctx, unLikeReq)
}

func (p *kFavoriteServiceClient) GetFavoritesByUserId(ctx context.Context, getFavoritesByUserIdReq *favorite.GetFavoritesByUserIdReq, callOptions ...callopt.Option) (r *favorite.GetFavoritesByUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoritesByUserId(ctx, getFavoritesByUserIdReq)
}

func (p *kFavoriteServiceClient) CountFavorite(ctx context.Context, countFavoriteReq *favorite.CountFavoriteReq, callOptions ...callopt.Option) (r *favorite.CountFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFavorite(ctx, countFavoriteReq)
}
