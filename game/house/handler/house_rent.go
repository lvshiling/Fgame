package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/house/pbutil"
	playerhouse "fgame/fgame/game/house/player"
	housetemplate "fgame/fgame/game/house/template"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_RECEIVE_RENT_TYPE), dispatch.HandlerFunc(handleHouseRent))
}

//处理房子租金
func handleHouseRent(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子租金消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMgs := msg.(*uipb.CSHouseReceiveRent)
	houseIndex := csMgs.GetHouseIndex()

	err = houseRent(tpl, houseIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Error("house:处理获取房子租金消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"houseIndex": houseIndex,
		}).Debug("house:处理获取房子租金消息完成")
	return nil
}

func houseRent(pl player.Player, houseIndex int32) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)
	// 房子是否存在
	houseInfo := houseManager.GetHouse(houseIndex)
	if houseInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
			}).Warn("house:房子租金失败，房子不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	houseType := houseInfo.GetHouseType()
	houseLevel := houseInfo.GetHouseLevel()

	houseTemplate := housetemplate.GetHouseTemplateService().GetHouseTemplate(houseIndex, houseType, houseLevel)
	if houseTemplate == nil {
		// 控制台日志，写全引起错误的信息
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子租金失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	// 已经损坏
	if houseInfo.IsBroken() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子租金失败，已经损坏")
		playerlogic.SendSystemMessage(pl, lang.HouseHadBroken)
		return
	}

	// 已经领取租金
	if houseInfo.IsRent() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"houseType":  houseType,
				"houseLevel": houseLevel,
			}).Warn("house:房子租金失败，已经领取租金")
		playerlogic.SendSystemMessage(pl, lang.HouseHadRent)
		return
	}

	rewSilver := int32(0)
	rewBindGold := int32(0)
	price := int64(houseTemplate.Rent)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch houseType {
	case housetypes.HouseTypeSilver:
		{
			rewSilver = int32(price)
			silverGetReason := commonlog.SilverLogReasonHouseRent
			silverGetReasonText := fmt.Sprintf(silverGetReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
			propertyManager.AddSilver(price, silverGetReason, silverGetReasonText)
		}
	case housetypes.HouseTypeBindGold:
		{
			rewBindGold = int32(price)
			bindGoldGetReason := commonlog.GoldLogReasonHouseRent
			bindGoldGetReasonText := fmt.Sprintf(bindGoldGetReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
			propertyManager.AddGold(price, true, bindGoldGetReason, bindGoldGetReasonText)
		}
	}

	//资源变化，同步给客户端显示
	propertylogic.SnapChangedProperty(pl)

	// 领取租金
	flag := houseManager.HouseReceiveRent(houseIndex)
	if !flag {
		panic(fmt.Errorf("house:领取租金应该成功,房子序号：%d", houseIndex))
	}

	rd := propertytypes.CreateRewData(0, 0, rewSilver, 0, rewBindGold) //通用客户端资源显示对象
	scMsg := pbutil.BuildSCHouseReceiveRent(houseIndex, rd)
	pl.SendMsg(scMsg)
	return
}
