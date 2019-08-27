package npc

import (
	coretemplate "fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/npc/npc"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
)

type HunCheObject struct {
	Period      int32
	OwerId      int64
	PlayerId    int64
	SpouseId    int64
	HunCheGrade marrytypes.MarryBanquetSubTypeHunChe
	SugarGrade  marrytypes.MarryBanquetSubTypeSugar
	MoveId      int32
	LastTime    int64
}

func (hco *HunCheObject) GetPeriod() int32 {
	return hco.Period
}

func (hco *HunCheObject) GetOwerId() int64 {
	return hco.OwerId
}

func (hco *HunCheObject) GetPlayerId() int64 {
	return hco.PlayerId
}

func (hco *HunCheObject) GetSpouseId() int64 {
	return hco.SpouseId
}

func (hco *HunCheObject) GetMoveId() int32 {
	return hco.MoveId
}

func CreateHunCheObject() *HunCheObject {
	o := &HunCheObject{}
	return o
}

func (hco *HunCheObject) IsReachGoal(curPos coretypes.Position) bool {
	tem := marrytemplate.GetMarryTemplateService().GetMarryMoveTeamplate(hco.GetMoveId())
	nextTem := tem.GetNextTemp()
	if nextTem == nil {
		return true
	}
	distance := utils.DistanceSquare(curPos, nextTem.GetPos())
	if distance > common.MIN_DISTANCE_SQUARE_ERROR {
		return false
	}
	return true
}

type HunCheNPC struct {
	*npc.NPCBase
	hunCheObject *HunCheObject
}

func (n *HunCheNPC) GetHunCheObject() *HunCheObject {
	return n.hunCheObject
}

//心跳
func (n *HunCheNPC) Heartbeat() {
	now := global.GetGame().GetTimeService().Now()
	banquetTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeHunChe, n.hunCheObject.HunCheGrade)
	sugarEach := int64(banquetTemplate.SugarEach)
	if now-n.hunCheObject.LastTime >= sugarEach {
		n.hunCheObject.LastTime = now
		//发送事件
		gameevent.Emit(marryeventtypes.EventTypeMarrySugarTime, nil, nil)
	}

	n.NPCBase.Heartbeat()
	return
}

func (n *HunCheNPC) ReachGoal() bool {
	moveId := n.hunCheObject.MoveId
	tem := marrytemplate.GetMarryTemplateService().GetMarryMoveTeamplate(moveId)
	nextMoveTemplate := tem.GetNextTemp()
	if nextMoveTemplate == nil {
		return false
	}
	n.hunCheObject.MoveId = int32(nextMoveTemplate.TemplateId())
	if nextMoveTemplate.GetNextTemp() == nil {
		//婚车完成
		gameevent.Emit(marryeventtypes.EventTypeMarryHunCheEnd, n, nil)
		n.GetScene().RemoveSceneObject(n, false)
	}
	return true
}

func CreateHunCheNPC(periodId int32, ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, playerId int64, spouseId int64, hunCheGrade int32, sugarGrade int32) *HunCheNPC {
	now := global.GetGame().GetTimeService().Now()
	idInScene := int32(0)
	id, _ := idutil.GetId()
	carGrade := marrytypes.MarryBanquetSubTypeHunChe(hunCheGrade)
	hunCheTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeHunChe, carGrade)
	to := coretemplate.GetTemplateService().Get(int(hunCheTemplate.BanquetBiology), (*template.BiologyTemplate)(nil))
	weddingCarTemplate := to.(*template.BiologyTemplate)

	moveTemp := marrytemplate.GetMarryTemplateService().GetMarryMoveFirstTeamplate()

	hunCheNPC := &HunCheNPC{}
	n := npc.NewNPCBase(hunCheNPC, ownerType, ownerId, ownerAllianceId, id, idInScene, weddingCarTemplate, moveTemp.GetPos(), 0)
	hunCheObj := CreateHunCheObject()
	hunCheObj.OwerId = ownerId
	hunCheObj.Period = periodId
	hunCheObj.MoveId = int32(moveTemp.Id)
	hunCheObj.PlayerId = playerId
	hunCheObj.SpouseId = spouseId
	hunCheObj.LastTime = now
	hunCheObj.HunCheGrade = carGrade
	hunCheObj.SugarGrade = marrytypes.MarryBanquetSubTypeSugar(sugarGrade)

	hunCheNPC.NPCBase = n
	hunCheNPC.hunCheObject = hunCheObj
	hunCheNPC.Calculate()
	return hunCheNPC
}
