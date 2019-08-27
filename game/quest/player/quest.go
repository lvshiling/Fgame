package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questentity "fgame/fgame/game/quest/entity"
	questtypes "fgame/fgame/game/quest/types"

	"github.com/pkg/errors"
)

//任务对象
type PlayerQuestObject struct {
	player             player.Player
	Id                 int64
	QuestId            int32
	QuestDataMap       map[int32]int32 //任务数据
	CollectItemDataMap map[int32]int32 //任务收集的物品
	QuestState         questtypes.QuestState
	UpdateTime         int64
	CreateTime         int64
	DeleteTime         int64
}

func NewPlayerQuestObject(pl player.Player) *PlayerQuestObject {
	pmo := &PlayerQuestObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerQuestObjectToEntity(pqo *PlayerQuestObject) (e *questentity.PlayerQuestEntity, err error) {
	questDataBytes, err := json.Marshal(pqo.QuestDataMap)
	if err != nil {
		return
	}
	collectItemDataBytes, err := json.Marshal(pqo.CollectItemDataMap)
	if err != nil {
		return
	}
	e = &questentity.PlayerQuestEntity{
		Id:              pqo.Id,
		PlayerId:        pqo.player.GetId(),
		QuestId:         pqo.QuestId,
		QuestData:       string(questDataBytes),
		CollectItemData: string(collectItemDataBytes),
		QuestState:      int32(pqo.QuestState),
		UpdateTime:      pqo.UpdateTime,
		CreateTime:      pqo.CreateTime,
		DeleteTime:      pqo.DeleteTime,
	}
	return
}

func (pqo *PlayerQuestObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerQuestObject) GetDBId() int64 {
	return pqo.Id
}

func (pqo *PlayerQuestObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerQuestObjectToEntity(pqo)
	return
}

func (pqo *PlayerQuestObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerQuestEntity)
	questDataMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pqe.QuestData), &questDataMap); err != nil {
		return err
	}
	collectItemDataMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pqe.CollectItemData), &collectItemDataMap); err != nil {
		return err
	}
	pqo.Id = pqe.Id
	pqo.QuestId = pqe.QuestId
	pqo.QuestDataMap = questDataMap
	pqo.CollectItemDataMap = collectItemDataMap
	pqo.QuestState = questtypes.QuestState(pqe.QuestState)
	pqo.UpdateTime = pqe.UpdateTime
	pqo.CreateTime = pqe.CreateTime
	pqo.DeleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerQuestObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Quest"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}
