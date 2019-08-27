package use

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenetypes "fgame/fgame/game/scene/types"
	shengtanscene "fgame/fgame/game/shengtan/scene"
	shengtantemplate "fgame/fgame/game/shengtan/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAlliance, itemtypes.ItemAllianceSubTypeJiuNiang, playerinventory.ItemUseHandleFunc(handleJiuLian))
}

func handleJiuLian(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shengtan:喝酒酿,玩家不在圣坛中")
		playerlogic.SendSystemMessage(pl, lang.ShengTanPlayerNoInScene)
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shengtan:喝酒酿,玩家不在圣坛中")
		playerlogic.SendSystemMessage(pl, lang.ShengTanPlayerNoInScene)
		return
	}
	sd, ok := s.SceneDelegate().(shengtanscene.ShengTanSceneData)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shengtan:喝酒酿,玩家不在圣坛中")
		playerlogic.SendSystemMessage(pl, lang.ShengTanPlayerNoInScene)
		return
	}
	//判断是否可以继续喝
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("shengtan:喝酒酿,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.ShengTanPlayerNoInScene)
		return
	}
	maxNum := shengtantemplate.GetShengTanTemplateService().GetShengTanTemplate().ExpAddItemLimit
	jiuNiangNum, _ := sd.GetJiuNiangNum()
	if jiuNiangNum+num > maxNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("shengtan:喝酒酿,使用酒酿超过上限")
		playerlogic.SendSystemMessage(pl, lang.ShengTanJiuNiangLimit)
		return
	}
	//添加酒酿
	expAdd := itemTemplate.TypeFlag1 * num
	sd.AddJiuNiang(num, expAdd)

	//仙盟频道
	noticeArgs := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeAllianceAltar, 0}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToGoNow, noticeArgs)
	noticeContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ShengTanJiuNiangUseNotice), pl.GetName(), itemTemplate.Name, joinLink)
	chatlogic.SystemBroadcastAllianceId(pl.GetAllianceId(), chattypes.MsgTypeText, []byte(noticeContent))
	return true, nil
}
