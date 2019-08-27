package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerbaby "fgame/fgame/game/baby/player"
	babytypes "fgame/fgame/game/baby/types"
)

func BuildSCBabyChangeName(babyId int64, newName string) *uipb.SCBabyChangeName {
	scMsg := &uipb.SCBabyChangeName{}
	scMsg.BabyId = &babyId
	scMsg.NewName = &newName

	return scMsg
}

func BuildSCBabyDongFang(isUseItem, isBron bool, returnItemMap map[int32]int32, baby *playerbaby.PlayerBabyObject) *uipb.SCBabyDongFang {
	scMsg := &uipb.SCBabyDongFang{}
	scMsg.IsUseItem = &isUseItem
	scMsg.IsBorn = &isBron
	if isBron {
		scMsg.BabyInfo = buildBabyInfo(baby)
	}
	for itemId, num := range returnItemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}

	return scMsg
}

func BuildSCBabyBornAccelerate() *uipb.SCBabyBornAccelerate {
	scMsg := &uipb.SCBabyBornAccelerate{}
	return scMsg
}

func BuildSCBabyBornNotice(baby *playerbaby.PlayerBabyObject) *uipb.SCBabyBornNotice {
	scMsg := &uipb.SCBabyBornNotice{}
	scMsg.BabyInfo = buildBabyInfo(baby)
	return scMsg
}

func BuildSCBabyEatTonic(addPro int32) *uipb.SCBabyEatTonic {
	scMsg := &uipb.SCBabyEatTonic{}
	scMsg.AddPro = &addPro
	return scMsg
}

func BuildSCBabyActivateSkill(babyId int64, talentList []*babytypes.TalentInfo) *uipb.SCBabyActivateSkill {
	scMsg := &uipb.SCBabyActivateSkill{}
	scMsg.BabyId = &babyId
	scMsg.TalentList = BuildTalentInfoList(talentList)

	return scMsg
}

func BuildSCBabyBornChaoSheng() *uipb.SCBabyBornChaoSheng {
	scMsg := &uipb.SCBabyBornChaoSheng{}
	return scMsg
}

func BuildSCBabyLearnUplevel(babyId int64, learnItemId, num int32, isAuto bool, baby *playerbaby.PlayerBabyObject) *uipb.SCBabyLearnUplevel {
	scMsg := &uipb.SCBabyLearnUplevel{}
	learnLevel := baby.GetLearnLevel()
	learnExp := baby.GetLearnExp()

	scMsg.BabyId = &babyId
	scMsg.ItemId = &learnItemId
	scMsg.Num = &num
	scMsg.IsAuto = &isAuto
	scMsg.LearnLevel = &learnLevel
	scMsg.LearnExp = &learnExp
	return scMsg
}

func BuildSCBabyInfo(pregnantInfo *playerbaby.PlayerPregnantObject, babyList []*playerbaby.PlayerBabyObject, slotListMap map[babytypes.ToySuitType][]*playerbaby.PlayerBabyToySlotObject) *uipb.SCBabyInfo {
	scMsg := &uipb.SCBabyInfo{}
	scMsg.DongfangInfo = buildDongFangInfo(pregnantInfo, int32(len(babyList)))
	scMsg.BabyInfoList = buildBabyInfoList(babyList)
	scMsg.ToySuitInfoList = buildToySuitInfoList(slotListMap)
	return scMsg
}

func BuildSCBabyRefreshSkill(babyId int64, talentList []*babytypes.TalentInfo) *uipb.SCBabyRefreshSkill {
	scMsg := &uipb.SCBabyRefreshSkill{}
	scMsg.BabyId = &babyId
	scMsg.TalentList = BuildTalentInfoList(talentList)
	return scMsg
}

func BuildSCBabyLockSkill(babyId int64, skillIndex, operationType int32) *uipb.SCBabyLockSkill {
	scMsg := &uipb.SCBabyLockSkill{}
	scMsg.SkillIndex = &skillIndex
	scMsg.BabyId = &babyId
	scMsg.Operation = &operationType

	return scMsg
}

func BuildSCBabyEquipToy(index, suitType int32) *uipb.SCBabyEquipToy {
	scMsg := &uipb.SCBabyEquipToy{}
	scMsg.Index = &index
	scMsg.SuitType = &suitType
	return scMsg
}

func BuildSCBabyToyUplevel(suitType, pos, level, result int32) *uipb.SCBabyToyUplevel {
	scMsg := &uipb.SCBabyToyUplevel{}
	scMsg.SlotId = &pos
	scMsg.Level = &level
	scMsg.SuitType = &suitType
	scMsg.Result = &result
	return scMsg
}

func BuildSCBabyZhuanShi(babyId int64, itemMap map[int32]int32, babyNum int32) *uipb.SCBabyZhuanShi {
	scMsg := &uipb.SCBabyZhuanShi{}
	scMsg.BabyId = &babyId
	scMsg.BabyNum = &babyNum

	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	return scMsg
}

func BuildSCBabyToySlotChanged(slotListMap map[babytypes.ToySuitType][]*playerbaby.PlayerBabyToySlotObject) *uipb.SCBabyToySlotChanged {
	scMsg := &uipb.SCBabyToySlotChanged{}
	scMsg.InfoList = buildToySuitInfoList(slotListMap)
	return scMsg
}

func BuildSCBabyBornSpouseNotice(emailId int64) *uipb.SCBabyBornSpouseNotice {
	scMsg := &uipb.SCBabyBornSpouseNotice{}
	scMsg.EmailId = &emailId
	return scMsg
}

func BuildSCBabyBornMessageNotice(bornTime, noticeTime int64) *uipb.SCBabyBornMessageNotice {
	scMsg := &uipb.SCBabyBornMessageNotice{}
	scMsg.BornTime = &bornTime
	scMsg.NoticeTime = &noticeTime
	return scMsg
}

func BuildSCBabyPowerNotice(power int64) *uipb.SCBabyPowerNotice {
	scBabyPowerNotice := &uipb.SCBabyPowerNotice{}
	scBabyPowerNotice.Power = &power
	return scBabyPowerNotice
}

func buildBabyInfoList(babyList []*playerbaby.PlayerBabyObject) (babyInfoList []*uipb.BabyInfo) {
	for _, baby := range babyList {
		babyInfoList = append(babyInfoList, buildBabyInfo(baby))
	}
	return
}

func buildBabyInfo(baby *playerbaby.PlayerBabyObject) *uipb.BabyInfo {
	lockTimes := baby.GetLockTimes()
	activateTimes := baby.GetActivateTimes()
	refreshTimes := baby.GetRefreshTimes()
	babyName := baby.GetBabyName()
	sexInt := int32(baby.GetBabySex())
	danbei := baby.GetDanBei()
	babyId := baby.GetDBId()
	learnLevel := baby.GetLearnLevel()
	learnExp := baby.GetLearnExp()
	quality := baby.GetQuality()

	info := &uipb.BabyInfo{}
	info.ActivateTimes = &activateTimes
	info.BabyName = &babyName
	info.Sex = &sexInt
	info.Danbei = &danbei
	info.Quality = &quality
	info.BabyId = &babyId
	info.LockTims = &lockTimes
	info.RefreshTimes = &refreshTimes
	info.LearnLevel = &learnLevel
	info.LearnExp = &learnExp
	info.TalentList = BuildTalentInfoList(baby.GetSkillList())

	return info
}

func buildDongFangInfo(pregnantInfo *playerbaby.PlayerPregnantObject, babyNum int32) *uipb.DongFangInfo {
	info := &uipb.DongFangInfo{}
	time := pregnantInfo.GetPregnantTime()
	tonic := pregnantInfo.GetTonicPro()
	chaosheng := pregnantInfo.GetChaoShengNum()

	info.PregnantTime = &time
	info.BabyNum = &babyNum
	info.TonicPro = &tonic
	info.ChaoshengNum = &chaosheng
	return info
}

func BuildDongFangInfo(pregnantInfo *babytypes.PregnantInfo) *uipb.DongFangInfo {
	info := &uipb.DongFangInfo{}
	time := pregnantInfo.PregnantTime
	tonic := pregnantInfo.TonicPro

	info.PregnantTime = &time
	info.TonicPro = &tonic
	return info
}

func buildToySuitInfoList(slotListMap map[babytypes.ToySuitType][]*playerbaby.PlayerBabyToySlotObject) (suitList []*uipb.ToySuitInfo) {
	for suitType, slotList := range slotListMap {
		typeInt := int32(suitType)
		info := &uipb.ToySuitInfo{}
		info.SuitType = &typeInt
		info.SlotInfo = buildToySlotList(slotList)

		suitList = append(suitList, info)
	}
	return suitList
}

func buildToySlotList(slotList []*playerbaby.PlayerBabyToySlotObject) (slotItemList []*uipb.ToySlotInfo) {
	for _, slot := range slotList {
		slotItemList = append(slotItemList, buildToySlotInfo(slot))
	}
	return slotItemList
}

func buildToySlotInfo(slot *playerbaby.PlayerBabyToySlotObject) *uipb.ToySlotInfo {
	slotItem := &uipb.ToySlotInfo{}
	slotId := int32(slot.GetSlotId())
	level := slot.GetLevel()
	itemId := slot.GetItemId()
	slotItem.SlotId = &slotId
	slotItem.Level = &level
	slotItem.ItemId = &itemId
	return slotItem
}

func BuildTalentInfoList(talentInfoList []*babytypes.TalentInfo) (talentList []*uipb.TalentInfo) {
	for _, talent := range talentInfoList {
		skillId := talent.SkillId
		status := int32(talent.Status)

		info := &uipb.TalentInfo{}
		info.SkillId = &skillId
		info.Status = &status

		talentList = append(talentList, info)
	}
	return talentList
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}
