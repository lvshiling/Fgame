package entity

//玩家每日消费记录数据
type PlayerCycleCostRecordEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	CostNum       int64 `gorm:"column:costNum"`
	PreDayCostNum int64 `gorm:"column:preDayCostNum"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerCycleCostRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerCycleCostRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerCycleCostRecordEntity) TableName() string {
	return "t_player_cycle_cost_record"
}
