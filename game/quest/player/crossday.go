package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/dao"
	questentity "fgame/fgame/game/quest/entity"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//任务模块跨天对象
type PlayerQuestCrossDayObject struct {
	player       player.Player
	id           int64
	crossDayTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerQuestCrossDayObject(pl player.Player) *PlayerQuestCrossDayObject {
	pmo := &PlayerQuestCrossDayObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerQuestCrossDayObjectToEntity(pqo *PlayerQuestCrossDayObject) (e *questentity.PlayerQuestCrossDayEntity, err error) {

	e = &questentity.PlayerQuestCrossDayEntity{
		Id:           pqo.id,
		CrossDayTime: pqo.crossDayTime,
		PlayerId:     pqo.player.GetId(),
		UpdateTime:   pqo.updateTime,
		CreateTime:   pqo.createTime,
		DeleteTime:   pqo.deleteTime,
	}
	return
}

func (pqo *PlayerQuestCrossDayObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerQuestCrossDayObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerQuestCrossDayObject) GetCrossDayTime() int64 {
	return pqo.crossDayTime
}

func (pqo *PlayerQuestCrossDayObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerQuestCrossDayObjectToEntity(pqo)
	return
}

func (pqo *PlayerQuestCrossDayObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerQuestCrossDayEntity)

	pqo.id = pqe.Id
	pqo.crossDayTime = pqe.CrossDayTime
	pqo.updateTime = pqe.UpdateTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerQuestCrossDayObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "quest_crossday"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}

//数据库加载
func (pqdm *PlayerQuestDataManager) loadQuestCrossDay() (err error) {
	//加载任务跨天信息
	questCrossDayEntity, err := dao.GetQuestDao().GetQuestCrossDay(pqdm.p.GetId())
	if err != nil {
		return
	}
	if questCrossDayEntity == nil {
		pqdm.initPlayerQuestCrossDayObject()
	} else {
		pqdm.playerCrossDayObject = NewPlayerQuestCrossDayObject(pqdm.p)
		pqdm.playerCrossDayObject.FromEntity(questCrossDayEntity)
	}
	return
}

//第一次初始化
func (pqdm *PlayerQuestDataManager) initPlayerQuestCrossDayObject() {
	ptmo := NewPlayerQuestCrossDayObject(pqdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ptmo.id = id
	//生成id
	ptmo.crossDayTime = now
	ptmo.createTime = now
	pqdm.playerCrossDayObject = ptmo
	ptmo.SetModified()
}

func (pqdm *PlayerQuestDataManager) GetQuestCrossDayTime() int64 {
	return pqdm.playerCrossDayObject.crossDayTime
}

func (pqdm *PlayerQuestDataManager) setCrossDayTime(now int64) {
	pqdm.playerCrossDayObject.crossDayTime = now
	pqdm.playerCrossDayObject.updateTime = now
	pqdm.playerCrossDayObject.SetModified()
}
