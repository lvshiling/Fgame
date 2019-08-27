package types

//外域boss类型
type OutlandBossType int32

const (
	OutlandBossTypeLocal OutlandBossType = 0 //本服外域boss
)

func (t OutlandBossType) Valid() bool {
	switch t {
	case OutlandBossTypeLocal:
		return true
	default:
		return false
	}
}
