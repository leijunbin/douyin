// Code generated by Kitex v1.9.1. DO NOT EDIT.

package thumbservice

import (
	"code.byted.org/kite/kitex/client"
	kitex "code.byted.org/kite/kitex/pkg/serviceinfo"
	"context"
	"github.com/yanhongsun/douyin/kitex_gen/like"
)

func serviceInfo() *kitex.ServiceInfo {
	return thumbServiceServiceInfo
}

var thumbServiceServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "ThumbService"
	handlerType := (*like.ThumbService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Likeyou":   kitex.NewMethodInfo(likeyouHandler, newThumbServiceLikeyouArgs, newThumbServiceLikeyouResult, false),
		"ThumbList": kitex.NewMethodInfo(thumbListHandler, newThumbServiceThumbListArgs, newThumbServiceThumbListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "like",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v1.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func likeyouHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.ThumbServiceLikeyouArgs)
	realResult := result.(*like.ThumbServiceLikeyouResult)
	success, err := handler.(like.ThumbService).Likeyou(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newThumbServiceLikeyouArgs() interface{} {
	return like.NewThumbServiceLikeyouArgs()
}

func newThumbServiceLikeyouResult() interface{} {
	return like.NewThumbServiceLikeyouResult()
}

func thumbListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.ThumbServiceThumbListArgs)
	realResult := result.(*like.ThumbServiceThumbListResult)
	success, err := handler.(like.ThumbService).ThumbList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newThumbServiceThumbListArgs() interface{} {
	return like.NewThumbServiceThumbListArgs()
}

func newThumbServiceThumbListResult() interface{} {
	return like.NewThumbServiceThumbListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Likeyou(ctx context.Context, request *like.LikeyouRequest) (r *like.LikeyouResponse, err error) {
	var _args like.ThumbServiceLikeyouArgs
	_args.Request = request
	var _result like.ThumbServiceLikeyouResult
	if err = p.c.Call(ctx, "Likeyou", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ThumbList(ctx context.Context, request *like.ThumbListResponse) (r *like.ThumbListResponse, err error) {
	var _args like.ThumbServiceThumbListArgs
	_args.Request = request
	var _result like.ThumbServiceThumbListResult
	if err = p.c.Call(ctx, "ThumbList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
