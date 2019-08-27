package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	titlelogic "fgame/fgame/game/title/logic"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_UP_STAR_TYPE), dispatch.HandlerFunc(handleTitleUpStar))
}

func handleTitleUpStar(s session.Session, msg interface{}) (err error) {
	log.Debug("title: 处理称号升星请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSTitleUpStar)
	titleId := csMsg.GetTitleId()

	err = titleUpStar(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("title:处理称号升星请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
			"error":    err,
		}).Debug("title:处理称号升星请求,成功")
	return
}

func titleUpStar(pl player.Player, titleId int32) (err error) {
	// 判断称号有没有激活
	// 模板那边校验称号时效性
	// 根据升星模板，判断称号升星上限
	// 根据升星模板，判断称号是否能升星
	// 消耗物品
	// 同步背包，同步属性
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeTitleUpstar) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("title: 称号升星失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	if !titleManager.IfTitleExist(titleId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:称号未激活")
		playerlogic.SendSystemMessage(pl, lang.TitleNotHas)
		return
	}

	titleTemp := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:模板不存在")
		playerlogic.SendSystemMessage(pl, lang.TitleTemplateNotExist)
		return
	}

	// 获取称号对象
	titleType := titleTemp.GetTitleType()
	titleObj := titleManager.GetTitleInfo(titleType, titleId)
	if titleObj.ValidTime != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:该称号不是永久称号")
		playerlogic.SendSystemMessage(pl, lang.TitleNotEver)
		return
	}

	starLev := titleObj.StarLev + 1

	// 获取升星模板
	upStarTemp := title.GetTitleService().GetTitleUpStarTemplate(int(titleId), starLev)
	if upStarTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:星级已经到达满级")
		playerlogic.SendSystemMessage(pl, lang.TitleStarLevelAlreadyTop)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := upStarTemp.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	// 消耗物品
	reason := commonlog.InventoryLogReasonTitleUpStar
	reasonText := fmt.Sprintf(reason.String(), int(titleType), starLev)
	flag := inventoryManager.BatchRemove(needItems, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("title: 消耗物品应该成功"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	// 称号升星判断
	_, bless, sucess := titlelogic.TitleUpstar(titleObj.StarNum, titleObj.StarBless, upStarTemp)
	flag = titleManager.Upstar(titleId, bless, sucess)
	if !flag {
		panic(fmt.Errorf("title: 称号升星应该成功"))
	}

	// 同步属性
	if sucess {
		titlelogic.TitlePropertyChanged(pl)
	}

	scMsg := pbutil.BuildSCTitleUpstar(titleId, titleObj.StarLev, titleObj.StarBless)
	pl.SendMsg(scMsg)

	return
}
