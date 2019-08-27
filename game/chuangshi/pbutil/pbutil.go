package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	chuangshidata "fgame/fgame/game/chuangshi/data"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	droppbutil "fgame/fgame/game/drop/pbutil"
)

func BuildSCChuangShiYuGaoInfo(isJoin bool, num int64) *uipb.SCChuangShiYuGao {
	scMsg := &uipb.SCChuangShiYuGao{}
	join := int32(0)
	if isJoin {
		join = 1
	}

	scMsg.IsJoin = &join
	scMsg.PeopleNum = &num
	return scMsg
}

func BuildSCBaoMingChuangShi(itemMap map[int32]int32) *uipb.SCBaoMingChuangShi {
	scMsg := &uipb.SCBaoMingChuangShi{}
	scMsg.DropInfo = droppbutil.BuildSimpleDropInfoList(itemMap)
	return scMsg
}

func BuildSCChuangShiPlayerInfoNotice(obj *playerchuangshi.PlayerChuangShiObject, guanZhiObj *playerchuangshi.PlayerChuangShiGuanZhiObject, signStatus, voteStatus int32) *uipb.SCChuangShiPlayerInfoNotice {
	scMsg := &uipb.SCChuangShiPlayerInfoNotice{}
	scMsg.PlayerInfo = buildPlayerChuangShiInfo(obj, guanZhiObj, signStatus, voteStatus)
	return scMsg
}

//
func BuildSCChuangShiShenWangBaoMing(status int32) *uipb.SCChuangShiShenWangBaoMing {
	scMsg := &uipb.SCChuangShiShenWangBaoMing{}
	scMsg.Status = &status
	return scMsg
}

func BuildSCChuangShiShengWangBaoMingList(signList []*chuangshidata.MemberInfo) *uipb.SCChuangShiShengWangBaoMingList {
	scMsg := &uipb.SCChuangShiShengWangBaoMingList{}
	for _, mem := range signList {
		scMsg.MemberInfoList = append(scMsg.MemberInfoList, buildChuangShiMemberInfo(mem))
	}

	return scMsg
}

func BuildSCChuangShiShengWangTouPiao(status int32, supportId int64) *uipb.SCChuangShiShengWangTouPiao {
	scMsg := &uipb.SCChuangShiShengWangTouPiao{}
	scMsg.Status = &status
	scMsg.SupportId = &supportId
	return scMsg
}

func BuildSCChuangShiInfo(campList []*chuangshidata.CampData, obj *playerchuangshi.PlayerChuangShiObject, guanZhiObj *playerchuangshi.PlayerChuangShiGuanZhiObject, signStatus, voteStatus int32) *uipb.SCChuangShiInfo {
	scMsg := &uipb.SCChuangShiInfo{}
	for _, camp := range campList {
		scMsg.CampList = append(scMsg.CampList, buildChuangShiCamp(camp))
	}
	scMsg.PlayerInfo = buildPlayerChuangShiInfo(obj, guanZhiObj, signStatus, voteStatus)
	return scMsg
}

func buildPlayerChuangShiInfo(obj *playerchuangshi.PlayerChuangShiObject, guanZhiObj *playerchuangshi.PlayerChuangShiGuanZhiObject, signStatus, voteStatus int32) *uipb.PlayerChuangShiInfo {
	typeInt := int32(obj.GetCampType())
	pos := int32(obj.GetPos())
	jifen := obj.GetJifen()
	diamonds := obj.GetDiamonds()
	weiWang := obj.GetWeiWang()

	info := &uipb.PlayerChuangShiInfo{}
	info.CampType = &typeInt
	info.Jifen = &jifen
	info.Diamonds = &diamonds
	info.Pos = &pos
	info.WeiWang = &weiWang
	info.SignStatus = &signStatus
	info.VoteStatus = &voteStatus
	info.GuanZhiInfo = buildPlayerGuanZhiInfo(guanZhiObj)

	return info
}

func buildPlayerGuanZhiInfo(guanZhiObj *playerchuangshi.PlayerChuangShiGuanZhiObject) *uipb.PlayerGuanZhiInfo {
	level := guanZhiObj.GetLevel()
	weiWang := guanZhiObj.GetWeiWang()
	rewLevel := guanZhiObj.GetRewLevel()

	info := &uipb.PlayerGuanZhiInfo{}
	info.Level = &level
	info.WeiWang = &weiWang
	info.ReceiveRewLevel = &rewLevel
	return info
}

func buildChuangShiCamp(camp *chuangshidata.CampData) *uipb.ChuangShiCamp {
	typeInt := int32(camp.CampType)
	info := &uipb.ChuangShiCamp{}
	info.Camp = &typeInt
	info.CityList = buildChuangShiCityList(camp.CityList)
	info.Power = &camp.Force
	info.Jifen = &camp.Jifen
	info.Diamonds = &camp.Diamonds
	if camp.KingMem != nil {
		info.KingMember = buildChuangShiMemberInfo(camp.KingMem)
	}
	return info
}

func BuildSCChuangShiShengWangTouPiaoList(voteList []*chuangshidata.VoteInfo) *uipb.SCChuangShiShengWangTouPiaoList {
	scMsg := &uipb.SCChuangShiShengWangTouPiaoList{}
	scMsg.MemberInfoList = buildChuangShiTouPiaoMemberInfoList(voteList)
	return scMsg
}

func BuildSCChuangShiCityRenMing() *uipb.SCChuangShiCityRenMing {
	scMsg := &uipb.SCChuangShiCityRenMing{}
	return scMsg
}

func BuildSCChuangShiCityPaySchedule() *uipb.SCChuangShiCityPaySchedule {
	scMsg := &uipb.SCChuangShiCityPaySchedule{}
	return scMsg
}

func BuildSCChuangShiCampPaySchedule() *uipb.SCChuangShiCampPaySchedule {
	scMsg := &uipb.SCChuangShiCampPaySchedule{}
	return scMsg
}

func BuildSCChuangShiMyPayReceive() *uipb.SCChuangShiMyPayReceive {
	scMsg := &uipb.SCChuangShiMyPayReceive{}
	return scMsg
}

func BuildSCChuangShiCampPayReceive() *uipb.SCChuangShiCampPayReceive {
	scMsg := &uipb.SCChuangShiCampPayReceive{}
	return scMsg
}

func BuildSCChuangShiCityJianShe() *uipb.SCChuangShiCityJianShe {
	scMsg := &uipb.SCChuangShiCityJianShe{}
	return scMsg
}

func BuildSCChuangShiEnterCity(cityId int64) *uipb.SCChuangShiEnterCity {
	scMsg := &uipb.SCChuangShiEnterCity{}
	scMsg.CityId = &cityId
	return scMsg
}

func BuildSCChuangShiPositionAdvance(success bool, level int32, weiWang int32) *uipb.SCChuangShiPositionAdvance {
	scMsg := &uipb.SCChuangShiPositionAdvance{}
	scMsg.Success = &success
	if !success {
		level--
	}
	scMsg.Level = &level
	scMsg.WeiWang = &weiWang
	return scMsg
}

func BuildSCChuangShiGuanZhiRew(rewLevel int32) *uipb.SCChuangShiGuanZhiRew {
	scMsg := &uipb.SCChuangShiGuanZhiRew{}
	scMsg.RewLevel = &rewLevel
	return scMsg
}

func BuildSCChuangShiGongChengTarget() *uipb.SCChuangShiGongChengTarget {
	scMsg := &uipb.SCChuangShiGongChengTarget{}
	return scMsg
}

func BuildSCChuangShiJoinCamp(campType int32) *uipb.SCChuangShiJoinCamp {
	scMsg := &uipb.SCChuangShiJoinCamp{}
	scMsg.Camp = &campType
	return scMsg
}

func BuildSCChuangShiCityTianQiSet(cityId int64, level, campType, cityType, index int32) *uipb.SCChuangShiCityTianQiSet {
	scMsg := &uipb.SCChuangShiCityTianQiSet{}
	scMsg.CityId = &cityId
	scMsg.Level = &level
	scMsg.CampType = &campType
	scMsg.CityType = &cityType
	scMsg.Index = &index
	return scMsg
}

func BuildSCChuangShiShenWangBroadcast(mem *chuangshidata.MemberInfo) *uipb.SCChuangShiShenWangBroadcast {
	scMsg := &uipb.SCChuangShiShenWangBroadcast{}
	scMsg.MemInfo = buildChuangShiMemberInfo(mem)
	return scMsg
}

func buildChuangShiTouPiaoMemberInfoList(voteList []*chuangshidata.VoteInfo) (infoList []*uipb.ChuangShiTouPiaoMemberInfo) {
	for _, vote := range voteList {
		info := &uipb.ChuangShiTouPiaoMemberInfo{}
		info.TicketNum = &vote.TicketNum
		info.MemberInfo = buildChuangShiMemberInfo(vote.Member)

		infoList = append(infoList, info)
	}

	return infoList
}

func buildChuangShiCityList(cityList []*chuangshidata.CityInfo) (infoList []*uipb.ChuangShiCity) {
	for _, city := range cityList {
		infoList = append(infoList, buildChuangShiCity(city))
	}
	return infoList
}

func buildChuangShiCity(city *chuangshidata.CityInfo) *uipb.ChuangShiCity {
	campType := int32(city.Camp)
	orignalCampType := int32(city.OrignalCamp)
	cityType := int32(city.CityType)

	info := &uipb.ChuangShiCity{}
	info.Camp = &campType
	info.OrignalCamp = &orignalCampType
	info.CityType = &cityType
	info.Index = &city.Index
	info.CityId = &city.CityId
	info.Jifen = &city.Jifen
	info.Diamonds = &city.Diamonds
	info.JianSheList = buildChuangShiJianSheList(city.JianSheList)
	if city.Member != nil {
		info.KingMember = buildChuangShiMemberInfo(city.Member)
	}
	return info
}

func buildChuangShiJianSheList(jianSheList []*chuangshidata.JianSheData) (infoList []*uipb.ChuangShiJianSheInfo) {
	for _, jianShe := range jianSheList {
		jianSheType := int32(jianShe.JianSheType)

		info := &uipb.ChuangShiJianSheInfo{}
		info.Type = &jianSheType
		info.Level = &jianShe.Level
		info.Exp = &jianShe.Exp
		info.SkillLevelSet = &jianShe.SkillLevelSet
		for skillLevel, activateTime := range jianShe.SkillMap {
			skill := &uipb.JianSheSkillInfo{}
			skill.Level = &skillLevel
			skill.ActivateTime = &activateTime

			info.SkillList = append(info.SkillList, skill)
		}

		infoList = append(infoList, info)
	}
	return infoList
}

func buildChuangShiMemberInfo(mem *chuangshidata.MemberInfo) *uipb.ChuangShiMemberInfo {
	info := &uipb.ChuangShiMemberInfo{}
	platform := mem.Platform
	serverId := mem.ServerId
	playerId := mem.PlayerId
	name := mem.PlayerName
	pos := int32(mem.Pos)
	allianceId := mem.AllianceId
	allianceName := mem.AllianceName
	alPos := int32(mem.AlPos)
	force := mem.Force
	campType := int32(mem.CampType)

	info.Platform = &platform
	info.ServerId = &serverId
	info.MemId = &playerId
	info.Name = &name
	info.Pos = &pos
	info.AllianceId = &allianceId
	info.AllianceName = &allianceName
	info.Force = &force
	info.AlPos = &alPos
	info.CampType = &campType
	return info
}
