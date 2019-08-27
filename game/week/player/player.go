package player

import (
	commonlog "fgame/fgame/common/log"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/week/dao"
	weekeventtypes "fgame/fgame/game/week/event/types"
	weektemplate "fgame/fgame/game/week/template"
	weektypes "fgame/fgame/game/week/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//玩家周卡管理器
type PlayerWeekManager struct {
	p          player.Player
	weekObject *PlayerWeekObject //周卡数据
}

func (m *PlayerWeekManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerWeekManager) Load() (err error) {

	//加载周卡数据
	weekEntity, err := dao.GetWeekDao().GetWeekEntity(m.p.GetId())
	if err != nil {
		return
	}

	if weekEntity != nil {
		obj := newPlayerWeekObject(m.p)
		obj.FromEntity(weekEntity)
		m.weekObject = obj
	} else {
		m.initWeekObj()
	}

	return nil
}

func (m *PlayerWeekManager) initWeekObj() {
	obj := newPlayerWeekObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.createTime = now
	for initType := weektypes.MinType; initType <= weektypes.MaxType; initType++ {
		obj.weekDataMap = make(map[weektypes.WeekType]*WeekData)
		obj.weekDataMap[initType] = initWeekData(initType)
	}
	m.weekObject = obj
	obj.SetModified()
}

//加载后
func (m *PlayerWeekManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerWeekManager) Heartbeat() {
}

//购买周卡
func (m *PlayerWeekManager) BuyWeek(weekType weektypes.WeekType) (flag bool) {
	weekData := m.GetWeekInfo(weekType)
	if weekData == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if weekData.IsWeek(now) {
		return
	}

	lastExpireTime := weekData.GetExpireTime()
	weekTemp := weektemplate.GetWeekTemplateService().GetWeekTemplate(weekType)
	expireTime := now + weekTemp.Duration
	weekData.setExpireTime(expireTime)
	m.weekObject.updateTime = now
	m.weekObject.SetModified()

	gameevent.Emit(weekeventtypes.EventTypeWeekBuy, m.p, weekType)

	weekReason := commonlog.WeekLogReasonBuy
	reasonText := fmt.Sprintf(weekReason.String(), weekType.String())
	logEventData := weekeventtypes.CreatePlayerWeekLogBuyLogEventData(lastExpireTime, int32(weekType), weekReason, reasonText)
	gameevent.Emit(weekeventtypes.EventTypeWeekLogBuy, m.p, logEventData)
	flag = true
	return
}

//领取奖励
func (m *PlayerWeekManager) ReceiveWeekRewards(weekType weektypes.WeekType) (flag bool) {
	weekData := m.GetWeekInfo(weekType)
	if weekData == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if !weekData.IsWeek(now) {
		return
	}

	if weekData.IsReceiveRewards(now) {
		return
	}

	weekData.cycleDayIncr(now)
	m.weekObject.updateTime = now
	m.weekObject.SetModified()
	flag = true
	return
}

func (m *PlayerWeekManager) GetWeekInfo(weekType weektypes.WeekType) *WeekData {
	return m.weekObject.weekDataMap[weekType]
}

func (m *PlayerWeekManager) GetWeekInfoMap() map[weektypes.WeekType]*WeekData {
	return m.weekObject.weekDataMap
}

func createPlayerWeekDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerWeekManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerWeekDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerWeekDataManager))
}
