package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	babyentity "fgame/fgame/game/baby/entity"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//怀孕对象
type PlayerPregnantObject struct {
	player       player.Player
	id           int64
	tonicPro     int32
	chaoshengNum int32
	pregnantTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerPregnantObject(pl player.Player) *PlayerPregnantObject {
	o := &PlayerPregnantObject{
		player: pl,
	}
	return o
}

func convertNewPlayerPregnantObjectToEntity(o *PlayerPregnantObject) (*babyentity.PlayerPregnantEntity, error) {

	e := &babyentity.PlayerPregnantEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		ChaoShengNum: o.chaoshengNum,
		TonicPro:     o.tonicPro,
		PregnantTime: o.pregnantTime,
		UpdateTime:   o.updateTime,
		DeleteTime:   o.deleteTime,
		CreateTime:   o.createTime,
	}
	return e, nil
}

func (o *PlayerPregnantObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerPregnantObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerPregnantObject) GetPregnantTime() int64 {
	return o.pregnantTime
}

func (o *PlayerPregnantObject) GetChaoShengNum() int32 {
	return o.chaoshengNum
}

func (o *PlayerPregnantObject) GetTonicPro() int32 {
	return o.tonicPro
}

func (o *PlayerPregnantObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerPregnantObjectToEntity(o)
	return e, err
}

func (o *PlayerPregnantObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*babyentity.PlayerPregnantEntity)

	o.id = pse.Id
	o.tonicPro = pse.TonicPro
	o.chaoshengNum = pse.ChaoShengNum
	o.pregnantTime = pse.PregnantTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerPregnantObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BABY"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//宝宝对象
type PlayerBabyObject struct {
	player             player.Player
	id                 int64
	name               string
	sex                types.SexType
	quality            int32
	skillList          []*babytypes.TalentInfo
	learnLevel         int32
	learnExp           int32
	activateTimes      int32
	refreshTimes       int32
	lockTimes          int32
	attrBeiShu         int32
	refreshCostItemNum int32
	updateTime         int64
	createTime         int64
	deleteTime         int64
}

func NewPlayerBabyObject(pl player.Player) *PlayerBabyObject {
	o := &PlayerBabyObject{
		player: pl,
	}
	return o
}

func convertNewPlayerBabyObjectToEntity(o *PlayerBabyObject) (*babyentity.PlayerBabyEntity, error) {

	data, _ := json.Marshal(o.skillList)

	e := &babyentity.PlayerBabyEntity{
		Id:                 o.id,
		PlayerId:           o.player.GetId(),
		Name:               o.name,
		Sex:                int32(o.sex),
		Quality:            o.quality,
		SkillList:          string(data),
		LearnLevel:         o.learnLevel,
		LearnExp:           o.learnExp,
		ActivateTimes:      o.activateTimes,
		RefreshTimes:       o.refreshTimes,
		LockTimes:          o.lockTimes,
		AttrBeiShu:         o.attrBeiShu,
		RefreshCostItemNum: o.refreshCostItemNum,
		UpdateTime:         o.updateTime,
		DeleteTime:         o.deleteTime,
		CreateTime:         o.createTime,
	}
	return e, nil
}

func (o *PlayerBabyObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerBabyObject) GetSkillList() []*babytypes.TalentInfo {
	return o.skillList
}

func (o *PlayerBabyObject) GetTalentSkillIdList() (skillList []int32) {
	for _, talent := range o.skillList {
		skillList = append(skillList, talent.SkillId)
	}
	return
}

func (o *PlayerBabyObject) GetActivateTimes() int32 {
	return o.activateTimes
}

func (o *PlayerBabyObject) GetLockTimes() int32 {
	return o.lockTimes
}

func (o *PlayerBabyObject) GetRefreshTimes() int32 {
	return o.refreshTimes
}

func (o *PlayerBabyObject) GetRefreshCostNum() int32 {
	return o.refreshCostItemNum
}

func (o *PlayerBabyObject) GetDanBei() int32 {
	return o.attrBeiShu
}

func (o *PlayerBabyObject) GetLearnLevel() int32 {
	return o.learnLevel
}

func (o *PlayerBabyObject) GetLearnExp() int32 {
	return o.learnExp
}

func (o *PlayerBabyObject) GetQuality() int32 {
	return o.quality
}

func (o *PlayerBabyObject) GetBabyName() string {
	return o.name
}

func (o *PlayerBabyObject) GetBabySex() types.SexType {
	return o.sex
}

func (o *PlayerBabyObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBabyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerBabyObjectToEntity(o)
	return e, err
}

func (o *PlayerBabyObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*babyentity.PlayerBabyEntity)

	var skillList []*babytypes.TalentInfo
	err := json.Unmarshal([]byte(pse.SkillList), &skillList)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.name = pse.Name
	o.sex = types.SexType(pse.Sex)
	o.quality = pse.Quality
	o.skillList = skillList
	o.learnLevel = pse.LearnLevel
	o.learnExp = pse.LearnExp
	o.activateTimes = pse.ActivateTimes
	o.refreshTimes = pse.RefreshTimes
	o.lockTimes = pse.LockTimes
	o.attrBeiShu = pse.AttrBeiShu
	o.refreshCostItemNum = pse.RefreshCostItemNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerBabyObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BABY"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//全部宝宝战力对象
type PlayerBabyPowerObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerBabyPowerObject(pl player.Player) *PlayerBabyPowerObject {
	o := &PlayerBabyPowerObject{
		player: pl,
	}
	return o
}

func convertNewPlayerBabyPowerObjectToEntity(o *PlayerBabyPowerObject) (*babyentity.PlayerBabyPowerEntity, error) {

	e := &babyentity.PlayerBabyPowerEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Power:      o.power,
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *PlayerBabyPowerObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerBabyPowerObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBabyPowerObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerBabyPowerObjectToEntity(o)
	return e, err
}

func (o *PlayerBabyPowerObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*babyentity.PlayerBabyPowerEntity)

	o.id = pse.Id
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerBabyPowerObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BabyPower"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
