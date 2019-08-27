package types

type ItemSkillType int32

const (
	ItemSkillTypeHuTiDun ItemSkillType = iota + 1
)

func (wt ItemSkillType) Valid() bool {
	switch wt {
	case ItemSkillTypeHuTiDun:
		return true
	}
	return false
}

var (
	itemSkillTypeMap = map[ItemSkillType]string{
		ItemSkillTypeHuTiDun: "金刚护体盾",
	}
)

func (spt ItemSkillType) String() string {
	return itemSkillTypeMap[spt]
}
