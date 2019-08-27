package marry

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	marrytypes "fgame/fgame/game/marry/types"
)

//婚烟数据
type MarryObject struct {
	Id                 int64
	ServerId           int32
	PlayerId           int64
	SpouseId           int64
	PlayerName         string
	SpouseName         string
	PlayerRingLevel    int32
	SpouseRingLevel    int32
	Role               int32
	SpouseRole         int32
	Sex                int32
	SpouseSex          int32
	Point              int32
	Ring               marrytypes.MarryRingType
	Status             marrytypes.MarryStatusType
	DevelopLevel       int32
	SpouseDevelopLevel int32
	PlayerSuit         map[int32]map[int32]int32
	SpouseSuit         map[int32]map[int32]int32
	UpdateTime         int64
	CreateTime         int64
	DeleteTime         int64
}

func NewMarryObject() *MarryObject {
	pso := &MarryObject{}
	return pso
}

func (mo *MarryObject) GetDBId() int64 {
	return mo.Id
}

//是否举办过婚礼
func (mo *MarryObject) HasHunLi() bool {
	if mo.Status == marrytypes.MarryStatusTypeEngagement || mo.Status == marrytypes.MarryStatusTypeMarried {
		return true
	}
	return false
}

func (mo *MarryObject) ToEntity() (e storage.Entity, err error) {
	pe := &marryentity.MarryEntity{}
	pe.Id = mo.Id
	pe.ServerId = mo.ServerId
	pe.PlayerId = mo.PlayerId
	pe.SpouseId = mo.SpouseId
	pe.PlayerName = mo.PlayerName
	pe.SpouseName = mo.SpouseName
	pe.PlayerRingLevel = mo.PlayerRingLevel
	pe.SpouseRingLevel = mo.SpouseRingLevel
	pe.Role = mo.Role
	pe.SpouseRole = mo.SpouseRole
	pe.Sex = mo.Sex
	pe.SpouseSex = mo.SpouseSex
	pe.Point = mo.Point
	pe.Ring = int32(mo.Ring)
	pe.Status = int32(mo.Status)
	pe.DevelopLevel = mo.DevelopLevel
	pe.SpouseDevelopLevel = mo.SpouseDevelopLevel
	pe.UpdateTime = mo.UpdateTime
	pe.CreateTime = mo.CreateTime
	pe.DeleteTime = mo.DeleteTime
	playerSuit, _ := json.Marshal(mo.PlayerSuit)
	pe.PlayerSuit = string(playerSuit)
	spouseSuitm, _ := json.Marshal(mo.SpouseSuit)
	pe.SpouseSuit = string(spouseSuitm)
	e = pe
	return
}

func (mo *MarryObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*marryentity.MarryEntity)
	mo.Id = pe.Id
	mo.ServerId = pe.ServerId
	mo.PlayerId = pe.PlayerId
	mo.SpouseId = pe.SpouseId
	mo.PlayerName = pe.PlayerName
	mo.SpouseName = pe.SpouseName
	mo.PlayerRingLevel = pe.PlayerRingLevel
	mo.SpouseRingLevel = pe.SpouseRingLevel
	mo.Role = pe.Role
	mo.SpouseRole = pe.SpouseRole
	mo.Sex = pe.Sex
	mo.SpouseSex = pe.SpouseSex
	mo.Ring = marrytypes.MarryRingType(pe.Ring)
	mo.Status = marrytypes.MarryStatusType(pe.Status)
	mo.DevelopLevel = pe.DevelopLevel
	mo.SpouseDevelopLevel = pe.SpouseDevelopLevel
	mo.UpdateTime = pe.UpdateTime
	mo.CreateTime = pe.CreateTime
	mo.DeleteTime = pe.DeleteTime
	mo.PlayerSuit = make(map[int32]map[int32]int32)
	err = json.Unmarshal([]byte(pe.PlayerSuit), &mo.PlayerSuit)
	if err != nil {
		return
	}
	mo.SpouseSuit = make(map[int32]map[int32]int32)
	err = json.Unmarshal([]byte(pe.SpouseSuit), &mo.SpouseSuit)
	if err != nil {
		return
	}
	return
}

func (mo *MarryObject) SetModified() {
	e, err := mo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
