// Code generated by Kitex v1.9.1. DO NOT EDIT.

package thumbservice

import (
	"code.byted.org/kite/kitex/byted"
	"code.byted.org/kite/kitex/client"
	"code.byted.org/kite/kitex/client/callopt"
	"context"
	"douyin/kitex_gen/like"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Likeyou(ctx context.Context, request *like.LikeyouRequest, callOptions ...callopt.Option) (r *like.LikeyouResponse, err error)
	ThumbList(ctx context.Context, request *like.ThumbListRequest, callOptions ...callopt.Option) (r *like.ThumbListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	config := byted.NewClientConfig()
	config.DestService = destService

	options = append(options, byted.ClientSuiteWithConfig(serviceInfo(), config))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kThumbServiceClient{
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

type kThumbServiceClient struct {
	*kClient
}

func (p *kThumbServiceClient) Likeyou(ctx context.Context, request *like.LikeyouRequest, callOptions ...callopt.Option) (r *like.LikeyouResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Likeyou(ctx, request)
}

func (p *kThumbServiceClient) ThumbList(ctx context.Context, request *like.ThumbListRequest, callOptions ...callopt.Option) (r *like.ThumbListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ThumbList(ctx, request)
}

// NewClientWithBytedConfig creates a client for the service defined in IDL.
func NewClientWithBytedConfig(destService string, config *byted.ClientConfig, opts ...client.Option) (Client, error) {
	if config == nil {
		config = byted.NewClientConfig()
	}
	config.DestService = destService

	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, byted.ClientSuiteWithConfig(serviceInfo(), config))
	options = append(options, opts...)
	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kThumbServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClientWithBytedConfig creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClientWithBytedConfig(destService string, config *byted.ClientConfig, opts ...client.Option) Client {
	kc, err := NewClientWithBytedConfig(destService, config, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}
