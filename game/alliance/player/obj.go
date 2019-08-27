package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//玩家仙盟对象
type PlayerAllianceObject struct {
	player                   player.Player
	id                       int64
	allianceId               int64
	allianceName             string
	allianceLevel            int32
	donateMap                map[alliancetypes.AllianceJuanXianType]int32
	currentGongXian          int64
	lastJuanXuanTime         int64
	sceneRewardMap           map[int32]int32
	warPoint                 int32
	lastAllianceSceneEndTime int64
	yaoPai                   int32
	lastYaoPaiUpdateTime     int64
	convertTimes             int32
	lastConvertUpdateTime    int64
	reliveTime               int32
	lastReliveTime           int64
	depotPoint               int32
	lastMemberCallTime       int64
	lastYuXiMemberCallTime   int64
	totalWinTime             int32
	updateTime               int64
	createTime               int64
	deleteTime               int64
}

func newPlayerAllianceObject(pl player.Player) *PlayerAllianceObject {
	o := &PlayerAllianceObject{
		player: pl,
	}
	return o
}

func convertPlayerAllianceObjectToEntity(o *PlayerAllianceObject) (e *allianceentity.PlayerAllianceEntity, err error) {
	donateMap, err := json.Marshal(o.donateMap)
	if err != nil {
		return
	}
	sceneRewardMap, err := json.Marshal(o.sceneRewardMap)
	if err != nil {
		return
	}
	e = &allianceentity.PlayerAllianceEntity{
		Id:                       o.id,
		PlayerId:                 o.player.GetId(),
		AllianceId:               o.allianceId,
		AllianceName:             o.allianceName,
		AllianceLevel:            o.allianceLevel,
		DonateMap:                string(donateMap),
		LastJuanXuanTime:         o.lastJuanXuanTime,
		CurrentGongXian:          o.currentGongXian,
		SceneRewardMap:           string(sceneRewardMap),
		WarPoint:                 o.warPoint,
		LastAllianceSceneEndTime: o.lastAllianceSceneEndTime,
		YaoPai:                   o.yaoPai,
		LastYaoPaiUpdateTime:     o.lastYaoPaiUpdateTime,
		ConvertTimes:             o.convertTimes,
		LastConvertUpdateTime:    o.lastConvertUpdateTime,
		ReliveTime:               o.reliveTime,
		LastReliveTime:           o.lastReliveTime,
		DepotPoint:               o.depotPoint,
		LastMemberCallTime:       o.lastMemberCallTime,
		LastYuXiMemberCallTime:   o.lastYuXiMemberCallTime,
		TotalWinTime:             o.totalWinTime,
		UpdateTime:               o.updateTime,
		CreateTime:               o.createTime,
		DeleteTime:               o.deleteTime,
	}
	return e, nil
}

func (o *PlayerAllianceObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAllianceObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *PlayerAllianceObject) GetAllianceName() string {
	return o.allianceName
}

func (o *PlayerAllianceObject) GetDonateMap() map[alliancetypes.AllianceJuanXianType]int32 {
	return o.donateMap
}

func (o *PlayerAllianceObject) GetSceneRewardMap() map[int32]int32 {
	return o.sceneRewardMap
}

func (o *PlayerAllianceObject) GetCurrentGongXian() int64 {
	return o.currentGongXian
}

func (o *PlayerAllianceObject) GetYaoPai() int32 {
	return o.yaoPai
}

func (o *PlayerAllianceObject) GetDepotPoint() int32 {
	return o.depotPoint
}

func (o *PlayerAllianceObject) GetConvertTimes() int32 {
	return o.convertTimes
}

func (o *PlayerAllianceObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerAllianceObject) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *PlayerAllianceObject) GetLastReliveTime() int64 {
	return o.lastReliveTime
}

func (o *PlayerAllianceObject) GetLastMemberCallTime() int64 {
	return o.lastMemberCallTime
}

func (o *PlayerAllianceObject) GetTotalWinTime() int32 {
	return o.totalWinTime
}

func (o *PlayerAllianceObject) RestReliveTime() {
	o.reliveTime = 0
}

func (o *PlayerAllianceObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerAllianceObjectToEntity(o)
	return e, err
}

func (o *PlayerAllianceObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*allianceentity.PlayerAllianceEntity)

	o.id = te.Id
	o.allianceId = te.AllianceId
	donateMap := make(map[alliancetypes.AllianceJuanXianType]int32)
	err := json.Unmarshal([]byte(te.DonateMap), &donateMap)
	if err != nil {
		return err
	}
	o.donateMap = donateMap

	sceneRewardMap := make(map[int32]int32)
	err = json.Unmarshal([]byte(te.SceneRewardMap), &sceneRewardMap)
	if err != nil {
		return err
	}
	o.sceneRewardMap = sceneRewardMap
	o.warPoint = te.WarPoint
	o.currentGongXian = te.CurrentGongXian
	o.allianceName = te.AllianceName
	o.allianceLevel = te.AllianceLevel
	o.lastJuanXuanTime = te.LastJuanXuanTime
	o.lastAllianceSceneEndTime = te.LastAllianceSceneEndTime
	o.yaoPai = te.YaoPai
	o.lastYaoPaiUpdateTime = te.LastYaoPaiUpdateTime
	o.convertTimes = te.ConvertTimes
	o.lastConvertUpdateTime = te.LastConvertUpdateTime
	o.reliveTime = te.ReliveTime
	o.lastReliveTime = te.LastReliveTime
	o.depotPoint = te.DepotPoint
	o.lastMemberCallTime = te.LastMemberCallTime
	o.lastYuXiMemberCallTime = te.LastYuXiMemberCallTime
	o.totalWinTime = te.TotalWinTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerAllianceObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerAlliance"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家仙术对象
type PlayerAllianceSkillObject struct {
	player     player.Player
	id         int64
	skillType  alliancetypes.AllianceSkillType
	level      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerAllianceSkillObject(pl player.Player) *PlayerAllianceSkillObject {
	o := &PlayerAllianceSkillObject{
		player: pl,
	}
	return o
}

func convertPlayerAllianceSkillObjectToEntity(o *PlayerAllianceSkillObject) (e *allianceentity.PlayerAllianceSkillEntity, err error) {
	e = &allianceentity.PlayerAllianceSkillEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		SkillType:  int32(o.skillType),
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerAllianceSkillObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAllianceSkillObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerAllianceSkillObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerAllianceSkillObjectToEntity(o)
	return e, err
}

func (o *PlayerAllianceSkillObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*allianceentity.PlayerAllianceSkillEntity)

	o.id = te.Id
	o.skillType = alliancetypes.AllianceSkillType(te.SkillType)
	o.level = te.Level
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerAllianceSkillObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceSkill"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerAllianceSkillObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerAllianceSkillObject) GetSkillType() alliancetypes.AllianceSkillType {
	return o.skillType
}
