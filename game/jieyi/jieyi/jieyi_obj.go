package jieyi

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	jieyientity "fgame/fgame/game/jieyi/entity"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/merge/merge"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

type JieYiObject struct {
	id             int64
	serverId       int32
	originServerId int32
	name           string
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewJieYiObject() *JieYiObject {
	o := &JieYiObject{}
	return o
}

func convertJieYiObjectToEntity(o *JieYiObject) (*jieyientity.JieYiEntity, error) {
	e := &jieyientity.JieYiEntity{
		Id:             o.id,
		ServerId:       o.serverId,
		OriginServerId: o.originServerId,
		Name:           o.name,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *JieYiObject) GetDBId() int64 {
	return o.id
}

func (o *JieYiObject) GetId() int64 {
	return o.id
}

func (o *JieYiObject) GetName() string {
	return o.name
}

func (o *JieYiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertJieYiObjectToEntity(o)
	return e, err
}

func (o *JieYiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*jieyientity.JieYiEntity)
	o.id = pse.Id
	o.serverId = pse.ServerId
	o.originServerId = pse.OriginServerId
	o.name = pse.Name
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *JieYiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JieYi"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

// 升序
type sortMemberObjList []*JieYiMemberObject

func (a sortMemberObjList) Len() int           { return len(a) }
func (a sortMemberObjList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortMemberObjList) Less(i, j int) bool { return a[i].jieYiTime < a[j].jieYiTime }

// 结义
type JieYi struct {
	jieYiObject           *JieYiObject
	jieYiMemberObjectList []*JieYiMemberObject // 结义成员
}

// 获取结义成员数量
func (j *JieYi) getMemberNum() int32 {
	return int32(len(j.jieYiMemberObjectList))
}

// 更新结义排名
func (j *JieYi) updateJieYiRank() {
	sort.Sort(sortMemberObjList(j.jieYiMemberObjectList))
	for i, mem := range j.jieYiMemberObjectList {
		mem.rank = int32(i) + 1
	}
}

// 获取结义id
func (j *JieYi) getJieYiId() int64 {
	return j.jieYiObject.GetDBId()
}

// 获取结义成员
func (j *JieYi) getJieYiMemberIndexAndObj(memberId int64) (int, *JieYiMemberObject) {
	for index, obj := range j.jieYiMemberObjectList {
		if memberId == obj.playerId {
			return index, obj
		}
	}
	return 0, nil
}

// 获取结义成员
func (j *JieYi) addJieYiMember(mem *JieYiMemberObject) {
	j.jieYiMemberObjectList = append(j.jieYiMemberObjectList, mem)
}

// 移除成员
func (j *JieYi) removeJieYiMember(memberId int64) {
	index, obj := j.getJieYiMemberIndexAndObj(memberId)
	if obj == nil {
		return
	}
	j.jieYiMemberObjectList = append(j.jieYiMemberObjectList[:index], j.jieYiMemberObjectList[index+1:]...)
	j.updateJieYiRank()
}

func (j *JieYi) isFull() bool {
	temp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	if temp == nil {
		return true
	}

	maxNum := temp.MaxPeopleNum
	num := j.getMemberNum()
	if num >= maxNum {
		return true
	}
	return false
}

func (j *JieYi) getLaoDa() int64 {
	if len(j.jieYiMemberObjectList) == 0 {
		return 0
	}
	return j.jieYiMemberObjectList[0].playerId
}

// 获取结义id
func (j *JieYi) GetJieYiObject() *JieYiObject {
	return j.jieYiObject
}

// 获取结义id
func (j *JieYi) GetJieYiMemberList() []*JieYiMemberObject {
	return j.jieYiMemberObjectList
}

// 获取排行
func (j *JieYi) GetJieYiMemberRank(memberId int64) int32 {
	for index, obj := range j.jieYiMemberObjectList {
		if memberId == obj.playerId {
			return int32(index + 1)
		}
	}
	return int32(0)
}

// 获取排行
func (j *JieYi) GetJieYiName() string {
	if merge.GetMergeService().GetMergeTime() > 0 {
		return fmt.Sprintf("s%d.%s", j.jieYiObject.originServerId, j.jieYiObject.GetName())
	}
	return j.jieYiObject.GetName()
}
