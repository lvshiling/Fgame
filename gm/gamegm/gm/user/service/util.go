package service

import (
	"context"
	"fmt"
)

func GetUserCenterPlatList(ctx context.Context) ([]int64, error) {
	rst := make([]int64, 0)
	userid := GmUserIdInContext(ctx)
	if userid < 0 {
		return rst, fmt.Errorf("用户id为空")
	}
	usservice := GmUserServiceInContext(ctx)
	if usservice == nil {
		return rst, fmt.Errorf("中心用户列表，用户服务为空")
	}

	userCenterPlatList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		return rst, err
	}
	return userCenterPlatList, nil
}

func GetUserSdkList(ctx context.Context) ([]int, error) {
	rst := make([]int, 0)
	userid := GmUserIdInContext(ctx)
	if userid < 0 {
		return rst, fmt.Errorf("用户id为空")
	}
	usservice := GmUserServiceInContext(ctx)
	if usservice == nil {
		return rst, fmt.Errorf("中心用户列表，用户服务为空")
	}

	userCenterPlatList, err := usservice.GetUserSdkTypeList(userid)
	if err != nil {
		return rst, err
	}
	return userCenterPlatList, nil
}
