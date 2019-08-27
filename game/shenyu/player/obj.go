package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenyuentity "fgame/fgame/game/shenyu/entity"

	"github.com/pkg/errors"
)

//玩家神域对象
type PlayerShenYuObject struct {
	player     player.Player
	id         int64
	keyNum     int32
	round      int32
	exp        int64           //废弃
	itemMap    map[int32]int32 //废弃
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerShenYuObject(pl player.Player) *PlayerShenYuObject {
	pdo := &PlayerShenYuObject{
		player: pl,
	}
	return pdo
}

func (pdo *PlayerShenYuObject) GetPlayerId() int64 {
	return pdo.player.GetId()
}

func (pdo *PlayerShenYuObject) GetDBId() int64 {
	return pdo.id
}

func convertObjectToEntity(pdo *PlayerShenYuObject) (*shenyuentity.PlayerShenYuEntity, error) {

	itemsBytes, err := json.Marshal(pdo.itemMap)
	if err != nil {
		return nil, err
	}

	e := &shenyuentity.PlayerShenYuEntity{
		Id:         pdo.id,
		PlayerId:   pdo.player.GetId(),
		KeyNum:     pdo.keyNum,
		Round:      pdo.round,
		Exp:        pdo.exp,
		ItemInfo:   string(itemsBytes),
		EndTime:    pdo.endTime,
		UpdateTime: pdo.updateTime,
		CreateTime: pdo.createTime,
		DeleteTime: pdo.deleteTime,
	}
	return e, nil
}

func (pdo *PlayerShenYuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(pdo)
	return
}

func (pdo *PlayerShenYuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shenyuentity.PlayerShenYuEntity)
	itemInfoMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pse.ItemInfo), &itemInfoMap); err != nil {
		return err
	}

	pdo.id = pse.Id
	pdo.keyNum = pse.KeyNum
	pdo.round = pse.Round
	pdo.endTime = pse.EndTime
	pdo.exp = pse.Exp
	pdo.itemMap = itemInfoMap
	pdo.updateTime = pse.UpdateTime
	pdo.createTime = pse.CreateTime
	pdo.deleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerShenYuObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenYu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}
