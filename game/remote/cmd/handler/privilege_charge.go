package handler

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	chargetemplate "fgame/fgame/game/charge/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/dao"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_PRIVILEGE_CHARGE_TYPE), cmd.CmdHandlerFunc(handlePrivilegeCharge))
}

func handlePrivilegeCharge(msg proto.Message) (err error) {
	log.Info("cmd:权限充值")
	cmdPrivilegeCharge := msg.(*cmdpb.CmdPrivilegeCharge)
	playerId := cmdPrivilegeCharge.GetPlayerId()
	gold := cmdPrivilegeCharge.GetGold()

	//TODO 限制最大
	if gold <= 0 {
		err = cmd.ErrorCodeCommonPlayerNoExist
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
			}).Warn("cmd:权限充值,参数错误")
		return
	}
	num := int32(1)
	if cmdPrivilegeCharge.Num != nil {
		num = cmdPrivilegeCharge.GetNum()
	}
	if num <= 0 {
		err = cmd.ErrorCodeCommonPlayerNoExist
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
				"num":      num,
			}).Warn("cmd:权限充值,参数错误")
		return
	}

	err = privilegeGold(playerId, gold, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
				"err":      err,
			}).Error("cmd:权限充值,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": playerId,
			"gold":     gold,
		}).Info("cmd:权限充值,成功")
	return
}

func privilegeGold(playerId int64, gold int64, num int32) (err error) {
	pe, err := dao.GetPlayerDao().QueryById(playerId)
	if err != nil {
		return
	}
	if pe == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
			}).Warn("cmd:权限充值,玩家不存在")
		return cmd.ErrorCodeCommonPlayerNoExist
	}
	sdkType := logintypes.SDKType(pe.SdkType)
	if !sdkType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
			}).Warn("cmd:权限充值,sdk无效")
		return cmd.ErrorCodePrivilegeChargeSDKWrong
	}
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplateByGold(sdkType, int32(gold))
	if chargeTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
			}).Warn("cmd:权限充值,充值模板不存在")
		return cmd.ErrorCodePrivilegeChargeGoldWrong
	}
	itemTemplate := item.GetItemService().GetChargeItemTemplate(chargeTemplate.SubType)
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"gold":     gold,
			}).Warn("cmd:权限充值,物品不存在")
		return cmd.ErrorCodePrivilegeChargeGoldWrong
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		attachment := make(map[int32]int32)
		attachment[int32(itemTemplate.TemplateId())] = num
		title := lang.GetLangService().ReadLang(lang.EmailSystemRewardTitle)
		content := lang.GetLangService().ReadLang(lang.EmailPrivilegeChargeContent)
		emaillogic.AddOfflineEmail(playerId, title, content, attachment)
		return
	}

	privilegeCharge(pl, int32(itemTemplate.TemplateId()), num)
	// charge.GetChargeService().PrivilegeCharge(playerId, gold)
	log.WithFields(
		log.Fields{
			"playerId": playerId,
			"gold":     gold,
		}).Info("cmd:权限充值,充值成功")

	return
}

func privilegeCharge(pl player.Player, itemId int32, num int32) {
	ctx := scene.WithPlayer(context.Background(), pl)
	attachment := make(map[int32]int32)
	attachment[itemId] = num
	msg := message.NewScheduleMessage(onPrivilegeCharge, ctx, attachment, nil)
	pl.Post(msg)
}

func onPrivilegeCharge(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}
	attachment := result.(map[int32]int32)
	title := lang.GetLangService().ReadLang(lang.EmailSystemRewardTitle)
	content := lang.GetLangService().ReadLang(lang.EmailPrivilegeChargeContent)
	// attachment := make(map[int32]int32)
	// attachment[itemId] = 1
	emaillogic.AddEmail(p, title, content, attachment)
	return nil
}
