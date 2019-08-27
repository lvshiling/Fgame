package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/house/house"
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
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_SELL_TYPE), dispatch.HandlerFunc(handleHouseSell))
}

//处理房子出售信息
func handleHouseSell(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子出售消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMgs := msg.(*uipb.CSHouseSell)
	houseIndex := csMgs.GetHouseIndex()
	logTime := csMgs.GetLogTime()

	err = houseSell(tpl, houseIndex, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Error("house:处理获取房子出售消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"houseIndex": houseIndex,
		}).Debug("house:处理获取房子出售消息完成")
	return nil
}

func houseSell(pl player.Player, houseIndex int32, logTime int64) (err error) {
	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)
	// 房子是否存在
	houseInfo := houseManager.GetHouse(houseIndex)
	if houseInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"houseIndex": houseIndex,
			}).Warn("house:房子出售失败，房子不存在")
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
			}).Warn("house:房子出售失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	price := int64(houseTemplate.HousePrice)
	// 不是满级,没有100%收益
	if !houseTemplate.IsMaxLevel() {
		// math.Ceil() 向上取整;  common.MAX_RATE：策划数据配置基本都是走万分比
		price = int64(math.Ceil(float64(price) * float64(houseTemplate.AdvanceSalePercent) / float64(common.MAX_RATE)))
	}

	rewSilver := int32(0)
	rewBindGold := int32(0)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch houseType {
	case housetypes.HouseTypeSilver:
		{
			rewSilver = int32(price)
			// 银两变化日志
			silverGetReason := commonlog.SilverLogReasonHouseSell
			silverGetReasonText := fmt.Sprintf(silverGetReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
			// 添加银两
			propertyManager.AddSilver(price, silverGetReason, silverGetReasonText)
		}
	case housetypes.HouseTypeBindGold:
		{
			rewBindGold = int32(price)
			// 元宝变化日志
			bindGoldGetReason := commonlog.GoldLogReasonHouseSell
			bindGoldGetReasonText := fmt.Sprintf(bindGoldGetReason.String(), houseIndex, houseType, houseLevel) //日志参数，写清楚变化的房子信息
			// 添加元宝
			propertyManager.AddGold(price, true, bindGoldGetReason, bindGoldGetReasonText)
		}
	}
	//资源变化，同步给客户端显示
	propertylogic.SnapChangedProperty(pl)

	// 房子出售
	flag := houseManager.HouseSell(houseIndex)
	if !flag {
		panic(fmt.Errorf("house:出售房子应该成功,房子序号：%d", houseIndex))
	}

	rd := propertytypes.CreateRewData(0, 0, rewSilver, 0, rewBindGold) //通用客户端资源显示对象
	logList := house.GetHouseService().GetLogByTime(logTime)           //1.logTime:客户端显示的最后一条日志时间；2.更新到最新的日志:取客户端显示的最后一条日志的时间到最新的日志
	scMsg := pbutil.BuildSCHouseSell(houseIndex, rd, logList)
	pl.SendMsg(scMsg)
	return
}
