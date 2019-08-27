package player

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/xianfu/dao"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fgame/fgame/pkg/idutil"
	"math"
)

//玩家秘境仙府数据管理器
type PlayerXinafuDataManager struct {
	pl player.Player
	//仙府对象集
	xfObjectMap map[xianfutypes.XianfuType]*PlayerXianfuObject
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXianfuDtatManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXinafuDataManager))
}

func CreatePlayerXinafuDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerXinafuDataManager{}
	m.pl = p
	m.xfObjectMap = make(map[xianfutypes.XianfuType]*PlayerXianfuObject)

	return m
}

func (m *PlayerXinafuDataManager) Player() player.Player {
	return m.pl
}

//加载秘境仙府信息
func (m *PlayerXinafuDataManager) Load() error {
	xfEntityArr, err := dao.GetXianfuDao().GetXianfuInfo(m.pl.GetId())
	if err != nil {
		return err
	}
	for _, xfentity := range xfEntityArr {
		newObj := CreateNewPlayerXianfuObject(m.pl)
		newObj.FromEntity(xfentity)
		m.xfObjectMap[newObj.xianfuType] = newObj
	}

	for typ := xianfutypes.MinXianfuType; typ <= xianfutypes.MaxXianfuType; typ++ {
		if _, ok := m.xfObjectMap[typ]; !ok {
			m.initNewXianfuObject(typ)
		}
	}

	return nil
}

//加载后处理
func (m *PlayerXinafuDataManager) AfterLoad() (err error) {
	now := global.GetGame().GetTimeService().Now()
	err = m.RefreshData(now)
	if err != nil {
		return
	}

	return
}

//心跳
func (m *PlayerXinafuDataManager) Heartbeat() {
}

//刷新秘境仙府数据
func (m *PlayerXinafuDataManager) RefreshData(nowTime int64) (err error) {
	err = m.refreshUseTimes(nowTime)
	if err != nil {
		return
	}

	m.refreshUpgradeState(nowTime)
	return
}

//刷新xianfuObj升级状态
func (m *PlayerXinafuDataManager) refreshUpgradeState(nowTime int64) {
	for _, obj := range m.xfObjectMap {
		if obj.isUpgradeDone() {
			obj.upgradeDone(nowTime)
		}
	}
}

//重置仙府挑战次数
func (m *PlayerXinafuDataManager) refreshUseTimes(nowTime int64) (err error) {
	for _, obj := range m.xfObjectMap {
		err = obj.refreshUseTimes(nowTime)
		if err != nil {
			return
		}
	}
	return
}

//生成玩家初始秘境仙府数据
func (m *PlayerXinafuDataManager) initNewXianfuObject(xfType xianfutypes.XianfuType) {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	newObj := CreateNewPlayerXianfuObject(m.pl)
	newObj.id = id
	if xfTemplate := xianfutemplate.GetXianfuTemplateService().GetFirstXianfu(xfType); xfTemplate != nil {
		newObj.xianfuId = int32(xfTemplate.TemplateId())
	}
	newObj.xianfuType = xfType
	newObj.state = xianfutypes.XianfuStateWaitedToUpgrade
	newObj.createTime = now

	m.xfObjectMap[xfType] = newObj
	newObj.SetModified()
}

//获取所有秘境仙府信息
func (m *PlayerXinafuDataManager) GetPlayerXianfuInfoList() map[xianfutypes.XianfuType]*PlayerXianfuObject {
	return m.xfObjectMap
}

//升级秘境仙府
func (m *PlayerXinafuDataManager) UpgradePlayerXianfu(xfType xianfutypes.XianfuType, nowTime int64) (err error) {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return
	}
	obj.state = xianfutypes.XianfuStateUpgrading
	obj.startTime = nowTime
	obj.updateTime = nowTime
	obj.SetModified()
	gameevent.Emit(xianfueventtypes.EventTypeXianFuStartUpgrade, m.pl, xfType)
	return
}

//加速升级秘境仙府
func (m *PlayerXinafuDataManager) FinishAccelerateUpgrade(xfType xianfutypes.XianfuType, nowTime int64) {
	obj := m.GetPlayerXianfuInfo(xfType)
	obj.upgradeDone(nowTime)
}

//是否处于升级中状态
func (m *PlayerXinafuDataManager) IsUpgrading(xfType xianfutypes.XianfuType) bool {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return false
	}

	if obj.state != xianfutypes.XianfuStateUpgrading {
		return false
	}
	return true
}

//获取剩余次数
func (m *PlayerXinafuDataManager) GetChallengeTimes(xfType xianfutypes.XianfuType) int32 {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return 0
	}
	maxChallengeTimes := xianfutemplate.GetXianfuTemplateService().GetBasicPlayTimes(xfType)
	return maxChallengeTimes - obj.useTimes
}

//完成扫荡或挑战
func (m *PlayerXinafuDataManager) UseTimes(xfType xianfutypes.XianfuType, num int32, nowTime int64) (err error) {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return
	}
	obj.useTimes += num
	obj.updateTime = nowTime
	obj.SetModified()

	m.EmitJoinEvent(xfType, num)
	return
}

//发送挑战事件
func (m *PlayerXinafuDataManager) EmitJoinEvent(xfType xianfutypes.XianfuType, num int32) {
	eventData := xianfueventtypes.CreateXianFuChallengeEventData(xfType, num)
	gameevent.Emit(xianfueventtypes.EventTypeXianFuChallenge, m.pl, eventData)
}

//发送完成事件
func (m *PlayerXinafuDataManager) EmitFinishEvent(xfType xianfutypes.XianfuType, num int32) {
	eventData := xianfueventtypes.CreateXianFuFinishEventData(xfType, num)
	gameevent.Emit(xianfueventtypes.EventTypeXianFuFinish, m.pl, eventData)
}

//发送扫荡事件
func (m *PlayerXinafuDataManager) EmitSweepEvent(xfType xianfutypes.XianfuType, num int32) {
	eventData := xianfueventtypes.CreateXianFuChallengeEventData(xfType, num)
	gameevent.Emit(xianfueventtypes.EventTypeXianFuSweep, m.pl, eventData)
}

//获取加速所需元宝
func (m *PlayerXinafuDataManager) GetAccelerateNeedGold(xfType xianfutypes.XianfuType) int32 {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return 0
	}
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(obj.xianfuId, xfType)
	nextXfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xfTemplate.GetNextId(), xfType)

	now := global.GetGame().GetTimeService().Now()
	costTime := nextXfTemplate.GetUpgradeTime()
	diff := costTime - (now - obj.GetStartTime())
	acceleratedTimeOfHour := math.Ceil(float64(diff) / float64(common.HOUR))

	totalNeedGold := int32(math.Ceil(float64(nextXfTemplate.GetSpeedUpNeedGold()*acceleratedTimeOfHour) / float64(common.MAX_RATE)))
	return totalNeedGold
}

//刷新波数
func (m *PlayerXinafuDataManager) RefreshGroup(group int32) {
	obj := m.GetPlayerXianfuInfo(xianfutypes.XianfuTypeExp)
	if obj.group >= group {
		return
	}
	nowTime := global.GetGame().GetTimeService().Now()
	obj.group = group
	obj.updateTime = nowTime
	obj.SetModified()
	return
}

//获取波数
func (m *PlayerXinafuDataManager) GetGroup(xfType xianfutypes.XianfuType) int32 {
	obj := m.GetPlayerXianfuInfo(xfType)
	return obj.group
}

//获取秘境仙府
func (m *PlayerXinafuDataManager) GetPlayerXianfuInfo(xfType xianfutypes.XianfuType) (obj *PlayerXianfuObject) {
	return m.xfObjectMap[xfType]
}

//获取秘境仙府Id
func (m *PlayerXinafuDataManager) GetXianfuId(xfType xianfutypes.XianfuType) int32 {
	obj := m.GetPlayerXianfuInfo(xfType)
	if obj == nil {
		return 0
	}
	return obj.xianfuId
}

//修正波数记录
func (m *PlayerXinafuDataManager) FixXianExpGroupRecord() {
	obj := m.GetPlayerXianfuInfo(xianfutypes.XianfuTypeExp)
	if obj == nil {
		return
	}

	xianfuTemp := xianfutemplate.GetXianfuTemplateService().GetXianfu(obj.xianfuId, obj.xianfuType)
	if xianfuTemp == nil {
		return
	}

	maxGroup := xianfuTemp.GetMapTemplate().GetNumGroup()
	if obj.group <= maxGroup {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.group = maxGroup
	obj.updateTime = now
	obj.SetModified()
	return
}
