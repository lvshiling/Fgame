package entity

type PlayerShangguzhilingEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	PlayerId      int64  `gorm:"column:playerId"`
	LingShouType  int32  `gorm:"column:lingShouType"`
	Level         int32  `gorm:"column:level"`         //喂养等级
	Experience    int64  `gorm:"column:experience"`    //喂养经验条
	LingWen       string `gorm:"column:lingwen"`       //灵纹
	UprankLevel   int32  `gorm:"column:uprankLevel"`   //进阶等级
	UprankBless   int64  `gorm:"column:uprankBless"`   //进阶祝福值
	UprankTimes   int32  `gorm:"column:uprankTimes"`   //进阶已尝试次数
	Linglian      string `gorm:"column:linglian"`      //灵炼
	LinglianTimes int32  `gorm:"column:linglianTimes"` //灵炼次数
	ReceiveTime   int64  `gorm:"column:receiveTime"`   //上一次领取奖励的时间
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (p *PlayerShangguzhilingEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerShangguzhilingEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerShangguzhilingEntity) TableName() string {
	return "t_player_shangguzhiling"
}
