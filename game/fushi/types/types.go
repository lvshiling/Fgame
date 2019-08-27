package types

// 八卦符石类型
type FuShiType int32

const (
	FuShiTypeJiaMu   FuShiType = iota // 甲木
	FuShiTypeYiMu                     // 乙木
	FuShiTypeBingHuo                  // 丙火
	FuShiTypeDingHuo                  // 丁火
	FuShiTypeWuTu                     // 戊土
	FuShiTypeJiTu                     // 己土
	FuShiTypeGengJin                  // 庚金
	FuShiTypeXinJin                   // 辛金
	FuShiTypeRenShui                  // 壬水
	FuShiTypeKuiSHui                  // 癸水
)

func (f FuShiType) Vaild() bool {
	switch f {
	case FuShiTypeJiaMu,
		FuShiTypeYiMu,
		FuShiTypeBingHuo,
		FuShiTypeDingHuo,
		FuShiTypeWuTu,
		FuShiTypeJiTu,
		FuShiTypeGengJin,
		FuShiTypeXinJin,
		FuShiTypeRenShui,
		FuShiTypeKuiSHui:
		return true
	default:
		return false
	}
}
