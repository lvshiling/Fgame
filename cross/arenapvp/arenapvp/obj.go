package arenapvp

import (
	"fgame/fgame/core/storage"
	arenapvpentity "fgame/fgame/cross/arenapvp/entity"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
)

//霸主数据
type ArenapvpBaZhuObject struct {
	Id             int64
	Platform       int32
	ServerId       int32
	PlayerPlatform int32
	PlayerServerId int32
	PlayerId       int64
	PlayerName     string
	RaceNumber     int32
	Role           int32
	Sex            int32
	WingId         int32
	WeaponId       int32
	FashionId      int32
	UpdateTime     int64
	CreateTime     int64
	DeleteTime     int64
}

func NewArenapvpBaZhuObject() *ArenapvpBaZhuObject {
	return &ArenapvpBaZhuObject{}
}

func CreateArenapvpBaZhuObjectWithPvpPlayerInfo(pvpPlayer *arenapvpdata.PvpPlayerInfo, platform int32, serverId int32, raceNumber int32) *ArenapvpBaZhuObject {
	baZhu := NewArenapvpBaZhuObject()
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	baZhu.Id = id
	baZhu.Platform = platform
	baZhu.ServerId = serverId
	baZhu.PlayerPlatform = pvpPlayer.Platform
	baZhu.PlayerServerId = pvpPlayer.ServerId
	baZhu.PlayerId = pvpPlayer.PlayerId
	baZhu.PlayerName = pvpPlayer.PlayerName
	baZhu.RaceNumber = raceNumber
	baZhu.Role = pvpPlayer.Role
	baZhu.Sex = pvpPlayer.Sex
	baZhu.WingId = pvpPlayer.WingId
	baZhu.WeaponId = pvpPlayer.WeaponId
	baZhu.FashionId = pvpPlayer.FashionId
	baZhu.CreateTime = now
	baZhu.SetModified()

	return baZhu
}

func (so *ArenapvpBaZhuObject) GetDBId() int64 {
	return so.Id
}

func (oo *ArenapvpBaZhuObject) ToEntity() (e storage.Entity, err error) {
	oe := &arenapvpentity.ArenapvpBaZhuEntity{}
	oe.Id = oo.Id
	oe.PlayerPlatform = oo.PlayerPlatform
	oe.PlayerServerId = oo.PlayerServerId
	oe.Platform = oo.Platform
	oe.ServerId = oo.ServerId
	oe.PlayerId = oo.PlayerId
	oe.Role = oo.Role
	oe.Sex = oo.Sex
	oe.WingId = oo.WingId
	oe.WeaponId = oo.WeaponId
	oe.FashionId = oo.FashionId
	oe.PlayerName = oo.PlayerName
	oe.RaceNumber = oo.RaceNumber
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *ArenapvpBaZhuObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*arenapvpentity.ArenapvpBaZhuEntity)
	oo.Id = oe.Id
	oo.Platform = oe.Platform
	oo.ServerId = oe.ServerId
	oo.PlayerPlatform = oe.PlayerPlatform
	oo.PlayerServerId = oe.PlayerServerId
	oo.PlayerId = oe.PlayerId
	oo.Sex = oe.Sex
	oo.Role = oe.Role
	oo.PlayerName = oe.PlayerName
	oo.RaceNumber = oe.RaceNumber
	oo.WingId = oe.WingId
	oo.WeaponId = oe.WeaponId
	oo.FashionId = oe.FashionId
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *ArenapvpBaZhuObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
