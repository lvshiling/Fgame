package alliance

import (
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//仙盟boss
type AllianceBossObject struct {
	id         int64
	serverId   int32
	allianceId int64
	summonTime int64
	bossLevel  int32
	bossExp    int32
	isSummon   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func createAllianceBossObject() *AllianceBossObject {
	o := &AllianceBossObject{}
	return o
}

func convertAllianceBossObjectToEntity(o *AllianceBossObject) (*allianceentity.AllianceBossEntity, error) {
	e := &allianceentity.AllianceBossEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		AllianceId: o.allianceId,
		SummonTime: o.summonTime,
		BossLevel:  o.bossLevel,
		BossExp:    o.bossExp,
		IsSummon:   o.isSummon,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *AllianceBossObject) GetId() int64 {
	return o.id
}

func (o *AllianceBossObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceBossObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceBossObject) GetServerId() int32 {
	return o.serverId
}

func (o *AllianceBossObject) GetSummonTime() int64 {
	return o.summonTime
}

func (o *AllianceBossObject) GetBossLevel() int32 {
	return o.bossLevel
}

func (o *AllianceBossObject) GetBossExp() int32 {
	return o.bossExp
}

func (o *AllianceBossObject) GetIsSummon() bool {
	return o.isSummon == 1
}

func (o *AllianceBossObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceBossObject) SetIsSummon(isSummon bool) {
	now := global.GetGame().GetTimeService().Now()
	if isSummon {
		o.isSummon = 1
		o.summonTime = now
		o.bossExp = 0
		o.bossLevel = 1
	} else {
		o.isSummon = 0
	}

	o.updateTime = now
	o.SetModified()
}

func (o *AllianceBossObject) IsCrossFive() (isCrossFive bool, err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := o.summonTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return false, err
	}
	if !flag {
		o.summonTime = 0
		o.isSummon = 0
		// o.bossExp = 0
		// o.bossLevel = 1
		o.updateTime = now
		o.SetModified()
	}
	return
}

func (o *AllianceBossObject) AddExp(exp int32) {
	if exp <= 0 {
		return
	}
	// o.IsCrossFive()
	// if o.GetIsSummon() {
	// 	return
	// }
	allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(o.bossLevel)
	if allianceBossTemplate == nil {
		return
	}
	nextAllianceBossTemplate := allianceBossTemplate.GetNextTemplate()
	if nextAllianceBossTemplate == nil {
		return
	}
	o.bossExp += exp
	for o.bossExp >= nextAllianceBossTemplate.Experience {
		o.bossExp -= nextAllianceBossTemplate.Experience
		allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(o.bossLevel)
		if allianceBossTemplate == nil {
			break
		}
		nextAllianceBossTemplate = allianceBossTemplate.GetNextTemplate()
		if nextAllianceBossTemplate == nil {
			break
		}
		o.bossLevel += 1
	}
	o.SetModified()
}

func (o *AllianceBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceBossObjectToEntity(o)
	return e, err
}

func (o *AllianceBossObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceBossEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.allianceId = ae.AllianceId
	o.summonTime = ae.SummonTime
	o.bossLevel = ae.BossLevel
	o.bossExp = ae.BossExp
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceBoss"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
