package types

import (
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/timeutils"
)

type EquipmentSlotInfo struct {
	SlotId int32           `json:"slotId"`
	Level  int32           `json:"level"`
	Star   int32           `json:"star"`
	ItemId int32           `json:"itemId"`
	Gems   map[int32]int32 `json:"gems"`
}

// 物品属性
type ItemPropertyData interface {
	IsExpire() bool                      //物品是否过期
	GetExpireType() NewItemLimitTimeType //过期类型
	GetExpireTime() int64                //过期时间(配置数据)
	GetExpireTimestamp() int64           //根据过期类型计算后的时间戳
	GetItemGetTime() int64               //物品获取时间
	InitBase()                           //兼容处理，初始化父类
	Copy() ItemPropertyData
}

// 基础物品属性
type ItemPropertyDataBase struct {
	ExpireType  NewItemLimitTimeType `json:"expireType"`  //过期类型
	ExpireTime  int64                `json:"expireTime"`  //过期时间
	ItemGetTime int64                `json:"itemGetTime"` //过期时间
}

func (base *ItemPropertyDataBase) InitBase() {}

func (base *ItemPropertyDataBase) GetItemGetTime() int64 {
	return base.ItemGetTime
}

func (base *ItemPropertyDataBase) GetExpireType() NewItemLimitTimeType {
	return base.ExpireType
}

func (base *ItemPropertyDataBase) GetExpireTime() int64 {
	return base.ExpireTime
}

func (base *ItemPropertyDataBase) GetExpireTimestamp() int64 {
	switch base.ExpireType {
	case NewItemLimitTimeTypeExpiredToday:
		{
			beginOfTime, _ := timeutils.BeginOfNow(base.ItemGetTime)
			return base.ExpireTime + beginOfTime
		}
	case NewItemLimitTimeTypeExpiredAfterTime:
		{
			return base.ExpireTime + base.ItemGetTime
		}
	case NewItemLimitTimeTypeExpiredDate:
		{
			return base.ExpireTime
		}
	}
	return 0
}

func (base *ItemPropertyDataBase) IsExpire() bool {
	now := global.GetGame().GetTimeService().Now()
	expire := base.GetExpireTimestamp()
	if expire <= 0 {
		return false
	}
	if now > expire {
		return true
	}
	return false
}

func (base *ItemPropertyDataBase) CopyBase() *ItemPropertyDataBase {
	copyBase := &ItemPropertyDataBase{}
	copyBase.ExpireType = base.ExpireType
	copyBase.ExpireTime = base.ExpireTime
	copyBase.ItemGetTime = base.ItemGetTime
	return copyBase
}

func (base *ItemPropertyDataBase) Copy() ItemPropertyData {
	return base.CopyBase()
}

func CreateItemPropertyDataBase(expireType NewItemLimitTimeType, expireTime, itemGetTime int64) *ItemPropertyDataBase {
	base := &ItemPropertyDataBase{}
	base.ExpireType = expireType
	base.ExpireTime = expireTime
	base.ItemGetTime = itemGetTime
	return base
}

func CreateDefaultItemPropertyDataBase() *ItemPropertyDataBase {
	now := global.GetGame().GetTimeService().Now()
	return CreateItemPropertyDataBase(NewItemLimitTimeTypeNone, 0, now)
}
