package shenmo

import (
	"fgame/fgame/core/storage"
	shenmopb "fgame/fgame/cross/shenmo/pb"
	"fgame/fgame/game/global"
	shenmoentity "fgame/fgame/game/shenmo/entity"
	"fgame/fgame/pkg/idutil"
)

//排行榜数据
type ShenMoRankData struct {
	serverId     int32
	allianceId   int64
	allianceName string
	jiFenNum     int32
}

func (t *ShenMoRankData) GetServerId() int32 {
	return t.serverId
}

func (t *ShenMoRankData) GetAllianceId() int64 {
	return t.allianceId
}

func (t *ShenMoRankData) GetAllianceName() string {
	return t.allianceName
}

func (t *ShenMoRankData) GetJiFenNum() int32 {
	return t.jiFenNum
}

//跨服数据转换
func convertFromRankInfo(rankInfo *shenmopb.ShenMoRankInfo) *ShenMoRankData {
	rankData := &ShenMoRankData{}
	rankData.serverId = rankInfo.ServerId
	rankData.allianceId = rankInfo.AllianceId
	rankData.allianceName = rankInfo.AllianceName
	rankData.jiFenNum = rankInfo.JiFenNum
	return rankData
}

//跨服数据转换
func convertFromRankInfoList(rankInfoList []*shenmopb.ShenMoRankInfo) (dataList []*ShenMoRankData) {
	dataList = make([]*ShenMoRankData, 0, len(rankInfoList))
	for _, rankInfo := range rankInfoList {
		dataList = append(dataList, convertFromRankInfo(rankInfo))
	}
	return dataList
}

//本服对象转换
func convertFromRankObjectList(rankObjList []*ShenMoRankObject, isThis bool) (dataList []*ShenMoRankData) {
	dataList = make([]*ShenMoRankData, 0, len(rankObjList))
	for _, rankObj := range rankObjList {
		dataList = append(dataList, convertFromRankObject(rankObj, isThis))
	}
	return dataList
}

//本服对象转换
func convertFromRankObject(rankObj *ShenMoRankObject, isThis bool) *ShenMoRankData {
	rankData := &ShenMoRankData{}
	rankData.serverId = rankObj.serverId
	rankData.allianceId = rankObj.allianceId
	rankData.allianceName = rankObj.allianceName
	if isThis {
		rankData.jiFenNum = rankObj.jiFenNum
	} else {
		rankData.jiFenNum = rankObj.lastJiFenNum
	}
	return rankData
}

//神魔战场排行榜数据
type ShenMoRankObject struct {
	id           int64
	serverId     int32
	allianceId   int64
	allianceName string
	jiFenNum     int32
	lastJiFenNum int32
	lastTime     int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

//本周记录排序
type ThisShenMoRankObjectList []*ShenMoRankObject

func (adl ThisShenMoRankObjectList) Len() int {
	return len(adl)
}

func (adl ThisShenMoRankObjectList) Less(i, j int) bool {
	if adl[i].jiFenNum == adl[j].jiFenNum {
		return adl[i].lastTime < adl[j].lastTime
	}
	return adl[i].jiFenNum < adl[j].jiFenNum
}

func (adl ThisShenMoRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//上周记录排序
type LastShenMoRankObjectList []*ShenMoRankObject

func (adl LastShenMoRankObjectList) Len() int {
	return len(adl)
}

func (adl LastShenMoRankObjectList) Less(i, j int) bool {
	return adl[i].lastJiFenNum < adl[j].lastJiFenNum
}

func (adl LastShenMoRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

func initShenMoRankObject(serverId int32,
	allianceId int64,
	allianceName string,
	jiFenNum int32) *ShenMoRankObject {

	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	o := &ShenMoRankObject{}
	o.id = id
	o.serverId = serverId
	o.allianceId = allianceId
	o.allianceName = allianceName
	o.jiFenNum = jiFenNum
	o.lastTime = now
	o.lastJiFenNum = 0
	o.createTime = now
	return o
}

func NewShenMoRankObject() *ShenMoRankObject {
	return &ShenMoRankObject{}
}

func (so *ShenMoRankObject) GetDBId() int64 {
	return so.id
}

func (oo *ShenMoRankObject) ToEntity() (e storage.Entity, err error) {
	oe := &shenmoentity.ShenMoRankEntity{}
	oe.Id = oo.id
	oe.ServerId = oo.serverId
	oe.AllianceId = oo.allianceId
	oe.AllianceName = oo.allianceName
	oe.JiFenNum = oo.jiFenNum
	oe.LastJiFenNum = oo.lastJiFenNum
	oe.LastTime = oo.lastTime
	oe.UpdateTime = oo.updateTime
	oe.CreateTime = oo.createTime
	oe.DeleteTime = oo.deleteTime
	e = oe
	return
}

func (oo *ShenMoRankObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*shenmoentity.ShenMoRankEntity)
	oo.id = oe.Id
	oo.serverId = oe.ServerId
	oo.allianceId = oe.AllianceId
	oo.allianceName = oe.AllianceName
	oo.jiFenNum = oe.JiFenNum
	oo.lastJiFenNum = oe.LastJiFenNum
	oo.lastTime = oe.LastTime
	oo.updateTime = oe.UpdateTime
	oo.createTime = oe.CreateTime
	oo.deleteTime = oe.DeleteTime
	return
}

func (oo *ShenMoRankObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

//神魔战场时间戳数据
type ShenMoRankTimeObject struct {
	id         int64
	serverId   int32
	thisTime   int64
	lastTime   int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewShenMoRankTimeObject() *ShenMoRankTimeObject {
	poo := &ShenMoRankTimeObject{}
	return poo
}

func (so *ShenMoRankTimeObject) GetDBId() int64 {
	return so.id
}

func (oo *ShenMoRankTimeObject) ToEntity() (e storage.Entity, err error) {
	oe := &shenmoentity.ShenMoRankTimeEntity{}
	oe.Id = oo.id
	oe.ServerId = oo.serverId
	oe.ThisTime = oo.thisTime
	oe.LastTime = oo.lastTime
	oe.UpdateTime = oo.updateTime
	oe.CreateTime = oo.createTime
	oe.DeleteTime = oo.deleteTime
	e = oe
	return
}

func (oo *ShenMoRankTimeObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*shenmoentity.ShenMoRankTimeEntity)
	oo.id = oe.Id
	oo.serverId = oe.ServerId
	oo.thisTime = oe.ThisTime
	oo.lastTime = oe.LastTime
	oo.updateTime = oe.UpdateTime
	oo.createTime = oe.CreateTime
	oo.deleteTime = oe.DeleteTime
	return
}

func (oo *ShenMoRankTimeObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
