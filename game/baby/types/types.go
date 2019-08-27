package types

// 天赋技能状态
type SkillStatusType int32

const (
	SkillStatusTypeUnLock SkillStatusType = iota //激活未锁定
	SkillStatusTypeLock                          //激活已锁定
)

func (t SkillStatusType) Valid() bool {
	switch t {
	case SkillStatusTypeUnLock,
		SkillStatusTypeLock:
		return true
	default:
		return false
	}
}

// 玩具套装类型
type ToySuitType int32

const (
	ToySuitTypeOne   ToySuitType = iota //套装1
	ToySuitTypeTwo                      //套装2
	ToySuitTypeThree                    //套装3
	ToySuitTypeFour                     //套装4
	ToySuitTypeFive                     //套装5
)

func (t ToySuitType) Valid() bool {
	switch t {
	case ToySuitTypeOne,
		ToySuitTypeTwo,
		ToySuitTypeThree,
		ToySuitTypeFour,
		ToySuitTypeFive:
		return true
	default:
		return false
	}
}

const (
	MinSuitType = ToySuitTypeOne
	MaxSuitType = ToySuitTypeFive
)

// 技能类型
type SkillType int32

const (
	SkillTypeSkill SkillType = iota //技能
	SkillTypeAttr                   //属性
)

func (t SkillType) Valid() bool {
	switch t {
	case SkillTypeSkill,
		SkillTypeAttr:
		return true
	default:
		return false
	}
}
