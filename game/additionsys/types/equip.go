package types

//强化类型
type StrengthenType int32

const (
	//强化
	StrengthenTypeUpgrade StrengthenType = iota + 1
)

func (est StrengthenType) Valid() bool {
	switch est {
	case StrengthenTypeUpgrade:
		return true
	}
	return false
}

type StrengthenResultType int32

const (
	//强化成功
	StrengthenResultTypeSuccess StrengthenResultType = iota
	//强化失败
	StrengthenResultTypeFailed
	//强化回退
	StrengthenResultTypeBack
)

type ShengJiResType int32 //附加系统升级返回

const (
	ShengJiResTypeSucceed  ShengJiResType = iota //升级成功
	ShengJiResTypeDefeated                       //升级失败
)
