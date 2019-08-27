package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	marrylogic "fgame/fgame/game/marry/logic"
	marryservice "fgame/fgame/game/marry/marry"
	pbutil "fgame/fgame/game/marry/pbutil"
	marryplayer "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MarryDingQingSuoYaoDeal), dispatch.HandlerFunc(handleMarryDingQingSuoYaoDeal))
}

type marryDingqingSuoYaoDeal struct {
	SuitId int32
	PosId  int32
}

//定情信物索要处理
func handleMarryDingQingSuoYaoDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:定情信物索要")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	suoYaoInfo := msg.(*uipb.CSMarryDingQingSuoYaoDealMsg)
	suitId := suoYaoInfo.GetSuitId()
	posId := suoYaoInfo.GetPosId()
	aggreeFlag := suoYaoInfo.GetDealFlag()
	if !aggreeFlag {
		marryDingQingSuoYaoDealFail(tpl, suitId, posId)
		return
	}
	marryDingQingSuoYaoDealSuccess(tpl, suitId, posId)
	log.Debug("marry:定情信物索要,成功")
	return
}

//拒绝索要
func marryDingQingSuoYaoDealFail(pl player.Player, suitId int32, posId int32) {
	spId := pl.GetSpouseId()
	spl := player.GetOnlinePlayerManager().GetPlayerById(spId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"spouseId": spId,
			}).Warn("marry:定情信物索要,配偶不在线")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotSpouseNotOnLine)
		return
	}
	//发送给伴侣更改数据
	postData := &marryDingqingSuoYaoDeal{}
	postData.PosId = posId
	postData.SuitId = suitId
	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(sendMarryDingQingSpouseDealMsgFail, splCtx, postData, nil)
	spl.Post(msg)
	//TODO:cjy 结果好像没用
	responMsg := pbutil.BuildSCMarryDingQingSuoYaoDealMsg(int32(1))
	pl.SendMsg(responMsg)
}

func sendMarryDingQingSpouseDealMsgFail(ctx context.Context, result interface{}, err error) error {
	info := result.(*marryDingqingSuoYaoDeal)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(info.SuitId, info.PosId)
	title := lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuFail)
	//TODO:cjy 信物名字是空的
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuFailContent), item.GetXinWuName())
	emaillogic.AddEmail(pl, title, content, nil)

	return nil
}

//同意索要
func marryDingQingSuoYaoDealSuccess(pl player.Player, suitId int32, posId int32) {
	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(suitId, posId)
	spId := pl.GetSpouseId()
	spl := player.GetOnlinePlayerManager().GetPlayerById(spId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"spouseId": spId,
			}).Warn("marry:定情信物索要,配偶不在线")

		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotSpouseNotOnLine)
		return
	}

	//伴侣已经拥有信物
	if marryservice.GetMarryService().ExistsSpouseDingQing(pl.GetId(), suitId, posId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"spouseId": spId,
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物索要,配偶已经有订情信物")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuAlreadyExists)
		return
	}

	itemMap := make(map[int32]int32)
	itemMap[item.GetItemId()] = 1
	// 背包里物品不足走消耗元宝
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	enoughFlag := true
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			enoughFlag = false
		}
	}
	if enoughFlag {
		// 足够则消耗物品
		reason := commonlog.InventoryLogReasonMarryDingQingTokenGive
		reasonText := fmt.Sprintf(reason.String(), item.GetXinWuName())
		flag := inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic("jieyi: 消耗物品应该成功")
		}
		// 物品改变推送
		inventorylogic.SnapInventoryChanged(pl)
	} else {
		//不足消耗钱
		flag := autoBuyDingQingItem(pl, item.GetItemId(), 1)
		if !flag {
			return
		}
	}

	//发送给伴侣更改数据
	postData := newMarryDingqingSuoYaoDealPostData(suitId, posId)
	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(sendMarryDingQingSpouseDealMsg, splCtx, postData, nil)
	spl.Post(msg)

	responMsg := pbutil.BuildSCMarryDingQingSuoYaoDealMsg(int32(1))
	pl.SendMsg(responMsg)
}

func sendMarryDingQingSpouseDealMsg(ctx context.Context, result interface{}, err error) error {
	info := result.(*marryDingqingSuoYaoDeal)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(info.SuitId, info.PosId)
	if item == nil {
		return nil
	}
	//TODO:cjy 异步有可能已经自己镶嵌了
	//TODO:cjy 有可能已经有定情信物
	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)
	if marryManager.ExistsDingQing(info.SuitId, info.PosId) { //如果存在
		xinWuMap := make(map[int32]int32)
		xinWuMap[item.GetItemId()] = 1
		existsTitle := lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuSuccess)
		//TODO:cjy 信物名字是空的
		existsContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuSuccessContent), item.GetXinWuName(), item.Attack)
		emaillogic.AddEmail(pl, existsTitle, existsContent, xinWuMap)
		return nil
	}

	marrylogic.AddPlayerDingQing(pl, info.SuitId, info.PosId)

	title := lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuSuccess)
	//TODO:cjy 信物名字是空的
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryXinWuDingQingSuoQuSuccessContent), item.GetXinWuName(), item.Attack)
	emaillogic.AddEmail(pl, title, content, nil)

	return nil
}

func newMarryDingqingSuoYaoDealPostData(suitId int32, posId int32) *marryDingqingSuoYaoDeal {
	rst := &marryDingqingSuoYaoDeal{
		SuitId: suitId,
		PosId:  posId,
	}
	return rst
}
