package client

import (
	"context"
	chuangshiapi "fgame/fgame/cross/chuangshi/api"
	chuangshipb "fgame/fgame/cross/chuangshi/pb"
	chuangshidata "fgame/fgame/game/chuangshi/data"

	"google.golang.org/grpc"
)

type ChuangshiClient interface {
	GetChuangShiWholeInfo(ctx context.Context) (resp *chuangshipb.ChuangShiWholeInfoResponse, err error)
	ShenWangSignUp(ctx context.Context, platform, serverId int32, playerId int64) (resp *chuangshipb.ChuangShiShenWangSignUpResponse, err error)
	ShenWangVote(ctx context.Context, supportId int64) (resp *chuangshipb.ChuangShiShenWangVoteResponse, err error)
	CityRenMing(ctx context.Context, playerId int64, cityId, beCommitId int64) (resp *chuangshipb.ChuangShiCityRenMingResponse, err error)
	CityPaySchedule(ctx context.Context, playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (resp *chuangshipb.ChuangShiCityPayScheduleResponse, err error)
	CampPaySchedule(ctx context.Context, playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (resp *chuangshipb.ChuangShiCampPayScheduleResponse, err error)
	CampPayReceive(ctx context.Context, playerId int64) (resp *chuangshipb.ChuangShiCampPayReceiveResponse, err error)
	CityJianShe(ctx context.Context, playerId, cityId int64, jianSheType, num int32) (resp *chuangshipb.ChuangShiCityJianSheResponse, err error)
	GongChengTargetFuShu(ctx context.Context, playerId, cityId int64) (resp *chuangshipb.ChuangShiGongChengTargetResponse, err error)
	JoinCamp(ctx context.Context, campType int32, memList ...*chuangshidata.MemberInfo) (resp *chuangshipb.ChuangShiJoinCampResponse, err error)
	CityTianQiSet(ctx context.Context, playerId, cityId int64, level int32) (resp *chuangshipb.ChuangShiCityTianQiSetResponse, err error)
}

type chuangshiClient struct {
	c      *grpc.ClientConn
	remote chuangshipb.ChuangshiClient
}

func (m *chuangshiClient) GetChuangShiWholeInfo(ctx context.Context) (resp *chuangshipb.ChuangShiWholeInfoResponse, err error) {
	req := &chuangshipb.ChuangShiWholeInfoRequest{}
	resp, err = m.remote.GetChuangShiWholeInfo(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) ShenWangSignUp(ctx context.Context, platform, serverId int32, playerId int64) (resp *chuangshipb.ChuangShiShenWangSignUpResponse, err error) {
	req := &chuangshipb.ChuangShiShenWangSignUpRequest{}
	req.PlayerId = playerId
	req.Platform = platform
	req.ServerId = serverId
	resp, err = m.remote.ShenWangSignUp(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) ShenWangVote(ctx context.Context, supportId int64) (resp *chuangshipb.ChuangShiShenWangVoteResponse, err error) {
	req := &chuangshipb.ChuangShiShenWangVoteRequest{}
	req.SupportId = supportId
	resp, err = m.remote.ShenWangVote(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CityRenMing(ctx context.Context, playerId int64, cityId, beCommitId int64) (resp *chuangshipb.ChuangShiCityRenMingResponse, err error) {
	req := &chuangshipb.ChuangShiCityRenMingRequest{}
	req.BeCommitId = beCommitId
	req.CityId = cityId
	req.MemberId = playerId
	resp, err = m.remote.CityRenMing(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CityPaySchedule(ctx context.Context, playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (resp *chuangshipb.ChuangShiCityPayScheduleResponse, err error) {
	req := &chuangshipb.ChuangShiCityPayScheduleRequest{}
	req.PlayerId = playerId
	req.ScheduleList = chuangshiapi.BuildCityPayScheduleList(paramList)

	resp, err = m.remote.CityPaySchedule(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CampPaySchedule(ctx context.Context, playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (resp *chuangshipb.ChuangShiCampPayScheduleResponse, err error) {
	req := &chuangshipb.ChuangShiCampPayScheduleRequest{}
	req.PlayerId = playerId
	req.ScheduleList = chuangshiapi.BuildCampPayScheduleList(paramList)

	resp, err = m.remote.CampPaySchedule(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CampPayReceive(ctx context.Context, playerId int64) (resp *chuangshipb.ChuangShiCampPayReceiveResponse, err error) {
	req := &chuangshipb.ChuangShiCampPayReceiveRequest{}
	req.PlayerId = playerId

	resp, err = m.remote.CampPayReceive(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CityJianShe(ctx context.Context, playerId, cityId int64, jianSheType, num int32) (resp *chuangshipb.ChuangShiCityJianSheResponse, err error) {
	req := &chuangshipb.ChuangShiCityJianSheRequest{}
	req.PlayerId = playerId
	req.CityId = cityId
	req.JianSheType = jianSheType
	req.Num = num

	resp, err = m.remote.CityJianShe(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) GongChengTargetFuShu(ctx context.Context, playerId, cityId int64) (resp *chuangshipb.ChuangShiGongChengTargetResponse, err error) {
	req := &chuangshipb.ChuangShiGongChengTargetRequest{}
	req.PlayerId = playerId
	req.CityId = cityId

	resp, err = m.remote.GongChengTargetFuShu(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) JoinCamp(ctx context.Context, campType int32, memList ...*chuangshidata.MemberInfo) (resp *chuangshipb.ChuangShiJoinCampResponse, err error) {
	req := &chuangshipb.ChuangShiJoinCampRequest{}
	req.CampType = campType
	req.MemList = chuangshiapi.BuildChuangShiMemberInfoListByData(memList)

	resp, err = m.remote.JoinCamp(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *chuangshiClient) CityTianQiSet(ctx context.Context, playerId, cityId int64, level int32) (resp *chuangshipb.ChuangShiCityTianQiSetResponse, err error) {
	req := &chuangshipb.ChuangShiCityTianQiSetRequest{}
	req.PlayerId = playerId
	req.CityId = cityId
	req.Level = level

	resp, err = m.remote.CityTianQiSet(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewChuangshiClient(conn *grpc.ClientConn) ChuangshiClient {
	m := &chuangshiClient{}
	m.c = conn
	m.remote = chuangshipb.NewChuangshiClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
