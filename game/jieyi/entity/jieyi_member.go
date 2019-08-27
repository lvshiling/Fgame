package entity

// 结义成员数据
type JieYiMemberEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	JieYiId     int64  `gorm:"column:jieYiId"`
	PlayerId    int64  `gorm:"column:playerId"`
	Name        string `gorm:"column:name"`
	Level       int32  `gorm:"column:level"`
	Role        int32  `gorm:"column:role"`
	Sex         int32  `gorm:"column:sex"`
	ZhuanSheng  int32  `gorm:"column:zhuanSheng"`
	Force       int64  `gorm:"column:force"`
	TokenType   int32  `gorm:"column:tokenType"`
	TokenLev    int32  `gorm:"column:tokenLev"`
	TokenPro    int32  `gorm:"column:tokenPro"`
	TokenNum    int32  `gorm:"column:tokenNum"`
	JieYiDaoJu  int32  `gorm:"column:jieYiDaoJu"`
	NameLev     int32  `gorm:"column:nameLev"`
	ShengWeiZhi int32  `gorm:"column:shengWeiZhi"`
	JieYiTime   int64  `gorm:"column:jieYiTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *JieYiMemberEntity) GetId() int64 {
	return e.Id
}

func (e *JieYiMemberEntity) TableName() string {
	return "t_jieyi_member"
}
