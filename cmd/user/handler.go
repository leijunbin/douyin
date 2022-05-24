package main

import (
	"context"
	"github.com/douyin/cmd/user/service"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		// TODO: create response
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.SetStatusMsg(&errno.ParamErr.ErrMsg)
		return resp, nil
	}
	// 创建新的创建用户服务, 调用床架用户任务函数
	err = service.NewUserRegisterService(ctx).CreateUser(req)
	if err != nil {
		// TODO: create response
		return resp, nil
	}
	// TODO: create a success response
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		// TODO: create response
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.SetStatusMsg(&errno.ParamErr.ErrMsg)
		return resp, nil
	}
	// 创建新的创建用户服务, 调用床架用户任务函数
	err = service.NewUserLoginService(ctx).CheckUser(req)
	if err != nil {
		// TODO: create response
		return resp, nil
	}
	// TODO: create a success response
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	// 判断你输入

	// 创建新的创建用户服务, 调用床架用户任务函数
	err = service.NewGetUserInfoService(ctx).QueryUser(req)
	if err != nil {
		// TODO: create response
		return resp, nil
	}
	// TODO: create a success response
	return resp, nil
}
