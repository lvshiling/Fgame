package use

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/template"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBossCallTicket, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleItemUse))
}

func handleItemUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("boosTicket:添加怪物,处于跨服场景")
		playerlogic.SendSystemMessage(pl, lang.BossTicketPlayerNoInWorld)
		return
	}

	s := pl.GetScene()
	if s == nil {
		log.Warn("bossTicket:场景不存在")
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeWorld {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("boosTicket:添加怪物,不是世界场景")
		playerlogic.SendSystemMessage(pl, lang.BossTicketPlayerNoInWorld)
		return
	}

	if s.MapTemplate().LimitPkMode == pktypes.PkStatePeach.Mask() {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("boosTicket:添加怪物,不是PVP场景")
		playerlogic.SendSystemMessage(pl, lang.BossTicketPlayerNoInPVPMap)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	monsterId := itemTemplate.TypeFlag1
	tempBiologyTemplate := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if tempBiologyTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"monsterId": monsterId,
			}).Warn("bossTicket:添加怪物错误，生物模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	biologyTemplate := tempBiologyTemplate.(*gametemplate.BiologyTemplate)

	if biologyTemplate.CanReborn() {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"biologyType": biologyTemplate.GetBiologyType().String(),
			}).Warn("bossTicket:添加怪物,Boss不可重生")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), 0, biologyTemplate, pl.GetPosition(), 0, 0)
	if n == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"biologyType": biologyTemplate.GetBiologyType().String(),
			}).Warn("bossTicket:添加怪物,npc不存在")
		return
	}
	s.AddSceneObject(n)

	// 公告
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, pl.GetName())
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, n.GetBiologyTemplate().Name)
	linkArgs := []int64{int64(chattypes.ChatLinkToWorldMap), int64(s.MapId())}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToKill, linkArgs)

	format := lang.GetLangService().ReadLang(lang.BossTicketBossBornNotice)
	content := fmt.Sprintf(format, playerName, bossName, joinLink)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return true, nil
}
