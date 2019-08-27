package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	activitytypes "fgame/fgame/game/activity/types"
	buffcommon "fgame/fgame/game/buff/common"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	skillcommon "fgame/fgame/game/skill/common"
)

func BuildSILogin(playerId int64) *crosspb.SILogin {
	siLogin := &crosspb.SILogin{}
	siLogin.PlayerId = &playerId
	return siLogin
}

var (
	//所有通用一个
	siHeartBeat = &crosspb.SIHeartBeat{}
)

func BuildSIHeartBeat() *crosspb.SIHeartBeat {
	return siHeartBeat
}

func BuildSIPlayerData(pl player.Player) *crosspb.SIPlayerData {
	siPlayerData := &crosspb.SIPlayerData{}
	siPlayerData.PlayerBasicData = BuildPlayerBasicData(pl)
	siPlayerData.PlayerShowData = BuildPlayerShowData(pl)
	siPlayerData.BasicPropertyData = BuildPlayerBasicPropertyData(pl)
	battleProperties := pl.GetAllSystemBattleProperties()
	siPlayerData.BattlePropertyData = BuildBattlePropertyData(battleProperties)
	siPlayerData.BuffList = BuildPlayerBuffDataList(pl)
	siPlayerData.SkillList = BuildPlayerSkillDataList(pl)
	siPlayerData.PkData = BuildPlayerPkData(pl)
	siPlayerData.TeamData = BuildPlayerTeamData(pl)
	siPlayerData.AllianceData = BuildPlayerAllianceData(pl)
	siPlayerData.CrossData = BuildPlayerCrossData(pl)
	siPlayerData.ArenaData = BuildPlayerArenaData(pl)
	siPlayerData.XueChiData = BuildPlayerXueChiData(pl)
	siPlayerData.ReliveData = BuildPlayerReliveData(pl)
	siPlayerData.BattleData = BuildPlayerBattleData(pl)
	siPlayerData.DenseWatData = BuildPlayerDenseWatData(pl)
	siPlayerData.ShenMoData = BuildPlayerShenMoData(pl)
	siPlayerData.ArenapvpData = BuildPlayerArenapvpData(pl)
	siPlayerData.JieYiData = BuildPlayerJieYiData(pl)
	siPlayerData.ChuangShiData = BuildPlayerChuangShiData(pl)
	siPlayerData.BossReliveDataList = BuildPlayerBossReliveDataList(pl)
	siPlayerData.TeShuSkillDataList = BuildPlayerTeShuSkillDataList(pl)

	power := pl.GetForce()
	siPlayerData.Power = &power
	lingTong := pl.GetLingTong()
	if lingTong != nil {
		siPlayerData.LingTongData = BuildLingTongData(lingTong)
	}
	siPlayerData.ActivityPkDataList = BuildPlayerActivityPkDataList(pl)
	siPlayerData.ActivityRankDataList = BuildPlayerActivityRankDataList(pl)
	return siPlayerData
}

func BuildLingTongData(lingTong scene.LingTong) *crosspb.LingTongData {
	lingTongData := &crosspb.LingTongData{}
	lingTongId := lingTong.GetLingTongId()
	lingTongData.LingTongId = &lingTongId
	name := lingTong.GetName()
	lingTongData.Name = &name
	lingTongData.BattlePropertyData = BuildBattlePropertyData(lingTong.GetAllSystemBattleProperties())
	lingTongData.LingTongShowData = BuildLingTongShowData(lingTong)
	return lingTongData
}

func BuildLingTongShowData(lingTong scene.LingTong) *crosspb.LingTongShowData {
	lingTongShowData := &crosspb.LingTongShowData{}
	fashionId := lingTong.GetLingTongFashionId()
	lingTongShowData.FashionId = &fashionId
	weaponId := lingTong.GetLingTongWeaponId()
	lingTongShowData.WeaponId = &weaponId
	weaponState := lingTong.GetLingTongWeaponState()
	lingTongShowData.WeaponState = &weaponState
	titleId := lingTong.GetLingTongTitleId()
	lingTongShowData.TitleId = &titleId
	wingId := lingTong.GetLingTongWingId()
	lingTongShowData.WingId = &wingId
	mountId := lingTong.GetLingTongMountId()
	lingTongShowData.MountId = &mountId
	shenFaId := lingTong.GetLingTongShenFaId()
	lingTongShowData.ShenFaId = &shenFaId
	lingYuId := lingTong.GetLingTongLingYuId()
	lingTongShowData.LingYuId = &lingYuId
	xianTiId := lingTong.GetLingTongXianTiId()
	lingTongShowData.XianTiId = &xianTiId
	return lingTongShowData
}

//玩家基础数据
func BuildPlayerBasicData(pl player.Player) *crosspb.PlayerBasicData {
	playerBasicData := &crosspb.PlayerBasicData{}
	userId := pl.GetUserId()
	playerId := pl.GetId()
	name := pl.GetOriginName()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	serverId := pl.GetServerId()
	platform := pl.GetPlatform()
	playerBasicData.Name = &name
	playerBasicData.UserId = &userId
	playerBasicData.PlayerId = &playerId
	playerBasicData.Role = &role
	playerBasicData.Sex = &sex
	playerBasicData.ServerId = &serverId
	guaJi := pl.IsGuaJiPlayer()
	playerBasicData.GuaJi = &guaJi
	playerBasicData.Platform = &platform
	return playerBasicData
}

//玩家展示数据
func BuildPlayerShowData(pl player.Player) *crosspb.PlayerShowData {
	playerShowData := &crosspb.PlayerShowData{}
	titleId := pl.GetTitleId()
	weaponId := pl.GetWeaponId()
	clothesId := pl.GetFashionId()
	rideId := pl.GetMountId()
	mountAdvanceId := pl.GetMountAdvanceId()
	mountHidden := pl.IsMountHidden()
	wingId := pl.GetWingId()
	shenFaId := pl.GetShenFaId()
	lingYuId := pl.GetLingYuId()
	spouse := pl.GetSpouse()
	realm := pl.GetRealm()
	baGua := pl.GetBaGua()
	weaponSate := pl.GetWeaponState()
	faBaoId := pl.GetFaBaoId()
	xianTiId := pl.GetXianTiId()
	petId := pl.GetPetId()
	playerShowData.TitleId = &titleId
	playerShowData.WeaponId = &weaponId
	playerShowData.WeaponState = &weaponSate
	playerShowData.ClothesId = &clothesId
	playerShowData.RideId = &rideId
	playerShowData.MountAdvanceId = &mountAdvanceId
	playerShowData.MountHidden = &mountHidden
	playerShowData.WingId = &wingId
	playerShowData.ShenFaId = &shenFaId
	playerShowData.LingYuId = &lingYuId
	playerShowData.Spouse = &spouse
	playerShowData.Realm = &realm
	playerShowData.FaBaoId = &faBaoId
	playerShowData.XianTiId = &xianTiId
	playerShowData.PetId = &petId
	playerShowData.BaGua = &baGua
	return playerShowData
}

//玩家属性
func BuildProperty(key int32, val int64) *crosspb.Property {
	prop := &crosspb.Property{}
	prop.Key = &key
	prop.Value = &val
	return prop
}

//玩家属性数据
func BuildPlayerBasicPropertyData(pl player.Player) *crosspb.PropertyData {
	propertyData := &crosspb.PropertyData{}
	for key, val := range pl.GetBaseProperties() {
		p := BuildProperty(key, val)
		propertyData.PropertyList = append(propertyData.PropertyList, p)
	}
	return propertyData
}

//玩家战斗属性
func BuildBattlePropertyData(battleProperties map[int32]int64) *crosspb.PropertyData {
	propertyData := &crosspb.PropertyData{}
	for key, val := range battleProperties {
		p := BuildProperty(key, val)
		propertyData.PropertyList = append(propertyData.PropertyList, p)
	}
	return propertyData
}

//玩家技能数据
func BuildPlayerSkillDataList(pl player.Player) (skillList []*crosspb.SkillData) {
	for _, ski := range pl.GetAllSkills() {
		skillData := BuildSkillData(ski)
		skillList = append(skillList, skillData)
	}
	return
}

func BuildSkillData(obj skillcommon.SkillObject) *crosspb.SkillData {
	level := obj.GetLevel()
	//TODO 修改技能使用时间
	lastTime := int64(0)
	skillId := obj.GetSkillId()
	skillData := &crosspb.SkillData{}
	skillData.Level = &level
	skillData.LastTime = &lastTime
	skillData.SkillId = &skillId
	if len(obj.GetTianFuList()) != 0 {
		for _, tianFuInfo := range obj.GetTianFuList() {
			skillData.TianFuList = append(skillData.TianFuList, buildTianFu(tianFuInfo))
		}
	}
	return skillData
}

func buildTianFu(tianFuInfo *skillcommon.TianFuInfo) *crosspb.SkillTianFu {
	skillTianFu := &crosspb.SkillTianFu{}
	tianFuId := tianFuInfo.TianFuId
	level := tianFuInfo.Level
	skillTianFu.TianFuId = &tianFuId
	skillTianFu.Level = &level
	return skillTianFu
}

//玩家buff数据
func BuildPlayerBuffDataList(pl player.Player) (buffList []*crosspb.BuffData) {
	for _, buf := range pl.GetBuffs() {
		buffData := BuildBuffData(buf)
		buffList = append(buffList, buffData)
	}
	return
}

//buff数据
func BuildBuffData(obj buffcommon.BuffObject) *crosspb.BuffData {
	ownerId := obj.GetOwnerId()
	buffId := obj.GetBuffId()
	groupId := obj.GetGroupId()
	startTime := obj.GetStartTime()
	lastTouchTime := obj.GetLastTouchTime()
	duration := obj.GetDuration()
	culTime := obj.GetCulTime()
	useTime := obj.GetUseTime()
	buffData := &crosspb.BuffData{}
	buffData.OwnerId = &ownerId
	buffData.BuffId = &buffId
	buffData.GroupId = &groupId
	buffData.StartTime = &startTime
	buffData.LastTouchTime = &lastTouchTime
	buffData.Duration = &duration
	buffData.CulTime = &culTime
	buffData.UseTime = &useTime
	return buffData
}

//pk数据
func BuildPlayerPkData(pl player.Player) *crosspb.PkData {

	pkValue := pl.GetPkValue()
	pkData := &crosspb.PkData{}
	pkData.PkValue = &pkValue

	return pkData
}

//队伍数据
func BuildPlayerTeamData(pl player.Player) *crosspb.TeamData {

	teamId := pl.GetTeamId()
	teamName := pl.GetTeamName()
	teamPurpose := int32(pl.GetTeamPurpose())
	teamData := &crosspb.TeamData{}
	teamData.TeamId = &teamId
	teamData.TeamName = &teamName
	teamData.TeamPurpose = &teamPurpose

	return teamData
}

//仙盟数据
func BuildPlayerAllianceData(pl player.Player) *crosspb.AllianceData {

	allianceId := pl.GetAllianceId()
	allianceName := pl.GetAllianceName()
	mengZhuId := pl.GetMengZhuId()
	memPos := int32(pl.GetMemPos())
	allianceData := &crosspb.AllianceData{}
	allianceData.AllianceId = &allianceId
	allianceData.AllianceName = &allianceName
	allianceData.MengZhuId = &mengZhuId
	allianceData.MemPos = &memPos
	return allianceData
}

// 结义数据
func BuildPlayerJieYiData(pl player.Player) *crosspb.JieYiData {
	jieYiId := pl.GetJieYiId()
	jieYiName := pl.GetJieYiName()
	rank := pl.GetJieYiRank()
	jieYiData := &crosspb.JieYiData{}
	jieYiData.JieYiId = &jieYiId
	jieYiData.JieYiName = &jieYiName
	jieYiData.Rank = &rank
	return jieYiData
}

// 阵营数据
func BuildPlayerChuangShiData(pl player.Player) *crosspb.ChuangShiData {
	pos := int32(pl.GetGuanZhi())
	campType := int32(pl.GetCamp())
	chuangShiData := &crosspb.ChuangShiData{}
	chuangShiData.Pos = &pos
	chuangShiData.CampType = &campType
	return chuangShiData
}

//活动数据
func BuildPlayerCrossData(pl player.Player) *crosspb.CrossData {
	crossType := int32(pl.GetCrossType())
	crossData := &crosspb.CrossData{}
	crossData.CrossType = &crossType
	crossArgs := pl.GetCrossArgs()
	crossData.Args = append(crossData.Args, crossArgs...)
	return crossData
}

//竞技场数据
func BuildPlayerArenaData(pl player.Player) *crosspb.PlayerArenaData {
	winTime := pl.GetArenaWinTime()
	reliveTime := pl.GetArenaReliveTime()
	playerArenaData := &crosspb.PlayerArenaData{}
	playerArenaData.WinTime = &winTime
	playerArenaData.ReliveTime = &reliveTime
	return playerArenaData
}

//竞技场pvp数据
func BuildPlayerArenapvpData(pl player.Player) *crosspb.PlayerArenapvpData {
	reliveTimes := pl.GetArenapvpReliveTimes()
	data := &crosspb.PlayerArenapvpData{}
	data.ReliveTimes = &reliveTimes
	return data
}

func BuildPlayerXueChiData(pl player.Player) *crosspb.XueChiData {
	xueChiData := &crosspb.XueChiData{}
	bloodLine := pl.GetBloodLine()
	blood := pl.GetBlood()
	xueChiData.BloodLine = &bloodLine
	xueChiData.Blood = &blood
	return xueChiData
}

func BuildPlayerReliveData(pl player.Player) *crosspb.PlayerReliveData {
	playerReliveData := &crosspb.PlayerReliveData{}
	culTime := pl.GetCulReliveTime()
	lastReliveTime := pl.GetLastReliveTime()
	playerReliveData.CulTime = &culTime
	playerReliveData.LastReliveTime = &lastReliveTime
	return playerReliveData
}

func BuildPlayerBattleData(pl player.Player) *crosspb.PlayerBattleData {
	playerBattleData := &crosspb.PlayerBattleData{}
	level := pl.GetLevel()
	zhuanSheng := pl.GetZhuanSheng()
	isHuiYuan := pl.IsHuiYuanPlus()
	playerBattleData.Level = &level
	playerBattleData.ZhuanSheng = &zhuanSheng
	playerBattleData.IsHuiYuan = &isHuiYuan
	return playerBattleData
}

func BuildPlayerDenseWatData(pl player.Player) *crosspb.DenseWatData {
	denseWatData := &crosspb.DenseWatData{}
	num := pl.GetDenseWatNum()
	endTime := pl.GetDenseWatEndTime()
	denseWatData.Num = &num
	denseWatData.EndTime = &endTime
	return denseWatData
}

func BuildPlayerShenMoData(pl player.Player) *crosspb.ShenMoData {
	shenMoData := &crosspb.ShenMoData{}
	gongXunNum := pl.GetShenMoGongXunNum()
	killNum := pl.GetShenMoKillNum()
	endTime := pl.GetShenMoEndTime()

	shenMoData.GongXunNum = &gongXunNum
	shenMoData.KillNum = &killNum
	shenMoData.EndTime = &endTime
	return shenMoData
}

func BuildPlayerActivityPkDataList(pl player.Player) (dataList []*crosspb.ActivityPkData) {
	for _, activityData := range pl.GetPlayerActvitityKillMap() {
		pkData := BuildPlayerActivityPkData(activityData.GetActivityType(), activityData.GetKilledNum(), activityData.GetLastKilledTime())
		dataList = append(dataList, pkData)
	}
	return dataList
}

func BuildPlayerActivityRankDataList(pl player.Player) (dataList []*crosspb.ActivityRankData) {
	for activityType, rankData := range pl.GetActivityRankMap() {
		activityTypeInt := int32(activityType)
		endTime := rankData.GetEndTime()
		data := &crosspb.ActivityRankData{}
		data.ActivityType = &activityTypeInt
		data.EndTime = &endTime
		for rankType, val := range rankData.GetRankMap() {
			playerRankData := BuildPlayerActivityRankData(activityType, rankType, val)
			data.PlayerRankDataList = append(data.PlayerRankDataList, playerRankData)
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func BuildPlayerActivityPkData(activityType activitytypes.ActivityType, killedNum int32, lastKilledTime int64) *crosspb.ActivityPkData {
	pkData := &crosspb.ActivityPkData{}
	activityTypeInt := int32(activityType)
	pkData.ActivityType = &activityTypeInt
	pkData.KilledNum = &killedNum
	pkData.LastKillTime = &lastKilledTime
	return pkData
}

func BuildPlayerActivityRankData(activityType activitytypes.ActivityType, rankType int32, val int64) *crosspb.PlayerActivityRankData {
	data := &crosspb.PlayerActivityRankData{}
	activityTypeInt := int32(activityType)
	data.ActivityType = &activityTypeInt
	data.RankType = &rankType
	data.Val = &val
	return data
}

//玩家系统战斗属性变更
func BuildSIPlayerSystemBattlePropertyChanged(battleProperties map[int32]int64, power int64) *crosspb.SIPlayerSystemBattlePropertyChanged {
	siPlayerSystemBattlePropertyChanged := &crosspb.SIPlayerSystemBattlePropertyChanged{}
	siPlayerSystemBattlePropertyChanged.BattlePropertyData = BuildBattlePropertyData(battleProperties)
	siPlayerSystemBattlePropertyChanged.Power = &power
	return siPlayerSystemBattlePropertyChanged
}

func BuildSIArenaPlayerWinTimeChanged(winTime int32) *crosspb.SIPlayerArenaDataChanged {
	siPlayerArenaDataChanged := &crosspb.SIPlayerArenaDataChanged{}
	playerArenaData := &crosspb.PlayerArenaData{}
	playerArenaData.WinTime = &winTime
	siPlayerArenaDataChanged.PlayerArenaData = playerArenaData
	return siPlayerArenaDataChanged
}

func BuildSIArenaPlayerReliveTimeChanged(reliveTime int32) *crosspb.SIPlayerArenaDataChanged {
	siPlayerArenaDataChanged := &crosspb.SIPlayerArenaDataChanged{}
	playerArenaData := &crosspb.PlayerArenaData{}
	playerArenaData.ReliveTime = &reliveTime
	siPlayerArenaDataChanged.PlayerArenaData = playerArenaData
	return siPlayerArenaDataChanged
}

func BuildPlayerFashionChanged(fashionId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.ClothesId = &fashionId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerFourGodKeyChanged(fourGodKey int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.FourGodKey = &fourGodKey
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerLingYuChanged(lingYu int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.LingYuId = &lingYu
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerModelChanged(model int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.Model = &model
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerMountChanged(mountId int32, advanceId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.RideId = &mountId
	playerShowData.MountAdvanceId = &advanceId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerRealmChanged(realm int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.Realm = &realm
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerShenFaChanged(shenFa int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.ShenFaId = &shenFa
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerSpouseChanged(spouse string) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.Spouse = &spouse
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerTitleChanged(title int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.TitleId = &title
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerWeaponChanged(weaponId int32, weaponState int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.WeaponId = &weaponId
	playerShowData.WeaponState = &weaponState
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerWeddingStatusChanged(weddingStatus int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.WeddingStatus = &weddingStatus
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerWingChanged(wingId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.WingId = &wingId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerRingTypeChanged(ringType int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.RingType = &ringType
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerSoulAwakenChanged(soulAwakenNum int32) *crosspb.SIPlayerBattleDataChanged {
	siPlayerBattleDataChanged := &crosspb.SIPlayerBattleDataChanged{}
	playerBattleData := &crosspb.PlayerBattleData{}
	playerBattleData.SoulAwakenNum = &soulAwakenNum
	siPlayerBattleDataChanged.PlayerBattleData = playerBattleData
	return siPlayerBattleDataChanged
}

func BuildPlayerLevelChanged(level int32) *crosspb.SIPlayerBattleDataChanged {
	siPlayerBattleDataChanged := &crosspb.SIPlayerBattleDataChanged{}
	playerBattleData := &crosspb.PlayerBattleData{}
	playerBattleData.Level = &level
	siPlayerBattleDataChanged.PlayerBattleData = playerBattleData
	return siPlayerBattleDataChanged
}

func BuildPlayerZhuanShengChanged(zhuanSheng int32) *crosspb.SIPlayerBattleDataChanged {
	siPlayerBattleDataChanged := &crosspb.SIPlayerBattleDataChanged{}
	playerBattleData := &crosspb.PlayerBattleData{}
	playerBattleData.ZhuanSheng = &zhuanSheng
	siPlayerBattleDataChanged.PlayerBattleData = playerBattleData
	return siPlayerBattleDataChanged
}
func BuildPlayerVipLevelChanged(vipLevel int32) *crosspb.SIPlayerBattleDataChanged {
	siPlayerBattleDataChanged := &crosspb.SIPlayerBattleDataChanged{}
	playerBattleData := &crosspb.PlayerBattleData{}
	playerBattleData.Vip = &vipLevel
	siPlayerBattleDataChanged.PlayerBattleData = playerBattleData
	return siPlayerBattleDataChanged
}

func BuildPlayerTeamChanged(teamId int64, teamName string, teamPurpose int32) *crosspb.SIPlayerTeamSync {
	siPlayerTeamSync := &crosspb.SIPlayerTeamSync{}
	teamData := &crosspb.TeamData{}
	teamData.TeamId = &teamId
	teamData.TeamName = &teamName
	teamData.TeamPurpose = &teamPurpose
	siPlayerTeamSync.TeamData = teamData
	return siPlayerTeamSync
}

func BuildPlayerAllianceChanged(allianceId int64, allianceName string, mengZhuId int64, memPos int32) *crosspb.SIPlayerAllianceSync {
	siPlayerAllianceSync := &crosspb.SIPlayerAllianceSync{}
	allianceData := &crosspb.AllianceData{}
	allianceData.AllianceId = &allianceId
	allianceData.AllianceName = &allianceName
	allianceData.MengZhuId = &mengZhuId
	allianceData.MemPos = &memPos
	siPlayerAllianceSync.AllianceData = allianceData
	return siPlayerAllianceSync
}

func BuildPlayerJieYiChanged(jieYiId int64, jieYiName string, rank int32) *crosspb.SIPlayerJieYiSync {
	siPlayerJieYiSync := &crosspb.SIPlayerJieYiSync{}
	jieYiData := &crosspb.JieYiData{}
	jieYiData.JieYiId = &jieYiId
	jieYiData.JieYiName = &jieYiName
	jieYiData.Rank = &rank

	siPlayerJieYiSync.JieYiData = jieYiData
	return siPlayerJieYiSync
}

func BuildPlayerFaBaoChanged(faBaoId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.FaBaoId = &faBaoId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerPetChanged(petId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.PetId = &petId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerXianTiChanged(xianTiId int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.XianTiId = &xianTiId
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildPlayerBaGuaChanged(baGua int32) *crosspb.SIPlayerShowDataChanged {
	siPlayerShowDataChanged := &crosspb.SIPlayerShowDataChanged{}
	playerShowData := &crosspb.PlayerShowData{}
	playerShowData.BaGua = &baGua
	siPlayerShowDataChanged.PlayerShowData = playerShowData
	return siPlayerShowDataChanged
}

func BuildLingTongFashionChanged(fashionId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.FashionId = &fashionId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongLingYuChanged(lingYuId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.LingYuId = &lingYuId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongMountChanged(mountId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.MountId = &mountId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongFaBaoChanged(fabaoId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.FaBaoId = &fabaoId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongShenFaChanged(shenFaId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.ShenFaId = &shenFaId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongWeaponChanged(weaponId int32, weaponState int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.WeaponId = &weaponId
	lingTongShowData.WeaponState = &weaponState
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongWingChanged(wingId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.WingId = &wingId
	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongXianTiChanged(xianTiId int32) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	lingTongShowData := &crosspb.LingTongShowData{}
	lingTongShowData.XianTiId = &xianTiId

	siLingTongDataChanged.LingTongShowData = lingTongShowData
	return siLingTongDataChanged
}

func BuildLingTongChanged(lingTongId int32, name string) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	siLingTongDataChanged.LingTongId = &lingTongId
	siLingTongDataChanged.Name = &name
	return siLingTongDataChanged
}

var (
	siLingTongDataRemove = &crosspb.SILingTongDataRemove{}
)

func BuildLingTongRemove() *crosspb.SILingTongDataRemove {

	return siLingTongDataRemove
}

func BuildLingTongDataInit(lingTong scene.LingTong) *crosspb.SILingTongDataInit {

	siLingTongDataInit := &crosspb.SILingTongDataInit{}
	siLingTongDataInit.LingTongData = BuildLingTongData(lingTong)
	return siLingTongDataInit
}

//玩家系统战斗属性变更
func BuildLingTongSystemBattlePropertyChanged(battleProperties map[int32]int64) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	siLingTongDataChanged.BattlePropertyData = BuildBattlePropertyData(battleProperties)

	return siLingTongDataChanged
}

//玩家名字变更
func BuildLingTongNameChanged(name string) *crosspb.SILingTongDataChanged {
	siLingTongDataChanged := &crosspb.SILingTongDataChanged{}
	siLingTongDataChanged.Name = &name

	return siLingTongDataChanged
}

//buff数据
func BuildSIBuffAdd(obj buffcommon.BuffObject) *crosspb.SIBuffAdd {
	siBuffAdd := &crosspb.SIBuffAdd{}
	siBuffAdd.BuffData = BuildBuffData(obj)
	return siBuffAdd
}

//buff数据
func BuildSIBuffUpdate(obj buffcommon.BuffObject) *crosspb.SIBuffUpdate {
	siBuffUpdate := &crosspb.SIBuffUpdate{}
	siBuffUpdate.BuffData = BuildBuffData(obj)
	return siBuffUpdate
}

//buff数据
func BuildSIBuffRemove(buffId int32) *crosspb.SIBuffRemove {
	siBuffRemove := &crosspb.SIBuffRemove{}
	siBuffRemove.BuffId = &buffId
	return siBuffRemove
}

// pvp
func BuildSIPlayerArenapvpDataChanged(reliveTimes int32) *crosspb.SIPlayerArenapvpDataChanged {
	siMsg := &crosspb.SIPlayerArenapvpDataChanged{}
	info := &crosspb.PlayerArenapvpData{}
	info.ReliveTimes = &reliveTimes
	siMsg.PlayerArenapvpData = info
	return siMsg
}

func BuildPlayerBossReliveDataList(pl player.Player) (dataList []*crosspb.PlayerBossReliveData) {
	for bossType, reliveData := range pl.GetPlayerBossReliveMap() {
		bossTypeInt := int32(bossType)
		data := &crosspb.PlayerBossReliveData{}
		data.BossType = &bossTypeInt
		reliveTime := reliveData.GetReliveTime()
		data.ReliveTime = &reliveTime
		dataList = append(dataList, data)
	}
	return dataList
}

func BuildSIPlayerTeShuSkillReset(pl player.Player) (siPlayerTeshuSkillReset *crosspb.SIPlayerTeshuSkillReset) {
	siPlayerTeshuSkillReset = &crosspb.SIPlayerTeshuSkillReset{}
	dataList := BuildPlayerTeShuSkillDataList(pl)
	siPlayerTeshuSkillReset.SkillList = dataList
	return siPlayerTeshuSkillReset
}

func BuildPlayerTeShuSkillDataList(pl player.Player) (dataList []*crosspb.TeShuSkillData) {
	for _, skillObj := range pl.GetTeShuSkills() {
		data := &crosspb.TeShuSkillData{}
		skillId := skillObj.GetSkillId()
		data.SkillId = &skillId
		chuFaRate := skillObj.GetChuFaRate()
		data.ChuFaRate = &chuFaRate
		diKangRate := skillObj.GetDiKangRate()
		data.DiKangRate = &diKangRate
		dataList = append(dataList, data)
	}
	return dataList
}
