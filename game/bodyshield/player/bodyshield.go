package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/bodyshield/bodyshield"
	"fgame/fgame/game/bodyshield/dao"
	bshieldentity "fgame/fgame/game/bodyshield/entity"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//护体盾对象
type PlayerBodyShieldObject struct {
	player         player.Player
	Id             int64
	PlayerId       int64
	AdvanceId      int
	JinjiadanLevel int32
	JinjiadanNum   int32
	JinjiadanPro   int32
	TimesNum       int32
	Bless          int32
	BlessTime      int64
	ShieldId       int32
	ShieldNum      int32
	ShieldPro      int32
	Power          int64
	SPower         int64
	UpdateTime     int64
	CreateTime     int64
	DeleteTime     int64
}

func NewPlayerBodyShieldObject(pl player.Player) *PlayerBodyShieldObject {
	o := &PlayerBodyShieldObject{
		player: pl,
	}
	return o
}

func convertNewPlayerBodyShieldObjectToEntity(o *PlayerBodyShieldObject) (*bshieldentity.PlayerBodyShieldEntity, error) {

	e := &bshieldentity.PlayerBodyShieldEntity{
		Id:             o.Id,
		PlayerId:       o.PlayerId,
		AdvancedId:     o.AdvanceId,
		JinjiadanLevel: o.JinjiadanLevel,
		JinjiadanNum:   o.JinjiadanNum,
		JinjiadanPro:   o.JinjiadanPro,
		TimesNum:       o.TimesNum,
		Bless:          o.Bless,
		BlessTime:      o.BlessTime,
		ShieldId:       o.ShieldId,
		ShieldNum:      o.ShieldNum,
		ShieldPro:      o.ShieldPro,
		Power:          o.Power,
		SPower:         o.SPower,
		UpdateTime:     o.UpdateTime,
		CreateTime:     o.CreateTime,
		DeleteTime:     o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerBodyShieldObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerBodyShieldObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerBodyShieldObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerBodyShieldObjectToEntity(o)
	return e, err
}

func (o *PlayerBodyShieldObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*bshieldentity.PlayerBodyShieldEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.AdvanceId = pse.AdvancedId
	o.JinjiadanLevel = pse.JinjiadanLevel
	o.JinjiadanNum = pse.JinjiadanNum
	o.JinjiadanPro = pse.JinjiadanPro
	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.ShieldId = pse.ShieldId
	o.ShieldNum = pse.ShieldNum
	o.ShieldPro = pse.ShieldPro
	o.Power = pse.Power
	o.SPower = pse.SPower
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerBodyShieldObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BodyShield"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家护体盾管理器
type PlayerBodyShieldDataManager struct {
	p player.Player
	//玩家护体盾对象
	playerBodyShieldObject *PlayerBodyShieldObject
}

func (m *PlayerBodyShieldDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerBodyShieldDataManager) Load() (err error) {
	//加载玩家护体盾信息
	bShieldentity, err := dao.GetBodyShieldDao().GetBodyShieldEntity(m.p.GetId())
	if err != nil {
		return
	}
	if bShieldentity == nil {
		m.initPlayerBodyShieldObject()
	} else {
		m.playerBodyShieldObject = NewPlayerBodyShieldObject(m.p)
		m.playerBodyShieldObject.FromEntity(bShieldentity)
	}

	return nil
}

//第一次初始化
func (m *PlayerBodyShieldDataManager) initPlayerBodyShieldObject() {
	o := NewPlayerBodyShieldObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.Id = id
	//生成id
	o.PlayerId = m.p.GetId()
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.BodyShield
	shieldId := playerCreateTemplate.Shield
	o.AdvanceId = int(advanceId)
	o.JinjiadanLevel = 0
	o.JinjiadanNum = int32(0)
	o.JinjiadanPro = 0
	o.TimesNum = int32(0)
	o.Bless = int32(0)
	o.BlessTime = int64(0)
	o.ShieldId = shieldId
	o.ShieldNum = 0
	o.ShieldPro = 0
	o.Power = int64(0)
	o.SPower = int64(0)
	o.CreateTime = now
	m.playerBodyShieldObject = o
	o.SetModified()
}

func (m *PlayerBodyShieldDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerBodyShieldObject.AdvanceId)
	nextNumber := number + 1
	bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(nextNumber)
	if bodyShieldTemplate == nil {
		return
	}
	if !bodyShieldTemplate.GetIsClear() {
		return
	}

	lastTime := m.playerBodyShieldObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerBodyShieldObject.Bless = 0
			m.playerBodyShieldObject.BlessTime = 0
			m.playerBodyShieldObject.TimesNum = 0
			m.playerBodyShieldObject.SetModified()
		}
	}
	return
}

//加载后
func (m *PlayerBodyShieldDataManager) AfterLoad() (err error) {
	m.refreshBless()
	return nil
}

//护体盾信息对象
func (m *PlayerBodyShieldDataManager) GetBodyShiedInfo() *PlayerBodyShieldObject {
	m.refreshBless()
	return m.playerBodyShieldObject
}

// func (m *PlayerBodyShieldDataManager) EatJinJiaDan(pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}
// 	if sucess {
// 		m.playerBodyShieldObject.JinjiadanLevel += 1
// 		m.playerBodyShieldObject.JinjiadanNum = 0
// 		m.playerBodyShieldObject.JinjiadanPro = pro
// 	} else {
// 		m.playerBodyShieldObject.JinjiadanNum += 1
// 		m.playerBodyShieldObject.JinjiadanPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerBodyShieldObject.UpdateTime = now
// 	m.playerBodyShieldObject.SetModified()
// 	return
// }

func (m *PlayerBodyShieldDataManager) EatJinJiaDan(level int32) {
	if m.playerBodyShieldObject.JinjiadanLevel == level || level <= 0 {
		return
	}
	jinJiaDanTemplate := bodyshield.GetBodyShieldService().GetBodyShieldJinJia(level)
	if jinJiaDanTemplate == nil {
		return
	}

	m.playerBodyShieldObject.JinjiadanLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	return
}

//心跳
func (m *PlayerBodyShieldDataManager) Heartbeat() {

}

//进阶
func (m *PlayerBodyShieldDataManager) BodyShieldAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(m.playerBodyShieldObject.AdvanceId + 1))
		if bodyShieldTemplate == nil {
			return
		}
		m.playerBodyShieldObject.AdvanceId += 1
		m.playerBodyShieldObject.TimesNum = 0
		m.playerBodyShieldObject.Bless = 0
		m.playerBodyShieldObject.BlessTime = 0

		gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvanced, m.p, m.playerBodyShieldObject.AdvanceId)
	} else {
		m.playerBodyShieldObject.TimesNum += addTimes
		if m.playerBodyShieldObject.Bless == 0 {
			m.playerBodyShieldObject.BlessTime = now
		}
		m.playerBodyShieldObject.Bless += pro
	}
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerBodyShieldDataManager) BodyShieldAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerBodyShieldObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(nextAdvancedId))
		if bodyShieldTemplate == nil {
			return
		}
		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.AdvanceId += canAddNum
	m.playerBodyShieldObject.TimesNum = 0
	m.playerBodyShieldObject.Bless = 0
	m.playerBodyShieldObject.BlessTime = 0
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvanced, m.p, m.playerBodyShieldObject.AdvanceId)
	return
}

//神盾尖刺培养
func (m *PlayerBodyShieldDataManager) ShieldFeed(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	if sucess {
		shieldTemplate := bodyshield.GetBodyShieldService().GetShield(m.playerBodyShieldObject.ShieldId + 1)
		if shieldTemplate == nil {
			return
		}
		m.playerBodyShieldObject.ShieldId += 1
		m.playerBodyShieldObject.ShieldNum = 0
		m.playerBodyShieldObject.ShieldPro = pro

		gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvanced, m.p, m.playerBodyShieldObject.ShieldId)
	} else {
		m.playerBodyShieldObject.ShieldNum += addTimes
		m.playerBodyShieldObject.ShieldPro += pro
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	return
}

//直升券神盾尖刺升阶
func (m *PlayerBodyShieldDataManager) ShieldFeedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerBodyShieldObject.ShieldId + 1
	for addAdvancedNum > 0 {
		shieldTemplate := bodyshield.GetBodyShieldService().GetShield(nextAdvancedId)
		if shieldTemplate == nil {
			return
		}

		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.ShieldId += int32(canAddNum)
	m.playerBodyShieldObject.ShieldNum = 0
	m.playerBodyShieldObject.ShieldPro = 0
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvanced, m.p, m.playerBodyShieldObject.ShieldId)
	return
}

//护体盾战斗力
func (m *PlayerBodyShieldDataManager) BodyShieldPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := m.playerBodyShieldObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.Power = power
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldPowerChanged, m.p, power)
	return
}

//神盾尖刺战斗力
func (m *PlayerBodyShieldDataManager) ShieldPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := m.playerBodyShieldObject.SPower
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.SPower = power
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
	gameevent.Emit(bodyshieldeventtypes.EventTypeShieldPowerChanged, m.p, power)
	return
}

func (m *PlayerBodyShieldDataManager) ToBodyShieldInfo() *bodyshieldtypes.BodyShieldInfo {
	bodyShieldInfo := &bodyshieldtypes.BodyShieldInfo{
		AdvancedId:     int32(m.playerBodyShieldObject.AdvanceId),
		JinjiadanLevel: m.playerBodyShieldObject.JinjiadanLevel,
		JinjiadanPro:   m.playerBodyShieldObject.JinjiadanPro,
	}
	return bodyShieldInfo
}

func (m *PlayerBodyShieldDataManager) ToShieldInfo() *bodyshieldtypes.ShieldInfo {
	shieldInfo := &bodyshieldtypes.ShieldInfo{
		ShieldId: m.playerBodyShieldObject.ShieldId,
		Progress: m.playerBodyShieldObject.ShieldPro,
	}
	return shieldInfo
}

//仅gm使用 护体盾进阶
func (m *PlayerBodyShieldDataManager) GmSetBodyShieldAdvanced(advancedId int) {
	m.playerBodyShieldObject.AdvanceId = advancedId
	m.playerBodyShieldObject.TimesNum = int32(0)
	m.playerBodyShieldObject.Bless = int32(0)
	m.playerBodyShieldObject.BlessTime = int64(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()

	gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvanced, m.p, m.playerBodyShieldObject.AdvanceId)
	return
}

//仅gm使用 护体盾食金甲丹等级
func (m *PlayerBodyShieldDataManager) GmSetBodyShieldJinJiaDanLevel(level int32) {

	m.playerBodyShieldObject.JinjiadanLevel = level
	m.playerBodyShieldObject.JinjiadanNum = 0
	m.playerBodyShieldObject.JinjiadanPro = 0

	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()
}

//仅gm使用 神盾尖刺进阶
func (m *PlayerBodyShieldDataManager) GmSetShieldAdvanced(shieldId int32) {
	m.playerBodyShieldObject.ShieldId = shieldId
	m.playerBodyShieldObject.ShieldNum = int32(0)
	m.playerBodyShieldObject.ShieldPro = int32(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerBodyShieldObject.UpdateTime = now
	m.playerBodyShieldObject.SetModified()

	gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvanced, m.p, m.playerBodyShieldObject.ShieldId)
	return
}

func CreatePlayerBodyShieldDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerBodyShieldDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerBShieldDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerBodyShieldDataManager))
}
