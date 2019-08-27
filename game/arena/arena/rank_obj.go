package arena

import (
	"fgame/fgame/core/storage"
	arenaentity "fgame/fgame/game/arena/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
)

//3v3排行榜数据
type ArenaRankObject struct {
	Id           int64
	ServerId     int32
	PlayerId     int64
	PlayerName   string
	WinCount     int32
	CurWinCount  int32
	LastWinCount int32
	LastTime     int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

//本周记录排序
type ThisArenaRankObjectList []*ArenaRankObject

func (adl ThisArenaRankObjectList) Len() int {
	return len(adl)
}

func (adl ThisArenaRankObjectList) Less(i, j int) bool {
	if adl[i].WinCount == adl[j].WinCount {
		return adl[i].LastTime < adl[j].LastTime
	}
	return adl[i].WinCount < adl[j].WinCount
}

func (adl ThisArenaRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//上周记录排序
type LastArenaRankObjectList []*ArenaRankObject

func (adl LastArenaRankObjectList) Len() int {
	return len(adl)
}

func (adl LastArenaRankObjectList) Less(i, j int) bool {
	return adl[i].LastWinCount < adl[j].LastWinCount
}

func (adl LastArenaRankObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

func initArenaRankObject(serverId int32, playerId int64, playerName string) *ArenaRankObject {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	o := &ArenaRankObject{}
	o.Id = id
	o.ServerId = serverId
	o.PlayerId = playerId
	o.PlayerName = playerName
	o.WinCount = 1
	o.CurWinCount = 1
	o.LastTime = now
	o.LastWinCount = 0
	o.CreateTime = now
	o.SetModified()
	return o
}

func NewArenaRankObject() *ArenaRankObject {
	return &ArenaRankObject{}
}

func (so *ArenaRankObject) GetDBId() int64 {
	return so.Id
}

func (so *ArenaRankObject) GetServerId() int32 {
	return so.ServerId
}

func (so *ArenaRankObject) GetPlayerId() int64 {
	return so.PlayerId
}

func (so *ArenaRankObject) GetPlayerName() string {
	return so.PlayerName
}

func (so *ArenaRankObject) GetWinCount() int32 {
	return so.WinCount
}

func (so *ArenaRankObject) GetLastWinCount() int32 {
	return so.LastWinCount
}

func (oo *ArenaRankObject) ToEntity() (e storage.Entity, err error) {
	oe := &arenaentity.ArenaRankEntity{}
	oe.Id = oo.Id
	oe.ServerId = oo.ServerId
	oe.PlayerId = oo.PlayerId
	oe.PlayerName = oo.PlayerName
	oe.WinCount = oo.WinCount
	oe.CurWinCount = oo.CurWinCount
	oe.LastWinCount = oo.LastWinCount
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *ArenaRankObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*arenaentity.ArenaRankEntity)
	oo.Id = oe.Id
	oo.ServerId = oe.ServerId
	oo.PlayerId = oe.PlayerId
	oo.PlayerName = oe.PlayerName
	oo.WinCount = oe.WinCount
	oo.CurWinCount = oe.CurWinCount
	oo.LastWinCount = oe.LastWinCount
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *ArenaRankObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
