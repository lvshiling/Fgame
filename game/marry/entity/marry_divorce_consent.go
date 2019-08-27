package entity

//协议离婚成功请求者已下线数据 (拥有扣除协议离婚亲密度)
type MarryDivorceConsentEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	PlayerId   int64 `gorm:"column:playerId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *MarryDivorceConsentEntity) GetId() int64 {
	return p.Id
}

func (p *MarryDivorceConsentEntity) TableName() string {
	return "t_marry_divorce_consent"
}
