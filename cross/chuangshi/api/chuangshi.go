package api

import (
	"context"
	"fgame/fgame/cross/chuangshi/chuangshi"
	chuangshipb "fgame/fgame/cross/chuangshi/pb"
	chuangshidata "fgame/fgame/game/chuangshi/data"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
)

//服务器服务
type ChuangShiServer struct {
}

//创世信息
func (ts *ChuangShiServer) GetChuangShiWholeInfo(ctx context.Context, req *chuangshipb.ChuangShiWholeInfoRequest) (resp *chuangshipb.ChuangShiWholeInfoResponse, err error) {
	campList := chuangshi.GetChuangShiService().GetChuangShiCampListList()

	resp = &chuangshipb.ChuangShiWholeInfoResponse{}
	resp.CampList = BuildChuangShiCampList(campList)
	return
}

// 神王报名
func (ts *ChuangShiServer) ShenWangSignUp(ctx context.Context, req *chuangshipb.ChuangShiShenWangSignUpRequest) (resp *chuangshipb.ChuangShiShenWangSignUpResponse, err error) {
	success, signList := chuangshi.GetChuangShiService().ChuangShiShenWangSignUp(chuangshitypes.ChuangShiCampType(req.CampType), req.PlayerId)

	resp = &chuangshipb.ChuangShiShenWangSignUpResponse{}
	resp.PlayerId = req.PlayerId
	resp.Success = success
	resp.SignList = BuildChuangShiSignInfoList(signList)
	resp.CampType = req.CampType
	return
}

// 神王投票
func (ts *ChuangShiServer) ShenWangVote(ctx context.Context, req *chuangshipb.ChuangShiShenWangVoteRequest) (resp *chuangshipb.ChuangShiShenWangVoteResponse, err error) {
	success, voteList := chuangshi.GetChuangShiService().ChuangShiShenWangVote(req.PlayerId, chuangshitypes.ChuangShiCampType(req.CampType), req.SupportId)

	resp = &chuangshipb.ChuangShiShenWangVoteResponse{}
	resp.Success = success
	resp.VoteList = BuildChuangShiVoteList(voteList)
	resp.CampType = req.CampType
	return
}

// 城主任命
func (ts *ChuangShiServer) CityRenMing(ctx context.Context, req *chuangshipb.ChuangShiCityRenMingRequest) (resp *chuangshipb.ChuangShiCityRenMingResponse, err error) {
	success := chuangshi.GetChuangShiService().CityRenMing(req.MemberId, req.CityId, req.BeCommitId)

	resp = &chuangshipb.ChuangShiCityRenMingResponse{}
	resp.Success = success
	return
}

// 城池工资分配
func (ts *ChuangShiServer) CityPaySchedule(ctx context.Context, req *chuangshipb.ChuangShiCityPayScheduleRequest) (resp *chuangshipb.ChuangShiCityPayScheduleResponse, err error) {
	paramList := chuangshidata.CrossConvertToCityPayScheduleParamList(req.GetScheduleList())
	chuangshi.GetChuangShiService().CityPaySchedule(req.PlayerId, paramList)

	resp = &chuangshipb.ChuangShiCityPayScheduleResponse{}
	return
}

// 阵营工资分配
func (ts *ChuangShiServer) CampPaySchedule(ctx context.Context, req *chuangshipb.ChuangShiCampPayScheduleRequest) (resp *chuangshipb.ChuangShiCampPayScheduleResponse, err error) {
	paramList := chuangshidata.CrossConvertToCampPayScheduleParamList(req.GetScheduleList())
	camp := chuangshi.GetChuangShiService().CampPaySchedule(req.PlayerId, paramList)

	resp = &chuangshipb.ChuangShiCampPayScheduleResponse{}
	resp.Camp = BuildChuangShiCamp(camp)
	return
}

// 阵营工资领取
func (ts *ChuangShiServer) CampPayReceive(ctx context.Context, req *chuangshipb.ChuangShiCampPayReceiveRequest) (resp *chuangshipb.ChuangShiCampPayReceiveResponse, err error) {
	camp := chuangshi.GetChuangShiService().CampPayReceive(req.PlayerId)

	resp = &chuangshipb.ChuangShiCampPayReceiveResponse{}
	resp.Camp = BuildChuangShiCamp(camp)
	return
}

// 城池建设
func (ts *ChuangShiServer) CityJianShe(ctx context.Context, req *chuangshipb.ChuangShiCityJianSheRequest) (resp *chuangshipb.ChuangShiCityJianSheResponse, err error) {
	success := chuangshi.GetChuangShiService().CityChengFangJianShe(req.PlayerId, req.CityId, chuangshitypes.ChuangShiCityJianSheType(req.JianSheType), req.Num)

	resp = &chuangshipb.ChuangShiCityJianSheResponse{}
	resp.Success = success
	return
}

// 设置攻城目标
func (ts *ChuangShiServer) GongChengTargetFuShu(ctx context.Context, req *chuangshipb.ChuangShiGongChengTargetRequest) (resp *chuangshipb.ChuangShiGongChengTargetResponse, err error) {

	success := chuangshi.GetChuangShiService().GongChengTargetFuShu(req.PlayerId, req.CityId)

	resp = &chuangshipb.ChuangShiGongChengTargetResponse{}
	resp.Success = success
	return
}

// 加入阵营
func (ts *ChuangShiServer) JoinCamp(ctx context.Context, req *chuangshipb.ChuangShiJoinCampRequest) (resp *chuangshipb.ChuangShiJoinCampResponse, err error) {

	campType := chuangshitypes.ChuangShiCampType(req.CampType)
	memList := chuangshidata.ConvertToMemberList(req.MemList)
	success := chuangshi.GetChuangShiService().JoinCamp(campType, memList)

	resp = &chuangshipb.ChuangShiJoinCampResponse{}
	resp.Success = success
	return
}

// 加入阵营
func (ts *ChuangShiServer) CityTianQiSet(ctx context.Context, req *chuangshipb.ChuangShiCityTianQiSetRequest) (resp *chuangshipb.ChuangShiCityTianQiSetResponse, err error) {
	success := chuangshi.GetChuangShiService().CityTianQiSet(req.PlayerId, req.CityId, req.Level)
	resp = &chuangshipb.ChuangShiCityTianQiSetResponse{}
	resp.Success = success
	return
}

func NewChuangShiServer() *ChuangShiServer {
	ss := &ChuangShiServer{}
	return ss
}
