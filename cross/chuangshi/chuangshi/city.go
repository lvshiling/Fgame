package chuangshi

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//城池数据
type CityData struct {
	city        *ChuangShiCityObject
	jianSheList []*ChuangShiCityJianSheObject
}

func (d *CityData) ifFuShu() bool {
	return d.city.typ == chuangshitypes.ChuangShiCityTypeFushu
}

func (d *CityData) GetCity() *ChuangShiCityObject {
	return d.city
}

func (d *CityData) GetChengFangJianSheList() []*ChuangShiCityJianSheObject {
	return d.jianSheList
}

func (d *CityData) GetChengFangJianShe(jianSheType chuangshitypes.ChuangShiCityJianSheType) *ChuangShiCityJianSheObject {
	for _, jianShe := range d.jianSheList {
		if jianShe.jianSheType != jianSheType {
			continue
		}
		return jianShe
	}
	return nil
}

func NewCityDatt() *CityData {
	return &CityData{}
}

type ChuangShiCityObject struct {
	camp         *Camp
	id           int64
	platform     int32
	serverId     int32
	campType     chuangshitypes.ChuangShiCampType
	originalCamp chuangshitypes.ChuangShiCampType
	typ          chuangshitypes.ChuangShiCityType
	index        int32
	ownerId      int64
	jifen        int64
	diamonds     int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewChuangShiCityObject(camp *Camp) *ChuangShiCityObject {
	o := &ChuangShiCityObject{}
	o.camp = camp
	return o
}

func convertChuangShiCityObjectToEntity(o *ChuangShiCityObject) (*chuangshientity.ChuangShiCityEntity, error) {
	e := &chuangshientity.ChuangShiCityEntity{
		Id:           o.id,
		Platform:     o.platform,
		ServerId:     o.serverId,
		CampType:     int32(o.campType),
		OriginalCamp: int32(o.originalCamp),
		Typ:          int32(o.typ),
		Index:        o.index,
		OwnerId:      o.ownerId,
		Jifen:        o.jifen,
		Diamonds:     o.diamonds,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}
func (o *ChuangShiCityObject) GetId() int64 {
	return o.id
}

func (o *ChuangShiCityObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiCityObject) GeOwnerId() int64 {
	return o.ownerId
}

func (o *ChuangShiCityObject) GetCampType() chuangshitypes.ChuangShiCampType {
	return o.campType
}

func (o *ChuangShiCityObject) GetOrignalCampType() chuangshitypes.ChuangShiCampType {
	return o.originalCamp
}

func (o *ChuangShiCityObject) GetIndex() int32 {
	return o.index
}

func (o *ChuangShiCityObject) GetDiamonds() int64 {
	return o.diamonds
}

func (o *ChuangShiCityObject) GetJifen() int64 {
	return o.jifen
}

func (o *ChuangShiCityObject) GetType() chuangshitypes.ChuangShiCityType {
	return o.typ
}

func (o *ChuangShiCityObject) GetCamp() *Camp {
	return o.camp
}

func (o *ChuangShiCityObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiCityObjectToEntity(o)
	return e, err
}

func (o *ChuangShiCityObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiCityEntity)
	o.id = pse.Id
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.campType = chuangshitypes.ChuangShiCampType(pse.CampType)
	o.originalCamp = chuangshitypes.ChuangShiCampType(pse.OriginalCamp)
	o.typ = chuangshitypes.ChuangShiCityType(pse.Typ)
	o.index = pse.Index
	o.ownerId = pse.OwnerId
	o.jifen = pse.Jifen
	o.diamonds = pse.Diamonds
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiCityObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiCity"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

//
//
type ChuangShiCityJianSheObject struct {
	id            int64
	platform      int32
	serverId      int32
	cityId        int64
	jianSheType   chuangshitypes.ChuangShiCityJianSheType
	jianSheExp    int32
	jianSheLevel  int32
	skillMap      map[int32]int64 //技能激活记录（天气台专用）
	skillLevelSet int32           //当前使用技能（天气台专用）
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewChuangShiCityJianSheObject() *ChuangShiCityJianSheObject {
	o := &ChuangShiCityJianSheObject{}
	return o
}

func convertChuangShiCityJianSheObjectToEntity(o *ChuangShiCityJianSheObject) (*chuangshientity.ChuangShiCityJianSheEntity, error) {

	data, err := json.Marshal(o.skillMap)
	if err != nil {
		return nil, err
	}

	e := &chuangshientity.ChuangShiCityJianSheEntity{
		Id:            o.id,
		Platform:      o.platform,
		ServerId:      o.serverId,
		CityId:        o.cityId,
		JianSheType:   int32(o.jianSheType),
		JianSheExp:    o.jianSheExp,
		JianSheLevel:  o.jianSheLevel,
		SkillLevelSet: o.skillLevelSet,
		SkillMap:      string(data),
		UpdateTime:    o.updateTime,
		CreateTime:    o.createTime,
		DeleteTime:    o.deleteTime,
	}
	return e, nil
}
func (o *ChuangShiCityJianSheObject) GetId() int64 {
	return o.id
}

func (o *ChuangShiCityJianSheObject) GetDBId() int64 {
	return o.id
}

func (o *ChuangShiCityJianSheObject) GeCityId() int64 {
	return o.cityId
}

func (o *ChuangShiCityJianSheObject) GetJianSheType() chuangshitypes.ChuangShiCityJianSheType {
	return o.jianSheType
}

func (o *ChuangShiCityJianSheObject) GetJianSheExp() int32 {
	return o.jianSheExp
}

func (o *ChuangShiCityJianSheObject) GetJianSheLevel() int32 {
	return o.jianSheLevel
}

func (o *ChuangShiCityJianSheObject) GetSkillLevelSet() int32 {
	return o.skillLevelSet
}

func (o *ChuangShiCityJianSheObject) GetSkillMap() map[int32]int64 {
	return o.skillMap
}

func (o *ChuangShiCityJianSheObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChuangShiCityJianSheObjectToEntity(o)
	return e, err
}

func (o *ChuangShiCityJianSheObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chuangshientity.ChuangShiCityJianSheEntity)

	dataMap := make(map[int32]int64)
	err := json.Unmarshal([]byte(pse.SkillMap), &dataMap)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.platform = pse.Platform
	o.serverId = pse.ServerId
	o.cityId = pse.CityId
	o.jianSheLevel = pse.JianSheLevel
	o.jianSheExp = pse.JianSheExp
	o.jianSheType = chuangshitypes.ChuangShiCityJianSheType(pse.JianSheType)
	o.skillLevelSet = pse.SkillLevelSet
	o.skillMap = dataMap
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChuangShiCityJianSheObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChuangShiCityJianShe"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

func (o *ChuangShiCityJianSheObject) IfActivate(level int32) bool {
	_, ok := o.skillMap[level]
	return ok
}
