package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	commonlog "fgame/fgame/common/log"
	gamevent "fgame/fgame/game/event"
	"fgame/fgame/game/shop/dao"
	shopentity "fgame/fgame/game/shop/entity"
	shopeventtypes "fgame/fgame/game/shop/event/types"
	"fgame/fgame/game/shop/shop"
)

//玩家当日商店购买道具对象
type PlayerShopObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	ShopId     int32
	DayCount   int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShopObject(pl player.Player) *PlayerShopObject {
	pmo := &PlayerShopObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerShopObjectToEntity(pso *PlayerShopObject) (*shopentity.PlayerShopEntity, error) {
	e := &shopentity.PlayerShopEntity{
		Id:         pso.Id,
		PlayerId:   pso.PlayerId,
		ShopId:     pso.ShopId,
		DayCount:   pso.DayCount,
		LastTime:   pso.LastTime,
		UpdateTime: pso.UpdateTime,
		CreateTime: pso.CreateTime,
		DeleteTime: pso.DeleteTime,
	}
	return e, nil
}

func (pso *PlayerShopObject) GetPlayerId() int64 {
	return pso.PlayerId
}

func (pso *PlayerShopObject) GetDBId() int64 {
	return pso.Id
}

func (pso *PlayerShopObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerShopObjectToEntity(pso)
	return e, err
}

func (pso *PlayerShopObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shopentity.PlayerShopEntity)

	pso.Id = pse.Id
	pso.PlayerId = pse.PlayerId
	pso.ShopId = pse.ShopId
	pso.DayCount = pse.DayCount
	pso.LastTime = pse.LastTime
	pso.UpdateTime = pse.UpdateTime
	pso.CreateTime = pse.CreateTime
	pso.DeleteTime = pse.DeleteTime
	return nil
}

func (pso *PlayerShopObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Shop: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pso.player.AddChangedObject(obj)
	return
}

//玩家当日商店购买限购道具管理器
type PlayerShopDataManager struct {
	p player.Player
	//玩家当日商店购买限购道具
	buyCountMap map[int32]*PlayerShopObject
}

func (psdm *PlayerShopDataManager) Player() player.Player {
	return psdm.p
}

//加载
func (psdm *PlayerShopDataManager) Load() (err error) {
	//加载玩家当日购买次数
	buyItems, err := dao.GetShopDao().GetShopList(psdm.p.GetId())
	if err != nil {
		return
	}

	//购买信息
	for _, item := range buyItems {
		pao := NewPlayerShopObject(psdm.p)
		pao.FromEntity(item)
		psdm.buyCountMap[pao.ShopId] = pao
	}
	return nil
}

//清数据
func (psdm *PlayerShopDataManager) clearAcrossDay(shopObj *PlayerShopObject, now int64) {
	shopObj.DayCount = 0
	shopObj.LastTime = 0
	shopObj.UpdateTime = now
	shopObj.SetModified()
	return
}

func (psdm *PlayerShopDataManager) deleteShopObj(shopObj *PlayerShopObject, now int64) {
	shopObj.DayCount = 0
	shopObj.LastTime = 0
	shopObj.UpdateTime = now
	shopObj.DeleteTime = now
	shopObj.SetModified()
	return
}

//刷新数据
func (psdm *PlayerShopDataManager) refresh() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for shopId, obj := range psdm.buyCountMap {
		//判断配置是否改过
		shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
		if shopTemplate == nil || shopTemplate.LimitCount == 0 {
			psdm.deleteShopObj(obj, now)
			delete(psdm.buyCountMap, shopId)
			continue
		}
		if obj.LastTime != 0 {
			flag, err := timeutils.IsSameFive(obj.LastTime, now)
			if err != nil {
				return err
			}
			if !flag {
				psdm.clearAcrossDay(obj, now)
			}
		}
	}

	return nil
}

//加载后
func (psdm *PlayerShopDataManager) AfterLoad() (err error) {
	err = psdm.refresh()
	return
}

//心跳
func (pmdm *PlayerShopDataManager) Heartbeat() {

}

//获取玩家当日商店购买道具
func (psdm *PlayerShopDataManager) GetShopBuyAll() map[int32]*PlayerShopObject {
	psdm.refresh()
	return psdm.buyCountMap
}

func (psdm *PlayerShopDataManager) GetShopBuyByShopId(shopId int32) *PlayerShopObject {
	if v, ok := psdm.buyCountMap[shopId]; ok {
		return v
	}
	return nil
}

func (psdm *PlayerShopDataManager) GetDayCountByShopId(shopId int32) (dayCount int32) {
	if v, ok := psdm.buyCountMap[shopId]; ok {
		return v.DayCount
	}
	return 0
}

//是否达到当日购买限制
func (psdm *PlayerShopDataManager) IfReachLimit(shopId int32, totalNum int32) (bool, error) {
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return false, nil
	}
	if shopTemplate.LimitCount == 0 {
		return false, nil
	}
	shopObj := psdm.GetShopBuyByShopId(shopId)
	if shopObj == nil {
		return false, nil
	}
	psdm.refreshDayCount(shopId)
	if shopObj.DayCount+totalNum > shopTemplate.LimitCount {
		return true, nil
	}
	return false, nil
}

func (psdm *PlayerShopDataManager) refreshDayCount(shopId int32) {
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.LimitCount == 0 {
		return
	}
	shopObj := psdm.GetShopBuyByShopId(shopId)
	if shopObj == nil {
		return
	}
	if shopObj.LastTime == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameFive(shopObj.LastTime, now)
	if !flag {
		psdm.clearAcrossDay(shopObj, now)
	}
}

func (psdm *PlayerShopDataManager) LeftDayCount(shopId int32) (isLimitBuy bool, num int32) {
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.LimitCount == 0 {
		return
	}
	shopObj := psdm.GetShopBuyByShopId(shopId)
	if shopObj == nil {
		return true, shopTemplate.LimitCount
	}
	psdm.refreshDayCount(shopId)
	return true, shopTemplate.LimitCount - shopObj.DayCount
}

func (psdm *PlayerShopDataManager) initShopObj(shopId int32, buyTimes int32) {
	if buyTimes <= 0 {
		return
	}
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	//不做每日限购
	if shopTemplate.LimitCount == 0 {
		return
	}
	pao := NewPlayerShopObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pao.Id = id
	//生成id
	pao.PlayerId = psdm.p.GetId()
	pao.ShopId = shopId
	if buyTimes > shopTemplate.LimitCount {
		pao.DayCount = shopTemplate.LimitCount
	} else {
		pao.DayCount = buyTimes
	}
	pao.CreateTime = now
	pao.LastTime = now
	psdm.buyCountMap[shopId] = pao
	pao.SetModified()
}

//更新对象
func (psdm *PlayerShopDataManager) UpdateObject(shopId int32, totalNum int32, autoBuy bool) {
	if totalNum <= 0 {
		return
	}
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}

	//日志
	costMoney := shopTemplate.GetConsumeMoney(totalNum)
	shopBuyLogReason := commonlog.ShopBuyLogReason
	shopStr := ""
	if autoBuy {
		shopStr = fmt.Sprintf("%s", "自动购买")
	} else {
		shopStr = fmt.Sprintf("%s", "入背包")
	}
	reasonText := fmt.Sprintf(shopBuyLogReason.String(), shopId, totalNum, costMoney, shopStr)
	data := shopeventtypes.CreatePlayerShopBuyItemLogEventData(shopId, totalNum, costMoney, shopBuyLogReason, reasonText)
	gamevent.Emit(shopeventtypes.EventTypeShopBuyItemLog, psdm.p, data)

	//不做每日限购
	if shopTemplate.LimitCount == 0 {
		return
	}
	shopObj := psdm.GetShopBuyByShopId(shopId)
	if shopObj != nil {
		now := global.GetGame().GetTimeService().Now()
		if shopObj.DayCount+totalNum > shopTemplate.LimitCount {
			shopObj.DayCount = shopTemplate.LimitCount
		} else {
			shopObj.DayCount += totalNum
		}
		shopObj.LastTime = now
		shopObj.UpdateTime = now
		shopObj.SetModified()
		return
	}
	psdm.initShopObj(shopId, totalNum)
	return
}

func (psdm *PlayerShopDataManager) GmClearDayCount() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range psdm.buyCountMap {
		obj.DayCount = 0
		obj.UpdateTime = now
		obj.LastTime = now
		obj.SetModified()
	}
}

func CreatePlayerShopDataManager(p player.Player) player.PlayerDataManager {
	psdm := &PlayerShopDataManager{}
	psdm.p = p
	psdm.buyCountMap = make(map[int32]*PlayerShopObject)
	return psdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerShopDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShopDataManager))
}
