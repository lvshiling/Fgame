package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/vip/dao"
	vipeventtypes "fgame/fgame/game/vip/event/types"
	viptemplate "fgame/fgame/game/vip/template"
	viptypes "fgame/fgame/game/vip/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//玩家VIP管理器
type PlayerVipDataManager struct {
	p               player.Player
	playerVipObject *PlayerVipObject
}

func (m *PlayerVipDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerVipDataManager) Load() (err error) {
	//加载玩家VIP信息
	vipEntity, err := dao.GetVipDao().GetVipEntity(m.p.GetId())
	if err != nil {
		return
	}
	if vipEntity == nil {
		m.initPlayerVipObject()
	} else {
		m.playerVipObject = NewPlayerVipObject(m.p)
		m.playerVipObject.FromEntity(vipEntity)
	}
	return nil
}

//第一次初始化
func (m *PlayerVipDataManager) initPlayerVipObject() {
	o := NewPlayerVipObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.vipLevel = int32(0)
	o.vipStar = int32(0)
	o.consumeLevel = int32(0)
	o.chargeNum = int64(0)
	o.freeGiftMap = map[int32]int32{}
	o.discountMap = map[int32]int32{}
	o.createTime = now
	m.playerVipObject = o
	o.SetModified()
}

//加载后
func (m *PlayerVipDataManager) AfterLoad() (err error) {
	m.resetLiBao()
	return nil
}

// 重置vip礼包
func (m *PlayerVipDataManager) resetLiBao() {
	now := global.GetGame().GetTimeService().Now()
	resetTime := constant.GetConstantService().GetVipLiBaoResetTime()
	//还没到重置
	if now < resetTime {
		return
	}
	if m.playerVipObject.createTime < resetTime {
		m.playerVipObject.discountMap = map[int32]int32{}
		m.playerVipObject.freeGiftMap = map[int32]int32{}
		m.playerVipObject.createTime = now
		m.playerVipObject.updateTime = now
		m.playerVipObject.SetModified()
	}
}

//心跳
func (m *PlayerVipDataManager) Heartbeat() {}

func (m *PlayerVipDataManager) checkVipLevel(curCharge int64) (isUplevel bool) {
	curLevel := m.playerVipObject.vipLevel
	curStar := m.playerVipObject.vipStar
	curTemp := viptemplate.GetVipTemplateService().GetVipTemplate(curLevel, curStar)
	if curTemp == nil {
		log.Warnf("vip模板不存在，等级:%d,星级:%d", curLevel, curStar)
		return
	}
	nextTemp := curTemp.GetNextTemplate()
	if nextTemp == nil {
		return
	}

	for {
		maxCost := int64(nextTemp.NeedValue)
		if curCharge < maxCost {
			break
		}

		curTemp = nextTemp
		nextTemp = nextTemp.GetNextTemplate()
		isUplevel = true
		if nextTemp == nil {
			break
		}
	}

	m.playerVipObject.vipLevel = curTemp.Level
	m.playerVipObject.vipStar = curTemp.Star

	return
}

// 更新累计充值记录
func (m *PlayerVipDataManager) UpdateChargeNum(num int64) {
	if num < 1 {
		panic(fmt.Errorf("num 不能小于1, num:%d", num))
	}

	curCharge := m.playerVipObject.chargeNum
	beforeGold := curCharge
	beforeLevel := m.playerVipObject.vipLevel
	curCharge += num

	isUplevel := m.checkVipLevel(curCharge)
	m.checkCostLevel(curCharge)
	m.playerVipObject.chargeNum = curCharge
	m.playerVipObject.SetModified()

	if isUplevel {
		m.p.SyncVip(m.playerVipObject.vipLevel)
		gameevent.Emit(vipeventtypes.EventTypeVipLevelChanged, m.p, nil)

		reason := commonlog.VipLogReasonUplevel
		data := vipeventtypes.CreatePlayerVipAdvancedLogEventData(beforeLevel, beforeGold, num, reason, reason.String())
		gameevent.Emit(vipeventtypes.EventTypeVipLevelChangedLog, m.p, data)
	}
}

//GM设置充值数量
func (m *PlayerVipDataManager) GMSetChargeNum(num int64) {
	curCharge := num
	m.playerVipObject.vipStar = 0
	m.playerVipObject.vipLevel = 0
	m.playerVipObject.consumeLevel = 0
	m.checkVipLevel(curCharge)
	m.p.SyncVip(m.playerVipObject.vipLevel)
	gameevent.Emit(vipeventtypes.EventTypeVipLevelChanged, m.p, nil)
	m.checkCostLevel(curCharge)

	m.playerVipObject.chargeNum = curCharge
	m.playerVipObject.SetModified()
}

//付费等级
func (m *PlayerVipDataManager) checkCostLevel(curCharge int64) {
	curLevel := m.playerVipObject.consumeLevel
	curTemp := viptemplate.GetVipTemplateService().GetCostLevelTemplate(curLevel)
	if curTemp == nil {
		log.Warnf("消费等级模板不存在，等级:%d", curLevel)
		return
	}

	nextTemp := curTemp.GetNextTemplate()
	if nextTemp == nil {
		return
	}

	for {
		maxCharge := int64(nextTemp.NeedValue)
		if curCharge < maxCharge {
			break
		}

		curTemp = nextTemp
		nextTemp = nextTemp.GetNextTemplate()
		if nextTemp == nil {
			break
		}
	}

	m.playerVipObject.consumeLevel = curTemp.Level
	return
}

// 获取vip等级
func (m *PlayerVipDataManager) GetVipLevel() (level int32, star int32) {
	return m.playerVipObject.vipLevel, m.playerVipObject.vipStar
}

// 获取消费等级
func (m *PlayerVipDataManager) GetCostLevel() int32 {
	return m.playerVipObject.consumeLevel
}

// 购买VIP礼包
func (m *PlayerVipDataManager) BuyVipGift(vipLevel int32) {
	_, ok := m.playerVipObject.discountMap[vipLevel]
	if ok {
		return
	}

	buyTimes := int32(1)
	m.playerVipObject.discountMap[vipLevel] = buyTimes
	m.playerVipObject.SetModified()
}

func (m *PlayerVipDataManager) IsCanBuyGift(level int32) bool {
	_, ok := m.playerVipObject.discountMap[level]
	if !ok {
		return true
	}

	return false
}

// 领取免费礼包
func (m *PlayerVipDataManager) ReceiveFreeGift(vipLevel int32) {
	_, ok := m.playerVipObject.freeGiftMap[vipLevel]
	if ok {
		return
	}

	isReceive := int32(1)
	m.playerVipObject.freeGiftMap[vipLevel] = isReceive
	m.playerVipObject.SetModified()
}

func (m *PlayerVipDataManager) IsCanReceiveFreeGift(level int32) bool {
	_, ok := m.playerVipObject.freeGiftMap[level]
	if !ok {
		return true
	}

	return false
}

// 获取礼包购买记录
func (m *PlayerVipDataManager) GetGiftRecord() (giftRecord map[int32]int32, freeGiftRecord map[int32]int32) {
	return m.playerVipObject.discountMap, m.playerVipObject.freeGiftMap
}

// 获取历史充值
func (m *PlayerVipDataManager) GetChargeNum() int64 {
	return m.playerVipObject.chargeNum
}

// GM充值礼包购买次数
func (m *PlayerVipDataManager) GMResetBuyGift() {
	m.playerVipObject.discountMap = map[int32]int32{}
	m.playerVipObject.SetModified()
}

func (m *PlayerVipDataManager) ToVipInfo() *viptypes.VipInfo {
	vipInfo := &viptypes.VipInfo{
		VipLevel: m.playerVipObject.vipLevel,
		VipStar:  m.playerVipObject.vipStar,
	}
	return vipInfo
}

func CreatePlayerVipDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerVipDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerVipDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerVipDataManager))
}
