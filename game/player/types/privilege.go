package types

type PrivilegeType int32

const (
	PrivilegeTypeNone   PrivilegeType = iota //没有权限
	PrivilegeTypeNormal                      //普通扶持
	PrivilegeTypeDev                         //研发扶持
)

func (t PrivilegeType) Valid() bool {
	switch t {
	case PrivilegeTypeNone,
		PrivilegeTypeNormal,
		PrivilegeTypeDev:
		return true
	}
	return false
}
