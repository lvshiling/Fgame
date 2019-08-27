package entity

//仙盟数据
type AllianceEntity struct {
	Id                       int64  `gorm:"primary_key;column:id"`
	ServerId                 int32  `gorm:"column:serverId"`
	OriginServerId           int32  `gorm:"column:originServerId"`
	Name                     string `gorm:"column:name"`
	Notice                   string `gorm:"column:notice"`
	Level                    int32  `gorm:"column:level"`
	JianShe                  int64  `gorm:"column:jianShe"`
	HuFu                     int64  `gorm:"column:huFu"`
	TotalForce               int64  `gorm:"column:totalForce"`
	MengzhuId                int64  `gorm:"column:mengzhuId"`
	CreateId                 int64  `gorm:"column:createId"`
	TransportTimes           int32  `gorm:"column:transportTimes"`
	LastTransportRefreshTime int64  `gorm:"column:lastTransportRefreshTime"`
	IsAutoAgree              int32  `gorm:"column:isAutoAgree"`
	IsAutoRemoveDepot        int32  `gorm:"column:isAutoRemoveDepot"`
	MaxRemoveZhuanSheng      int32  `gorm:"column:maxRemoveZhuanSheng"`
	MaxRemoveQuality         int32  `gorm:"column:maxRemoveQuality"`
	LastMergeTime            int64  `gorm:"column:lastMergeTime"`
	// CampType                 int32  `gorm:"column:campType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *AllianceEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceEntity) TableName() string {
	return "t_alliance"
}
