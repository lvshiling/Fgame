package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	marryentity "fgame/fgame/game/marry/entity"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//结婚对象
type PlayerMarryObject struct {
	player             player.Player
	Id                 int64
	PlayerId           int64
	SpouseId           int64
	SpouseName         string
	Status             marrytypes.MarryStatusType
	Ring               marrytypes.MarryRingType
	RingLevel          int32
	RingNum            int32
	RingExp            int32
	TreeLevel          int32
	TreeNum            int32
	TreeExp            int32
	IsProposal         int32
	WedStatus          marrytypes.MarryWedStatusSelfType
	developLevel       int32
	developExp         int32
	coupleDevelopLevel int32
	UpdateTime         int64
	CreateTime         int64
	DeleteTime         int64
	MarryCount         int32
}

func NewPlayerMarryObject(pl player.Player) *PlayerMarryObject {
	pso := &PlayerMarryObject{
		player: pl,
	}
	return pso
}

func (pmo *PlayerMarryObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerMarryObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerMarryObject) GetDevelopLevel() int32 {
	return pmo.developLevel
}

func (pmo *PlayerMarryObject) GetDevelopExp() int32 {
	return pmo.developExp
}

func (pmo *PlayerMarryObject) HasHunLi() bool {
	if pmo.Status == marrytypes.MarryStatusTypeEngagement || pmo.Status == marrytypes.MarryStatusTypeMarried {
		return true
	}
	return false
}

func (pmo *PlayerMarryObject) ToEntity() (e storage.Entity, err error) {
	e = &marryentity.PlayerMarryEntity{
		Id:                 pmo.Id,
		PlayerId:           pmo.PlayerId,
		SpouseId:           pmo.SpouseId,
		SpouseName:         pmo.SpouseName,
		Status:             int32(pmo.Status),
		Ring:               int32(pmo.Ring),
		RingLevel:          pmo.RingLevel,
		RingNum:            pmo.RingNum,
		RingExp:            pmo.RingExp,
		TreeLevel:          pmo.TreeLevel,
		TreeNum:            pmo.TreeNum,
		TreeExp:            pmo.TreeExp,
		IsProposal:         pmo.IsProposal,
		WedStatus:          int32(pmo.WedStatus),
		DevelopExp:         pmo.developExp,
		DevelopLevel:       pmo.developLevel,
		CoupleDevelopLevel: pmo.coupleDevelopLevel,
		UpdateTime:         pmo.UpdateTime,
		CreateTime:         pmo.CreateTime,
		DeleteTime:         pmo.DeleteTime,
		MarryCount:         pmo.MarryCount,
	}
	return e, err
}

func (pmo *PlayerMarryObject) FromEntity(e storage.Entity) error {
	pme, _ := e.(*marryentity.PlayerMarryEntity)
	pmo.Id = pme.Id
	pmo.PlayerId = pme.PlayerId
	pmo.SpouseId = pme.SpouseId
	pmo.SpouseName = pme.SpouseName
	pmo.Status = marrytypes.MarryStatusType(pme.Status)
	pmo.Ring = marrytypes.MarryRingType(pme.Ring)
	pmo.RingLevel = pme.RingLevel
	pmo.RingNum = pme.RingNum
	pmo.RingExp = pme.RingExp
	pmo.TreeLevel = pme.TreeLevel
	pmo.TreeNum = pme.TreeNum
	pmo.TreeExp = pme.TreeExp
	pmo.IsProposal = pme.IsProposal
	pmo.WedStatus = marrytypes.MarryWedStatusSelfType(pme.WedStatus)
	pmo.developExp = pme.DevelopExp
	pmo.developLevel = pme.DevelopLevel
	pmo.coupleDevelopLevel = pme.CoupleDevelopLevel
	pmo.UpdateTime = pme.UpdateTime
	pmo.CreateTime = pme.CreateTime
	pmo.DeleteTime = pme.DeleteTime
	pmo.MarryCount = pme.MarryCount
	return nil
}

func (pmo *PlayerMarryObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Marry"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//玩家查看过喜帖对象
type PlayerViewWedCardObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	CardId     int64
	ViewTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerViewWedCardObject(pl player.Player) *PlayerViewWedCardObject {
	pso := &PlayerViewWedCardObject{
		player: pl,
	}
	return pso
}

func (pvwco *PlayerViewWedCardObject) GetPlayerId() int64 {
	return pvwco.PlayerId
}

func (pvwco *PlayerViewWedCardObject) GetDBId() int64 {
	return pvwco.Id
}

func (pvwco *PlayerViewWedCardObject) ToEntity() (e storage.Entity, err error) {
	e = &marryentity.PlayerViewWedCardEntity{
		Id:         pvwco.Id,
		PlayerId:   pvwco.PlayerId,
		CardId:     pvwco.CardId,
		ViewTime:   pvwco.ViewTime,
		UpdateTime: pvwco.UpdateTime,
		CreateTime: pvwco.CreateTime,
		DeleteTime: pvwco.DeleteTime,
	}
	return e, err
}

func (pvwco *PlayerViewWedCardObject) FromEntity(e storage.Entity) error {
	pvwe, _ := e.(*marryentity.PlayerViewWedCardEntity)
	pvwco.Id = pvwe.Id
	pvwco.PlayerId = pvwe.PlayerId
	pvwco.CardId = pvwe.CardId
	pvwco.ViewTime = pvwe.ViewTime
	pvwco.UpdateTime = pvwe.UpdateTime
	pvwco.CreateTime = pvwe.CreateTime
	pvwco.DeleteTime = pvwe.DeleteTime
	return nil
}

func (pvwco *PlayerViewWedCardObject) SetModified() {
	e, err := pvwco.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ViewWedCard"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pvwco.player.AddChangedObject(obj)
	return
}

//玩家豪气值对象
type PlayerMarryHeroismObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Heroism    int32
	OutOfTime  int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMarryHeroismObject(pl player.Player) *PlayerMarryHeroismObject {
	pso := &PlayerMarryHeroismObject{
		player: pl,
	}
	return pso
}

func (pmho *PlayerMarryHeroismObject) GetPlayerId() int64 {
	return pmho.PlayerId
}

func (pmho *PlayerMarryHeroismObject) GetDBId() int64 {
	return pmho.Id
}

func (pmho *PlayerMarryHeroismObject) ToEntity() (e storage.Entity, err error) {
	e = &marryentity.PlayerMarryHeroismEntity{
		Id:         pmho.Id,
		PlayerId:   pmho.PlayerId,
		Heroism:    pmho.Heroism,
		OutOfTime:  pmho.OutOfTime,
		UpdateTime: pmho.UpdateTime,
		CreateTime: pmho.CreateTime,
		DeleteTime: pmho.DeleteTime,
	}
	return e, err
}

func (pmho *PlayerMarryHeroismObject) FromEntity(e storage.Entity) error {
	pmhe, _ := e.(*marryentity.PlayerMarryHeroismEntity)
	pmho.Id = pmhe.Id
	pmho.PlayerId = pmhe.PlayerId
	pmho.Heroism = pmhe.Heroism
	pmho.OutOfTime = pmhe.OutOfTime
	pmho.UpdateTime = pmhe.UpdateTime
	pmho.CreateTime = pmhe.CreateTime
	pmho.DeleteTime = pmhe.DeleteTime
	return nil
}

func (pmho *PlayerMarryHeroismObject) SetModified() {
	e, err := pmho.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "MarryHeroism"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmho.player.AddChangedObject(obj)
	return
}

//玩家推送婚礼按钮记录
type PlayerPushWedRecordObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	WedId       int64
	HunCheTime  int64
	BanquetTime int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerPushWedRecordObject(pl player.Player) *PlayerPushWedRecordObject {
	pso := &PlayerPushWedRecordObject{
		player: pl,
	}
	return pso
}

func (pmho *PlayerPushWedRecordObject) GetPlayerId() int64 {
	return pmho.PlayerId
}

func (pmho *PlayerPushWedRecordObject) GetDBId() int64 {
	return pmho.Id
}

func (pmho *PlayerPushWedRecordObject) ToEntity() (e storage.Entity, err error) {
	e = &marryentity.PlayerPushWedRecordEntity{
		Id:          pmho.Id,
		PlayerId:    pmho.PlayerId,
		WedId:       pmho.WedId,
		HunCheTime:  pmho.HunCheTime,
		BanquetTime: pmho.BanquetTime,
		UpdateTime:  pmho.UpdateTime,
		CreateTime:  pmho.CreateTime,
		DeleteTime:  pmho.DeleteTime,
	}
	return e, err
}

func (pmho *PlayerPushWedRecordObject) FromEntity(e storage.Entity) error {
	pmhe, _ := e.(*marryentity.PlayerPushWedRecordEntity)
	pmho.Id = pmhe.Id
	pmho.PlayerId = pmhe.PlayerId
	pmho.WedId = pmhe.WedId
	pmho.HunCheTime = pmhe.HunCheTime
	pmho.BanquetTime = pmhe.BanquetTime
	pmho.UpdateTime = pmhe.UpdateTime
	pmho.CreateTime = pmhe.CreateTime
	pmho.DeleteTime = pmhe.DeleteTime
	return nil
}

func (pmho *PlayerPushWedRecordObject) SetModified() {
	e, err := pmho.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "player_pushwed_record"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmho.player.AddChangedObject(obj)
	return
}

//纪念次数
type PlayerMarryJiNianObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	JiNianType  marrytypes.MarryBanquetSubTypeWed
	JiNianCount int32
	SendFlag    int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerMarryJiNianObject(pl player.Player) *PlayerMarryJiNianObject {
	pso := &PlayerMarryJiNianObject{
		player: pl,
	}
	return pso
}

func (pmho *PlayerMarryJiNianObject) GetPlayerId() int64 {
	return pmho.PlayerId
}

func (pmho *PlayerMarryJiNianObject) GetDBId() int64 {
	return pmho.Id
}

func (pmho *PlayerMarryJiNianObject) ToEntity() (e storage.Entity, err error) {
	e = &marryentity.PlayerMarryJiNianEntity{
		Id:          pmho.Id,
		PlayerId:    pmho.PlayerId,
		JiNianType:  int32(pmho.JiNianType),
		JiNianCount: pmho.JiNianCount,
		UpdateTime:  pmho.UpdateTime,
		CreateTime:  pmho.CreateTime,
		DeleteTime:  pmho.DeleteTime,
		SendFlag:    pmho.SendFlag,
	}
	return e, err
}

func (pmho *PlayerMarryJiNianObject) FromEntity(e storage.Entity) error {
	pmhe, _ := e.(*marryentity.PlayerMarryJiNianEntity)
	pmho.Id = pmhe.Id
	pmho.PlayerId = pmhe.PlayerId
	pmho.JiNianType = marrytypes.MarryBanquetSubTypeWed(pmhe.JiNianType)
	pmho.JiNianCount = pmhe.JiNianCount
	pmho.UpdateTime = pmhe.UpdateTime
	pmho.CreateTime = pmhe.CreateTime
	pmho.DeleteTime = pmhe.DeleteTime
	pmho.SendFlag = pmhe.SendFlag
	return nil
}

func (pmho *PlayerMarryJiNianObject) SetModified() {
	e, err := pmho.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "t_player_marry_jinian"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmho.player.AddChangedObject(obj)
	return
}

//纪念次数时装获取次数
type PlayerMarryJiNianSjObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SjGetFlag  bool
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMarryJiNianSjObject(pl player.Player) *PlayerMarryJiNianSjObject {
	pso := &PlayerMarryJiNianSjObject{
		player: pl,
	}
	return pso
}

func (pmho *PlayerMarryJiNianSjObject) GetPlayerId() int64 {
	return pmho.PlayerId
}

func (pmho *PlayerMarryJiNianSjObject) GetDBId() int64 {
	return pmho.Id
}

func (pmho *PlayerMarryJiNianSjObject) ToEntity() (e storage.Entity, err error) {
	rst := &marryentity.PlayerMarryJiNianSjEntity{
		Id:         pmho.Id,
		PlayerId:   pmho.PlayerId,
		UpdateTime: pmho.UpdateTime,
		CreateTime: pmho.CreateTime,
		DeleteTime: pmho.DeleteTime,
	}
	if pmho.SjGetFlag {
		rst.SjGetFlag = 1
	}
	e = rst
	return e, err
}

func (pmho *PlayerMarryJiNianSjObject) FromEntity(e storage.Entity) error {
	pmhe, _ := e.(*marryentity.PlayerMarryJiNianSjEntity)
	pmho.Id = pmhe.Id
	pmho.PlayerId = pmhe.PlayerId
	pmho.UpdateTime = pmhe.UpdateTime
	pmho.CreateTime = pmhe.CreateTime
	pmho.DeleteTime = pmhe.DeleteTime
	if pmhe.SjGetFlag > 0 {
		pmho.SjGetFlag = true
	}
	return nil
}

func (pmho *PlayerMarryJiNianSjObject) SetModified() {
	e, err := pmho.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "t_player_marry_jinian_sj"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmho.player.AddChangedObject(obj)
	return
}

type PlayerMarryDingQingObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SuitMap    map[int32]map[int32]int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMarryDingQingObject(pl player.Player) *PlayerMarryDingQingObject {
	pso := &PlayerMarryDingQingObject{
		player: pl,
	}
	return pso
}

func (dq *PlayerMarryDingQingObject) GetPlayerId() int64 {
	return dq.PlayerId
}

func (dq *PlayerMarryDingQingObject) GetDBId() int64 {
	return dq.Id
}

func (dq *PlayerMarryDingQingObject) ToEntity() (e storage.Entity, err error) {
	msg, err := json.Marshal(dq.SuitMap)
	if err != nil {
		return
	}
	e = &marryentity.PlayerMarryDingQingEntity{
		Id:         dq.Id,
		PlayerId:   dq.PlayerId,
		Suit:       string(msg),
		UpdateTime: dq.UpdateTime,
		CreateTime: dq.CreateTime,
		DeleteTime: dq.DeleteTime,
	}
	return e, err
}

func (dq *PlayerMarryDingQingObject) FromEntity(e storage.Entity) error {
	suitMap := make(map[int32]map[int32]int32)

	pmhe, _ := e.(*marryentity.PlayerMarryDingQingEntity)
	if len(pmhe.Suit) > 0 {
		err := json.Unmarshal([]byte(pmhe.Suit), &suitMap)
		if err != nil {
			return err
		}
	}
	dq.Id = pmhe.Id
	dq.PlayerId = pmhe.PlayerId
	dq.UpdateTime = pmhe.UpdateTime
	dq.CreateTime = pmhe.CreateTime
	dq.DeleteTime = pmhe.DeleteTime
	dq.SuitMap = suitMap
	return nil
}

func (dq *PlayerMarryDingQingObject) SetModified() {
	e, err := dq.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "t_player_marry_dingqing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	dq.player.AddChangedObject(obj)
	return
}
