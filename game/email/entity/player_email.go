package entity

type PlayerEmailEntity struct {
	Id              int64  `gorm:"primary_key;column:id"`  //邮件Id
	PlayerId        int64  `gorm:"column:playerId"`        //玩家Id
	IsRead          int32  `gorm:"column:isRead"`          //是否已读
	IsGetAttachment int32  `gorm:"column:isGetAttachment"` //是否已领取附件
	Title           string `gorm:"column:title"`           //邮件标题
	Content         string `gorm:"column:content"`         //邮件内容
	AttachementInfo string `gorm:"column:attachementInfo"` //附件信息
	UpdateTime      int64  `gorm:"column:updateTime"`      //更新时间
	CreateTime      int64  `gorm:"column:createTime"`      //创建时间
	DeleteTime      int64  `gorm:"column:deleteTime"`      //删除时间

}

func (e *PlayerEmailEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerEmailEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerEmailEntity) TableName() string {
	return "t_player_email"
}
