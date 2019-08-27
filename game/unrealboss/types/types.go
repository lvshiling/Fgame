package types

//幻境boss类型
type UnrealBossType int32

const (
	UnrealBossTypeUnreal UnrealBossType = 0 //幻境boss
)

func (t UnrealBossType) Valid() bool {
	switch t {
	case UnrealBossTypeUnreal:
		return true
	default:
		return false
	}
}
