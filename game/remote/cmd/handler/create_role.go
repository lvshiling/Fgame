package handler

import (
	"fgame/fgame/account/login/types"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	playerdao "fgame/fgame/game/player/dao"
	playerentity "fgame/fgame/game/player/entity"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_CREATE_ROLE_TYPE), cmd.CmdHandlerFunc(handleCreateRole))
}

// 系统公告
func handleCreateRole(msg proto.Message) (err error) {
	log.Info("cmd:请求发送创建角色")
	cmdMsg := msg.(*cmdpb.CmdCreateRole)
	userId := cmdMsg.GetUserId()
	sdkType := types.SDKType(cmdMsg.GetSdkType())
	if !sdkType.Valid() {
		err = cmd.ErrorCodeCommonArgumentInvalid
		log.WithFields(
			log.Fields{
				"userId":  userId,
				"sdkType": sdkType,
			}).Warn("cmd:请求发送创建角色,参数无效")
		return
	}
	err = createPlayer(userId, sdkType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"userId":  userId,
				"sdkType": sdkType,
				"err":     err,
			}).Error("cmd:请求发送创建角色,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"userId":  userId,
			"sdkType": sdkType,
		}).Info("cmd:请求发送创建角色,成功")

	return
}

//后台创建角色
func createPlayer(userId int64, sdkType types.SDKType) (err error) {
	log.WithFields(
		log.Fields{
			"userId":  userId,
			"sdkType": sdkType,
		}).Info("player:后台用户正在创建角色")

	//当前服务器
	serverId := global.GetGame().GetServerIndex()

	pe, err := playerdao.GetPlayerDao().QueryByUserId(userId, serverId)
	if err != nil {
		return
	}
	if pe != nil {
		err = cmd.ErrorCodeCommonPlayerExist
		return
	}

	name := ""
	//重复10次
	for i := 0; i < 10; i++ {
		name = dummytemplate.GetDummyTemplateService().GetRandomDummyName()
		pe, err = playerdao.GetPlayerDao().QueryByName(serverId, name)
		if err != nil {
			return
		}
		if pe == nil {
			break
		}
	}
	if len(name) == 0 {
		err = cmd.ErrorCodeCommonPlayerNoSuitableName
		return
	}
	role := playertypes.RandomRole()
	sex := playertypes.RandomSex()
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	defaultSystemCompensate := int32(1)
	pe = &playerentity.PlayerEntity{
		Id:               id,
		UserId:           userId,
		ServerId:         serverId,
		SdkType:          int32(sdkType),
		OriginServerId:   serverId,
		Name:             name,
		Role:             int32(role),
		Sex:              int32(sex),
		CreateTime:       now,
		SystemCompensate: defaultSystemCompensate,
	}
	if err = global.GetGame().GetDB().DB().Save(pe).Error; err != nil {
		return
	}
	log.WithFields(
		log.Fields{
			"userId":  userId,
			"sdkType": sdkType,
		}).Info("player:用户创建角色成功")

	return
}
