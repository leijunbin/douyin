package handlers

import (
	"context"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryUser(c *gin.Context) {
	var queryVar UserInfoParam
	// 参数绑定, 失败报错
	if err := c.BindQuery(&queryVar); err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
	}

	/*
		TODO: 如果token过期, 提示重新登录
		if token {
		    SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		}
	*/

	// 调用远程服务
	userInfo, err := rpc.GetUserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: queryVar.UserID,
		Token:  queryVar.Token,
	})

	if err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
	}
	SendUserInfoResponse(c, errno.Success, userInfo)

}
