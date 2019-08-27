package types

// 打宝塔操作类型
type TowerOperationType int32

const (
	TowerOperationTypeBegin TowerOperationType = iota //打宝开始
	TowerOperationTypeEnd                             //打宝结束
)

func (t TowerOperationType) Valid() bool {
	switch t {
	case TowerOperationTypeBegin,
		TowerOperationTypeEnd:
		return true
	}
	return false
}
