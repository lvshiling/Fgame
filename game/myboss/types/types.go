package types

type MyBossType int32

const (
	MyBossTypeCommon MyBossType = iota + 1 //普通个人BOSS
	MyBossTypeVip                          //VIP个人BOSS
)

func (t MyBossType) Valid() bool {
	switch t {
	case MyBossTypeCommon,
		MyBossTypeVip:
		return true
	}

	return false
}
