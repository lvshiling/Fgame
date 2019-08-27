package api

import (
	"fgame/fgame/cross/chuangshi/chuangshi"
	chuangshipb "fgame/fgame/cross/chuangshi/pb"
	chuangshidata "fgame/fgame/game/chuangshi/data"
)

func BuildChuangShiCampList(campList []*chuangshi.Camp) (infoList []*chuangshipb.ChuangShiCamp) {
	for _, camp := range campList {
		infoList = append(infoList, BuildChuangShiCamp(camp))
	}
	return infoList
}

func BuildChuangShiCamp(camp *chuangshi.Camp) *chuangshipb.ChuangShiCamp {
	campType := int32(camp.GetCampObj().GetCampType())
	shenWangStatus := int32(camp.GetCampObj().GetShenWangStatus())
	info := &chuangshipb.ChuangShiCamp{}
	info.SignList = BuildChuangShiSignInfoList(camp.GetShenWangSignList())
	info.VoteList = BuildChuangShiVoteList(camp.GetShenWangVoteList())
	info.CampType = campType
	info.MemberList = BuildChuangShiMemberInfoList(camp.GetMemberList())
	info.CityList = buildChuangShiCityList(camp.GetCityList())
	info.ShenWangStatus = shenWangStatus
	return info
}

func BuildChuangShiVoteList(voteList []*chuangshi.ChuangShiVoteInfo) (infoList []*chuangshipb.ChuangShiVote) {
	for _, vote := range voteList {
		infoList = append(infoList, buildChuangShiVote(vote))
	}
	return infoList
}

func buildChuangShiVote(vote *chuangshi.ChuangShiVoteInfo) *chuangshipb.ChuangShiVote {
	info := &chuangshipb.ChuangShiVote{}
	info.TicketNum = vote.Vote.GetTicketNum()
	info.Member = buildChuangShiMemberInfo(vote.Member)
	return info
}

func buildChuangShiCityList(cityDataList []*chuangshi.CityData) (infoList []*chuangshipb.ChuangShiCity) {
	for _, cityData := range cityDataList {
		infoList = append(infoList, buildChuangShiCity(cityData))
	}
	return infoList
}

func buildChuangShiCity(cityData *chuangshi.CityData) *chuangshipb.ChuangShiCity {
	city := cityData.GetCity()
	jianSheList := cityData.GetChengFangJianSheList()

	campType := int32(city.GetCampType())
	orignalCampType := int32(city.GetOrignalCampType())
	typeInt := int32(city.GetType())
	index := city.GetIndex()
	cityId := city.GetId()
	diamonds := city.GetDiamonds()
	jifen := city.GetJifen()

	info := &chuangshipb.ChuangShiCity{}
	info.OrignalCamp = orignalCampType
	info.Camp = campType
	info.CityType = typeInt
	info.Index = index
	info.CityId = cityId
	info.Diamonds = diamonds
	info.Jifen = jifen
	info.JianSheList = buildChuangShiJianSheList(jianSheList)
	info.KingMember = buildChuangShiMemberInfo(city.GetCamp().GetMember(city.GeOwnerId()))
	return info
}

func BuildChuangShiSignInfoList(singList []*chuangshi.ChuangShiSignInfo) (infoList []*chuangshipb.ChuangShiMemberInfo) {
	for _, sign := range singList {
		infoList = append(infoList, buildChuangShiMemberInfo(sign.Member))
	}
	return infoList
}

func BuildChuangShiMemberInfoList(memList []*chuangshi.ChuangShiMemberObject) (infoList []*chuangshipb.ChuangShiMemberInfo) {
	for _, mem := range memList {
		infoList = append(infoList, buildChuangShiMemberInfo(mem))
	}
	return infoList
}

func buildChuangShiMemberInfo(obj *chuangshi.ChuangShiMemberObject) *chuangshipb.ChuangShiMemberInfo {
	info := &chuangshipb.ChuangShiMemberInfo{}
	if obj == nil {
		return info
	}

	platform := obj.GetPlatform()
	serverId := obj.GetServerId()
	playerId := obj.GetPlayerId()
	name := obj.GetPlayerName()
	pos := int32(obj.GetPos())
	allianceId := obj.GetAllianceId()
	allianceName := obj.GetAllianceName()
	alPos := int32(obj.GetAlPos())
	force := obj.GetForce()
	campType := int32(obj.GetCampType())

	info.Platform = platform
	info.ServerId = serverId
	info.PlayerId = playerId
	info.Name = name
	info.Pos = pos
	info.AllianceId = allianceId
	info.AllianceName = allianceName
	info.AlPos = alPos
	info.Force = force
	info.CampType = campType
	return info
}

func buildChuangShiJianSheList(jianSheList []*chuangshi.ChuangShiCityJianSheObject) (infoList []*chuangshipb.CityJianShe) {
	for _, jianShe := range jianSheList {

		info := &chuangshipb.CityJianShe{}

		info.JianSheType = int32(jianShe.GetJianSheType())
		info.Level = jianShe.GetJianSheLevel()
		info.Exp = jianShe.GetJianSheExp()
		info.SkillLevelSet = jianShe.GetSkillLevelSet()
		for skillLevel, activateTime := range jianShe.GetSkillMap() {
			skill := &chuangshipb.JianSheSkill{}
			skill.Level = skillLevel
			skill.ActivateTime = activateTime

			info.SkillMap = append(info.SkillMap, skill)
		}

		infoList = append(infoList, info)
	}
	return infoList
}

func BuildChuangShiMemberInfoListByData(memList []*chuangshidata.MemberInfo) (infoList []*chuangshipb.ChuangShiMemberInfo) {
	for _, mem := range memList {
		infoList = append(infoList, buildChuangShiMemberInfoByData(mem))
	}
	return infoList
}

func buildChuangShiMemberInfoByData(data *chuangshidata.MemberInfo) *chuangshipb.ChuangShiMemberInfo {
	info := &chuangshipb.ChuangShiMemberInfo{}
	platform := data.Platform
	serverId := data.ServerId
	playerId := data.PlayerId
	name := data.PlayerName
	pos := int32(data.Pos)
	allianceId := data.AllianceId
	allianceName := data.AllianceName
	alPos := int32(data.AlPos)
	force := data.Force
	campType := int32(data.CampType)

	info.Platform = platform
	info.ServerId = serverId
	info.PlayerId = playerId
	info.Name = name
	info.Pos = pos
	info.AllianceId = allianceId
	info.AllianceName = allianceName
	info.AlPos = alPos
	info.Force = force
	info.CampType = campType
	return info
}

func BuildCityPayScheduleList(paramList []*chuangshidata.CityPayScheduleParam) (infoList []*chuangshipb.CityPaySchedule) {
	for _, param := range paramList {
		d := &chuangshipb.CityPaySchedule{}
		d.AlPos = int32(param.AlPos)
		d.Ratio = param.Ratio

		infoList = append(infoList, d)
	}

	return
}

func BuildCampPayScheduleList(paramList []*chuangshidata.CamPayScheduleParam) (infoList []*chuangshipb.CampPaySchedule) {
	for _, param := range paramList {
		d := &chuangshipb.CampPaySchedule{}
		d.CityId = param.CityId
		d.Ratio = param.Ratio

		infoList = append(infoList, d)
	}

	return
}
