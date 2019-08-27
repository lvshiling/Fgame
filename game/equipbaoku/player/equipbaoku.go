package player

import (
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/equipbaoku/dao"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家装备宝库管理器
type PlayerEquipBaoKuDataManager struct {
	p player.Player
	//玩家装备宝库对象
	playerEquipBaoKuObjectMap map[equipbaokutypes.BaoKuType]*PlayerEquipBaoKuObject
	//玩家当日宝库商店购买限购道具
	buyCountMap map[int32]*PlayerEquipBaoKuShopObject
}

func (m *PlayerEquipBaoKuDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerEquipBaoKuDataManager) Load() (err error) {
	m.playerEquipBaoKuObjectMap = make(map[equipbaokutypes.BaoKuType]*PlayerEquipBaoKuObject)

	//加载玩家装备宝库信息
	equipBaoKuEntity, err := dao.GetEquipBaoKuDao().GetPlayerEquipBaoKuEntity(m.p.GetId())
	if err != nil {
		return
	}

	//加载玩家当日购买次数
	buyItems, err := dao.GetEquipBaoKuDao().GetEquipBaoKuShopList(m.p.GetId())
	if err != nil {
		return
	}

	//玩家装备宝库信息
	for _, e := range equipBaoKuEntity {
		obj := NewPlayerEquipBaoKuObject(m.p)
		obj.FromEntity(e)
		m.playerEquipBaoKuObjectMap[obj.typ] = obj
	}

	//购买信息
	for _, item := range buyItems {
		pao := NewPlayerEquipBaoKuShopObject(m.p)
		pao.FromEntity(item)
		m.buyCountMap[pao.ShopId] = pao
	}

	return nil
}

//第一次初始化
func (m *PlayerEquipBaoKuDataManager) initPlayerEquipBaoKuObject(typ equipbaokutypes.BaoKuType) *PlayerEquipBaoKuObject {
	obj := NewPlayerEquipBaoKuObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.typ = typ
	obj.luckyPoints = int32(0)
	obj.attendPoints = int32(0)
	obj.totalAttendTimes = int32(0)
	obj.lastSystemRefreshTime = now
	obj.createTime = now

	return obj
}

//加载后
func (m *PlayerEquipBaoKuDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerEquipBaoKuDataManager) Heartbeat() {

}

//获取装备宝库信息
func (m *PlayerEquipBaoKuDataManager) GetEquipBaoKuObj(typ equipbaokutypes.BaoKuType) *PlayerEquipBaoKuObject {
	_ = m.RefreshLuckyPoints()
	obj := m.getPlayerBaoKuObject(typ)
	return obj
}

//刷新装备宝库幸运值
func (m *PlayerEquipBaoKuDataManager) RefreshLuckyPoints() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.playerEquipBaoKuObjectMap {
		isSame, err := timeutils.IsSameFive(obj.lastSystemRefreshTime, now)
		if err != nil {
			return err
		}
		if !isSame {
			beforePoints := obj.luckyPoints
			obj.luckyPoints = 0
			obj.lastSystemRefreshTime = now
			obj.updateTime = now
			obj.SetModified()
			data := &equipbaokutypes.BaoKuData{
				Typ:         obj.typ,
				LuckyPoints: beforePoints,
			}
			gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuLuckyBox, m.p, data)
		}
	}
	return
}

func (m *PlayerEquipBaoKuDataManager) getPlayerBaoKuObject(typ equipbaokutypes.BaoKuType) *PlayerEquipBaoKuObject {
	obj, ok := m.playerEquipBaoKuObjectMap[typ]
	if !ok {
		obj = m.initPlayerEquipBaoKuObject(typ)
		obj.SetModified()
		m.playerEquipBaoKuObjectMap[typ] = obj
	}
	return obj
}

//TODO: cjb添加日志
//探索装备宝库
func (m *PlayerEquipBaoKuDataManager) AttendEquipBaoKu(luckyVal, attendVal, attendCount int32, changeType commontypes.ChangeType, typ equipbaokutypes.BaoKuType) (flag bool) {
	_ = m.RefreshLuckyPoints()
	obj := m.getPlayerBaoKuObject(typ)
	if luckyVal > 0 {
		befLuckyNum := obj.luckyPoints
		obj.luckyPoints += luckyVal
		//宝库幸运值日志
		luckyReason := commonlog.EquipBaoKuLogReasonLuckyPointsChange
		luckyReasonText := fmt.Sprintf(luckyReason.String(), typ.GetBaoKuName(), changeType.String())
		luckyData := equipbaokueventtypes.CreatePlayerEquipBaoKuLuckyPointsLogEventData(befLuckyNum, obj.luckyPoints, nil, luckyReason, luckyReasonText, typ)
		gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuLuckyPointsLog, m.p, luckyData)
	}
	if attendVal > 0 {
		befAttendNum := obj.attendPoints
		obj.attendPoints += attendVal
		//宝库积分日志
		jiFenReason := commonlog.EquipBaoKuLogReasonAttendPointsChange
		jiFenReasonText := fmt.Sprintf(jiFenReason.String(), typ.GetBaoKuName(), changeType.String())
		attendData := equipbaokueventtypes.CreatePlayerEquipBaoKuAttendPointsLogEventData(befAttendNum, obj.attendPoints, 0, 0, jiFenReason, jiFenReasonText, typ)
		gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuAttendPointsLog, m.p, attendData)
	}
	if attendCount > 0 {
		obj.totalAttendTimes += attendCount

		//宝库探索
		gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuAttend, m.p, attendCount)
	}

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}

//消耗装备宝库幸运值
func (m *PlayerEquipBaoKuDataManager) SubEquipBaoKuLuckyPoints(luckyVal int32, typ equipbaokutypes.BaoKuType) (flag bool) {
	_ = m.RefreshLuckyPoints()
	obj := m.getPlayerBaoKuObject(typ)
	if luckyVal > obj.luckyPoints {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.luckyPoints -= luckyVal
	obj.updateTime = now
	obj.SetModified()

	flag = true
	return
}

//TODO cjb:记录日志
//消耗装备宝库积分
func (m *PlayerEquipBaoKuDataManager) SubEquipBaoKuAttendPoints(attendVal int32, typ equipbaokutypes.BaoKuType) (flag bool) {
	obj := m.getPlayerBaoKuObject(typ)
	if attendVal > obj.attendPoints {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.attendPoints -= attendVal
	obj.updateTime = now
	obj.SetModified()

	flag = true
	return
}

//获取装备宝库掉落包
func (m *PlayerEquipBaoKuDataManager) GetEquipBaoKuDrop(times int32, typ equipbaokutypes.BaoKuType) (rewList []*droptemplate.DropItemData) {
	equipBaoKuTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetEquipBaoKuByLevAndZhuanNum(m.p.GetLevel(), m.p.GetZhuanSheng(), typ)
	obj := m.getPlayerBaoKuObject(typ)
	curTotal := obj.totalAttendTimes
	for index := int32(0); index < times; index++ {
		curTotal += 1
		dropId := equipBaoKuTemplate.DropId

		rewDropByTimesMap := equipBaoKuTemplate.GetRewDropMap()
		timesDescList := equipBaoKuTemplate.GetDropTimesDescList()

		vipType := viptypes.CostLevelRuleTypeMaterialBaoKu
		if typ == equipbaokutypes.BaoKuTypeEquip {
			vipType = viptypes.CostLevelRuleTypeEquipBaoKu
		}
		ruleTimesMap := viplogic.CountDropTimesWithCostLevel(m.p, vipType, timesDescList)
		for _, times := range timesDescList {
			ruleTimes := ruleTimesMap[times]
			ret := curTotal % ruleTimes
			if ret == 0 {
				dropId = rewDropByTimesMap[int32(times)]
				break
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropBaoKuItemLevel(dropId)
		if dropData != nil {
			rewList = append(rewList, dropData)
		}
	}

	return
}

//宝库商店清数据
func (m *PlayerEquipBaoKuDataManager) clearAcrossDay(shopObj *PlayerEquipBaoKuShopObject, now int64) {
	shopObj.DayCount = 0
	shopObj.LastTime = 0
	shopObj.UpdateTime = now
	shopObj.SetModified()
	return
}

func (m *PlayerEquipBaoKuDataManager) deleteShopObj(shopObj *PlayerEquipBaoKuShopObject, now int64) {
	shopObj.DayCount = 0
	shopObj.LastTime = 0
	shopObj.UpdateTime = now
	shopObj.DeleteTime = now
	shopObj.SetModified()
	return
}

//刷新数据
func (m *PlayerEquipBaoKuDataManager) RefreshEquipBaoKuShop() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for shopId, obj := range m.buyCountMap {
		//判断配置是否改过
		shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
		if shopTemplate == nil {
			m.deleteShopObj(obj, now)
			delete(m.buyCountMap, shopId)
			continue
		}
		if obj.LastTime != 0 {
			flag, err := timeutils.IsSameFive(obj.LastTime, now)
			if err != nil {
				return err
			}
			if !flag {
				m.clearAcrossDay(obj, now)
			}
		}
	}

	return nil
}

//获取玩家当日商店购买道具
func (m *PlayerEquipBaoKuDataManager) GetEquipBaoKuShopBuyAll() map[int32]*PlayerEquipBaoKuShopObject {
	m.RefreshEquipBaoKuShop()
	return m.buyCountMap
}

func (m *PlayerEquipBaoKuDataManager) GetEquipBaoKuShopBuyByShopId(shopId int32) *PlayerEquipBaoKuShopObject {
	if v, ok := m.buyCountMap[shopId]; ok {
		return v
	}
	return nil
}

func (m *PlayerEquipBaoKuDataManager) GetDayCountByShopId(shopId int32) (dayCount int32) {
	if v, ok := m.buyCountMap[shopId]; ok {
		return v.DayCount
	}
	return 0
}

//是否达到当日购买限制
func (m *PlayerEquipBaoKuDataManager) IfReachLimit(shopId int32, totalNum int32) (bool, error) {
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil {
		return false, nil
	}
	if shopTemplate.MaxCount == 0 {
		return false, nil
	}
	shopObj := m.GetEquipBaoKuShopBuyByShopId(shopId)
	if shopObj == nil {
		return false, nil
	}
	m.refreshDayCount(shopId)
	if shopObj.DayCount+totalNum > shopTemplate.MaxCount {
		return true, nil
	}
	return false, nil
}

func (m *PlayerEquipBaoKuDataManager) refreshDayCount(shopId int32) {
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.MaxCount == 0 {
		return
	}
	shopObj := m.GetEquipBaoKuShopBuyByShopId(shopId)
	if shopObj == nil {
		return
	}
	if shopObj.LastTime == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameFive(shopObj.LastTime, now)
	if !flag {
		m.clearAcrossDay(shopObj, now)
	}
}

func (m *PlayerEquipBaoKuDataManager) LeftDayCount(shopId int32) (isLimitBuy bool, num int32) {
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.MaxCount == 0 {
		return
	}
	shopObj := m.GetEquipBaoKuShopBuyByShopId(shopId)
	if shopObj == nil {
		return true, shopTemplate.MaxCount
	}
	m.refreshDayCount(shopId)
	return true, shopTemplate.MaxCount - shopObj.DayCount
}

func (m *PlayerEquipBaoKuDataManager) initEquipBaoKuShopObj(shopId int32, buyTimes int32) {
	if buyTimes <= 0 {
		return
	}
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.MaxCount == 0 {
		return
	}
	pao := NewPlayerEquipBaoKuShopObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pao.Id = id
	//生成id
	pao.PlayerId = m.p.GetId()
	pao.ShopId = shopId
	if buyTimes > shopTemplate.MaxCount {
		pao.DayCount = shopTemplate.MaxCount
	} else {
		pao.DayCount = buyTimes
	}
	pao.CreateTime = now
	pao.LastTime = now
	m.buyCountMap[shopId] = pao
	pao.SetModified()
}

//更新对象
func (m *PlayerEquipBaoKuDataManager) UpdateEquipBaoKuShopObject(shopId int32, totalNum int32) {
	if totalNum <= 0 {
		return
	}
	shopTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetBaoKuJiFenTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}

	//不做每日限购
	if shopTemplate.MaxCount != 0 {
		shopObj := m.GetEquipBaoKuShopBuyByShopId(shopId)
		if shopObj != nil {
			now := global.GetGame().GetTimeService().Now()
			if shopObj.DayCount+totalNum > shopTemplate.MaxCount {
				shopObj.DayCount = shopTemplate.MaxCount
			} else {
				shopObj.DayCount += totalNum
			}
			shopObj.LastTime = now
			shopObj.UpdateTime = now
			shopObj.SetModified()

		} else {
			m.initEquipBaoKuShopObj(shopId, totalNum)
		}
	}
	//宝库兑换
	gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuBuy, m.p, totalNum)
	return
}

//GM设置幸运值
func (m *PlayerEquipBaoKuDataManager) GMSetLuckyPoints(val int32, typ equipbaokutypes.BaoKuType) (err error) {
	obj := m.getPlayerBaoKuObject(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.luckyPoints = val
	obj.updateTime = now
	obj.SetModified()
	return
}

//GM设置参与积分
func (m *PlayerEquipBaoKuDataManager) GMSetAttendPoints(val int32, typ equipbaokutypes.BaoKuType) (err error) {
	obj := m.getPlayerBaoKuObject(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.attendPoints = val
	obj.updateTime = now
	obj.SetModified()
	return
}

//GM设置抽奖次数
func (m *PlayerEquipBaoKuDataManager) GMSetTotalAttendTimes(val int32, typ equipbaokutypes.BaoKuType) (err error) {
	obj := m.getPlayerBaoKuObject(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.totalAttendTimes = val
	obj.updateTime = now
	obj.SetModified()
	return
}

//GM清理兑换次数
func (m *PlayerEquipBaoKuDataManager) GmClearDayCount() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.buyCountMap {
		obj.DayCount = 0
		obj.UpdateTime = now
		obj.LastTime = now
		obj.SetModified()
	}
}

func CreatePlayerEquipBaoKuDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerEquipBaoKuDataManager{}
	m.p = p
	m.buyCountMap = make(map[int32]*PlayerEquipBaoKuShopObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerEquipBaoKuDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerEquipBaoKuDataManager))
}
