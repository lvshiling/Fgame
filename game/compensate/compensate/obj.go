package compensate

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	compensateentity "fgame/fgame/game/compensate/entity"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//全服补偿对象
type CompensateObject struct {
	id             int64
	serverId       int32
	title          string
	content        string
	attachment     []*droptemplate.DropItemData
	roleLevel      int32
	roleCreateTime int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewCompensateObject() *CompensateObject {
	o := &CompensateObject{}
	return o
}

func convertNewCompensateObjectToEntity(o *CompensateObject) (*compensateentity.CompensateEntity, error) {
	attachmentList, err := json.Marshal(o.attachment)
	if err != nil {
		return nil, err
	}

	e := &compensateentity.CompensateEntity{
		Id:             o.id,
		ServerId:       o.serverId,
		Titlte:         o.title,
		Content:        o.content,
		Attachment:     string(attachmentList),
		RoleLevel:      o.roleLevel,
		RoleCreateTime: o.roleCreateTime,
		UpdateTime:     o.updateTime,
		DeleteTime:     o.deleteTime,
		CreateTime:     o.createTime,
	}
	return e, nil
}

func (o *CompensateObject) GetDBId() int64 {
	return o.id
}

func (o *CompensateObject) GetCompensateId() int64 {
	return o.id
}

func (o *CompensateObject) GetTitle() string {
	return o.title
}

func (o *CompensateObject) GetConetent() string {
	return o.content
}

func (o *CompensateObject) GetAttachment() []*droptemplate.DropItemData {
	return o.attachment
}

func (o *CompensateObject) GetRoleCreateTime() int64 {
	return o.roleCreateTime
}

func (o *CompensateObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *CompensateObject) GetRoleLevel() int32 {
	return o.roleLevel
}

func (o *CompensateObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewCompensateObjectToEntity(o)
	return e, err
}

func (o *CompensateObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*compensateentity.CompensateEntity)

	var attachmentList []*droptemplate.DropItemData
	err := json.Unmarshal([]byte(pse.Attachment), &attachmentList)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.title = pse.Titlte
	o.content = pse.Content
	o.attachment = attachmentList
	o.roleLevel = pse.RoleLevel
	o.roleCreateTime = pse.RoleCreateTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *CompensateObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "compensate"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
