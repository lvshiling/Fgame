package entity

//聊天设置
type ChatSettingEntity struct {
	Id               int64 `gorm:"primary_key;column:id"`
	ServerId         int32 `gorm:"column:serverId"`
	WorldVipLevel    int32 `gorm:"column:worldVipLevel"`
	WorldLevel       int32 `gorm:"column:worldLevel"`
	AllianceVipLevel int32 `gorm:"column:allianceVipLevel"`
	AllianceLevel    int32 `gorm:"column:allianceLevel"`
	PrivateVipLevel  int32 `gorm:"column:privateVipLevel"`
	PrivateLevel     int32 `gorm:"column:privateLevel"`
	TeamVipLevel     int32 `gorm:"column:teamVipLevel"`
	TeamLevel        int32 `gorm:"column:teamLevel"`
	UpdateTime       int64 `gorm:"column:updateTime"`
	CreateTime       int64 `gorm:"column:createTime"`
	DeleteTime       int64 `gorm:"column:deleteTime"`
}

func (e *ChatSettingEntity) GetId() int64 {
	return e.Id
}

func (e *ChatSettingEntity) TableName() string {
	return "t_chat_setting"
}
