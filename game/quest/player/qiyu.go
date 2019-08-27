package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/dao"
	questentity "fgame/fgame/game/quest/entity"
	questtemplate "fgame/fgame/game/quest/template"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//奇遇对象
type PlayerQiYuObject struct {
	player      player.Player
	id          int64
	qiyuId      int32
	level       int32
	zhuan       int32
	fei         int32
	isFinish    int32
	isReceive   int32
	isHadNotice int32
	endTime     int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func NewPlayerQiYuObject(pl player.Player) *PlayerQiYuObject {
	pmo := &PlayerQiYuObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerQiYuObjectToEntity(o *PlayerQiYuObject) (e *questentity.PlayerQiYuEntity, err error) {

	e = &questentity.PlayerQiYuEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		IsFinish:    o.isFinish,
		IsReceive:   o.isReceive,
		IsHadNotice: o.isHadNotice,
		QiYuId:      o.qiyuId,
		Level:       o.level,
		Zhuan:       o.zhuan,
		Fei:         o.fei,
		EndTime:     o.endTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return
}

func (o *PlayerQiYuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerQiYuObject) GetQiYuId() int32 {
	return o.qiyuId
}

func (o *PlayerQiYuObject) GetIsFinish() int32 {
	return o.isFinish
}

func (o *PlayerQiYuObject) GetIsReceive() int32 {
	return o.isReceive
}

func (o *PlayerQiYuObject) GetIsHadNotice() int32 {
	return o.isHadNotice
}

func (o *PlayerQiYuObject) GetEndTime() int64 {
	return o.endTime
}

func (o *PlayerQiYuObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerQiYuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerQiYuObjectToEntity(o)
	return
}

func (o *PlayerQiYuObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerQiYuEntity)

	o.id = pqe.Id
	o.qiyuId = pqe.QiYuId
	o.isFinish = pqe.IsFinish
	o.isReceive = pqe.IsReceive
	o.isHadNotice = pqe.IsHadNotice
	o.level = pqe.Level
	o.fei = pqe.Fei
	o.zhuan = pqe.Zhuan
	o.endTime = pqe.EndTime
	o.createTime = pqe.CreateTime
	o.deleteTime = pqe.DeleteTime
	return nil
}

func (o *PlayerQiYuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "QiYu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//数据库加载
func (m *PlayerQuestDataManager) loadQiYu() (err error) {
	m.playerQiYuMap = make(map[int32]*PlayerQiYuObject)
	//加载奇遇信息
	qiyuEntityList, err := dao.GetQuestDao().GetQiYuList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range qiyuEntityList {
		obj := NewPlayerQiYuObject(m.p)
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		m.playerQiYuMap[obj.qiyuId] = obj
	}
	return
}

func (m *PlayerQuestDataManager) AddQiYuQuest(qiyuId int32, level, zhuan, fei int32) (qiyu *PlayerQiYuObject, flag bool) {
	obj, ok := m.playerQiYuMap[qiyuId]
	if !ok {
		obj = m.initQiYu(qiyuId, level, zhuan, fei)
	} else {
		if obj.level == level && obj.zhuan == zhuan && obj.fei == fei {
			flag = false
			return
		}

		now := global.GetGame().GetTimeService().Now()
		qiyuTemplate := questtemplate.GetQuestTemplateService().GetQiYuTemplate(qiyuId)
		obj.level = level
		obj.zhuan = zhuan
		obj.fei = fei
		obj.isHadNotice = 0
		obj.endTime = qiyuTemplate.GetEndTime(now)
		obj.updateTime = now
		obj.SetModified()
	}

	qiyu = obj
	flag = true
	return
}

func (m *PlayerQuestDataManager) initQiYu(qiyuId, level, zhuan, fei int32) *PlayerQiYuObject {
	now := global.GetGame().GetTimeService().Now()
	qiyuTemplate := questtemplate.GetQuestTemplateService().GetQiYuTemplate(qiyuId)

	obj := NewPlayerQiYuObject(m.p)
	id, _ := idutil.GetId()
	obj.id = id
	obj.qiyuId = qiyuId
	obj.isFinish = 0
	obj.level = level
	obj.zhuan = zhuan
	obj.fei = fei
	obj.endTime = qiyuTemplate.GetEndTime(now)
	obj.createTime = now
	obj.SetModified()
	m.playerQiYuMap[qiyuId] = obj

	return obj
}

func (m *PlayerQuestDataManager) GetQiYuMap() map[int32]*PlayerQiYuObject {
	return m.playerQiYuMap
}

func (m *PlayerQuestDataManager) GetQiYu(qiyuId int32) *PlayerQiYuObject {
	obj, ok := m.playerQiYuMap[qiyuId]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerQuestDataManager) IsFinishQiYu(qiyuId int32) (flag bool) {
	obj := m.GetQiYu(qiyuId)
	if obj == nil {
		return
	}

	if obj.isFinish == 0 {
		return
	}

	flag = true
	return
}

func (m *PlayerQuestDataManager) IsReceiveQiYu(qiyuId int32) (flag bool) {
	obj := m.GetQiYu(qiyuId)
	if obj == nil {
		return
	}

	if obj.isReceive == 0 {
		return
	}

	flag = true
	return
}

func (m *PlayerQuestDataManager) QiYuFinish(qiyuId int32) {
	obj := m.GetQiYu(qiyuId)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if now > obj.endTime {
		return
	}
	obj.isFinish = 1
	obj.updateTime = now
	obj.SetModified()

	return
}

func (m *PlayerQuestDataManager) ReceiveQiYu(qiyuId int32) {
	obj := m.GetQiYu(qiyuId)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.isReceive = 1
	obj.updateTime = now
	obj.SetModified()

	return
}

func (m *PlayerQuestDataManager) NoticeQiYu(qiyuId int32) {
	obj := m.GetQiYu(qiyuId)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.isHadNotice = 1
	obj.updateTime = now
	obj.SetModified()
	return
}
