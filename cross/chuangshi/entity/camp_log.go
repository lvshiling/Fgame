package entity

// TODO xzk27 战报
//  方阵营神王【玩家名称六字】将【阵营名称】的【城池名称六字】设为了攻城目标，请各位仙友届时准时参与，共谋大业
// b)【阵营名称】的神王【玩家名称六字】将我方的【城池名称六字】设为了攻城目标，请各位仙友届时准时参与，共卫家园
// c)我方阵营总仙友众志成城，成功攻破【阵营名称】的【城池名称六字】，一统仙界的日子指日可待
// d)我方【城池名称六】被【阵营名称】攻破，家园沦陷，此等大辱岂可坐视不理，望众仙友下次城战共同参与，夺回家园

// a)我方阵营神王【玩家名称六字】对阵营工资进行了分配，各位玩家可前往查看   【前往查看】，
// b)我方阵营城主【玩家名称六字】对城池工资进行了分配，各位玩家可前往查看   【前往查看】，

// a)我方阵营神王【玩家名称六字】任命【玩家名称六字】为【城池名称六字】的城主，任命后城主享有分配城池工资的特权
// b)我方阵营【玩家名称六字】成功当选为新一任的神王，享有分配阵营工资、进攻城池、任命城主等特权，望各位仙友协力扶持，共谋大业
// d)城池名称六字】的城主【角色名称六字】由于卸任了仙盟盟主职位，因此该城池城主自动移交给该仙盟新一任盟主【角色名称六字】

//阵营战报
type ChuangShiCampLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	Platform   int32  `gorm:"column:platform"`
	ServerId   int32  `gorm:"column:serverId"`
	CampType   int32  `gorm:"column:campType"` //阵营类型
	Content    string `gorm:"column:content"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *ChuangShiCampLogEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiCampLogEntity) TableName() string {
	return "t_chuangshi_camp_log"
}
