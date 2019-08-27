package types

//藏经阁boss类型
type CangJingGeBossType int32

const (
	CangJingGeBossTypeLocal CangJingGeBossType = 0 //本服藏经阁boss
)

func (t CangJingGeBossType) Valid() bool {
	switch t {
	case CangJingGeBossTypeLocal:
		return true
	default:
		return false
	}
}
