package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/jieyi/dao"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
)

type PlayerJieYiDataManager struct {
	pl                player.Player
	playerJieYiObject *PlayerJieYiObject
}

// 加载
func (m *PlayerJieYiDataManager) Load() (err error) {
	jieYiEntity, err := dao.GetJieYiDao().GetPlayerJieYiEntity(m.pl.GetId())
	if err != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj := NewJieYiObject(m.pl)
	if jieYiEntity == nil {
		id, _ := idutil.GetId()
		obj.id = id
		obj.tokenType = jieyitypes.JieYiTokenTypeInvalid
		obj.createTime = now
		obj.SetModified()
	} else {
		obj.FromEntity(jieYiEntity)
	}

	m.playerJieYiObject = obj
	return nil
}

func (m *PlayerJieYiDataManager) Player() player.Player {
	return m.pl
}

//加载后
func (m *PlayerJieYiDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerJieYiDataManager) Heartbeat() {
}

//同步结义数据
func (m *PlayerJieYiDataManager) SyncJieYi(daoJu jieyitypes.JieYiDaoJuType, token jieyitypes.JieYiTokenType, jieYiId int64, name string, rank int32) {
	now := global.GetGame().GetTimeService().Now()

	//结义成功
	if m.playerJieYiObject.jieYiId == 0 && jieYiId != 0 {
		//结义成功
		event.Emit(jieyieventtypes.JieYiEventTypeJieYiSuccess, m.pl, nil)
	}

	//扣除声威,更新上次离开结义时间
	if m.playerJieYiObject.jieYiId != 0 && jieYiId == 0 {
		m.playerJieYiObject.rank = 0
		m.playerJieYiObject.lastLeaveTime = now
	}

	if m.playerJieYiObject.jieYiId != jieYiId {
		m.playerJieYiObject.jieYiId = jieYiId
	}

	if m.playerJieYiObject.jieYiDaoJu < daoJu {
		beformDaoJu := m.playerJieYiObject.jieYiDaoJu
		m.playerJieYiObject.jieYiDaoJu = daoJu
		//发时装
		event.Emit(jieyieventtypes.JieYiEventTypeDaoJuTypeChange, m.pl, daoJu)
		// 道具改变日志
		reason := commonlog.JieYiLogReasonDaoJuTypeChange
		reasonText := fmt.Sprintf(reason.String(), beformDaoJu.String(), daoJu.String())
		logData := jieyieventtypes.CreatePlayerJieYiDaoJuChangeLogEventData(beformDaoJu, daoJu, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeDaoJuTypeChangeLog, m.pl, logData)
	}
	if m.playerJieYiObject.tokenType < token {
		beformToken := m.playerJieYiObject.tokenType
		m.playerJieYiObject.tokenType = token
		// 信物改变日志
		reason := commonlog.JieYiLogReasonTokenTypeChange
		reasonText := fmt.Sprintf(reason.String(), beformToken.String(), token.String())
		logData := jieyieventtypes.CreatePlayerJieYiTokenTypeChangeLogEventData(beformToken, token, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeTokenTypeChangeLog, m.pl, logData)
	}

	if m.playerJieYiObject.name != name {
		m.playerJieYiObject.name = name
	}

	if m.playerJieYiObject.rank != rank {
		m.playerJieYiObject.rank = rank
	}
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
	event.Emit(jieyieventtypes.JieYiEventTypeJieYiChange, m.pl, nil)
}

// 玩家解除结义成功
func (m *PlayerJieYiDataManager) JieChuSuccess() {
	m.SyncJieYi(m.playerJieYiObject.GetDaoJuType(), m.playerJieYiObject.GetTokenType(), 0, "", 0)
}

// 邀请成功
func (m *PlayerJieYiDataManager) InviteSucess() {
	now := global.GetGame().GetTimeService().Now()
	m.playerJieYiObject.lastInviteTime = now
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
}

// 玩家是否结义
func (m *PlayerJieYiDataManager) IsJieYi() bool {
	return m.playerJieYiObject.jieYiId != 0
}

// 玩家是否结义
func (m *PlayerJieYiDataManager) GetJieYiId() int64 {
	return m.playerJieYiObject.jieYiId
}

func (m *PlayerJieYiDataManager) GetJieYiName() string {
	return m.playerJieYiObject.name
}

func (m *PlayerJieYiDataManager) GetJieYiRank() int32 {
	return m.playerJieYiObject.rank
}

// 获取玩家结义对象
func (m *PlayerJieYiDataManager) GetPlayerJieYiObj() *PlayerJieYiObject {
	return m.playerJieYiObject
}

// 玩家是否激活信物
func (m *PlayerJieYiDataManager) IsTokenActivite() bool {
	return m.playerJieYiObject.tokenType != jieyitypes.JieYiTokenTypeInvalid
}

// 玩家信物升级未必成功
func (m *PlayerJieYiDataManager) TokenUpLevel(pro int32, sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		beformLevel := m.playerJieYiObject.tokenLev
		m.playerJieYiObject.tokenLev++
		m.playerJieYiObject.tokenNum = 0
		m.playerJieYiObject.tokenPro = 0

		event.Emit(jieyieventtypes.JieYiEventTypeTokenLevelChange, m.pl, m.playerJieYiObject.tokenLev)
		// 信物等级改变日志
		curLevel := m.playerJieYiObject.tokenLev
		reason := commonlog.JieYiLogReasonTokenLevelChange
		reasonText := fmt.Sprintf(reason.String(), beformLevel, curLevel)
		logData := jieyieventtypes.CreatePlayerJieYiTokenLevelChangeLogEventData(beformLevel, curLevel, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeTokenLevelChangeLog, m.pl, logData)
	} else {
		m.playerJieYiObject.tokenNum++
		m.playerJieYiObject.tokenPro += pro
	}
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
}

// 玩家是否能再次结义
func (m *PlayerJieYiDataManager) IsCanJieYi() bool {
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	now := global.GetGame().GetTimeService().Now()
	if now-m.playerJieYiObject.lastLeaveTime > constantTemp.JieChuCD {
		return true
	}
	return false
}

// 玩家是否能再次邀请
func (m *PlayerJieYiDataManager) IsCanInvite() bool {
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	now := global.GetGame().GetTimeService().Now()
	if now-m.playerJieYiObject.lastInviteTime > constantTemp.YaoQingCD {
		return true
	}
	return false
}

// 玩家发布结义成功
func (m *PlayerJieYiDataManager) LeaveWordSuccess() {
	now := global.GetGame().GetTimeService().Now()
	m.playerJieYiObject.lastPostTime = now
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
}

// 声威值能否掉落
func (m *PlayerJieYiDataManager) IsCanDropShengWeiZhi() bool {
	now := global.GetGame().GetTimeService().Now()
	temp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	dropCD := temp.DropCD
	if (now - m.playerJieYiObject.lastDropTime) > dropCD {
		return true
	}
	return false
}

// 增加声威值
func (m *PlayerJieYiDataManager) AddShengWeiZhi(num int32) {
	// 未结义不增加声威值
	if !m.IsJieYi() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	beformNum := m.playerJieYiObject.shengWeiZhi
	curNum := m.playerJieYiObject.shengWeiZhi + num
	if curNum > int32(constantTemp.ShengWeiMax) {
		curNum = int32(constantTemp.ShengWeiMax)
	}
	m.playerJieYiObject.shengWeiZhi = curNum
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()

	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, m.pl, nil)
	// 声威值改变日志
	reason := commonlog.JieYiLogReasonGM
	reasonText := reason.String()
	logData := jieyieventtypes.CreatePlayerJieYiShengWeiZhiChangeLogEventData(beformNum, curNum, reason, reasonText)
	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChangeLog, m.pl, logData)
}

func (m *PlayerJieYiDataManager) ShengWeiDrop() (flag bool, dropNum int32, dropLev int32) {
	now := global.GetGame().GetTimeService().Now()
	if !m.IsCanDropShengWeiZhi() {
		return
	}

	shengWei := m.playerJieYiObject.shengWeiZhi
	beformNameLev := m.playerJieYiObject.nameLev
	beformNum := m.playerJieYiObject.shengWeiZhi
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	startNameTenp := jieyitemplate.GetJieYiTemplateService().GetJieYiNameTemplate(beformNameLev)
	// 声威值掉落
	if mathutils.RandomHit(common.MAX_RATE, int(constantTemp.DropRate)) && shengWei != 0 {
		randDropPercent := mathutils.RandomRange(int(constantTemp.DropPercentMin), int(constantTemp.DropPercentMax))
		dropPercent := float64(randDropPercent) / float64(common.MAX_RATE)
		dropShengWei := math.Ceil(float64(shengWei) * dropPercent)
		m.playerJieYiObject.shengWeiZhi -= int32(dropShengWei)
		m.playerJieYiObject.lastDropTime = now
		m.playerJieYiObject.updateTime = now
		sysReturn := float64(constantTemp.DropSystemReturn) / float64(common.MAX_RATE)
		dropNum += int32(math.Ceil(float64(dropShengWei) * (1 - sysReturn)))

		// 声威值改变日志
		curNum := m.playerJieYiObject.shengWeiZhi
		reason := commonlog.JieYiLogReasonShengWeiZhiChange
		reasonText := fmt.Sprintf(reason.String(), beformNum, curNum)
		logData := jieyieventtypes.CreatePlayerJieYiShengWeiZhiChangeLogEventData(beformNum, curNum, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChangeLog, m.pl, logData)
	}

	// 威名等级掉落
	if mathutils.RandomHit(common.MAX_RATE, int(startNameTenp.DropPercent)) && beformNameLev != 0 {
		// 随机掉落等级
		minDropLev := startNameTenp.DeathMinLevel
		maxDropLev := startNameTenp.DeathMaxLevel
		randDropLev := int32(mathutils.RandomRange(int(minDropLev), int(maxDropLev)))
		dropLev = randDropLev
		for i := randDropLev; i > 0; i-- {
			nameTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiNameTemplate(m.playerJieYiObject.nameLev)
			dropNum += int32(nameTemp.StarCount)
			m.playerJieYiObject.nameLev--
		}
		m.playerJieYiObject.lastDropTime = now
		m.playerJieYiObject.updateTime = now

		// 威名等级改变日志
		curLevel := m.playerJieYiObject.nameLev
		reason := commonlog.JieYiLogReasonNameLevChange
		reasonText := fmt.Sprintf(reason.String(), beformNameLev, curLevel)
		logData := jieyieventtypes.CreatePlayerJieYiNameLevelChangeLogEventData(beformNameLev, curLevel, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeNameLevelChangeLog, m.pl, logData)
	}
	m.playerJieYiObject.SetModified()
	flag = true
	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, m.pl, nil)
	return
}

// 玩家威名升级未必成功
func (m *PlayerJieYiDataManager) NameUpLevel(success bool, shengWeiZhi int32) {
	now := global.GetGame().GetTimeService().Now()
	if success {
		beformLevel := m.playerJieYiObject.nameLev
		m.playerJieYiObject.nameLev++
		m.playerJieYiObject.namePro = 0
		m.playerJieYiObject.nameNum = 0

		event.Emit(jieyieventtypes.JieYiEventTypeNameUpLev, m.pl, m.playerJieYiObject.nameLev)
		// 威名等级改变日志
		curLevel := m.playerJieYiObject.nameLev
		reason := commonlog.JieYiLogReasonNameLevChange
		reasonText := fmt.Sprintf(reason.String(), beformLevel, curLevel)
		logData := jieyieventtypes.CreatePlayerJieYiNameLevelChangeLogEventData(beformLevel, curLevel, reason, reasonText)
		event.Emit(jieyieventtypes.JieYiEventTypeNameLevelChangeLog, m.pl, logData)
	} else {
		m.playerJieYiObject.nameNum++
	}
	beformNum := m.playerJieYiObject.shengWeiZhi
	m.playerJieYiObject.shengWeiZhi -= shengWeiZhi
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()

	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, m.pl, nil)
	// 声威值改变日志
	curNum := m.playerJieYiObject.shengWeiZhi
	reason := commonlog.JieYiLogReasonShengWeiZhiChange
	reasonText := fmt.Sprintf(reason.String(), beformNum, curNum)
	logData := jieyieventtypes.CreatePlayerJieYiShengWeiZhiChangeLogEventData(beformNum, curNum, reason, reasonText)
	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChangeLog, m.pl, logData)
}

func (m *PlayerJieYiDataManager) GetNameLevel() int32 {
	return m.playerJieYiObject.nameLev
}

func (m *PlayerJieYiDataManager) GetShengWeiZhi() int32 {
	return m.playerJieYiObject.shengWeiZhi
}

// gm清理邀请时间
func (m *PlayerJieYiDataManager) GmClearLastInvite() {
	now := global.GetGame().GetTimeService().Now()
	m.playerJieYiObject.lastInviteTime = 0
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
}

// gm清理结义所有时间
func (m *PlayerJieYiDataManager) GmClearAllTime() {
	now := global.GetGame().GetTimeService().Now()
	m.playerJieYiObject.lastPostTime = 0
	m.playerJieYiObject.lastDropTime = 0
	m.playerJieYiObject.lastLeaveTime = 0
	m.playerJieYiObject.lastQiuYuanTime = 0
	m.playerJieYiObject.lastInviteTime = 0
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
}

// gm设置声威值
func (m *PlayerJieYiDataManager) GmSetShengWeiZhi(num int32) {
	now := global.GetGame().GetTimeService().Now()
	beformNum := m.playerJieYiObject.shengWeiZhi
	m.playerJieYiObject.shengWeiZhi = num
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()

	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, m.pl, nil)
	// 声威值改变日志
	curNum := m.playerJieYiObject.shengWeiZhi
	reason := commonlog.JieYiLogReasonGM
	reasonText := reason.String()
	logData := jieyieventtypes.CreatePlayerJieYiShengWeiZhiChangeLogEventData(beformNum, curNum, reason, reasonText)
	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChangeLog, m.pl, logData)
}

// gm设置威名等级
func (m *PlayerJieYiDataManager) GmSetNameLev(level int32) {
	now := global.GetGame().GetTimeService().Now()
	beformLevel := m.playerJieYiObject.nameLev
	m.playerJieYiObject.nameLev = level
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
	event.Emit(jieyieventtypes.JieYiEventTypeShengWeiZhiChange, m.pl, nil)
	event.Emit(jieyieventtypes.JieYiEventTypeNameUpLev, m.pl, m.playerJieYiObject.nameLev)

	// 威名等级改变日志
	curLevel := m.playerJieYiObject.nameLev
	reason := commonlog.JieYiLogReasonGM
	reasonText := reason.String()
	logData := jieyieventtypes.CreatePlayerJieYiNameLevelChangeLogEventData(beformLevel, curLevel, reason, reasonText)
	event.Emit(jieyieventtypes.JieYiEventTypeNameLevelChangeLog, m.pl, logData)
}

// gm设置信物等级
func (m *PlayerJieYiDataManager) GmSetTokenLev(level int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerJieYiObject.tokenLev = level
	m.playerJieYiObject.updateTime = now
	m.playerJieYiObject.SetModified()
	event.Emit(jieyieventtypes.JieYiEventTypeTokenLevelChange, m.pl, level)
}

func CreatePlayerJieYiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerJieYiDataManager{}
	m.pl = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerJieYiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerJieYiDataManager))
}
