package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	activitytypes "fgame/fgame/game/activity/types"
	alliancecommon "fgame/fgame/game/alliance/common"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	buffcommon "fgame/fgame/game/buff/common"
	densewatcommon "fgame/fgame/game/densewat/common"
	jieyicommon "fgame/fgame/game/jieyi/common"
	pkcommon "fgame/fgame/game/pk/common"
	pktypes "fgame/fgame/game/pk/types"
	playercommon "fgame/fgame/game/player/common"
	playertypes "fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	relivecommon "fgame/fgame/game/relive/common"
	"fgame/fgame/game/scene/scene"
	shenmocommon "fgame/fgame/game/shenmo/common"
	skillcommon "fgame/fgame/game/skill/common"
	teamcommon "fgame/fgame/game/team/common"
	teamtypes "fgame/fgame/game/team/types"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	xuechicommon "fgame/fgame/game/xuechi/common"
)

//基础数据
func ConvertFromPlayerBasicData(data *crosspb.PlayerBasicData) playercommon.PlayerCommonObject {
	id := data.GetPlayerId()
	userId := data.GetUserId()
	name := data.GetName()
	roleType := playertypes.RoleType(data.GetRole())
	sexType := playertypes.SexType(data.GetSex())
	serverId := data.GetServerId()
	guaJi := data.GetGuaJi()
	platform := data.GetPlatform()
	obj := playercommon.NewBasicPlayerCommonObject(
		id,
		userId,
		serverId,
		name,
		roleType,
		sexType,
		guaJi,
		platform,
	)
	return obj
}

//展示数据
func ConvertFromPlayerShowData(data *crosspb.PlayerShowData) *battle.PlayerShowObject {
	fashionId := data.GetClothesId()
	mountId := data.GetRideId()
	titleId := data.GetTitleId()
	weaponId := data.GetWeaponId()
	wingId := data.GetWingId()
	weaponState := data.GetWeaponState()
	shenFaId := data.GetShenFaId()
	lingYuId := data.GetLingYuId()
	fourGodKey := data.GetFourGodKey()
	realm := data.GetRealm()
	spouse := data.GetSpouse()
	spouseId := int64(0)
	weddingStatus := data.GetWeddingStatus()
	ringType := data.GetRingType()
	ringLevel := int32(0)
	faBaoId := data.GetFaBaoId()
	petId := data.GetPetId()
	xianTiId := data.GetXianTiId()
	baGua := data.GetBaGua()
	mountHidden := data.GetMountHidden()
	mountAdvanceId := data.GetMountAdvanceId()
	developLevel := int32(0)
	shenYuKey := int32(0)
	obj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		0,
		developLevel,
		shenYuKey,
	)

	return obj
}

//buff数据
func ConvertFromBuffDataList(dataList []*crosspb.BuffData) (objList []buffcommon.BuffObject) {
	for _, data := range dataList {
		obj := ConvertFromBuffData(data)
		objList = append(objList, obj)
	}
	return objList
}

//buff数据
func ConvertFromBuffData(data *crosspb.BuffData) buffcommon.BuffObject {
	ownerId := data.GetOwnerId()
	buffId := data.GetBuffId()
	groupId := data.GetGroupId()
	startTime := data.GetStartTime()
	useTime := data.GetUseTime()
	culTime := data.GetCulTime()
	lastTouchTime := data.GetLastTouchTime()
	duration := data.GetDuration()

	obj := buffcommon.NewBuffObject(
		ownerId,
		buffId,
		groupId,
		startTime,
		useTime,
		culTime,
		lastTouchTime,
		duration,
		nil,
	)
	return obj
}

//技能数据
func ConvertFromSkillDataList(dataList []*crosspb.SkillData) (objList []skillcommon.SkillObject) {
	for _, data := range dataList {
		obj := ConvertFromSkillData(data)
		objList = append(objList, obj)
	}
	return objList
}

//技能数据
//buff数据
func ConvertFromSkillData(data *crosspb.SkillData) skillcommon.SkillObject {
	level := data.GetLevel()
	skillId := data.GetSkillId()
	tianFuList := convertFromSkillTianFu(data.TianFuList)
	obj := skillcommon.CreateSkillObject(
		skillId,
		level,
		tianFuList,
	)
	return obj
}

func convertFromSkillTianFu(dataList []*crosspb.SkillTianFu) (tianFuList []*skillcommon.TianFuInfo) {
	if len(dataList) == 0 {
		return
	}
	tianFuList = make([]*skillcommon.TianFuInfo, 0, 3)
	for _, data := range dataList {
		tianFuInfo := &skillcommon.TianFuInfo{
			TianFuId: data.GetTianFuId(),
			Level:    data.GetLevel(),
		}
		tianFuList = append(tianFuList, tianFuInfo)
	}
	return
}

//pk数据
func ConvertFromPkData(data *crosspb.PkData) pkcommon.PlayerPkObject {
	pkValue := data.GetPkValue()
	state := pktypes.PkStatePeach
	camp := pktypes.PkCommonCampDefault
	obj := pkcommon.NewPlayerPkObject(state, camp, pkValue)
	return obj
}

//战斗属性
func ConvertFromBattleProperty(propertyData *crosspb.PropertyData) map[int32]int64 {

	ps := make(map[int32]int64)

	for _, property := range propertyData.GetPropertyList() {
		typ := propertytypes.BattlePropertyType(property.GetKey())
		if !typ.IsValid() {
			//TODO 警告
			continue
		}
		val := property.GetValue()
		ps[property.GetKey()] = val
	}
	return ps
}

//基本属性
func ConvertFromBaseProperty(propertyData *crosspb.PropertyData) map[int32]int64 {
	ps := make(map[int32]int64)
	for _, property := range propertyData.GetPropertyList() {
		typ := propertytypes.BasePropertyType(property.GetKey())
		if !typ.IsValid() {
			//TODO 警告
			continue
		}
		val := property.GetValue()
		ps[property.GetKey()] = val
	}
	return ps
}

func ConvertFromAllianceData(allianceData *crosspb.AllianceData) alliancecommon.PlayerAllianceObject {
	allianceId := allianceData.GetAllianceId()
	allianceName := allianceData.GetAllianceName()
	mengZhuId := allianceData.GetMengZhuId()
	memPos := alliancetypes.AlliancePosition(allianceData.GetMemPos())
	obj := alliancecommon.CreatePlayerAllianceObject(allianceId, allianceName, mengZhuId, memPos)
	return obj
}

func ConvertFromJieYiData(jieYiData *crosspb.JieYiData) jieyicommon.PlayerJieYiObject {
	jieYiName := jieYiData.GetJieYiName()
	rank := jieYiData.GetRank()
	jieYiId := jieYiData.GetJieYiId()
	obj := jieyicommon.CreatePlayerJieYiObject(jieYiId, jieYiName, rank)
	return obj
}

// func ConvertCommonPlayerChuangShiObject(chuangShiData *crosspb.ChuangShiData) chuangshidata.CommonPlayerChuangShiObject {
// 	campType := chuangshitypes.ChuangShiCampType(chuangShiData.GetCampType())
// 	pos := chuangshitypes.ChuangShiGuanZhi(chuangShiData.GetPos())
// 	obj := chuangshidata.CreatePlayerChuangShiObjetc(pos, campType)
// 	return obj
// }

func ConvertFromTeamData(teamData *crosspb.TeamData) teamcommon.PlayerTeamObject {
	teamId := teamData.GetTeamId()
	teamName := teamData.GetTeamName()
	teamPurpose := teamtypes.TeamPurposeType(teamData.GetTeamPurpose())
	obj := teamcommon.CreatePlayerTeamObject(teamId, teamName, teamPurpose)
	return obj
}

func ConvertFromArenaData(arenaData *crosspb.PlayerArenaData) *battle.PlayerArenaObject {
	reliveTime := arenaData.GetReliveTime()
	winTime := arenaData.GetWinTime()
	obj := battle.CreatePlayerArenaObject(reliveTime, winTime)
	return obj
}

func ConvertFromArenapvpData(arenapvpData *crosspb.PlayerArenapvpData) *battle.PlayerArenapvpObject {
	reliveTimes := arenapvpData.GetReliveTimes()
	obj := battle.CreatePlayerArenapvpObject(reliveTimes)
	return obj
}

func ConvertFromXueChiData(xueChiData *crosspb.XueChiData) *xuechicommon.PlayerXueChiObject {
	blood := xueChiData.GetBlood()
	bloodLine := xueChiData.GetBloodLine()
	obj := xuechicommon.CreatePlayerXueChiObject(blood, bloodLine)
	return obj
}

func ConvertFromReliveData(reliveData *crosspb.PlayerReliveData) *relivecommon.PlayerReliveObject {
	culTime := reliveData.GetCulTime()
	lastReliveTime := reliveData.GetLastReliveTime()
	obj := relivecommon.CreatePlayerReliveObject(culTime, lastReliveTime)
	return obj
}

func ConvertFromBattleData(battleData *crosspb.PlayerBattleData) *battlecommon.PlayerBattleObject {
	vip := battleData.GetVip()
	level := battleData.GetLevel()
	zhuanSheng := battleData.GetZhuanSheng()
	soulAwakenNum := battleData.GetSoulAwakenNum()
	isHuiYuan := battleData.GetIsHuiYuan()
	obj := battlecommon.CreatePlayerBattleObject(vip, level, zhuanSheng, soulAwakenNum, isHuiYuan)
	return obj
}

func ConvertFromDenseWatData(denseWatData *crosspb.DenseWatData) *densewatcommon.PlayerDenseWatObject {
	num := denseWatData.GetNum()
	endTime := denseWatData.GetEndTime()
	obj := densewatcommon.CreatePlayerDenseWatObject(num, endTime)
	return obj
}

func ConvertFromActivityPkDataList(activityPkDataList []*crosspb.ActivityPkData) (killDataList []*scene.PlayerActvitiyKillData) {
	for _, activityPkData := range activityPkDataList {
		killData := ConvertFromActivityPkData(activityPkData)
		killDataList = append(killDataList, killData)
	}
	return
}

func ConvertFromActivityRankDataList(activityRankDataList []*crosspb.ActivityRankData) (rankDataList []*scene.PlayerActvitiyRankData) {
	for _, rankData := range activityRankDataList {
		playerRankData := ConvertFromActivityRankData(rankData)
		rankDataList = append(rankDataList, playerRankData)
	}
	return
}

func ConvertFromActivityPkData(activityPkData *crosspb.ActivityPkData) *scene.PlayerActvitiyKillData {
	obj := scene.CreatePlayerActvitiyKillData(activitytypes.ActivityType(activityPkData.GetActivityType()), activityPkData.GetKilledNum(), activityPkData.GetLastKillTime())
	return obj
}

func ConvertFromActivityRankData(activityRankData *crosspb.ActivityRankData) *scene.PlayerActvitiyRankData {
	activityType := activitytypes.ActivityType(activityRankData.GetActivityType())
	endTime := activityRankData.GetEndTime()
	rankMap := make(map[int32]int64)
	for _, randData := range activityRankData.PlayerRankDataList {
		rankMap[randData.GetRankType()] = randData.GetVal()
	}
	obj := scene.CreatePlayerActvitiyRankData(activityType, rankMap, endTime)
	return obj
}

func ConvertFromShenMoData(shenMoData *crosspb.ShenMoData) *shenmocommon.PlayerShenMoObject {
	gongXunNum := shenMoData.GetGongXunNum()
	killNum := shenMoData.GetKillNum()
	endTime := shenMoData.GetEndTime()

	obj := shenmocommon.CreatePlayerShenMoObject(gongXunNum, killNum, endTime)
	return obj
}

func ConvertFromBossReliveDataList(reliveDataPbList []*crosspb.PlayerBossReliveData) (reliveDataList []*scene.PlayerBossReliveData) {
	for _, reliveDataPb := range reliveDataPbList {
		reliveData := ConvertFromBossReliveData(reliveDataPb)
		reliveDataList = append(reliveDataList, reliveData)
	}
	return
}
func ConvertFromBossReliveData(reliveData *crosspb.PlayerBossReliveData) *scene.PlayerBossReliveData {
	bossType := reliveData.GetBossType()
	reliveTime := reliveData.GetReliveTime()

	obj := scene.CreatePlayerBossReliveData(worldbosstypes.BossType(bossType), reliveTime)
	return obj
}

func ConvertFromTeShuSkillDataList(teShuSkillPbList []*crosspb.TeShuSkillData) (teShuSkillList []*scene.TeshuSkillObject) {
	for _, teShuSkillPb := range teShuSkillPbList {
		teShuSkillData := ConvertFromTeShuSkillData(teShuSkillPb)
		teShuSkillList = append(teShuSkillList, teShuSkillData)
	}
	return
}
func ConvertFromTeShuSkillData(teShuSkillData *crosspb.TeShuSkillData) *scene.TeshuSkillObject {
	skillId := teShuSkillData.GetSkillId()
	chuFaRate := teShuSkillData.GetChuFaRate()
	diKangRate := teShuSkillData.GetDiKangRate()

	obj := scene.CreateTeshuSkillObject(skillId, chuFaRate, diKangRate)
	return obj
}
