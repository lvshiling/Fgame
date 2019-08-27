package model

type MailApply struct {
	Id               int64  `gorm:"primary_key;column:id"`
	MailType         int    `gorm:"column:mailType"`
	ServerId         int    `gorm:"column:serverId"`
	Title            string `gorm:"column:title"`
	Content          string `gorm:"column:content"`
	Playerlist       string `gorm:"column:playerlist"`
	Proplist         string `gorm:"column:proplist"`
	FreezTime        int    `gorm:"column:freezTime"`
	EffectDays       int    `gorm:"column:effectDays"`
	RoleStartTime    int64  `gorm:"column:roleStartTime"`
	RoleEndTime      int64  `gorm:"column:roleEndTime"`
	MinLevel         int    `gorm:"column:minLevel"`
	MaxLevel         int    `gorm:"column:maxLevel"`
	UpdateTime       int64  `gorm:"column:updateTime"`
	CreateTime       int64  `gorm:"column:createTime"`
	DeleteTime       int64  `gorm:"column:deleteTime"`
	MailUser         int64  `gorm:"column:mailUser"`
	MailTime         int64  `gorm:"column:mailTime"`
	MailState        int    `gorm:"column:mailState"`
	ApproveUser      int64  `gorm:"column:approveUser"`
	ApproveTime      int64  `gorm:"column:approveTime"`
	ApproveReason    string `gorm:"column:approveReason"`
	SendFlag         int    `gorm:"column:sendFlag"`
	SdkType          int    `gorm:"column:sdkType"`
	CenterPlatformId int64  `gorm:"column:centerPlatformId"`
	BindFlag         int    `gorm:"column:bindFlag"`
	Remark           string `gorm:"column:remark"`
}

func (m *MailApply) TableName() string {
	return "t_mail_apply"
}

func (m *MailApply) IfCanApprove() bool {
	if m.MailState == 1 {
		return true
	}
	return false
}
