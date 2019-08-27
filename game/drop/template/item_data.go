package template

import (
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/timeutils"
)

//掉落包结果
type DropItemData struct {
	ItemId         int32                               `json:"itemId"`
	Num            int32                               `json:"num"`
	Level          int32                               `json:"level"`          //物品等级
	Upstar         int32                               `json:"upstar"`         //金装强化星级
	AttrList       []int32                             `json:"attrList"`       //金装附件属性列表
	BindType       itemtypes.ItemBindType              `json:"bindType"`       //绑定类型
	IsRandomAttr   bool                                `json:"isRandomAttr"`   //是否随机金装属性
	ExpireType     inventorytypes.NewItemLimitTimeType `json:"expireType"`     //过期类型
	ExpireTime     int64                               `json:"expireTime"`     //过期时间
	ItemGetTime    int64                               `json:"itemGetTime"`    //获取时间
	Quality        int32                               `json:"quality"`        //宝宝品质
	Sex            playertypes.SexType                 `json:"sex"`            //宝宝性别
	TalentList     []*babytypes.TalentInfo             `json:"talentList"`     //宝宝天赋
	Danbei         int32                               `json:"danbei"`         //宝宝单倍
	OpenLightLevel int32                               `json:"openLightLevel"` //开光等级
	OpenTimes      int32                               `json:"openMin"`        //开光次数
}

func CreateItemData(itemId, num, level int32, bindType itemtypes.ItemBindType) *DropItemData {
	now := global.GetGame().GetTimeService().Now()
	attrList := []int32{}
	upstar := int32(0)
	isRandomAttr := false
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := now
	quality := int32(0)
	danbei := int32(0)
	talentList := []*babytypes.TalentInfo{}
	sex := playertypes.SexType(0)
	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, 0, 0, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateBaoBaoCardItemData(itemId, num, danbei, quality int32, sex playertypes.SexType, talentList []*babytypes.TalentInfo) *DropItemData {
	now := global.GetGame().GetTimeService().Now()
	level := int32(0)
	attrList := []int32{}
	upstar := int32(0)
	isRandomAttr := false
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := now
	bindType := itemtypes.ItemBindTypeUnBind

	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, 0, 0, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateItemDataWithExpire(itemId, num, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, expireTime, itemGetTime int64) *DropItemData {
	attrList := []int32{}
	upstar := int32(0)
	isRandomAttr := false
	quality := int32(0)
	danbei := int32(0)
	talentList := []*babytypes.TalentInfo{}
	sex := playertypes.SexType(0)
	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, 0, 0, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateItemDataWithData(data *DropItemData) *DropItemData {
	itemId := data.GetItemId()
	num := data.GetNum()
	level := data.GetLevel()
	bindType := data.GetBindType()
	upstar := data.GetUpstar()
	attrList := data.GetAttrList()
	isRandomAttr := data.GetIsRandomAttr()
	expireType := data.GetExpireType()
	expireTime := data.GetExpireTime()
	itemGetTime := data.GetItemGetTime()
	quality := data.GetQuality()
	danbei := data.GetDanbei()
	talentList := data.GetTalentList()
	sex := data.GetSex()
	openLightLevel := data.GetOpenLightLevel()
	openTimes := data.GetOpenTimes()
	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, openLightLevel, openTimes, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateItemDataWithGoldPropertyData(itemId, num, level int32, bindType itemtypes.ItemBindType, upstar int32, attrList []int32, openLightLevel int32, openTimes int32, isRandomAttr bool, expireType inventorytypes.NewItemLimitTimeType, expireTime, itemGetTime int64) *DropItemData {
	quality := int32(0)
	danbei := int32(0)
	talentList := []*babytypes.TalentInfo{}
	sex := playertypes.SexType(0)
	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, openLightLevel, openTimes, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateItemDataWithPropertyData(itemId, num, level int32, bindType itemtypes.ItemBindType, upstar int32, attrList []int32, isRandomAttr bool, expireType inventorytypes.NewItemLimitTimeType, expireTime, itemGetTime int64) *DropItemData {
	quality := int32(0)
	danbei := int32(0)
	talentList := []*babytypes.TalentInfo{}
	sex := playertypes.SexType(0)
	return CreateItemDataBasic(itemId, num, level, bindType, upstar, attrList, 0, 0, isRandomAttr, expireType, expireTime, itemGetTime, danbei, quality, sex, talentList)
}

func CreateItemDataBasic(itemId, num, level int32, bindType itemtypes.ItemBindType, upstar int32, attrList []int32, openLightLevel int32, openTimes int32, isRandomAttr bool, expireType inventorytypes.NewItemLimitTimeType, expireTime, itemGetTime int64, danbei, quality int32, sex playertypes.SexType, talentList []*babytypes.TalentInfo) *DropItemData {
	newAttrList := make([]int32, len(attrList))
	copy(newAttrList, attrList)

	newTalentList := make([]*babytypes.TalentInfo, 0, len(talentList))
	for _, talent := range talentList {
		newTalent := babytypes.NewTalentInfo(talent.SkillId, talent.Status, talent.Type)
		newTalentList = append(newTalentList, newTalent)
	}

	d := &DropItemData{
		ItemId:         itemId,
		Num:            num,
		Level:          level,
		BindType:       bindType,
		Upstar:         upstar,
		AttrList:       newAttrList,
		IsRandomAttr:   isRandomAttr,
		ExpireType:     expireType,
		ExpireTime:     expireTime,
		ItemGetTime:    itemGetTime,
		Quality:        quality,
		Sex:            sex,
		Danbei:         danbei,
		TalentList:     newTalentList,
		OpenLightLevel: openLightLevel,
		OpenTimes:      openTimes,
	}

	return d
}

func (d *DropItemData) GetItemId() int32 {
	return d.ItemId
}

func (d *DropItemData) GetNum() int32 {
	return d.Num
}

func (d *DropItemData) GetLevel() int32 {
	return d.Level
}

func (d *DropItemData) GetBindType() itemtypes.ItemBindType {
	return d.BindType
}

func (d *DropItemData) GetUpstar() int32 {
	return d.Upstar
}

func (d *DropItemData) GetQuality() int32 {
	return d.Quality
}

func (d *DropItemData) GetDanbei() int32 {
	return d.Danbei
}

func (d *DropItemData) GetSex() playertypes.SexType {
	return d.Sex
}

func (d *DropItemData) GetTalentList() []*babytypes.TalentInfo {
	return d.TalentList
}

func (d *DropItemData) GetAttrList() []int32 {
	return d.AttrList
}

func (d *DropItemData) GetExpireType() inventorytypes.NewItemLimitTimeType {
	return d.ExpireType
}

func (d *DropItemData) GetExpireTime() int64 {
	return d.ExpireTime
}

func (d *DropItemData) GetExpireTimestamp() int64 {
	switch d.ExpireType {
	case inventorytypes.NewItemLimitTimeTypeExpiredToday:
		{
			beginOfTime, _ := timeutils.BeginOfNow(d.ItemGetTime)
			return d.ExpireTime + beginOfTime
		}
	case inventorytypes.NewItemLimitTimeTypeExpiredAfterTime:
		{
			return d.ExpireTime + d.ItemGetTime
		}
	case inventorytypes.NewItemLimitTimeTypeExpiredDate:
		{
			return d.ExpireTime
		}
	}
	return 0
}

func (d *DropItemData) GetItemGetTime() int64 {
	return d.ItemGetTime
}

func (d *DropItemData) GetIsRandomAttr() bool {
	return d.IsRandomAttr
}

func (d *DropItemData) IsMerge(target *DropItemData) bool {
	if d.ItemId != target.ItemId {
		return false
	}

	if d.Level != target.Level {
		return false
	}

	if d.BindType != target.BindType {
		return false
	}

	if d.ExpireType != target.ExpireType && d.ExpireTime != target.ExpireTime {
		return false
	}

	return true
}

func (d *DropItemData) GetOpenLightLevel() int32 {
	return d.OpenLightLevel
}

func (d *DropItemData) GetOpenTimes() int32 {
	return d.OpenTimes
}
