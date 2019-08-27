package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/email/dao"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

func init() {
	player.RegisterPlayerDataManager(types.PlayerEmailDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerEmailDataManager))
}

//玩家邮件数据管理器
type PlayerEmailDataManager struct {
	p player.Player
	//玩家邮件列表
	playerEmailList []*PlayerEmailObject
}

//玩家
func (m *PlayerEmailDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerEmailDataManager) Load() (err error) {
	//加载玩家邮件
	emailsEntityArr, err := dao.GetEmailDao().GetEmails(m.Player().GetId())
	if err != nil {
		return
	}

	for _, emailEntity := range emailsEntityArr {
		newPlayerEmailObj := NewPlayerEmailObject(m.Player())
		newPlayerEmailObj.FromEntity(emailEntity)
		// if err != nil {
		// 	return
		// }
		m.playerEmailList = append(m.playerEmailList, newPlayerEmailObj)
	}

	return
}

//加载后
func (m *PlayerEmailDataManager) AfterLoad() error {
	//删除过期邮件
	newArr := make([]*PlayerEmailObject, 0, len(m.playerEmailList))
	for _, emailObj := range m.playerEmailList {
		if emailObj.isTimeOut() {
			now := global.GetGame().GetTimeService().Now()
			emailObj.deleteTime = now
			emailObj.SetModified()
			continue
		}
		newArr = append(newArr, emailObj)
	}
	m.playerEmailList = newArr

	//是否数量达到上限
	for m.IsOutMaxLimitCount() {
		canDelEmail := m.findCanDelEmail()
		if canDelEmail == nil {
			panic("add a new email: never reach here")
		}

		m.DelEmail(canDelEmail.id)
	}

	return nil
}

//心跳
func (m *PlayerEmailDataManager) Heartbeat() {

}

//获取邮件列表
func (m *PlayerEmailDataManager) GetEmails() []*PlayerEmailObject {
	return m.playerEmailList
}

func (m *PlayerEmailDataManager) GetEmail(mailId int64) (int32, *PlayerEmailObject) {
	for index, emailObj := range m.playerEmailList {
		if emailObj.id == mailId {
			return int32(index), emailObj
		}
	}
	return -1, nil
}

//邮件是否已读
func (m *PlayerEmailDataManager) IsRead(emailId int64) bool {
	_, email := m.GetEmail(emailId)
	if email == nil {
		return false
	}
	if email.isRead == 0 {
		return false
	}
	return true
}

//设置邮件已读
func (m *PlayerEmailDataManager) ReadEmail(readEmailId int64) (err error) {
	_, emailObj := m.GetEmail(readEmailId)
	if emailObj == nil {
		return
	}

	if emailObj.isRead == 0 {
		emailObj.isRead = 1
		emailObj.updateTime = global.GetGame().GetTimeService().Now()
		emailObj.SetModified()
	}
	return
}

//获取已读邮件
func (m *PlayerEmailDataManager) GetReadEmails() map[int64]*PlayerEmailObject {
	readEmails := make(map[int64]*PlayerEmailObject)
	for _, emailObje := range m.playerEmailList {
		if emailObje.isRead != 0 {
			readEmails[emailObje.GetEmailId()] = emailObje
		}
	}

	return readEmails
}

//是否【无or领取】附件
func (m *PlayerEmailDataManager) HasNotOrReceiveAttachment(emailId int64) bool {
	_, obj := m.GetEmail(emailId)
	if obj == nil {
		return false
	}

	if len(obj.attachmentInfo) == 0 {
		return true
	}
	if obj.isGetAttachment != 0 {
		return true
	}
	return false
}

//删除邮件
func (m *PlayerEmailDataManager) DelEmail(emailId int64) (err error) {
	index, emailObj := m.GetEmail(emailId)
	if emailObj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()

	emailObj.deleteTime = now
	emailObj.SetModified()
	m.playerEmailList = append(m.playerEmailList[:index], m.playerEmailList[index+1:]...)

	return
}

//设置邮件领取附件
func (m *PlayerEmailDataManager) ReceiveEmailAttachment(emailId int64) (err error) {
	_, emailObj := m.GetEmail(emailId)
	if emailObj == nil {
		return
	}
	emailObj.updateTime = global.GetGame().GetTimeService().Now()
	emailObj.isRead = 1
	emailObj.isGetAttachment = 1
	emailObj.SetModified()

	return
}

//获取未领取附件邮件列表
func (m *PlayerEmailDataManager) GetNotReceiveAttachmentEmails() []*PlayerEmailObject {
	var attachmentEmails []*PlayerEmailObject
	for _, emailObje := range m.playerEmailList {
		if !m.HasNotOrReceiveAttachment(emailObje.GetEmailId()) {
			attachmentEmails = append(attachmentEmails, emailObje)
		}
	}

	return attachmentEmails
}

//是否达到上限
func (m *PlayerEmailDataManager) IsOutMaxLimitCount() bool {
	maxLimitCount := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEmailNumLimit)
	return len(m.playerEmailList) > int(maxLimitCount)
}

//添加邮件1
func (m *PlayerEmailDataManager) AddNewEmail(title string, content string, attachmentInfo []*droptemplate.DropItemData) (newEmailObj *PlayerEmailObject) {
	now := global.GetGame().GetTimeService().Now()
	newEmailObj = m.AddEmail(title, content, now, attachmentInfo)
	return
}

//添加邮件2
func (m *PlayerEmailDataManager) AddEmail(title string, content string, createTime int64, attachmentInfo []*droptemplate.DropItemData) (newEmailObj *PlayerEmailObject) {
	emailId, _ := idutil.GetId()

	newEmailObj = NewPlayerEmailObject(m.Player())
	newEmailObj.id = emailId
	newEmailObj.isRead = 0
	newEmailObj.isGetAttachment = 0
	newEmailObj.title = title
	newEmailObj.content = content
	newEmailObj.attachmentInfo = attachmentInfo
	newEmailObj.createTime = createTime
	newEmailObj.SetModified()

	m.addEmail(newEmailObj)
	return
}

//添加邮件对象
func (m *PlayerEmailDataManager) addEmail(newEmailObj *PlayerEmailObject) {
	m.playerEmailList = append(m.playerEmailList, newEmailObj)
}

//获取允许删除的邮件
func (m *PlayerEmailDataManager) findCanDelEmail() *PlayerEmailObject {

	for _, emailObj := range m.playerEmailList {
		if emailObj.isRead != 0 {
			if m.HasNotOrReceiveAttachment(emailObj.GetEmailId()) {
				return emailObj
			}
		}
	}
	if len(m.playerEmailList) != 0 {
		return m.playerEmailList[0]
	}
	return nil
}

func (m *PlayerEmailDataManager) GMClearEmail() {
	now := global.GetGame().GetTimeService().Now()

	for _, emailObj := range m.GetEmails() {
		emailObj.deleteTime = now
		emailObj.SetModified()
	}

	m.playerEmailList = []*PlayerEmailObject{}
}

func CreatePlayerEmailDataManager(p player.Player) player.PlayerDataManager {
	maxLimitCount := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEmailNumLimit)

	m := &PlayerEmailDataManager{}
	m.p = p
	m.playerEmailList = make([]*PlayerEmailObject, 0, maxLimitCount)

	return m
}
