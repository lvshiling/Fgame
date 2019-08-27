package service

import (
	"context"
	"errors"
	pb "fgame/fgame/game/remote/cmd/pb"
	rmmodel "fgame/fgame/gm/gamegm/remote/model"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/golang/protobuf/proto"
)

type IUserRemoteService interface {
	ForbidPlayer(p_serverId int32, p_play *rmmodel.ForbidPlayer) error
	ForbidPlayerChat(p_serverId int32, p_play *rmmodel.ForbidPlayerChat) error

	UnForbidPlayer(p_serverId int32, p_playerid int64) error
	UnForbidPlayerChat(p_serverId int32, p_playerid int64) error

	IgnorePlayerChat(p_serverId int32, p_play *rmmodel.IgnorePlayerChat) error
	UnIgnorePlayer(p_serverId int32, p_playerid int64) error

	ChatSet(p_serverId int32, p_worldVipLevel int32, p_worldLevel int32, p_allianceVipLevel int32, p_allianceLevel int32, p_privateVipLevel int32, p_privateLevel int32, p_teamVip int32, p_teamPlayerLevel int32) error
	RegisterServerSet(p_serverId int32, p_open int32) error

	SendServerCompensate(p_serverId int32, p_title string, p_content string, p_attachmentStr string, p_needLevel int32, p_needCreateTime int64, p_bindFlag int) error
	SendPlayerCompensate(p_serverId int32, p_playerIdList string, p_title string, p_content string, p_attachmentStr string, p_bindFlag int) error

	PrivilegeCharge(p_serverId int32, p_playerId int64, p_gold int64, p_num int32) error
	PrivilegeChargeSet(p_serverId int32, p_playerId int64, p_type int32) error

	BroadcastNotice(p_serverId int32, p_content string, p_beginTime int64, p_endTime int64, p_intervalTime int64) error
	KickOut(p_serverId int32, p_playerId int64, p_reason string) error

	SetMarryBanquetHouTaiType(p_serverId int32, houtai int32) error

	CreateRole(p_serverId int32, p_userId int64, p_sdkType int32) error

	CustomRecycleGold(p_serverId int32, p_gold int64) error
	Ping(p_serverKeyId int32, p_gameServerId int32, p_platformId int32) error

	FirstChargeReset(p_serverId int32) error

	//修改公告
	ModifyAllianceGongGao(serverId int32, allianceId int64, gongGaoName string) error
	DismissAlliance(serverId int32, allianceId int64) error
}

type userRemoteService struct {
}

func (m *userRemoteService) ForbidPlayer(p_serverId int32, p_play *rmmodel.ForbidPlayer) error {
	request := &pb.CmdForbidPlayer{
		PlayerId:     &p_play.PlayerId,
		ForbidReason: &p_play.ForbidReason,
		ForbidName:   &p_play.ForbidName,
		ForbidTime:   &p_play.ForbidTime,
	}

	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) ForbidPlayerChat(p_serverId int32, p_play *rmmodel.ForbidPlayerChat) error {
	request := &pb.CmdForbidPlayerChat{
		PlayerId:     &p_play.PlayerId,
		ForbidReason: &p_play.ForbidReason,
		ForbidName:   &p_play.ForbidName,
		ForbidTime:   &p_play.ForbidTime,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) UnForbidPlayer(p_serverId int32, p_playerid int64) error {

	request := &pb.CmdUnforbidPlayer{
		PlayerId: &p_playerid,
	}

	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) UnForbidPlayerChat(p_serverId int32, p_playerid int64) error {

	request := &pb.CmdUnforbidPlayerChat{
		PlayerId: &p_playerid,
	}

	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) IgnorePlayerChat(p_serverId int32, p_play *rmmodel.IgnorePlayerChat) error {
	request := &pb.CmdIgnorePlayerChat{
		PlayerId:     &p_play.PlayerId,
		ForbidReason: &p_play.ForbidReason,
		ForbidName:   &p_play.ForbidName,
		ForbidTime:   &p_play.ForbidTime,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) UnIgnorePlayer(p_serverId int32, p_playerid int64) error {
	request := &pb.CmdUnignorePlayerChat{
		PlayerId: &p_playerid,
	}

	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) ChatSet(p_serverId int32, p_worldVipLevel int32, p_worldLevel int32, p_allianceVipLevel int32, p_allianceLevel int32, p_privateVipLevel int32, p_privateLevel int32, p_teamVip int32, p_teamPlayerLevel int32) error {
	request := &pb.CmdChatSet{
		WorldVipLevel:    &p_worldVipLevel,
		WorldLevel:       &p_worldLevel,
		AllianceVipLevel: &p_allianceVipLevel,
		AllianceLevel:    &p_allianceLevel,
		PrivateVipLevel:  &p_privateVipLevel,
		PrivateLevel:     &p_privateLevel,
		TeamLevel:        &p_teamPlayerLevel,
		TeamVipLevel:     &p_teamVip,
	}

	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) RegisterServerSet(p_serverId int32, p_open int32) error {
	request := &pb.CmdRegisterSet{
		Open: &p_open,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) SendServerCompensate(p_serverId int32, p_title string, p_content string, p_attachmentStr string, p_needLevel int32, p_needCreateTime int64, p_bindFlag int) error {
	bindFlag := false
	if p_bindFlag > 0 {
		bindFlag = true
	}
	request := &pb.CmdSendServerCompensate{
		Title:          &p_title,
		Content:        &p_content,
		AttachmentStr:  &p_attachmentStr,
		NeedLevel:      &p_needLevel,
		NeedCreateTime: &p_needCreateTime,
		Bind:           &bindFlag,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) SendPlayerCompensate(p_serverId int32, p_playerIdList string, p_title string, p_content string, p_attachmentStr string, p_bindFlag int) error {
	bindFlag := false
	if p_bindFlag > 0 {
		bindFlag = true
	}
	request := &pb.CmdSendPlayerCompensate{
		Title:         &p_title,
		Content:       &p_content,
		PlayerIdList:  &p_playerIdList,
		AttachmentStr: &p_attachmentStr,
		Bind:          &bindFlag,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) PrivilegeCharge(p_serverId int32, p_playerId int64, p_gold int64, p_num int32) error {
	request := &pb.CmdPrivilegeCharge{
		PlayerId: &p_playerId,
		Gold:     &p_gold,
		Num:      &p_num,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) PrivilegeChargeSet(p_serverId int32, p_playerId int64, p_type int32) error {
	request := &pb.CmdPrivilegeSet{
		PlayerId:  &p_playerId,
		Privilege: &p_type,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) BroadcastNotice(p_serverId int32, p_content string, p_beginTime int64, p_endTime int64, p_intervalTime int64) error {
	request := &pb.CmdBroadcastNotice{
		Content:      &p_content,
		BeginTime:    &p_beginTime,
		EndTime:      &p_endTime,
		IntervalTime: &p_intervalTime,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) KickOut(p_serverId int32, p_playerId int64, p_reason string) error {
	request := &pb.CmdKickoutPlayer{
		PlayerId:      &p_playerId,
		KickoutReason: &p_reason,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) SetMarryBanquetHouTaiType(p_serverId int32, houtai int32) error {
	request := &pb.CmdMarrySet{
		HouTaiType: &houtai,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) CreateRole(p_serverId int32, p_userId int64, p_sdkType int32) error {
	request := &pb.CmdCreateRole{
		UserId:  &p_userId,
		SdkType: &p_sdkType,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) CustomRecycleGold(p_serverId int32, p_gold int64) error {
	request := &pb.CmdCustomRecycleGold{
		Gold: &p_gold,
	}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) Ping(p_serverKeyId int32, p_gameServerId int32, p_platformId int32) error {
	request := &pb.CmdPing{
		ServerId:   &p_gameServerId,
		PlatformId: &p_platformId,
	}
	return m.doCmd(p_serverKeyId, request)
}

func (m *userRemoteService) FirstChargeReset(p_serverId int32) error {
	request := &pb.CmdFirstChargeReset{}
	return m.doCmd(p_serverId, request)
}

func (m *userRemoteService) ModifyAllianceGongGao(serverId int32, allianceId int64, gongGaoName string) error {
	request := &pb.CmdAllainceNotice{}
	request.AllianceId = &allianceId
	request.NoticeStr = &gongGaoName
	return m.doCmd(serverId, request)
}

func (m *userRemoteService) DismissAlliance(serverId int32, allianceId int64) error {
	request := &pb.CmdAllianceDismiss{}
	request.AllianceId = &allianceId

	return m.doCmd(serverId, request)
}

func (m *userRemoteService) doCmd(p_serverId int32, msg proto.Message) error {
	client, err := GetFClientRemote(p_serverId)
	if err != nil {
		return nil
	}
	ctx := context.Background()

	rst, err := client.DoCmd(ctx, msg)
	if err != nil {
		return err
	}
	if rst.ErrorCode != 0 {
		return errors.New(rst.ErrorMsg)
	}
	return nil
}

func NewUserRemoteServer() IUserRemoteService {
	rst := &userRemoteService{}
	return rst
}

type contextKey string

const (
	userRemoteServiceKey = contextKey("UserRemoteService")
)

func WithUserRemoteService(ctx context.Context, ls IUserRemoteService) context.Context {
	return context.WithValue(ctx, userRemoteServiceKey, ls)
}

func UserRemoteServiceInContext(ctx context.Context) IUserRemoteService {
	us, ok := ctx.Value(userRemoteServiceKey).(IUserRemoteService)
	if !ok {
		return nil
	}
	return us
}

func SetupUserRemoteServiceHandler(ls IUserRemoteService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithUserRemoteService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
