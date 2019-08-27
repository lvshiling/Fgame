package player

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/common/message"
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/heartbeat"
	coreutils "fgame/fgame/core/utils"
	playeractivity "fgame/fgame/game/activity/player"
	additionsyscommon "fgame/fgame/game/additionsys/common"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	alliancecommon "fgame/fgame/game/alliance/common"
	playeralliance "fgame/fgame/game/alliance/player"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitypes "fgame/fgame/game/anqi/types"
	playerarena "fgame/fgame/game/arena/player"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	playerbaby "fgame/fgame/game/baby/player"
	babytypes "fgame/fgame/game/baby/types"
	baguacommon "fgame/fgame/game/bagua/common"
	playerbagua "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	"fgame/fgame/game/buff/buff"
	playerbuff "fgame/fgame/game/buff/player"
	playercache "fgame/fgame/game/cache/player"
	cdcommon "fgame/fgame/game/cd/common"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	playercross "fgame/fgame/game/cross/player"
	crosssession "fgame/fgame/game/cross/session"
	crosstypes "fgame/fgame/game/cross/types"
	playerdensewat "fgame/fgame/game/densewat/player"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	playerring "fgame/fgame/game/ring/player"
	dianxingcommon "fgame/fgame/game/dianxing/common"
	playerdianxing "fgame/fgame/game/dianxing/player"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	fabaocommon "fgame/fgame/game/fabao/common"
	playerfabao "fgame/fgame/game/fabao/player"
	playerfashion "fgame/fgame/game/fashion/player"
	playerfourgod "fgame/fgame/game/fourgod/player"
	playerfuncopen "fgame/fgame/game/funcopen/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	jieyicommom "fgame/fgame/game/jieyi/common"
	playerjieyi "fgame/fgame/game/jieyi/player"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutypes "fgame/fgame/game/lingyu/types"
	playerlucky "fgame/fgame/game/lucky/player"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	playermassacre "fgame/fgame/game/massacre/player"
	massacretypes "fgame/fgame/game/massacre/types"
	"fgame/fgame/game/merge/merge"
	mountcommon "fgame/fgame/game/mount/common"
	playermount "fgame/fgame/game/mount/player"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	playerpk "fgame/fgame/game/pk/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/dao"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playerrealm "fgame/fgame/game/realm/player"
	playerrelive "fgame/fgame/game/relive/player"
	playerscene "fgame/fgame/game/scene/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatypes "fgame/fgame/game/shenfa/types"
	playershenmo "fgame/fgame/game/shenmo/player"
	playershenyu "fgame/fgame/game/shenyu/player"
	shihunfancommon "fgame/fgame/game/shihunfan/common"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	skillcommon "fgame/fgame/game/skill/common"
	playerskill "fgame/fgame/game/skill/player"
	xianzuncardcommon "fgame/fgame/game/xianzuncard/common"
	ringcommon "fgame/fgame/game/ring/common"
	"fgame/fgame/game/skill/skill"
	skilltemplate "fgame/fgame/game/skill/template"
	playersoul "fgame/fgame/game/soul/player"
	soultypes "fgame/fgame/game/soul/types"
	sysskillcommon "fgame/fgame/game/systemskill/common"
	playersystemskill "fgame/fgame/game/systemskill/player"
	teamcommon "fgame/fgame/game/team/common"
	playerteam "fgame/fgame/game/team/player"
	teamtypes "fgame/fgame/game/team/types"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotypes "fgame/fgame/game/tianmo/types"
	playertianshu "fgame/fgame/game/tianshu/player"
	playertitle "fgame/fgame/game/title/player"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	playervip "fgame/fgame/game/vip/player"
	viptypes "fgame/fgame/game/vip/types"
	playerweapon "fgame/fgame/game/weapon/player"
	weapontypes "fgame/fgame/game/weapon/types"
	wingcommon "fgame/fgame/game/wing/common"
	playerwing "fgame/fgame/game/wing/player"
	playerworldboss "fgame/fgame/game/worldboss/player"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	xianticommon "fgame/fgame/game/xianti/common"
	playerxianti "fgame/fgame/game/xianti/player"
	playerxuechi "fgame/fgame/game/xuechi/player"
	xueduncommon "fgame/fgame/game/xuedun/common"
	playerxuedun "fgame/fgame/game/xuedun/player"
	accounttypes "fgame/fgame/login/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Player struct {
	//定时
	hbRunner heartbeat.HeartbeatTaskRunner
	//锁
	m sync.Mutex
	//数据更新
	*PlayerUpdater
	//状态
	*fsm.SubjectBase
	//用户id
	userId int64
	//角色id
	playerId int64
	//第三方用户id
	platformUserId string
	//服务器id
	serverId int32
	//实名认证
	realNameState accounttypes.RealNameState
	//设备
	devicePlatformType logintypes.DevicePlatformType
	//sdk
	sdkType logintypes.SDKType
	//挂机
	guaJi bool
	gm    bool
	//外部对话
	s gamesession.Session
	//跨服对话
	crossSession crosssession.SendSession
	//跨服对话心跳
	crossTimer *time.Timer
	crossDone  chan struct{}
	//基本数据
	po *PlayerObject
	//玩家数据管理器
	playerDataManagerMap map[types.PlayerDataManagerType]player.PlayerDataManager
	//消息队列
	msgQueue *MessageQueue
	//加载过的
	loaded bool
	//buff管理器
	*buff.BuffDataManager
	//战斗管理器
	*battle.PlayerBattleManager
	//玩家系统属性
	*battle.SystemPropertyManager
	//属性管理器
	*battle.PlayerBattlePropertyManager
	//场景管理器
	*battle.PlayerSceneManager
	//pk管理器
	*battle.PlayerPKManager
	//技能动作
	*battle.SkillActionManager
	//展示管理器
	*battle.PlayerShowManager
	//竞技管理器
	*battle.PlayerArenaManager
	//竞技pvp管理器
	*battle.PlayerArenapvpManager
	//技能管理器
	*skill.SkillManager
	//血池管理器
	*battle.PlayerXueChiManager
	//复活管理器
	*battle.PlayerReliveManager
	//采集管理器
	*battle.PlayerCollectManager
	//打宝塔管理器
	*battle.PlayerTowerManager
	//天书管理器
	*battle.PlayerTianShuManager
	//幸运管理器
	*battle.PlayerLuckyManager
	//幻境BOSS管理器
	*battle.PlayerUnrealBossManager
	//外域BOSS管理器
	*battle.PlayerOutlandBossManager
	//队伍
	*battle.PlayerTeamManager
	//仙盟
	*battle.PlayerAllianceManager
	*battle.PlayerJieYiManager
	*battle.PlayerZhenYingManager
	*battle.PlayerBossReliveManager
	//神兽攻城管理器
	*battle.PlayerGodSiegeManager
	//挂机管理器
	*battle.PlayerGuaJiManager
	//状态数据管理器
	*battle.StateDataManager
	//灵童管理器
	*battle.PlayerLingTongShowManager
	//金银密窟
	*battle.PlayerDenseWatManager
	//神魔战场
	*battle.PlayerShenMoManager
	//活动数据
	*battle.PlayerActivityManager
	//特殊技能
	*battle.TeShuSkillManager
	//移动
	*battle.MoveAction
	//cd组管理器
	cdGroupManager *cdcommon.CDGroupManager

	//登出
	logouting bool
	//是否无间炼狱排队
	isLianYuLineUp bool
	//是否神魔战场排队
	isShenMoLineUp bool
	//是否排队
	isLineup bool

	done chan struct{}
	//测试使用
	randomLogoutState fsm.State

	//整合到技能管理器
	skillActionTime int64
	skillTime       int64
}

//-----------------------------battle object 接口实现 ----------------------------------------
//更新系统属性
func (p *Player) UpdateBattleProperty(mask uint64) {
	//更新属性
	ppdm := p.getPlayerPropertyManager()
	//更新战斗属性
	ppdm.UpdateBattleProperty(mask)
}

//获取系统属性
func (p *Player) GetSystemBattleProperty(t propertytypes.BattlePropertyType) int64 {
	m := p.getPlayerPropertyManager()
	return m.GetBattleProperty(t)
}

//使用技能
func (p *Player) UseSkill(skillId int32) bool {
	flag := p.SkillManager.UseSkill(skillId)
	if !flag {
		return false
	}
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)

	p.SkillActionManager.AddSkillAction(skillId)
	//技能动作
	if skillTemplate.IsPositive() {
		p.GuaJiAttack()
		now := global.GetGame().GetTimeService().Now()
		p.skillTime = now
		p.skillActionTime = int64(skillTemplate.ActionTime)
	}
	return true
}

func (p *Player) GetSkillTime() int64 {
	return p.skillTime
}

func (p *Player) GetSkillActionTime() int64 {
	return p.skillActionTime
}

//----------------------------------scene object接口 ---------------------------------

//进入场景
func (p *Player) EnteringScene() (flag bool) {
	p.m.Lock()
	defer p.m.Unlock()
	flag = player.GetPlayerStateMachine().Trigger(p, player.EventPlayerEnterScene)
	return
}

//进入游戏
func (p *Player) EnterGame() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerGaming)
	return flag
}

//退出场景
func (p *Player) LeaveScene() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLeaveScene)
	if !flag {
		return flag
	}
	p.msgQueue.Pause()

	return flag
}

func (p *Player) SetCrossSession(sess crosssession.SendSession) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.crossSession == nil {
		p.crossSession = sess
		p.startCrossHeartbeat()
		p.crossDone = make(chan struct{})
	}
	return
}

func (p *Player) startCrossHeartbeat() {
	p.crossTimer.Reset(crossTimer)
	go func() {
	Loop:
		for {
			select {
			case <-p.crossTimer.C:
				gameevent.Emit(crosseventtypes.EventTypePlayerCrossHeartbeat, p, nil)
				//发送跨服心跳
				p.crossTimer.Reset(crossTimer)
			case <-p.crossDone:
				break Loop
			}
		}
	}()
}

//进入跨服
func (p *Player) EnterCross() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerEnterCross)
	return flag
}

//进入跨服
func (p *Player) Cross() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerCrossing)
	return flag
}

//退出跨服
func (p *Player) LeaveCross() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLeaveCross)
	if !flag {
		return flag
	}
	p.msgQueue.Pause()
	return flag
}

//退出跨服
func (p *Player) GetCrossSession() crosssession.SendSession {
	return p.crossSession
}

func (p *Player) IsCross() bool {
	switch p.CurrentState() {
	case player.PlayerStateCrossing:
		return true
	}
	return false
}

//------------------------------玩家接口 --------------------------------
func (p *Player) EnterAuth() bool {
	p.m.Lock()
	defer p.m.Unlock()
	//随机卡死状态

	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerAuth)
	if !flag {
		return false
	}
	return true
}

func (p *Player) EnterLoadingRoleList() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLoadingRoleList)
	if !flag {
		return false
	}
	return true
}

func (p *Player) EnterCreateRole() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerCreatingRole)
	if !flag {
		return false
	}
	return true
}

func (p *Player) EnterWaitingSelectRole() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerWaitingSelectRole)
	if !flag {
		return false
	}
	return true
}

func (p *Player) Auth(playerId int64) bool {
	p.m.Lock()
	defer p.m.Unlock()
	// if pe == nil {
	// 	panic(fmt.Errorf("player:基本信息应该不为空"))
	// }

	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerSelectRole)
	if !flag {
		return false
	}
	p.playerId = playerId
	p.PlayerUpdater = NewPlayerUpdater()
	p.initManagers()
	// p.po = NewPlayerObject(p)
	// p.po.FromEntity(pe)

	return true
}

//初始化管理器数据
func (p *Player) initManagers() {
	p.playerDataManagerMap = make(map[types.PlayerDataManagerType]player.PlayerDataManager)
	for typ, factory := range player.GetPlayerDataManagerMap() {
		p.playerDataManagerMap[typ] = factory.CreatePlayerDataManager(p)
	}
}

func (p *Player) EnterLoad() bool {
	p.m.Lock()
	defer p.m.Unlock()

	flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLoading)
	if !flag {
		return false
	}

	return true
}

//玩家加载数据
func (p *Player) Load() (err error) {
	p.m.Lock()
	defer p.m.Unlock()
	//加载用户数据
	pe, err := dao.GetPlayerDao().QueryByUserId(p.userId, p.serverId)
	if err != nil {
		return
	}
	p.po = NewPlayerObject(p)
	p.po.FromEntity(pe)

	//所有数据完成
	for t, m := range p.playerDataManagerMap {
		// beforeTime := global.GetGame().GetTimeService().Now()
		err = m.Load()
		if err != nil {
			return errors.WithMessage(err, t.String())
		}
		// costTime := global.GetGame().GetTimeService().Now() - beforeTime
		// log.WithFields(
		// 	log.Fields{
		// 		"playerId": p.GetId(),
		// 		"module":   t.String(),
		// 		"costTime": costTime,
		// 	}).Info("player:加载数据")
	}

	return
}

//加载后操作
func (p *Player) AfterLoad() (flag bool, err error) {
	p.m.Lock()
	defer p.m.Unlock()
	flag = player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLoaded)
	if !flag {
		return false, nil
	}
	//cd组
	p.cdGroupManager = cdcommon.NewCDGroupManager()

	for t, pdm := range p.playerDataManagerMap {
		err = pdm.AfterLoad()
		if err != nil {
			return false, errors.WithMessage(err, t.String())
		}
	}

	//TODO 更新红名值
	//buff管理器
	buffs := p.getPlayerBuffManager().GetBuffs()
	p.BuffDataManager = buff.CreateBuffDataManagerWithBuffs(p, buffs)
	vipManager := p.getPlayerVipManager()
	vip, _ := vipManager.GetVipLevel()

	propertyManager := p.getPlayerPropertyManager()
	level := propertyManager.GetLevel()
	zhuanSheng := propertyManager.GetZhuanSheng()
	soulAwakenNum := p.getPlayerSoulManager().GetAwakenNum()
	isHuiYuanPlus := p.getPlayerHuiYuanManager().IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	battleObject := battlecommon.CreatePlayerBattleObject(vip, level, zhuanSheng, soulAwakenNum, isHuiYuanPlus)
	//战斗管理器
	p.PlayerBattleManager = battle.CreatePlayerBattleManagerWithObject(p, false, battleObject, scenetypes.FactionTypePlayer)
	//获取数据
	playerPropertyManager := p.getPlayerPropertyManager()
	hp := playerPropertyManager.GetHP()
	tp := playerPropertyManager.GetTP()

	systemProperties := playerPropertyManager.GetChangedBattlePropertiesAndReset()
	//系统属性管理器
	p.SystemPropertyManager = battle.CreateSystemPropertyManagerWithData(p, systemProperties)
	//战斗属性管理器
	p.PlayerBattlePropertyManager = battle.CreatePlayerBattlePropertyManager(p, hp, tp, 0)

	//pk管理器
	playerPkManager := p.getPkDataManager()
	p.PlayerPKManager = battle.CreatePlayerPKManagerWithObject(p, playerPkManager)
	allSkillList := p.getPlayerSkillManager().GetAllSkill()
	p.SkillManager = skill.CreateSkillManager(p, p.GetCDGroupManager(), allSkillList)
	p.SkillActionManager = battle.CreateSkillActionManager(p)

	//打宝塔管理器
	p.PlayerTowerManager = battle.CreatePlayerTowerManager(p)

	//场景管理器
	playerSceneManager := p.getPlayerSceneManager()
	playerSceneObject := playerSceneManager.GetPlayerScene()
	p.PlayerSceneManager = battle.CreatePlayerSceneManagerWithObject(p, playerSceneObject)

	fashionManager := p.getPlayerFashionManager()
	fashionId := fashionManager.GetFashionId()
	weaponManager := p.getPlayerWeaponManager()
	weaponId := weaponManager.GetWeaponWear()
	weaponState := weaponManager.GetWeaponState(weaponId)
	titleManager := p.getPlayerTitleManager()
	titleId := titleManager.GetTitleId()
	wingManager := p.getPlayerWingManager()
	wingId := wingManager.GetWingId()
	mountManager := p.getPlayerMountManager()
	mountId := int32(0)
	mountAdvanceId := int32(0)
	if p.IsFuncOpen(funcopentypes.FuncOpenTypeMount) {
		mountId = mountManager.GetMountId()
		mountAdvanceId = mountManager.GetMountAdvancedId()
	}

	mountHidden := mountManager.IsHidden()
	shenFaManager := p.getPlayerShenfaManager()
	shenFaId := shenFaManager.GetShenFaId()
	lingYuManager := p.getPlayerLingyuManager()
	lingYuId := lingYuManager.GetLingYuId()
	realmManager := p.getPlayerRealmManager()
	realmLevel := realmManager.GetTianJieTaLevel()
	fourGodManager := p.getPlayerFourGodManager()
	key := fourGodManager.GetKeyNum()
	marryManager := p.getPlayerMarryManager()
	spouse := marryManager.GetSpouseName()
	spouseId := marryManager.GetSpouseId()
	weddingStatus := int32(marryManager.GetWedStatus())
	ringType := marryManager.GetRingType()
	ringLevel := marryManager.GetRingLevel()
	faBaoManager := p.getPlayerFaBaoManager()
	faBaoId := faBaoManager.GetFaBaoId()
	baGuaManager := p.getPlayerBaGuaManager()
	baGua := baGuaManager.GetLevel()
	petId := int32(0)
	xianTiManager := p.getPlayerXianTiManager()
	xianTiId := xianTiManager.GetXianTiId()
	flyPetId := int32(0)
	developLevel := marryManager.GetMarryDevelopLevel()
	shenYuManager := p.getPlayerShenYuManager()
	shenYuKey := shenYuManager.GetKeyNum()
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		int32(weaponState),
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		key,
		realmLevel,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	//展示
	p.PlayerShowManager = battle.CreatePlayerShowManagerWithObject(p, showObj)
	arenaManager := p.getPlayerArenaManager()
	arenaReliveTime := arenaManager.GetPlayerArenaObject().GetReliveTime()
	arenaWinTime := arenaManager.GetPlayerArenaObject().GetWinCount()
	arenaObj := battle.CreatePlayerArenaObject(arenaReliveTime, arenaWinTime)
	p.PlayerArenaManager = battle.CreatePlayerArenaManagerWithObject(p, arenaObj)

	arenapvpManager := p.getPlayerArenapvpManager()
	plArenapvpObj := arenapvpManager.GetPlayerArenapvpObj()
	arenapvpObj := battle.CreatePlayerArenapvpObject(plArenapvpObj.GetReliveTimes())
	p.PlayerArenapvpManager = battle.CreatePlayerArenapvpManagerWithObject(p, arenapvpObj)

	xueChiManager := p.getPlayerXueChiManager()
	bloodLine := xueChiManager.GetXueChi().BloodLine
	blood := xueChiManager.GetXueChi().Blood
	p.PlayerXueChiManager = battle.CreatePlayerXueChiManager(p, blood, bloodLine)
	reliveManager := p.getPlayerReliveManager()
	culTime := reliveManager.GetCulTime()
	lastReliveTime := reliveManager.GetLastReliveTime()
	p.PlayerReliveManager = battle.CreatePlayerReliveManager(p, culTime, lastReliveTime)
	p.PlayerCollectManager = battle.CreatePlayerCollectManager(p)
	teamObj := teamcommon.CreatePlayerTeamObject(0, "", teamtypes.TeamPurposeTypeNormal)
	p.PlayerTeamManager = battle.CreatePlayerTeamManagerWithObject(p, teamObj)
	allianceManager := p.getPlayerAllianceManager()
	allianceId := allianceManager.GetAllianceId()
	allianceName := allianceManager.GetAllianceName()
	mengZhuId := allianceManager.GetMengZhuId()
	memPos := allianceManager.GetPlayerAlliancePos()
	allianceObj := alliancecommon.CreatePlayerAllianceObject(allianceId, allianceName, mengZhuId, memPos)
	p.PlayerAllianceManager = battle.CreatePlayerAllianceManagerWithObject(p, allianceObj)
	jieYiManager := p.getPlayerJieYiManager()
	jieYiName := jieYiManager.GetJieYiName()
	jieYiRank := jieYiManager.GetJieYiRank()
	jieYiId := jieYiManager.GetJieYiId()

	jieYiObj := jieyicommom.CreatePlayerJieYiObject(jieYiId, jieYiName, jieYiRank)
	p.PlayerJieYiManager = battle.CreatePlayerJieYiManagerWithObject(p, jieYiObj)

	// chuangShiManager := p.getPlayerChuangShiManager()
	// chuangShiInfo := chuangShiManager.GetPlayerChuangShiInfo()
	campType := chuangshitypes.RandomChuangShiCamp()
	guanZhiPos := chuangshitypes.RandomChuangShiGuanZhi()
	p.PlayerZhenYingManager = battle.CreatePlayerZhenYingManager(p, campType, guanZhiPos)

	worldbossManager := p.getPlayerWorldbossManager()
	bossReliveList := worldbossManager.GetBossReliveList()
	//TODO
	p.PlayerBossReliveManager = battle.CreatePlayerBossReliveManager(p, bossReliveList)
	tianshuManager := p.getPlayerTianShuManager()
	p.PlayerTianShuManager = battle.CreatePlayerTianShuManagerWithData(p, tianshuManager.GetTianShuAll())
	p.PlayerGodSiegeManager = battle.CreatePlayerGodSiegeManager(p)
	luckyManager := p.getPlayerLuckyManager()
	p.PlayerLuckyManager = battle.CreatePlayerLuckyManagerWithData(p, luckyManager.GetLuckyInfoAll())
	p.PlayerGuaJiManager = battle.CreatePlayerGuaJiManager(p)
	p.StateDataManager = battle.CreateStateDateManager(p)
	unrealManager := p.getPlayerUnrealBossManager()
	curPilaoNum := unrealManager.GetCurPilaoNum()
	p.PlayerUnrealBossManager = battle.CreatePlayerUnrealBossManagerWithData(p, curPilaoNum)
	outlandBossManager := p.getPlayerOutlandBossManager()
	curZhuoQiNum := outlandBossManager.GetCurZhuoQiNum()
	p.PlayerOutlandBossManager = battle.CreatePlayerOutlandBossManagerWithData(p, curZhuoQiNum)

	p.PlayerLingTongShowManager = battle.CreatePlayerLingTongShowManager(p)
	//金银密窟
	denseWatManager := p.getPlayerDenseWatManager()
	num := denseWatManager.GetDenseWatInfo().GetNum()
	endTime := denseWatManager.GetDenseWatInfo().GetEndTime()
	p.PlayerDenseWatManager = battle.CreatePlayerDenseWatManager(p, num, endTime)

	//神魔战场
	shenMoManager := p.getPlayerShenMoManager()
	gongXunNum := shenMoManager.GetShenMoInfo().GetGongXunNum()
	killNum := shenMoManager.GetShenMoInfo().GetKillNum()
	shenMoEndTime := shenMoManager.GetShenMoInfo().GetEndTime()
	p.PlayerShenMoManager = battle.CreatePlayerShenMoManager(p, gongXunNum, killNum, shenMoEndTime)

	activityManager := p.getPlayerActivityManager()
	activityPkDataList := activityManager.GetActivityPkDataList()
	activityRankDataList := activityManager.GetActivityRankDataList()
	activityCollectDataList := activityManager.GetActivityCollectDataList()
	p.PlayerActivityManager = battle.CreatePlayerActivityManager(p, activityPkDataList, activityRankDataList, activityCollectDataList, nil)

	goldequipManager := p.getPlayerGoldEquiManager()
	teShuSkillList := goldequipManager.GetTeShuSkillList()
	p.TeShuSkillManager = battle.CreateTeShuSkillManager(p, teShuSkillList)
	p.MoveAction = battle.CreateMoveAction(p)

	p.loaded = true

	//检查离线时间
	p.checkOfflineTime()

	err = p.checkOnlineTime()
	if err != nil {
		return
	}
	//刷新今日充值
	p.refreshTodayCharge()

	//更新登陆ip
	p.updateLoginIp()

	//定时任务
	p.hbRunner.AddTask(CreateOnlineTimeChangedTask(p))

	//更新战斗属性
	p.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)
	//同步缓存
	p.syncCache()
	return
}

//同步缓存
func (p *Player) syncCache() {
	cacheManager := p.GetPlayerDataManager(types.PlayerCacheDataManagerType).(*playercache.PlayerCacheDataManager)
	cacheManager.SyncCache()

	return
}

func (p *Player) updateLoginIp() {
	now := global.GetGame().GetTimeService().Now()
	ip, _ := coreutils.SplitIp(p.s.Ip())
	p.po.Ip = ip
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) checkOnlineTime() error {
	now := global.GetGame().GetTimeService().Now()
	//是否跨天
	flag, err := timeutils.IsSameDay(p.po.LastLogoutTime, now)
	if err != nil {
		return err
	}
	if !flag {
		p.po.TodayOnlineTime = 0
	}
	p.po.Online = 1
	p.po.SetModified()
	return nil
}

//检查离线时间
func (p *Player) checkOfflineTime() {
	now := global.GetGame().GetTimeService().Now()
	p.po.LastLoginTime = now
	//沉迷了
	if p.IsWallow() {
		//计算离线时间
		offlineTime := now - p.po.LastLogoutTime + p.po.OfflineTime
		//离线时间超过了 清空
		if offlineTime >= fiveHourWallowTime {
			p.po.OfflineTime = 0
			p.po.OnlineTime = 0
		} else {
			p.po.OfflineTime = offlineTime
		}
	}
	p.po.SetModified()
	return
}

//刷新每日充值信息
func (p *Player) refreshTodayCharge() {
	now := global.GetGame().GetTimeService().Now()

	diff, _ := timeutils.DiffDay(now, p.po.ChargeTime)
	if diff != 0 {
		if diff == 1 {
			p.po.YesterdayChargeMoney = p.po.TodayChargeMoney
		} else {
			p.po.YesterdayChargeMoney = 0
		}
		p.po.TodayChargeMoney = 0
		p.po.ChargeTime = now
		p.po.SetModified()
	}
}

func (p *Player) IsWallow() bool {
	if p.realNameState == accounttypes.RealNameStateUp18 {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	onlineTime := p.po.OnlineTime + (now - p.po.LastLoginTime)
	if onlineTime >= threeHourWallowTime {
		return true
	}
	return false
}

func (p *Player) GetLastLogoutTime() int64 {
	if p.po == nil {
		return 0
	}
	return p.po.LastLogoutTime
}

func (p *Player) GetCreateTime() int64 {
	if p.po == nil {
		return 0
	}
	return p.po.CreateTime
}

func (p *Player) GetTodayOnlineTime() int64 {
	if p.po == nil {
		return 0
	}
	now := global.GetGame().GetTimeService().Now()
	//是否跨天
	flag, err := timeutils.IsSameDay(p.po.LastLoginTime, now)
	if err != nil {
		panic(fmt.Errorf("player:%s", err.Error()))
	}
	if flag {
		return p.po.TodayOnlineTime + now - p.po.LastLoginTime
	}
	begin, err := timeutils.BeginOfNow(now)
	if err != nil {
		panic(fmt.Errorf("player:%s", err.Error()))
	}
	return now - begin
}

func (p *Player) getOnlineTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	return p.po.OnlineTime + (now - p.po.LastLoginTime)
}

func (p *Player) GetWallowState() types.WallowState {
	if p.realNameState == accounttypes.RealNameStateUp18 {
		return types.WallowStateNone
	}
	onlineTime := p.getOnlineTime()
	if onlineTime >= fiveHourWallowTime {
		return types.WallowStateFiveHour
	} else if onlineTime >= threeHourWallowTime {
		return types.WallowStateThreeHour
	}
	return types.WallowStateNone
}

func (p *Player) GetOnlineTime() int64 {
	return p.po.OnlineTime
}

func (p *Player) GetTotalOnlineTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	return p.po.TotalOnlineTime + (now - p.po.LastLoginTime)
}

func (p *Player) Post(msg message.Message) {
	//TODO 可能造成死锁  做个超时处理
	p.msgQueue.Post(msg)
}

//tick
func (p *Player) Tick() {
	p.msgQueue.Tick()
}

//玩家定时心跳
func (p *Player) Heartbeat() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"error": terr,
					"stack": string(debug.Stack()),
				}).Error("player:Heartbeat,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()
	//创建更新操作
	p.Update()

	//TODO
	for _, manager := range p.playerDataManagerMap {
		manager.Heartbeat()
	}

	p.hbRunner.Heartbeat()

	if p.IsCross() {
		return
	}
	if p.GetScene() == nil {
		return
	}
	p.PlayerPKManager.Heartbeat()
	p.BuffDataManager.Heartbeat()
	//战斗
	p.PlayerBattleManager.Heartbeat()
	p.SkillActionManager.Heartbeat()
	p.MoveAction.Heartbeat()
	p.PlayerBattlePropertyManager.Heartbeat()
	p.PlayerXueChiManager.Heartbeat()
	p.PlayerReliveManager.Heartbeat()
	p.PlayerGuaJiManager.Heartbeat()
}

func (p *Player) IsLianYuLineUp() bool {
	return p.isLianYuLineUp
}

func (p *Player) LianYuLineUp(isLianYuLine bool) {
	p.m.Lock()
	defer p.m.Unlock()
	p.isLianYuLineUp = isLianYuLine
}

func (p *Player) IsLineUp() bool {
	return p.isLineup
}

func (p *Player) SetLineUp(linup bool) {
	p.m.Lock()
	defer p.m.Unlock()
	p.isLineup = linup
}

func (p *Player) IsShenMoLineUp() bool {
	return p.isShenMoLineUp
}

func (p *Player) ShenMoLineUp(isLianYuLine bool) {
	p.m.Lock()
	defer p.m.Unlock()
	p.isShenMoLineUp = isLianYuLine
}

func (p *Player) IsLogouting() bool {
	return p.logouting
}

func (p *Player) Logout() bool {
	p.m.Lock()
	defer p.m.Unlock()
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家正在登出")
	switch p.CurrentState() {
	case player.PlayerStateInit,
		player.PlayerStateAuth,
		player.PlayerStateLoadingRoleList,
		player.PlayerStateWaitingSelectRole,
		player.PlayerStateCreatingRole,
		player.PlayerStateSelectRole,
		player.PlayerStateLoading:
		defer func() {
			gameevent.Emit(playereventtypes.EventTypePlayerLogoutBeforeLoaded, p, nil)
		}()
		//以防panic
		flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLogout)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"state":    p.CurrentState(),
				}).Info("player:登出失败")
			return false
		}
		break
	case player.PlayerStateLoaded,
		player.PlayerStateEnterScene,
		player.PlayerStateLeaveScene,
		player.PlayerStateEnterCross,
		player.PlayerStateLeaveCross:
		defer func() {
			gameevent.Emit(playereventtypes.EventTypePlayerBeforeLogout, p, nil)
		}()
		// p.logouting = true
		//以防panic
		flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLogout)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"state":    p.CurrentState(),
				}).Info("player:登出失败")
			return false
		}
		return true
	case player.PlayerStateGaming:
		p.logouting = true
		gameevent.Emit(playereventtypes.EventTypePlayerExitSceneBeforeLogout, p, nil)
		return true
	case player.PlayerStateCrossing:
		p.logouting = true
		gameevent.Emit(playereventtypes.EventTypePlayerExitCrossBeforeLogout, p, nil)
		return true
	case player.PlayerStateLogouting,
		player.PlayerStateLogouted:
		return true
	}

	return true
}

func (p *Player) LogoutCross() {
	p.m.Lock()
	defer p.m.Unlock()
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家退出跨服")
	if p.crossSession != nil {
		p.crossSession.Close(false)
		close(p.crossDone)
		p.crossSession = nil
		p.crossTimer.Stop()
	}
	switch p.CurrentState() {
	case player.PlayerStateGaming:
		gameevent.Emit(playereventtypes.EventTypePlayerLogoutCrossInGame, p, nil)
		break
	case player.PlayerStateCrossing:
		gameevent.Emit(playereventtypes.EventTypePlayerLogoutCrossInCross, p, nil)
		break
	case player.PlayerStateEnterCross,
		player.PlayerStateLeaveCross:
		gameevent.Emit(playereventtypes.EventTypePlayerLogoutCrossInGlobal, p, nil)
		break
	}

}

func (p *Player) Done() <-chan struct{} {
	return p.done
}

func (p *Player) LogoutSave() bool {
	if !p.logout() {
		return false
	}
	return p.logoutSave()
}
func (p *Player) logout() bool {
	p.m.Lock()
	defer p.m.Unlock()
	if p.logouting {
		flag := player.GetPlayerStateMachine().Trigger(p, player.EventPlayerLogout)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"state":    p.CurrentState(),
				}).Info("player:登出失败")
			return false
		}
	}
	return true
}

func (p *Player) logoutSave() bool {

	defer func() {
		close(p.done)
	}()
	if p.po == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("player:玩家角色还没加载,登出成功")

		return true
	}
	if !p.loaded {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("player:玩家角色还没加载,登出成功")
		return true
	}

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色登出,成功")

	//TODO 改在事件处理
	p.PlayerPKManager.Logout()

	now := global.GetGame().GetTimeService().Now()
	onlineTime := now - p.po.LastLoginTime
	//设置在线时间
	flag, _ := timeutils.IsSameDay(now, p.po.LastLoginTime)
	if !flag {
		beginNow, _ := timeutils.BeginOfNow(now)
		p.po.TodayOnlineTime = now - beginNow
	} else {
		p.po.TodayOnlineTime += onlineTime
	}
	p.po.Online = 0
	p.po.TotalOnlineTime += onlineTime
	p.po.OnlineTime += onlineTime
	p.po.LastLogoutTime = now
	//修改
	p.po.SetModified()
	p.syncCache()

	gameevent.Emit(playereventtypes.EventTypePlayerLogout, p, nil)
	//保存
	p.Update()
	return true
}

func (p *Player) SendMsg(msg proto.Message) error {
	p.Session().Send(msg)
	return nil
}

func (p *Player) SendCrossMsg(msg proto.Message) {
	if p.crossSession == nil {
		return
	}
	p.crossSession.Send(msg)
}

//关闭
func (p *Player) Close(err error) {

	log.WithFields(
		log.Fields{
			"userId":   p.GetUserId(),
			"playerId": p.GetId(),
			"err":      err,
		}).Error("player:玩家错误关闭")
	p.Session().Close(true)
}

func (p *Player) Session() gamesession.Session {
	return p.s
}

func (p *Player) GetPlayerObject() *PlayerObject {
	return p.po
}

func (p *Player) GetUserId() int64 {
	return p.userId
}

func (p *Player) GetPlatformUserId() string {
	return p.platformUserId
}

func (p *Player) GetIp() string {
	return p.s.Ip()
}

func (p *Player) GetPlayerDataManager(typ types.PlayerDataManagerType) (pdm player.PlayerDataManager) {
	pdm, exist := p.playerDataManagerMap[typ]
	if !exist {
		return nil
	}
	return pdm
}

func (p *Player) getPlayerSceneManager() (ppdm *playerscene.PlayerSceneDataManager) {
	pdm := p.GetPlayerDataManager(types.PlayerSceneDataManagerType)
	if pdm == nil {
		return
	}
	ppdm, ok := pdm.(*playerscene.PlayerSceneDataManager)
	if !ok {
		return nil
	}
	return
}
func (p *Player) getPlayerPropertyManager() (ppdm *playerproperty.PlayerPropertyDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerPropertyDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerproperty.PlayerPropertyDataManager)
	if !ok {
		return nil
	}
	return
}
func (p *Player) getPlayerFashionManager() (m *playerfashion.PlayerFashionDataManager) {
	tm := p.playerDataManagerMap[types.PlayerFashionDataManagerType]
	m, ok := tm.(*playerfashion.PlayerFashionDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerSkillManager() (ppdm *playerskill.PlayerSkillDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerSkillDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerskill.PlayerSkillDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPkDataManager() *playerpk.PlayerPkDataManager {
	pdm, _ := p.playerDataManagerMap[types.PlayerPkDataManagerType]
	ppdm, _ := pdm.(*playerpk.PlayerPkDataManager)
	return ppdm
}

func (p *Player) getPlayerBuffManager() (ppdm *playerbuff.PlayerBuffDataManager) {
	pdm := p.GetPlayerDataManager(types.PlayerBuffDataManagerType)
	if pdm == nil {
		return
	}
	ppdm, ok := pdm.(*playerbuff.PlayerBuffDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerMountManager() (m *playermount.PlayerMountDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerMountDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playermount.PlayerMountDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerBodyshieldManager() (m *playerbodyshield.PlayerBodyShieldDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerBShieldDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerbodyshield.PlayerBodyShieldDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerAnqiManager() (m *playeranqi.PlayerAnqiDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerAnqiDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playeranqi.PlayerAnqiDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerMassacreManager() (m *playermassacre.PlayerMassacreDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerMassacreDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playermassacre.PlayerMassacreDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerDianXingManager() (m *playerdianxing.PlayerDianXingDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerDianXingDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerdianxing.PlayerDianXingDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerTianMoManager() (m *playertianmo.PlayerTianMoDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerTianMoDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playertianmo.PlayerTianMoDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerShiHunFanManager() (m *playershihunfan.PlayerShiHunFanDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playershihunfan.PlayerShiHunFanDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerWingManager() (m *playerwing.PlayerWingDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerWingDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerwing.PlayerWingDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerFaBaoManager() (m *playerfabao.PlayerFaBaoDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerFaBaoDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerfabao.PlayerFaBaoDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerXueDunManager() (m *playerxuedun.PlayerXueDunDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerXueDunDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerxuedun.PlayerXueDunDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerXianTiManager() (m *playerxianti.PlayerXianTiDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerXianTiDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerxianti.PlayerXianTiDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerWeaponManager() (m *playerweapon.PlayerWeaponDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerWeaponDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerweapon.PlayerWeaponDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerTitleManager() (m *playertitle.PlayerTitleDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerTitleDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playertitle.PlayerTitleDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerInventoryManager() (m *playerinventory.PlayerInventoryDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerInventoryDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerinventory.PlayerInventoryDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerGoldEquiManager() (m *playergoldequip.PlayerGoldEquipDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playergoldequip.PlayerGoldEquipDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerWushuangManager() (m *playerwushuangweapon.PlayerWushuangWeaponDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerWushuangWeaponDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerXianZunCardManager() (m *playerxianzuncard.PlayerXianZunCardDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerXianZunCardManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerxianzuncard.PlayerXianZunCardDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerRingManager() (m *playerring.PlayerRingDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerRingDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerring.PlayerRingDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerSoulManager() (m *playersoul.PlayerSoulDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerSoulDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playersoul.PlayerSoulDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerHuiYuanManager() (m *playerhuiyuan.PlayerHuiYuanManager) {
	tm := p.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerhuiyuan.PlayerHuiYuanManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerAllianceManager() (ppdm *playeralliance.PlayerAllianceDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerAllianceDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playeralliance.PlayerAllianceDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerJieYiManager() (ppdm *playerjieyi.PlayerJieYiDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerJieYiDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerjieyi.PlayerJieYiDataManager)
	if !ok {
		return nil
	}
	return
}
func (p *Player) getPlayerChuangShiManager() (ppdm *playerchuangshi.PlayerChuangShiDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerChuangShiDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerchuangshi.PlayerChuangShiDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerActivityManager() (ppdm *playeractivity.PlayerActivityDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerActivityDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playeractivity.PlayerActivityDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerTeamManager() (ppdm *playerteam.PlayerTeamDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerTeamDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerteam.PlayerTeamDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerMarryManager() (m *playermarry.PlayerMarryDataManager) {
	md := p.GetPlayerDataManager(types.PlayerMarryDataManagerType)
	if md == nil {
		return
	}
	m, ok := md.(*playermarry.PlayerMarryDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerFourGodManager() (ppdm *playerfourgod.PlayerFourGodDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerFourGodDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerfourgod.PlayerFourGodDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerCrossManager() (ppdm *playercross.PlayerCrossDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerCrossDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playercross.PlayerCrossDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerVipManager() (ppdm *playervip.PlayerVipDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerVipDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playervip.PlayerVipDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerTianShuManager() (ppdm *playertianshu.PlayerTianShuDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerTianShuDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playertianshu.PlayerTianShuDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerLuckyManager() (ppdm *playerlucky.PlayerLuckyDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerLuckyDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerlucky.PlayerLuckyDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerUnrealBossManager() (ppdm *playerunrealboss.PlayerUnrealBossDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerUnrealBossDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playerunrealboss.PlayerUnrealBossDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerOutlandBossManager() (ppdm *playeroutlandboss.PlayerOutlandBossDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerOutlandBossDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playeroutlandboss.PlayerOutlandBossDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) GetCDGroupManager() *cdcommon.CDGroupManager {
	return p.cdGroupManager
}

func (p *Player) getPlayerShenfaManager() (m *playershenfa.PlayerShenfaDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerShenfaDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playershenfa.PlayerShenfaDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerLingyuManager() (m *playerlingyu.PlayerLingyuDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerLingyuDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerlingyu.PlayerLingyuDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerRealmManager() (m *playerrealm.PlayerRealmDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerRealmDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerrealm.PlayerRealmDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerBaGuaManager() (m *playerbagua.PlayerBaGuaDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerBaGuaDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerbagua.PlayerBaGuaDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerArenaManager() (m *playerarena.PlayerArenaDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerArenaDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerarena.PlayerArenaDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerWorldbossManager() (m *playerworldboss.PlayerWorldbossManager) {
	tm := p.GetPlayerDataManager(types.PlayerWorldbossManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerworldboss.PlayerWorldbossManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerArenapvpManager() (m *playerarenapvp.PlayerArenapvpDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerArenapvpDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerarenapvp.PlayerArenapvpDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerXueChiManager() (m *playerxuechi.PlayerXueChiDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerXueChiDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerxuechi.PlayerXueChiDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerReliveManager() (m *playerrelive.PlayerReliveDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerReliveDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerrelive.PlayerReliveDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerDenseWatManager() (m *playerdensewat.PlayerDenseWatDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerDenseWatDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerdensewat.PlayerDenseWatDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerShenMoManager() (m *playershenmo.PlayerShenMoDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerShenMoWarDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playershenmo.PlayerShenMoDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerLingTongDevManager() (m *playerlingtongdev.PlayerLingTongDevDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerlingtongdev.PlayerLingTongDevDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerLingTongManager() (m *playerlingtong.PlayerLingTongDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerLingTongDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerlingtong.PlayerLingTongDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerSystemSkillManager() (m *playersystemskill.PlayerSystemSkillDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playersystemskill.PlayerSystemSkillDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerAdditionSysManager() (m *playeradditionsys.PlayerAdditionSysDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playeradditionsys.PlayerAdditionSysDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerBabyMananger() (m *playerbaby.PlayerBabyDataManager) {
	tm := p.GetPlayerDataManager(types.PlayerBabyDataManagerType)
	if tm == nil {
		return
	}
	m, ok := tm.(*playerbaby.PlayerBabyDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) getPlayerShenYuManager() (ppdm *playershenyu.PlayerShenYuDataManager) {
	pdm, exist := p.playerDataManagerMap[types.PlayerShenYuDataManagerType]
	if !exist {
		return nil
	}
	ppdm, ok := pdm.(*playershenyu.PlayerShenYuDataManager)
	if !ok {
		return nil
	}
	return
}

func (p *Player) RealNameAuth(realNameState accounttypes.RealNameState) {
	p.realNameState = realNameState
}

func (p *Player) GetRealNameState() accounttypes.RealNameState {
	return p.realNameState
}

func (p *Player) GetSDKType() logintypes.SDKType {
	return p.sdkType
}

func (p *Player) GetDevicePlatformType() logintypes.DevicePlatformType {
	return p.devicePlatformType
}

const (
	threeHourWallowTime = int64(3 * time.Hour / time.Millisecond)
	fiveHourWallowTime  = int64(5 * time.Hour / time.Millisecond)
)

func (p *Player) GetContext() context.Context {
	return p.Session().Context()
}

func (p *Player) GetId() int64 {
	return p.playerId
}

//快捷操作
func (p *Player) GetRole() types.RoleType {
	return p.po.Role
}

func (p *Player) GetSex() types.SexType {
	return p.po.Sex
}

func (p *Player) GetName() string {
	if p.po != nil {
		if merge.GetMergeService().GetMergeTime() > 0 {
			return fmt.Sprintf("s%d.%s", p.po.OriginServerId, p.po.Name)
		}
		return p.po.Name
	}
	return ""
}

func (p *Player) GetOriginName() string {
	if p.po != nil {
		return p.po.Name
	}
	return ""
}

func (p *Player) GetServerId() int32 {
	if p.serverId != 0 {
		return p.serverId
	}
	return p.po.OriginServerId
}

func (p *Player) GetPlatform() int32 {
	return global.GetGame().GetPlatform()
}

//坐骑
func (p *Player) GetMountInfo() *mountcommon.MountInfo {
	m := p.getPlayerMountManager()
	mountInfo := m.ToMountInfo()
	return mountInfo
}

//护体盾
func (p *Player) GetBodyshieldInfo() *bodyshieldtypes.BodyShieldInfo {
	m := p.getPlayerBodyshieldManager()
	bodyShieldInfo := m.ToBodyShieldInfo()
	return bodyShieldInfo
}

//暗器
func (p *Player) GetAnqiInfo() *anqitypes.AnqiInfo {
	m := p.getPlayerAnqiManager()
	anqiInfo := m.ToAnqiInfo()
	return anqiInfo
}

//神盾尖刺
func (p *Player) GetShieldInfo() *bodyshieldtypes.ShieldInfo {
	m := p.getPlayerBodyshieldManager()
	shieldInfo := m.ToShieldInfo()
	return shieldInfo
}

//战翼
func (p *Player) GetWingInfo() *wingcommon.WingInfo {
	m := p.getPlayerWingManager()
	wingInfo := m.ToWingInfo()
	return wingInfo
}

//护体仙羽
func (p *Player) GetFeatherInfo() *wingcommon.FeatherInfo {
	m := p.getPlayerWingManager()
	featherInfo := m.ToFeatherInfo()
	return featherInfo
}

//冰魂
func (p *Player) GetAllWeaponInfo() *weapontypes.AllWeaponInfo {
	m := p.getPlayerWeaponManager()
	weaponInfo := m.ToAllWeaponInfo()
	return weaponInfo
}

//装备
func (p *Player) GetEquipmentSlotList() []*inventorytypes.EquipmentSlotInfo {
	m := p.getPlayerInventoryManager()
	return m.ToEquipmentSlotList()

}

//获取技能列表
func (p *Player) GetSkillList() []*skillcommon.SkillObjectImpl {
	skillList := p.SkillManager.ToAllSkillInfo()
	return skillList

}

//元神金装
func (p *Player) GetGoldEquipSlotList() []*goldequiptypes.GoldEquipSlotInfo {
	m := p.getPlayerGoldEquiManager()
	return m.ToGoldEquipSlotList()

}

//无双神器
func (p *Player) GetWushuangListInfo() []*wushuangweapontypes.WushuangInfo {
	m := p.getPlayerWushuangManager()
	return m.ToWushuangListInfo()
}

//仙尊特权卡
func (p *Player) GetXianZunCard() []*xianzuncardcommon.XianZunCardInfo {
	m := p.getPlayerXianZunCardManager()
	return m.GetXianZunCardTypeList()
}

//特戒
func (p *Player) GetRingInfo() []*ringcommon.RingInfo {
	m := p.getPlayerRingManager()
	return m.GetRingInfoList()
}

//古魂
func (p *Player) GetAllSoulInfo() *soultypes.AllSoulInfo {
	m := p.getPlayerSoulManager()
	return m.ToAllSoulInfo()
}

//结婚
func (p *Player) GetMarryInfo() *marrytypes.MarryInfo {
	m := p.getPlayerMarryManager()
	marryInfo := m.ToMarryInfo()
	return marryInfo
}

//身法
func (p *Player) GetShenfaInfo() *shenfatypes.ShenfaInfo {
	m := p.getPlayerShenfaManager()
	wingInfo := m.ToShenfaInfo()
	return wingInfo
}

//领域
func (p *Player) GetLingyuInfo() *lingyutypes.LingyuInfo {
	m := p.getPlayerLingyuManager()
	lingyuInfo := m.ToLingyuInfo()
	return lingyuInfo
}

//vip
func (p *Player) GetVipInfo() *viptypes.VipInfo {
	m := p.getPlayerVipManager()
	vipInfo := m.ToVipInfo()
	return vipInfo
}

//戮仙刃
func (p *Player) GetMassacreInfo() *massacretypes.MassacreInfo {
	m := p.getPlayerMassacreManager()
	massacreInfo := m.ToMassacreInfo()
	return massacreInfo
}

//法宝
func (p *Player) GetFaBaoInfo() *fabaocommon.FaBaoInfo {
	m := p.getPlayerFaBaoManager()
	faBaoInfo := m.ToFaBaoInfo()
	return faBaoInfo
}

//血盾
func (p *Player) GetXueDunInfo() *xueduncommon.XueDunInfo {
	m := p.getPlayerXueDunManager()
	xueDunInfo := m.ToXueDunInfo()
	return xueDunInfo
}

//仙体
func (p *Player) GetXianTiInfo() *xianticommon.XianTiInfo {
	m := p.getPlayerXianTiManager()
	xianTiInfo := m.ToXianTiInfo()
	return xianTiInfo
}

//八卦
func (p *Player) GetBaGuaInfo() *baguacommon.BaGuaInfo {
	m := p.getPlayerBaGuaManager()
	baGuaInfo := m.ToBaGuaInfo()
	return baGuaInfo
}

//点星
func (p *Player) GetDianXingInfo() *dianxingcommon.DianXingInfo {
	m := p.getPlayerDianXingManager()
	dianXingInfo := m.ToDianXingInfo()
	return dianXingInfo
}

//天魔体
func (p *Player) GetTianMoTiInfo() *tianmotypes.TianMoInfo {
	m := p.getPlayerTianMoManager()
	tianMoInfo := m.ToTianMoInfo()
	return tianMoInfo
}

//噬魂幡
func (p *Player) GetShiHunFanInfo() *shihunfancommon.ShiHunFanInfo {
	m := p.getPlayerShiHunFanManager()
	shiHunFanInfo := m.ToShiHunFanInfo()
	return shiHunFanInfo
}

//灵童信息养成类
func (p *Player) GetAllLingTongDevInfo() *lingtongdevcommon.AllLingTongDevInfo {
	m := p.getPlayerLingTongDevManager()
	allLingTongDevInfo := m.ToAllLingTongDevInfo()
	return allLingTongDevInfo
}

//灵童信息
func (p *Player) GetLingTongInfo() *lingtongcommon.LingTongInfo {
	m := p.getPlayerLingTongManager()
	lingTongInfo := m.ToLingTongInfo()
	return lingTongInfo
}

//系统技能信息
func (p *Player) GetAllSystemSkillInfo() *sysskillcommon.AllSystemSkillInfo {
	m := p.getPlayerSystemSkillManager()
	allSystemSkillInfo := m.ToAllSystemSkillInfo()
	return allSystemSkillInfo
}

//附加系统类信息
func (p *Player) GetAllAdditionSysInfo() *additionsyscommon.AllAdditionSysInfo {
	m := p.getPlayerAdditionSysManager()
	allAdditionSysInfo := m.ToAllAdditionSysInfo()
	return allAdditionSysInfo
}

//怀孕信息
func (p *Player) GetPregnantInfo() *babytypes.PregnantInfo {
	m := p.getPlayerBabyMananger()
	pregnantInfo := m.ToPregnantInfo()
	return pregnantInfo
}

//基础属性列表
func (p *Player) GetBaseProperties() map[int32]int64 {
	return p.getPlayerPropertyManager().ToBaseProperties()
}

func (p *Player) GetBattleProperties() map[int32]int64 {
	return p.getPlayerPropertyManager().ToBattleProperties()
}

//本服没有竞技场
func (p *Player) SetArenaTeam(teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {

}

func (p *Player) GetCrossType() crosstypes.CrossType {
	m := p.getPlayerCrossManager()
	return m.GetCrossType()
}

func (p *Player) GetCrossArgs() []string {
	m := p.getPlayerCrossManager()
	return m.GetCrossArgs()
}

//功能开启
func (p *Player) IsFuncOpen(fot funcopentypes.FuncOpenType) bool {
	pdm, _ := p.playerDataManagerMap[types.PlayerFuncOpenDataManagerType]
	ppdm, _ := pdm.(*playerfuncopen.PlayerFuncOpenDataManager)
	return ppdm.IsOpen(fot)
}

//是否开场动画
func (p *Player) IsOpenVideo() bool {
	return p.po.IsOpenVideo != 0
}

//开场动画
func (p *Player) OpenVideo() {
	p.po.IsOpenVideo = 1
}

//设置权限
func (p *Player) SetPrivilege(typ types.PrivilegeType) {
	if !typ.Valid() {
		return
	}
	if typ == p.po.PrivilegeType {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	p.po.PrivilegeType = typ
	p.po.UpdateTime = now
	p.po.SetModified()
	//添加日志
}

// 性别变更
func (p *Player) ChangeSex() types.SexType {
	var newSex types.SexType
	if p.po.Sex == types.SexTypeMan {
		newSex = types.SexTypeWoman
	} else {
		newSex = types.SexTypeMan
	}

	now := global.GetGame().GetTimeService().Now()
	p.po.Sex = newSex
	p.po.UpdateTime = now
	p.po.SetModified()
	gameevent.Emit(playereventtypes.EventTypePlayerSexChanged, p, nil)
	return newSex
}

// 姓名变更
func (p *Player) ChangeName(newName string) {
	now := global.GetGame().GetTimeService().Now()
	p.po.Name = newName
	p.po.UpdateTime = now
	p.po.SetModified()
	gameevent.Emit(playereventtypes.EventTypePlayerNameChanged, p, newName)
}

//获取权限
func (p *Player) GetPrivilege() types.PrivilegeType {
	return p.po.PrivilegeType
}

//重载
func (p *Player) GetScene() scene.Scene {
	if p.PlayerSceneManager == nil {
		return nil
	}
	return p.PlayerSceneManager.GetScene()
}

const (
	queueCapacity = 10000
	maxTime       = time.Microsecond * 10
)

//仅Gm命令使用
func (p *Player) GmSetForbid(forbidText string) {
	now := global.GetGame().GetTimeService().Now()
	p.po.Forbid = 1
	p.po.ForbidText = forbidText
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) Forbid(forbidReason string, forbidName string, forbidTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.po.Forbid = 1
	p.po.ForbidText = forbidReason
	p.po.ForbidName = forbidName

	if forbidTime == 0 {
		p.po.ForbidEndTime = 0
	} else {
		p.po.ForbidEndTime = forbidTime + now
	}
	p.po.ForbidTime = now
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) IsForbid() bool {
	if p.po.Forbid == 0 {
		return false
	}
	if p.po.ForbidEndTime == 0 {
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	return p.po.ForbidEndTime > now
}

func (p *Player) Unforbid() {
	now := global.GetGame().GetTimeService().Now()
	p.po.Forbid = 0
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) ForbidChat(forbidChatReason string, forbidChatName string, forbidChatTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.po.ForbidChat = 1
	p.po.ForbidChatText = forbidChatReason
	p.po.ForbidChatName = forbidChatName
	p.po.ForbidChatTime = now
	if forbidChatTime == 0 {
		p.po.ForbidChatEndTime = 0
	} else {
		p.po.ForbidChatEndTime = forbidChatTime + now
	}
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) UnforbidChat() {
	now := global.GetGame().GetTimeService().Now()
	p.po.ForbidChat = 0
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) IsForbidChat() bool {
	if p.po.ForbidChat == 0 {
		return false
	}
	if p.po.ForbidChatEndTime == 0 {
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	return p.po.ForbidChatEndTime > now
}

func (p *Player) IgnoreChat(forbidChatReason string, forbidChatName string, forbidChatTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.po.IgnoreChat = 1
	p.po.IgnoreChatText = forbidChatReason
	p.po.IgnoreChatName = forbidChatName
	p.po.IgnoreChatTime = now
	if forbidChatTime == 0 {
		p.po.IgnoreChatEndTime = 0
	} else {
		p.po.IgnoreChatEndTime = forbidChatTime + now
	}

	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) UnignoreChat() {
	now := global.GetGame().GetTimeService().Now()
	p.po.IgnoreChat = 0
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) IsIgnoreChat() bool {
	if p.po.IgnoreChat == 0 {
		return false
	}
	if p.po.IgnoreChatEndTime == 0 {
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	return p.po.IgnoreChatEndTime > now
}

func (p *Player) AddChargeInfo(goldNum, money int64) {
	if goldNum <= 0 || money <= 0 {
		panic(fmt.Errorf("goldNum 或 money 不能小于1"))
	}

	p.refreshTodayCharge()

	now := global.GetGame().GetTimeService().Now()
	p.po.TotalChargeGold += goldNum
	p.po.TotalChargeMoney += money
	p.po.TodayChargeMoney += money
	p.po.ChargeTime = now
	p.po.UpdateTime = now
	p.po.SetModified()
	return
}

func (p *Player) GetChargeGoldNum() int64 {
	return p.po.TotalChargeGold + p.po.TotalPrivilegeChargeGold
}

func (p *Player) AddPrivilegeChargeInfo(goldNum int64) {
	if goldNum <= 0 {
		panic(fmt.Errorf("goldNum 不能小于1"))
	}

	now := global.GetGame().GetTimeService().Now()
	p.po.TotalPrivilegeChargeGold += goldNum
	p.po.UpdateTime = now
	p.po.SetModified()
	return
}

func (p *Player) IsSystemCompensate() bool {
	return p.po.SystemCompensate != 0
}

func (p *Player) SendSystemCompensate() {
	now := global.GetGame().GetTimeService().Now()
	p.po.SystemCompensate = 1
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) GMSetSystemCompensate(status bool) {
	if status {
		p.po.SystemCompensate = 1
	} else {
		p.po.SystemCompensate = 0
	}
	now := global.GetGame().GetTimeService().Now()
	p.po.UpdateTime = now
	p.po.SetModified()
}

func (p *Player) IsGetNewReward() bool {
	return p.po.GetNewReward != 0
}

func (p *Player) GetNewReward() {
	now := global.GetGame().GetTimeService().Now()
	p.po.GetNewReward = 1
	p.po.UpdateTime = now
	p.po.SetModified()
	return
}

func (p *Player) IsGuaJiPlayer() bool {
	return p.guaJi
}

func (p *Player) IsGm() bool {
	return p.gm
}

const (
	crossTimer = time.Second * 30
	debugState = false
)

var (
	currentLogoutState = player.PlayerStateInit
)

//创建玩家
func NewPlayer(s gamesession.Session, sdkType logintypes.SDKType, devicePlatformType logintypes.DevicePlatformType, platformUserId string, serverId int32, userId int64, realNameState accounttypes.RealNameState, guaJi bool, gm bool) player.Player {
	p := &Player{
		s:                  s,
		sdkType:            sdkType,
		devicePlatformType: devicePlatformType,
		platformUserId:     platformUserId,
		userId:             userId,
		serverId:           serverId,
		realNameState:      realNameState,
		guaJi:              guaJi,
		gm:                 gm,
	}
	p.SubjectBase = fsm.NewSubjectBase(player.PlayerStateInit)
	p.msgQueue = NewMessageQueue(p, queueCapacity, maxTime)
	p.done = make(chan struct{})
	p.crossTimer = time.NewTimer(crossTimer)
	p.crossTimer.Stop()
	p.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	if debugState {
		numState := int32(player.PlayerStateLogouted) - int32(player.PlayerStateInit)
		p.randomLogoutState = fsm.State(rand.Int31n(numState) + int32(player.PlayerStateInit))
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"state":    p.randomLogoutState,
			}).Info("player:初始化随机退出状态")
		currentLogoutState = fsm.State(currentLogoutState + 1)
		if currentLogoutState > player.PlayerStateLogouted {
			currentLogoutState = player.PlayerStateEnterCross
		}
	}
	return p
}
