package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	droppbutil "fgame/fgame/game/drop/pbutil"
	propertypbutil "fgame/fgame/game/property/pbutil"
	propertytypes "fgame/fgame/game/property/types"
	playerquest "fgame/fgame/game/quest/player"
)

func BuildSCQuestList(questList []*playerquest.PlayerQuestObject, finish bool) *uipb.SCQuestList {
	scQuestList := &uipb.SCQuestList{}
	scQuestList.QuestList = buildQuestList(questList)
	scQuestList.Finish = &finish
	return scQuestList
}

func BuildSCQuestAccept(questId int32) *uipb.SCQuestAccept {
	scQuestAccept := &uipb.SCQuestAccept{}
	scQuestAccept.QuestId = &questId
	return scQuestAccept
}

func BuildSCQuestCommit(questId int32, isDouble bool) *uipb.SCQuestCommit {
	scQuestCommit := &uipb.SCQuestCommit{}
	scQuestCommit.QuestId = &questId
	scQuestCommit.Double = &isDouble
	return scQuestCommit
}

func BuildSCQuestGather(questId int32, itemId int32, num int32) *uipb.SCQuestGather {
	scQuestGather := &uipb.SCQuestGather{}
	scQuestGather.QuestId = &questId
	scQuestGather.ItemId = &itemId
	scQuestGather.Num = &num
	return scQuestGather
}

func BuildSCQuestNPCDialog(questId int32, npcId int32) *uipb.SCQuestNPCDialog {
	scQuestNPCDialog := &uipb.SCQuestNPCDialog{}
	scQuestNPCDialog.QuestId = &questId
	scQuestNPCDialog.NpcId = &npcId
	return scQuestNPCDialog
}

func BuildSCQuestUpdate(quest *playerquest.PlayerQuestObject) *uipb.SCQuestUpdate {
	scQuestUpdate := &uipb.SCQuestUpdate{}
	questInfo := buildQuest(quest)
	scQuestUpdate.QuestList = []*uipb.QuestInfo{questInfo}

	return scQuestUpdate
}

func BuildSCQuestListUpdate(questList []*playerquest.PlayerQuestObject) *uipb.SCQuestUpdate {
	scQuestUpdate := &uipb.SCQuestUpdate{}

	scQuestUpdate.QuestList = buildQuestList(questList)
	return scQuestUpdate
}

func BuildSCQuestFeiXie(questId int32, mapId int32, pos types.Position) *uipb.SCQuestFeiXie {
	scQuestFeiXie := &uipb.SCQuestFeiXie{}
	scQuestFeiXie.QuestId = &questId
	scQuestFeiXie.MapId = &mapId
	scQuestFeiXie.Pos = buildQuestPos(pos)
	return scQuestFeiXie
}

func BuildSCQuestTuMoUseToken(token int32, questId int32) *uipb.SCQuestTuMoUseToken {
	scQuestTuMoUseToken := &uipb.SCQuestTuMoUseToken{}
	scQuestTuMoUseToken.Token = &token
	scQuestTuMoUseToken.QuestId = &questId
	return scQuestTuMoUseToken
}

func BuildSCQuestTuMoNumGet(num int32, buyNum int32, extraNum int32) *uipb.SCQuestTuMoNumGet {
	scQuestTuMoNumGet := &uipb.SCQuestTuMoNumGet{}
	scQuestTuMoNumGet.Num = &num
	scQuestTuMoNumGet.BuyNum = &buyNum
	scQuestTuMoNumGet.ExtraNum = &extraNum
	return scQuestTuMoNumGet
}

func BuildSCQuestTuMoDiscard(questId int32) *uipb.SCQuestTuMoDiscard {
	scQuestTuMoDiscard := &uipb.SCQuestTuMoDiscard{}
	scQuestTuMoDiscard.QuestId = &questId
	return scQuestTuMoDiscard
}

func BuildSCQuestTuMoBuyNum(buyNum int32, extraNum int32) *uipb.SCQuestTuMoBuyNum {
	scQuestTuMoBuyNum := &uipb.SCQuestTuMoBuyNum{}
	scQuestTuMoBuyNum.BuyNum = &buyNum
	scQuestTuMoBuyNum.ExtraNum = &extraNum
	return scQuestTuMoBuyNum
}

func BuildSCQuestTuMoFinishAll(itemMapList []map[int32]int32, num int32) *uipb.SCQuestTuMoFinishAll {
	scQuestTuMoFinishAll := &uipb.SCQuestTuMoFinishAll{}

	scQuestTuMoFinishAll.Num = &num
	for _, itemMap := range itemMapList {
		for itemId, num := range itemMap {
			scQuestTuMoFinishAll.ItemList = append(scQuestTuMoFinishAll.ItemList, buildItem(itemId, num))
		}
	}
	return scQuestTuMoFinishAll
}

func BuildSCQuestTuMoImmediateFinish(questId int32, itemMap map[int32]int32) *uipb.SCQuestTuMoImmediate {
	scQuestTuMoImmediate := &uipb.SCQuestTuMoImmediate{}
	scQuestTuMoImmediate.QuestId = &questId
	for itemId, num := range itemMap {
		scQuestTuMoImmediate.ItemList = append(scQuestTuMoImmediate.ItemList, buildItem(itemId, num))
	}
	return scQuestTuMoImmediate
}

func BuildSCQuestDailyFinishAll(dailyTag int32, itemMap map[int32]int32) *uipb.SCQuestDailyFinishAll {
	scQuestDailyFinishAll := &uipb.SCQuestDailyFinishAll{}
	scQuestDailyFinishAll.DailyTag = &dailyTag
	for itemId, num := range itemMap {
		scQuestDailyFinishAll.ItemList = append(scQuestDailyFinishAll.ItemList, buildItem(itemId, num))
	}
	return scQuestDailyFinishAll
}

func BuildSCQuestDailyFinishOnce(dailyTag int32, itemMap map[int32]int32) *uipb.SCQuestDailyFinishOnce {
	scMsg := &uipb.SCQuestDailyFinishOnce{}
	scMsg.DailyTag = &dailyTag
	for itemId, num := range itemMap {
		scMsg.ItemList = append(scMsg.ItemList, buildItem(itemId, num))
	}
	return scMsg
}

func BuildSCQuestDailySeq(dailyTag int32, seqId int32, times int32) *uipb.SCQuestDailySeq {
	scQuestDailySeq := &uipb.SCQuestDailySeq{}
	scQuestDailySeq.DailyTag = &dailyTag
	scQuestDailySeq.SeqId = &seqId
	scQuestDailySeq.Times = &times
	return scQuestDailySeq
}

func BuildSCQuestKaiFuMuBiaoGet(kaiFuMuBiaoMap map[int32]*playerquest.PlayerKaiFuMuBiaoObject) *uipb.SCQuestKaiFuMuBiaoGet {
	scQuestKaiFuMuBiaoGet := &uipb.SCQuestKaiFuMuBiaoGet{}
	for _, kaiFuMuBiaoObj := range kaiFuMuBiaoMap {
		scQuestKaiFuMuBiaoGet.KaiFuMuBiaoList = append(scQuestKaiFuMuBiaoGet.KaiFuMuBiaoList, buildKaiFuMuBiao(kaiFuMuBiaoObj))
	}
	return scQuestKaiFuMuBiaoGet
}

func BuildSCQuestKaiFuMuBiaoReceive(kaiFuTime int32) *uipb.SCQuestkaiFuMuBiaoReceive {
	scQuestkaiFuMuBiaoReceive := &uipb.SCQuestkaiFuMuBiaoReceive{}
	scQuestkaiFuMuBiaoReceive.KaiFuTime = &kaiFuTime
	return scQuestkaiFuMuBiaoReceive
}

func BuildSCQuestKaiFuMuBiaoFinishNumChanged(kaiFuMuBiaoMap map[int32]*playerquest.PlayerKaiFuMuBiaoObject, kaiFuDayList []int32) *uipb.SCQuestKaiFuMuBiaoGet {
	scQuestKaiFuMuBiaoGet := &uipb.SCQuestKaiFuMuBiaoGet{}
	for kaiFuDay, kaiFuMuBiaoObj := range kaiFuMuBiaoMap {
		flag := utils.ContainInt32(kaiFuDayList, kaiFuDay)
		if !flag {
			continue
		}
		scQuestKaiFuMuBiaoGet.KaiFuMuBiaoList = append(scQuestKaiFuMuBiaoGet.KaiFuMuBiaoList, buildKaiFuMuBiao(kaiFuMuBiaoObj))
	}
	return scQuestKaiFuMuBiaoGet
}

func BuildSCQuestQiYuNotice(qiyu *playerquest.PlayerQiYuObject) *uipb.SCQuestQiYuNotice {
	scMsg := &uipb.SCQuestQiYuNotice{}
	scMsg.InfoList = append(scMsg.InfoList, buildQiYuInfo(qiyu))
	return scMsg
}

func BuildSCQuestQiYuNoticeAll(qiyuMap map[int32]*playerquest.PlayerQiYuObject) *uipb.SCQuestQiYuNotice {
	scMsg := &uipb.SCQuestQiYuNotice{}
	scMsg.InfoList = buildQiYuInfoList(qiyuMap)
	return scMsg
}

func BuildSCQuestQiYuReceive(qiyuId, isReceive int32) *uipb.SCQuestQiYuReceive {
	scMsg := &uipb.SCQuestQiYuReceive{}
	scMsg.QiyuId = &qiyuId
	scMsg.IsReceive = &isReceive
	return scMsg
}

func BuildSCQuestDailyCommitRew(questId, dailyTag int32, itemMap map[int32]int32, rd *propertytypes.RewData) *uipb.SCQuestDailyCommitRew {
	scMsg := &uipb.SCQuestDailyCommitRew{}
	scMsg.QuestId = &questId
	scMsg.DailyTag = &dailyTag
	scMsg.DropList = droppbutil.BuildSimpleDropInfoList(itemMap)
	scMsg.RewInfo = propertypbutil.BuildRewProperty(rd)

	return scMsg
}

func buildQiYuInfoList(objList map[int32]*playerquest.PlayerQiYuObject) (infoList []*uipb.QiYuInfo) {
	for _, qiyu := range objList {
		infoList = append(infoList, buildQiYuInfo(qiyu))
	}
	return infoList
}

func buildQiYuInfo(obj *playerquest.PlayerQiYuObject) *uipb.QiYuInfo {
	info := &uipb.QiYuInfo{}
	isReceive := obj.GetIsReceive()
	endTime := obj.GetEndTime()
	qiyuId := obj.GetQiYuId()
	isHadNotice := obj.GetIsHadNotice()

	info.IsReceive = &isReceive
	info.EndTime = &endTime
	info.QiyuId = &qiyuId
	info.IsHadNotice = &isHadNotice
	return info
}

func buildKaiFuMuBiao(obj *playerquest.PlayerKaiFuMuBiaoObject) *uipb.QuestKaiFuMuBiao {
	questKaiFuMuBiao := &uipb.QuestKaiFuMuBiao{}
	kaiFuDay := obj.GetKaiFuDay()
	finishNum := obj.GetFinishNum()
	isReceived := obj.IsGroupRewardGet()

	questKaiFuMuBiao.KaiFuTime = &kaiFuDay
	questKaiFuMuBiao.FinishNum = &finishNum
	questKaiFuMuBiao.GroupReceived = &isReceived
	return questKaiFuMuBiao
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func buildQuestList(questList []*playerquest.PlayerQuestObject) (questInfoList []*uipb.QuestInfo) {
	for _, quest := range questList {
		questInfoList = append(questInfoList, buildQuest(quest))
	}
	return
}

func buildQuest(quest *playerquest.PlayerQuestObject) *uipb.QuestInfo {
	questInfo := &uipb.QuestInfo{}
	questId := quest.QuestId
	questState := int32(quest.QuestState)
	questInfo.QuestId = &questId
	questInfo.State = &questState
	questInfo.DataList = buildQuestDataList(quest.QuestDataMap)
	questInfo.CollectItemList = buildQuestDataList(quest.CollectItemDataMap)
	return questInfo
}

func buildQuestDataList(dataMap map[int32]int32) (dataList []*uipb.QuestData) {
	for k, v := range dataMap {
		data := buildQuestData(k, v)
		dataList = append(dataList, data)
	}
	return
}
func buildQuestData(k, v int32) (data *uipb.QuestData) {
	data = &uipb.QuestData{
		Key:   &k,
		Value: &v,
	}
	return
}

func buildQuestPos(pos types.Position) *uipb.QuestPosition {
	targetPosition := &uipb.QuestPosition{}
	x := float32(pos.X)
	y := float32(pos.Y)
	z := float32(pos.Z)
	targetPosition.PosX = &x
	targetPosition.PosY = &y
	targetPosition.PosZ = &z

	return targetPosition
}
