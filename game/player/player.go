package player

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/common/message"
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/session"
	additionsyscommon "fgame/fgame/game/additionsys/common"
	anqitypes "fgame/fgame/game/anqi/types"
	babytypes "fgame/fgame/game/baby/types"
	baguacommon "fgame/fgame/game/bagua/common"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	"fgame/fgame/game/cache/cache"
	crosssession "fgame/fgame/game/cross/session"
	crosstypes "fgame/fgame/game/cross/types"
	dianxingcommon "fgame/fgame/game/dianxing/common"
	gameevent "fgame/fgame/game/event"
	fabaocommon "fgame/fgame/game/fabao/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	lingyutypes "fgame/fgame/game/lingyu/types"
	marrytypes "fgame/fgame/game/marry/types"
	massacretypes "fgame/fgame/game/massacre/types"
	mountcommon "fgame/fgame/game/mount/common"
	playercommon "fgame/fgame/game/player/common"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	ringcommon "fgame/fgame/game/ring/common"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	gamessession "fgame/fgame/game/session"
	shenfatypes "fgame/fgame/game/shenfa/types"
	shihunfancommon "fgame/fgame/game/shihunfan/common"
	skillcommon "fgame/fgame/game/skill/common"
	soultypes "fgame/fgame/game/soul/types"
	sysskillcommon "fgame/fgame/game/systemskill/common"
	tianmotypes "fgame/fgame/game/tianmo/types"
	viptypes "fgame/fgame/game/vip/types"
	weapontypes "fgame/fgame/game/weapon/types"
	wingcommon "fgame/fgame/game/wing/common"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	xianticommon "fgame/fgame/game/xianti/common"
	xianzuncardcommon "fgame/fgame/game/xianzuncard/common"
	xueduncommon "fgame/fgame/game/xuedun/common"
	accounttype "fgame/fgame/login/types"
	"sync"

	"github.com/golang/protobuf/proto"
)

func PlayerInSession(s session.Session) Player {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl, ok := pl.(Player)
	if !ok {
		return nil
	}
	return tpl
}

//玩家更新器
type PlayerUpdater interface {
	AddChangedObject(obj types.PlayerDataEntity)
}

//玩家
type Player interface {
	//场景玩家
	scene.Player
	//玩家存储
	PlayerUpdater
	//状态
	fsm.Subject
	GetUserId() int64
	GetPlatformUserId() string
	GetIp() string
	RealNameAuth(accounttype.RealNameState)
	GetRealNameState() accounttype.RealNameState
	GetSDKType() logintypes.SDKType
	GetDevicePlatformType() logintypes.DevicePlatformType
	GetWallowState() playertypes.WallowState
	GetOnlineTime() int64
	GetTotalOnlineTime() int64
	IsGm() bool
	//获取上次登出时间
	GetLastLogoutTime() int64
	//获取创建时间
	GetCreateTime() int64
	//获取今天在线时间
	GetTodayOnlineTime() int64
	//进入认证
	EnterAuth() bool
	//进入加载角色列表
	EnterLoadingRoleList() bool
	//进入创建角色
	EnterCreateRole() bool
	//进入等候创建角色
	EnterWaitingSelectRole() bool
	//认证
	Auth(playerId int64) bool
	//进入加载
	EnterLoad() bool
	//加载
	Load() error
	//加载后
	AfterLoad() (bool, error)
	CanProcess(msg message.Message) bool
	//进入跨服
	EnterCross() bool
	//跨服中
	Cross() bool
	//退出跨服
	LeaveCross() bool
	//设置跨服
	SetCrossSession(crosssession.SendSession)

	//是否正在跨服
	IsCross() bool
	//是否无间炼狱排队
	IsLianYuLineUp() bool
	//设置无间炼狱排队
	LianYuLineUp(isLianYuLine bool)

	//是否神魔战场排队
	IsShenMoLineUp() bool
	//设置神魔战场排队
	ShenMoLineUp(isLianYuLine bool)

	//通用排队
	//是否排队
	IsLineUp() bool
	//设置排队
	SetLineUp(isLineup bool)

	IsLogouting() bool
	Logout() bool
	LogoutSave() bool
	LogoutCross()
	Done() <-chan struct{}
	GetPlayerDataManager(typ types.PlayerDataManagerType) PlayerDataManager
	//发送跨服消息
	SendCrossMsg(proto.Message)
	//游戏对话
	Session() gamessession.Session
	//跨服对话
	GetCrossSession() crosssession.SendSession
	//获取活动类型
	GetCrossType() crosstypes.CrossType
	//获取跨服参数
	GetCrossArgs() []string
	//是否开启
	IsFuncOpen(fpt funcopentypes.FuncOpenType) bool
	//更新系统战斗属性
	UpdateBattleProperty(mask uint64)
	//TODO 分开
	//------------------玩家信息接口------------------------
	//获取坐骑信息
	GetMountInfo() *mountcommon.MountInfo
	//获取护盾信息
	GetBodyshieldInfo() *bodyshieldtypes.BodyShieldInfo
	//获取暗器信息
	GetAnqiInfo() *anqitypes.AnqiInfo
	//战翼
	GetWingInfo() *wingcommon.WingInfo
	//兵魂
	GetAllWeaponInfo() *weapontypes.AllWeaponInfo
	//获取装备信息
	GetEquipmentSlotList() []*inventorytypes.EquipmentSlotInfo
	//获取元神金装信息
	GetGoldEquipSlotList() []*goldequiptypes.GoldEquipSlotInfo
	//获取古魂信息
	GetAllSoulInfo() *soultypes.AllSoulInfo
	//获取结婚信息
	GetMarryInfo() *marrytypes.MarryInfo
	//获取神盾尖刺信息
	GetShieldInfo() *bodyshieldtypes.ShieldInfo
	//获取护体仙羽
	GetFeatherInfo() *wingcommon.FeatherInfo
	//获取基础属性
	GetBaseProperties() map[int32]int64
	//获取战斗属性
	GetBattleProperties() map[int32]int64
	//获取身法属性
	GetShenfaInfo() *shenfatypes.ShenfaInfo
	//获取领域属性
	GetLingyuInfo() *lingyutypes.LingyuInfo
	//获取技能列表
	GetSkillList() []*skillcommon.SkillObjectImpl
	//获取vip
	GetVipInfo() *viptypes.VipInfo
	//获取戮仙刃信息
	GetMassacreInfo() *massacretypes.MassacreInfo
	//获取法宝信息
	GetFaBaoInfo() *fabaocommon.FaBaoInfo
	//获取血盾信息
	GetXueDunInfo() *xueduncommon.XueDunInfo
	//获取仙体信息
	GetXianTiInfo() *xianticommon.XianTiInfo
	//获取八卦秘境信息
	GetBaGuaInfo() *baguacommon.BaGuaInfo
	//获取点星系统信息
	GetDianXingInfo() *dianxingcommon.DianXingInfo
	//获取天魔体信息
	GetTianMoTiInfo() *tianmotypes.TianMoInfo
	//获取噬魂幡信息
	GetShiHunFanInfo() *shihunfancommon.ShiHunFanInfo
	//获取灵童养成类信息
	GetAllLingTongDevInfo() *lingtongdevcommon.AllLingTongDevInfo
	//获取灵童信息
	GetLingTongInfo() *lingtongcommon.LingTongInfo
	//获取系统技能信息
	GetAllSystemSkillInfo() *sysskillcommon.AllSystemSkillInfo
	//获取附加系统类信息
	GetAllAdditionSysInfo() *additionsyscommon.AllAdditionSysInfo
	//获取附加系统类信息
	GetPregnantInfo() *babytypes.PregnantInfo
	//获取无双装备信息
	GetWushuangListInfo() []*wushuangweapontypes.WushuangInfo
	//获取仙尊特权卡信息
	GetXianZunCard() []*xianzuncardcommon.XianZunCardInfo
	//获取特戒信息
	GetRingInfo() []*ringcommon.RingInfo

	//仅Gm命令使用
	GmSetForbid(forbidText string)
	//封号
	Forbid(forbidReason string, forbidName string, forbidTime int64)
	Unforbid()
	ForbidChat(forbidChatReason string, forbidChatName string, forbidChatTime int64)
	UnforbidChat()
	IgnoreChat(forbidChatReason string, forbidChatName string, forbidChatTime int64)
	UnignoreChat()
	IsForbid() bool
	IsForbidChat() bool
	IsIgnoreChat() bool
	SetPrivilege(types.PrivilegeType)
	GetPrivilege() types.PrivilegeType
	//是否开场动画
	IsOpenVideo() bool
	OpenVideo()
	//充值
	AddChargeInfo(goldNum, money int64)
	AddPrivilegeChargeInfo(goldNum int64)
	GetChargeGoldNum() int64
	//是否是新手
	IsGetNewReward() bool
	GetNewReward()
	//性别变更
	ChangeSex() types.SexType
	//姓名变更
	ChangeName(newName string)
	//进阶系统补偿
	IsSystemCompensate() bool
	SendSystemCompensate()
	GMSetSystemCompensate(status bool)
}

func ConvertFromPlayer(pl Player) *playercommon.PlayerInfo {
	info := &playercommon.PlayerInfo{
		PlayerId:           pl.GetId(),
		Name:               pl.GetName(),
		Role:               pl.GetRole(),
		Sex:                pl.GetSex(),
		Level:              pl.GetLevel(),
		Force:              pl.GetForce(),
		AllianceId:         pl.GetAllianceId(),
		AllianceName:       pl.GetAllianceName(),
		TeamId:             pl.GetTeamId(),
		OnlineState:        playertypes.PlayerOnlineStateOnline,
		BaseProperty:       pl.GetBaseProperties(),
		BattleProperty:     pl.GetBattleProperties(),
		EquipmentList:      pl.GetEquipmentSlotList(),
		GoldEquipList:      pl.GetGoldEquipSlotList(),
		MountInfo:          pl.GetMountInfo(),
		WingInfo:           pl.GetWingInfo(),
		BodyShieldInfo:     pl.GetBodyshieldInfo(),
		AnqiInfo:           pl.GetAnqiInfo(),
		AllSoulInfo:        pl.GetAllSoulInfo(),
		AllWeaponInfo:      pl.GetAllWeaponInfo(),
		FashionId:          pl.GetFashionId(),
		ShieldInfo:         pl.GetShieldInfo(),
		FeatherInfo:        pl.GetFeatherInfo(),
		MarryInfo:          pl.GetMarryInfo(),
		ShenfaInfo:         pl.GetShenfaInfo(),
		LingyuInfo:         pl.GetLingyuInfo(),
		RealmLevel:         pl.GetRealm(),
		SkillList:          pl.GetSkillList(),
		VipInfo:            pl.GetVipInfo(),
		MassacreInfo:       pl.GetMassacreInfo(),
		FaBaoInfo:          pl.GetFaBaoInfo(),
		XueDunInfo:         pl.GetXueDunInfo(),
		XianTiInfo:         pl.GetXianTiInfo(),
		BaGuaInfo:          pl.GetBaGuaInfo(),
		DianXingInfo:       pl.GetDianXingInfo(),
		TianMoTiInfo:       pl.GetTianMoTiInfo(),
		ShiHunFanInfo:      pl.GetShiHunFanInfo(),
		AllLingTongDevInfo: pl.GetAllLingTongDevInfo(),
		LingTongInfo:       pl.GetLingTongInfo(),
		AllSystemSkillInfo: pl.GetAllSystemSkillInfo(),
		AllAdditionSysInfo: pl.GetAllAdditionSysInfo(),
		PregnantInfo:       pl.GetPregnantInfo(),
		WushuangList:       pl.GetWushuangListInfo(),
		XianZunCardList:    pl.GetXianZunCard(),
		RingList:           pl.GetRingInfo(),
	}
	if pl.IsHuiYuanPlus() {
		info.IsHuiYuan = 1
	} else {
		info.IsHuiYuan = 0
	}
	return info
}

func ConvertFromRobotPlayer(pl Player) *playercommon.PlayerInfo {
	info := &playercommon.PlayerInfo{
		PlayerId:      pl.GetId(),
		Name:          pl.GetName(),
		Role:          pl.GetRole(),
		Sex:           pl.GetSex(),
		Level:         pl.GetLevel(),
		Force:         pl.GetForce(),
		TeamId:        pl.GetTeamId(),
		OnlineState:   playertypes.PlayerOnlineStateOnline,
		AllWeaponInfo: pl.GetAllWeaponInfo(),
		FashionId:     pl.GetFashionId(),
	}
	return info
}

type PlayerService interface {
	Heartbeat()
	RecommentPlayersExclude(playerIdMap map[int64]struct{}) (pList []Player)
	RecommentSpouses(pl Player, excludePlayers map[int64]struct{}) (pList []Player)
	BatchGetPlayerInfo(playerIdList []int64) (infoList []*playercommon.PlayerInfo, err error)
	GetPlayerInfo(playerId int64) (info *playercommon.PlayerInfo, err error)
	GetPlayerInfoByName(name string, serverId int32) (info *playercommon.PlayerInfo, err error)
}

type playerService struct {
}

func (ps *playerService) RecommentPlayersExclude(playerIdMap map[int64]struct{}) (pList []Player) {
	pList = GetOnlinePlayerManager().RecommentPlayersExclude(playerIdMap)
	return
}

func (ps *playerService) RecommentSpouses(pl Player, excludePlayers map[int64]struct{}) (pList []Player) {
	pList = GetOnlinePlayerManager().RecommentSpouses(pl, excludePlayers)
	return
}

func (ps *playerService) BatchGetPlayerInfo(playerIdList []int64) (infoList []*playercommon.PlayerInfo, err error) {
	for _, playerId := range playerIdList {
		info, err := ps.GetPlayerInfo(playerId)
		if err != nil {
			return nil, err
		}
		if info == nil {
			continue
		}
		infoList = append(infoList, info)
	}

	return
}

func (ps *playerService) GetPlayerInfo(playerId int64) (info *playercommon.PlayerInfo, err error) {
	// p := GetOnlinePlayerManager().GetPlayerById(playerId)
	// if p == nil {
	info, err = cache.GetCacheService().GetPlayerInfoByPlayerId(playerId)
	if err != nil {
		return
	}
	return
	// }
	// info = ConvertFromPlayer(p)
	return
}

func (ps *playerService) GetPlayerInfoByName(name string, serverId int32) (info *playercommon.PlayerInfo, err error) {
	// p := GetOnlinePlayerManager().GetPlayerByName(name)
	// if p == nil {
	info, err = cache.GetCacheService().GetPlayerInfoByName(name, serverId)
	if err != nil {
		return
	}
	return
	// }
	// info = ConvertFromPlayer(p)
	return
}

func (ps *playerService) Heartbeat() {
	gameevent.Emit(playereventtypes.EventTypeOnlineNumSync, nil, nil)
}

func GetPlayerService() PlayerService {
	return ps
}

var (
	once sync.Once
	ps   *playerService
)

func Init() (err error) {
	once.Do(func() {
		ps = &playerService{}
	})
	return nil
}
