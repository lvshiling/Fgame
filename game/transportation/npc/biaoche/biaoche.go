package npc

import (
	"fgame/fgame/core/storage"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	transportationentity "fgame/fgame/game/transportation/entity"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	"fmt"

	transportationtemplate "fgame/fgame/game/transportation/template"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

type TransportationObject struct {
	id                     int64
	playerId               int64
	serverId               int32
	allianceId             int64
	owerName               string
	transportMoveId        int32
	transportType          transportationtypes.TransportationType
	state                  transportationtypes.TransportStateType
	robName                string
	LastDistressUpdateTime int64
	updateTime             int64
	createTime             int64
	deleteTime             int64
}

func (o *TransportationObject) GetState() transportationtypes.TransportStateType {
	return o.state
}

func (o *TransportationObject) GetTransportType() transportationtypes.TransportationType {
	return o.transportType
}

func (o *TransportationObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TransportationObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *TransportationObject) GetRobName() string {
	return o.robName
}

func (o *TransportationObject) GetTransportMoveId() int32 {
	return o.transportMoveId
}

func (o *TransportationObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *TransportationObject) GetTransportId() int64 {
	return o.id
}

func (o *TransportationObject) SetRobName(robName string) {
	o.robName = robName
}

func CreateTransportationObject() *TransportationObject {
	o := &TransportationObject{}
	return o
}

func convertTransportationObjectToEntity(o *TransportationObject) (*transportationentity.TransportationEntity, error) {
	e := &transportationentity.TransportationEntity{
		Id:                     o.id,
		PlayerId:               o.playerId,
		ServerId:               o.serverId,
		AllianceId:             o.allianceId,
		TransportMoveId:        o.transportMoveId,
		TransportType:          int32(o.transportType),
		State:                  int32(o.state),
		RobName:                o.robName,
		OwerName:               o.owerName,
		LastDistressUpdateTime: o.LastDistressUpdateTime,
		UpdateTime:             o.updateTime,
		CreateTime:             o.createTime,
		DeleteTime:             o.deleteTime,
	}
	return e, nil
}

func (o *TransportationObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertTransportationObjectToEntity(o)
	return e, err
}

func (o *TransportationObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*transportationentity.TransportationEntity)
	o.id = ae.Id
	o.playerId = ae.PlayerId
	o.serverId = ae.ServerId
	o.allianceId = ae.AllianceId
	o.transportMoveId = ae.TransportMoveId
	o.transportType = transportationtypes.TransportationType(ae.TransportType)
	o.state = transportationtypes.TransportStateType(ae.State)
	o.robName = ae.RobName
	o.owerName = ae.OwerName
	o.LastDistressUpdateTime = ae.LastDistressUpdateTime
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *TransportationObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Transportation"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *TransportationObject) IsDistressCD() bool {
	if o.LastDistressUpdateTime == 0 {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	cd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDistressSignalCD))
	dif := now - o.LastDistressUpdateTime

	if dif < cd {
		return true
	}

	return false
}

func (o *TransportationObject) IsReachGoal(curPos coretypes.Position) bool {
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplate(o.GetTransportMoveId())
	nextTem := tem.GetNextTemp()
	if nextTem == nil {
		return true
	}
	distance := utils.DistanceSquare(curPos, nextTem.GetPosition())
	if distance > common.MIN_DISTANCE_SQUARE_ERROR {
		return false
	}
	return true
}

type BiaocheNPC struct {
	scene.NPC
	transportationObject *TransportationObject
}

func (n *BiaocheNPC) GetName() string {
	biaoCheTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(n.transportationObject.transportType)
	biaoCheName := fmt.Sprintf("%s的%s", n.transportationObject.owerName, biaoCheTemp.GetBiaocheTemplate().Name)
	return biaoCheName
}

func (n *BiaocheNPC) GetTransportationObject() *TransportationObject {
	return n.transportationObject
}

func (n *BiaocheNPC) Finish() {
	obj := n.transportationObject
	now := global.GetGame().GetTimeService().Now()
	obj.state = transportationtypes.TransportStateTypeFinish
	obj.updateTime = now
	obj.SetModified()
	return
}

func (n *BiaocheNPC) Fail(robName string) {
	obj := n.transportationObject
	now := global.GetGame().GetTimeService().Now()
	obj.state = transportationtypes.TransportStateTypeFail
	obj.robName = robName
	obj.updateTime = now
	obj.SetModified()
	return
}

func (n *BiaocheNPC) ReachGoal() bool {
	transportMoveId := n.transportationObject.transportMoveId
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplate(transportMoveId)
	nextMoveTemplate := tem.GetNextTemp()
	if nextMoveTemplate == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	n.transportationObject.transportMoveId = int32(nextMoveTemplate.TemplateId())
	n.transportationObject.updateTime = now
	n.transportationObject.SetModified()

	if nextMoveTemplate.GetNextTemp() == nil {
		//镖车完成
		gameevent.Emit(transportationeventtypes.EventTypeTransportationFinish, n, nil)

		n.GetScene().RemoveSceneObject(n, false)
	}

	return true
}

func (n *BiaocheNPC) Remove() {
	obj := n.transportationObject
	now := global.GetGame().GetTimeService().Now()
	obj.deleteTime = now
	obj.SetModified()
	return
}

func (n *BiaocheNPC) DistressSignal() {
	now := global.GetGame().GetTimeService().Now()
	n.transportationObject.LastDistressUpdateTime = now
	n.transportationObject.SetModified()
	return
}

func CreateBiaoCheNPCWithObj(transportationObject *TransportationObject, allianceId int64) *BiaocheNPC {
	ownerType := scenetypes.OwnerTypePlayer
	ownerId := transportationObject.playerId
	ownerAllianceId := allianceId
	if transportationObject.transportType == transportationtypes.TransportationTypeAlliance {
		ownerId = transportationObject.allianceId
		ownerType = scenetypes.OwnerTypeAlliance
	}

	idInScene := int32(0)
	id, _ := idutil.GetId()
	tempTransportation := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(transportationObject.GetTransportType())
	biologyTemplate := tempTransportation.GetBiaocheTemplate()
	moveTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplate(transportationObject.GetTransportMoveId())

	biaocheNPC := &BiaocheNPC{}
	n := npc.NewNPCBase(biaocheNPC, ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, moveTemp.GetPosition(), 0)
	// n := npc.CreateNPC(ownerType, ownerId, id, idInScene, biologyTemplate, moveTemp.GetPosition(), 0)
	biaocheNPC.NPC = n
	biaocheNPC.transportationObject = transportationObject
	biaocheNPC.Calculate()

	return biaocheNPC
}

func CreateBiaoCheNPC(playerId, allianceId int64, playerAllianceId int64, owerName string, typ transportationtypes.TransportationType) *BiaocheNPC {

	obj := CreateTransportationObject()
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	initMoveTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplateFirst()

	obj.id = id
	obj.playerId = playerId
	obj.serverId = global.GetGame().GetServerIndex()
	obj.allianceId = allianceId
	obj.transportMoveId = int32(initMoveTemp.TemplateId())
	obj.transportType = typ
	obj.state = transportationtypes.TransportStateTypeRuning
	obj.owerName = owerName
	obj.createTime = now
	obj.SetModified()
	biaoChe := CreateBiaoCheNPCWithObj(obj, playerAllianceId)

	return biaoChe
}
