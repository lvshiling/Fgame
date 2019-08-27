package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/exception"
	additionsyspbutil "fgame/fgame/game/additionsys/pbutil"
	anqipbutil "fgame/fgame/game/anqi/pbutil"
	babypbutil "fgame/fgame/game/baby/pbutil"
	bodyshieldpbutil "fgame/fgame/game/bodyshield/pbutil"
	dianxingpbutil "fgame/fgame/game/dianxing/pbutil"
	goldequippbutil "fgame/fgame/game/goldequip/pbutil"
	inventorypbutil "fgame/fgame/game/inventory/pbutil"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	lingtongpbutil "fgame/fgame/game/lingtong/pbutil"
	lingtongdevpbutil "fgame/fgame/game/lingtongdev/pbutil"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	lingyupbutil "fgame/fgame/game/lingyu/pbutil"
	marrytypes "fgame/fgame/game/marry/types"
	mountpbutil "fgame/fgame/game/mount/pbutil"
	"fgame/fgame/game/player"
	playercommon "fgame/fgame/game/player/common"
	playertypes "fgame/fgame/game/player/types"
	propertypbutil "fgame/fgame/game/property/pbutil"
	ringpbutil "fgame/fgame/game/ring/pbutil"
	"fgame/fgame/game/scene/scene"
	shenfapbutil "fgame/fgame/game/shenfa/pbutil"
	soulpbutil "fgame/fgame/game/soul/pbutil"
	sysskillpbutil "fgame/fgame/game/systemskill/pbutil"
	weaponpbutil "fgame/fgame/game/weapon/pbutil"
	wingpbutil "fgame/fgame/game/wing/pbutil"
	wushuangweaponpbutil "fgame/fgame/game/wushuangweapon/pbutil"
)

var (
	//所有通用一个
	scEnterSelectJob = &uipb.SCEnterSelectJob{}
	scSelectJob      = &uipb.SCSelectJob{}
)

func BuildSCEnterSelectJob() *uipb.SCEnterSelectJob {
	return scEnterSelectJob
}

func BuildSCSelectJob() *uipb.SCSelectJob {
	return scSelectJob
}

func BuildSCSystemMessage(content string, args ...string) *uipb.SCSystemMessage {
	scSystemMessage := &uipb.SCSystemMessage{}
	scSystemMessage.Content = &content
	scSystemMessage.Args = args
	return scSystemMessage
}

func BuildSCException(content string, code exception.ExceptionCode) *uipb.SCException {
	scException := &uipb.SCException{}
	scException.Content = &content
	codeInt := int32(code)
	scException.Code = &codeInt
	return scException
}

func BuildSCPlayerInfo(p player.Player, allianceFlag bool, tradeFlag bool) *uipb.SCPlayerInfo {
	info := &uipb.SCPlayerInfo{}
	playerId := p.GetId()
	name := p.GetName()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	onlineTime := int64(p.GetOnlineTime())
	info.PlayerId = &playerId
	info.Name = &name
	info.Role = &role
	info.Sex = &sex
	info.OnlineTime = &onlineTime
	totalOnlineTime := p.GetTotalOnlineTime()
	info.TotalOnlineTime = &totalOnlineTime
	isOpenVideo := p.IsOpenVideo()
	info.IsOpenVideo = &isOpenVideo
	createTime := p.GetCreateTime()
	info.CreateTime = &createTime
	todayOnlineTime := p.GetTodayOnlineTime()
	info.TodayOnlineTime = &todayOnlineTime
	privilege := int32(p.GetPrivilege())
	info.PrivilegeTypeInt = &privilege
	info.AllianceDepotFlag = &allianceFlag
	info.TradeFlag = &tradeFlag
	return info
}

func BuildPlayerBasicInfoByPlayer(p scene.Player, isBlacked bool) *uipb.PlayerBasicInfo {
	playerId := p.GetId()
	name := p.GetName()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	level := p.GetLevel()
	force := p.GetForce()
	onlineState := int32(playertypes.PlayerOnlineStateOnline)
	if isBlacked {
		onlineState = int32(playertypes.PlayerOnlineStateOffline)
	}

	basicInfo := &uipb.PlayerBasicInfo{}
	basicInfo.PlayerId = &playerId
	basicInfo.Name = &name
	basicInfo.Role = &role
	basicInfo.Sex = &sex
	basicInfo.Level = &level
	basicInfo.Force = &force
	basicInfo.OnlineState = &onlineState
	allianceId := p.GetAllianceId()
	basicInfo.AllianceId = &allianceId
	teamId := p.GetTeamId()
	basicInfo.TeamId = &teamId
	allianceName := p.GetAllianceName()
	basicInfo.AllianceName = &allianceName
	fashionId := p.GetFashionId()
	basicInfo.FashionId = &fashionId
	weaponId := p.GetWeaponId()
	basicInfo.WeaponId = &weaponId
	mountId := p.GetMountId()
	basicInfo.MountId = &mountId
	wingId := p.GetWingId()
	basicInfo.WingId = &wingId
	spouseId := p.GetSpouseId()
	spouseName := p.GetSpouse()
	basicInfo.SpouseId = &spouseId
	basicInfo.SpouseName = &spouseName
	realmLevel := p.GetRealm()
	basicInfo.RealmLevel = &realmLevel
	lingTong := p.GetLingTong()
	lingTongId := int32(0)
	lingTongWeaponId := int32(0)
	lingTongFashionId := int32(0)
	lingTongWingId := int32(0)
	if lingTong != nil {
		//TODO 获取灵童id
		// lingTongId = lingTong.GetLingTongMountId
		lingTongWeaponId = lingTong.GetLingTongWeaponId()
		lingTongFashionId = lingTong.GetLingTongFashionId()
		lingTongWingId = lingTong.GetLingTongWingId()
	}
	basicInfo.LingTongFashionId = &lingTongFashionId
	basicInfo.LingTongWeaponId = &lingTongWeaponId
	basicInfo.LingTongId = &lingTongId
	basicInfo.LingTongWingId = &lingTongWingId
	ring := p.GetRingType()
	ringLevel := int32(0)
	if ring > 0 {
		//特殊处理
		ring = ring - 1
		itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, marrytypes.MarryRingType(ring).ItemWedRingSubType())
		if itemTemplate != nil {
			ring = int32(itemTemplate.TemplateId())
			ringLevel = p.GetRingLevel()
		}
	}
	basicInfo.Ring = &ring
	basicInfo.RingLevel = &ringLevel

	isHuiYuan := p.IsHuiYuanPlus()
	basicInfo.IsHuiYuan = &isHuiYuan

	return basicInfo
}

func BuildPlayerBasicInfo(info *playercommon.PlayerInfo, isBlacked bool) *uipb.PlayerBasicInfo {
	playerId := info.PlayerId
	name := info.Name
	role := int32(info.Role)
	sex := int32(info.Sex)
	level := info.Level
	force := info.Force
	onlineState := int32(playertypes.PlayerOnlineStateOnline)
	if player.GetOnlinePlayerManager().GetPlayerById(info.PlayerId) == nil {
		onlineState = int32(playertypes.PlayerOnlineStateOffline)
	}

	if isBlacked {
		onlineState = int32(playertypes.PlayerOnlineStateOffline)
	}
	basicInfo := &uipb.PlayerBasicInfo{}
	basicInfo.PlayerId = &playerId
	basicInfo.Name = &name
	basicInfo.Role = &role
	basicInfo.Sex = &sex
	basicInfo.Level = &level
	basicInfo.Force = &force
	basicInfo.OnlineState = &onlineState
	allianceId := info.AllianceId
	basicInfo.AllianceId = &allianceId
	teamId := info.TeamId
	basicInfo.TeamId = &teamId
	allianceName := info.AllianceName
	basicInfo.AllianceName = &allianceName
	fashionId := info.FashionId
	basicInfo.FashionId = &fashionId
	weaponId := info.AllWeaponInfo.Wear
	basicInfo.WeaponId = &weaponId
	mountId := info.MountInfo.GetMountId()
	basicInfo.MountId = &mountId
	wingId := info.WingInfo.GetWingId()
	basicInfo.WingId = &wingId
	spouseId := info.MarryInfo.SpouseId
	spouseName := info.MarryInfo.SpouseName
	basicInfo.SpouseId = &spouseId
	basicInfo.SpouseName = &spouseName
	realmLevel := info.RealmLevel
	basicInfo.RealmLevel = &realmLevel
	xianZunCardList := make([]int32, 0, 8)
	for _, card := range info.XianZunCardList {
		xianZunCardList = append(xianZunCardList, card.Typ)
	}
	basicInfo.XianZunCard = xianZunCardList
	ring := int32(0)
	ringLevel := int32(0)

	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, marrytypes.MarryRingType(info.MarryInfo.Ring).ItemWedRingSubType())
	if itemTemplate != nil {
		ring = int32(itemTemplate.TemplateId())
		ringLevel = info.MarryInfo.RLevel
	}
	basicInfo.Ring = &ring
	basicInfo.RingLevel = &ringLevel

	//灵童
	lingTongId := info.LingTongInfo.LingTongId
	lingTongFashionId := info.LingTongInfo.FashionId
	lingTongWingId := int32(0)
	lingTongWeaponId := int32(0)
	for _, lingTongDevInfo := range info.AllLingTongDevInfo.LingTongDevList {
		classType := lingtongdevtypes.LingTongDevSysType(lingTongDevInfo.ClassType)
		if classType == lingtongdevtypes.LingTongDevSysTypeLingYi {
			lingTongWingId = lingTongDevInfo.GetSeqId()
		}
		if classType == lingtongdevtypes.LingTongDevSysTypeLingBing {
			lingTongWeaponId = lingTongDevInfo.GetSeqId()
		}
	}

	basicInfo.LingTongId = &lingTongId
	basicInfo.LingTongFashionId = &lingTongFashionId
	basicInfo.LingTongWeaponId = &lingTongWeaponId
	basicInfo.LingTongWingId = &lingTongWingId

	isHuiYuan := false
	if info.IsHuiYuan == 1 {
		isHuiYuan = true
	}
	basicInfo.IsHuiYuan = &isHuiYuan
	return basicInfo
}

func BuildSCPlayerBasicInfoGet(info *playercommon.PlayerInfo) *uipb.SCPlayerBasicInfoGet {

	infoGet := &uipb.SCPlayerBasicInfoGet{}
	infoGet.PlayerBasicInfo = BuildPlayerBasicInfo(info, false)

	return infoGet
}

func BuildscPlayerOpenVedio() *uipb.SCPlayerOpenVedio {
	scPlayerOpenVedio := &uipb.SCPlayerOpenVedio{}
	return scPlayerOpenVedio
}

func BuildSCPlayerBasicInfoBatchGet(infoList []*playercommon.PlayerInfo) *uipb.SCPlayerBasicInfoBatchGet {
	infoBatchGet := &uipb.SCPlayerBasicInfoBatchGet{}
	for _, info := range infoList {
		infoBatchGet.PlayerBasicInfoList = append(infoBatchGet.PlayerBasicInfoList, BuildPlayerBasicInfo(info, false))
	}

	return infoBatchGet
}

func BuildSCPlayerCountData() *uipb.SCPlayerCountData {
	scMsg := &uipb.SCPlayerCountData{}
	return scMsg
}

func BuildSCPlayerInfoGet(info *playercommon.PlayerInfo) *uipb.SCPlayerInfoGet {

	infoGet := &uipb.SCPlayerInfoGet{}
	infoGet.BasicInfo = BuildPlayerBasicInfo(info, false)
	infoGet.BasePropertyList = propertypbutil.BuildProperties(info.BaseProperty)
	infoGet.BattlePropertyList = propertypbutil.BuildProperties(info.BattleProperty)
	infoGet.EquipmentList = inventorypbutil.BuildEquipmentSlotInfoList(info.EquipmentList)
	infoGet.GoldEquipList = goldequippbutil.BuildGoldEquipSlotList(info.GoldEquipList)
	infoGet.MountInfo = mountpbutil.BuildMountInfo(info.MountInfo)
	infoGet.WingInfo = wingpbutil.BuildWingInfo(info.WingInfo)
	infoGet.BodyShieldInfo = bodyshieldpbutil.BuildBodyShieldInfo(info.BodyShieldInfo)
	infoGet.AnqiInfo = anqipbutil.BuildAnqiInfo(info.AnqiInfo)
	infoGet.AllSoulInfo = soulpbutil.BuildAllSoulInfo(info.AllSoulInfo)
	infoGet.AllWeaponInfo = weaponpbutil.BuildAllWeaponInfo(info.AllWeaponInfo)
	infoGet.ShenfaInfo = shenfapbutil.BuildShenfaInfo(info.ShenfaInfo)
	infoGet.LingyuInfo = lingyupbutil.BuildLingyuInfo(info.LingyuInfo)
	infoGet.ShieldInfo = bodyshieldpbutil.BuildShieldInfo(info.ShieldInfo)
	infoGet.FeatherInfo = wingpbutil.BuildFeatherInfo(info.FeatherInfo)
	infoGet.AllLingTongDevInfo = lingtongdevpbutil.BuildAllLingTongDevInfo(info.AllLingTongDevInfo)
	infoGet.LingTongInfo = lingtongpbutil.BuildLingTongCacheInfo(info.LingTongInfo)
	infoGet.AllSystemSkillInfo = sysskillpbutil.BuildAllSystemSkillInfo(info.AllSystemSkillInfo)
	infoGet.AllAdditionSysInfo = additionsyspbutil.BuildAllAdditionSysInfo(info.AllAdditionSysInfo)
	infoGet.PregnantInfo = babypbutil.BuildDongFangInfo(info.PregnantInfo)
	infoGet.DianXingInfo = dianxingpbutil.BuildDianXingCacheInfo(info.DianXingInfo)
	infoGet.WushuangSlotList = wushuangweaponpbutil.BuildWushuangBodyPosList(info.WushuangList)
	infoGet.RingInfoList = ringpbutil.BuildSCRingInfoList(info.RingList)

	fashionId := info.FashionId
	infoGet.FashionId = &fashionId
	return infoGet
}

func BuildSCPlayerSexChanged(sex playertypes.SexType) *uipb.SCPlayerSexChanged {
	scMsg := &uipb.SCPlayerSexChanged{}
	sexInt := int32(sex)
	scMsg.NewSex = &sexInt

	return scMsg
}

func BuildSCPlayerNameChanged(name string) *uipb.SCPlayerNameChanged {
	scMsg := &uipb.SCPlayerNameChanged{}
	scMsg.NewName = &name

	return scMsg
}
