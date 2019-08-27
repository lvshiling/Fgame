package types

//各系统类型
type WardrobeSysType int32

const (
	//坐骑
	WardrobeSysTypeMount WardrobeSysType = 1 + iota
	//战翼
	WardrobeSysTypeWing
	//兵魂
	WardrobeSysTypeWeapon
	//法宝
	WardrobeSysTypeFaBao
	//身法
	WardrobeSysTypeShenFa
	//仙体
	WardrobeSysTypeXianTi
	//领域
	WardrobeSysTypeField
	//时装
	WardrobeSysTypeFashion
	//称号
	WardrobeSysTypeTitle
)

// //衣橱套装类型
// type WardrobeType int32

// const (
// 	//元旦套装
// 	WardrobeTypeOne WardrobeType = iota
// 	//城战套装
// 	WardrobeTypeTwo
// 	WardrobeTypeThree
// 	WardrobeTypeFour
// 	WardrobeTypeFive
// 	WardrobeTypeSix
// 	WardrobeTypeSeven
// 	WardrobeTypeEight
// 	WardrobeTypeNine
// 	WardrobeTypeTen
// 	WardrobeTypeEleven
// 	WardrobeTypeTwelve
// 	WardrobeTypeThirteen
// )

// func (t WardrobeType) Valid() (flag bool) {
// 	switch t {
// 	case WardrobeTypeOne,
// 		WardrobeTypeTwo,
// 		WardrobeTypeThree,
// 		WardrobeTypeFour,
// 		WardrobeTypeFive,
// 		WardrobeTypeSix,
// 		WardrobeTypeSeven,
// 		WardrobeTypeEight,
// 		WardrobeTypeNine,
// 		WardrobeTypeTen,
// 		WardrobeTypeEleven,
// 		WardrobeTypeTwelve,
// 		WardrobeTypeThirteen:
// 		return true
// 	default:
// 		return
// 	}
// 	return
// }

// var (
// 	WardrobeTypeMin = WardrobeTypeOne
// 	WardrobeTypeMax = WardrobeTypeThirteen
// )

// type WardrobeSubType interface {
// 	SubType() int32
// 	Valid() bool
// }

// type WardrobeSubTypeFactory interface {
// 	CreateWardrobeSubType(subType int32) WardrobeSubType
// }

// type WardrobeSubTypeFactoryFunc func(subType int32) WardrobeSubType

// func (f WardrobeSubTypeFactoryFunc) CreateWardrobeSubType(subType int32) WardrobeSubType {
// 	return f(subType)
// }

// //无子类型
// type WardrobeCommonSubType int32

// const (
// 	//默认子类型
// 	WardrobeCommonSubTypeDefault WardrobeCommonSubType = iota
// )

// func (t WardrobeCommonSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeCommonSubType) Valid() bool {
// 	switch t {
// 	case
// 		WardrobeCommonSubTypeDefault:
// 		return true
// 	}
// 	return false
// }
// func CreateWardrobeCommonSubType(subType int32) WardrobeSubType {
// 	return WardrobeCommonSubType(subType)
// }

// //套装1子类型
// type WardrobeOneSubType int32

// const (
// 	WardrobeOneSubTypeOne WardrobeOneSubType = iota
// 	WardrobeOneSubTypeTwo
// 	WardrobeOneSubTypeThree
// 	WardrobeOneSubTypeFour
// )

// func (t WardrobeOneSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeOneSubType) Valid() bool {
// 	switch t {
// 	case WardrobeOneSubTypeOne,
// 		WardrobeOneSubTypeTwo,
// 		WardrobeOneSubTypeThree,
// 		WardrobeOneSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeOneSubType(subType int32) WardrobeSubType {
// 	return WardrobeOneSubType(subType)
// }

// //套装2子类型
// type WardrobeTwoSubType int32

// const (
// 	WardrobeTwoSubTypeOne WardrobeTwoSubType = iota
// 	WardrobeTwoSubTypeTwo
// 	WardrobeTwoSubTypeThree
// 	WardrobeTwoSubTypeFour
// )

// func (t WardrobeTwoSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeTwoSubType) Valid() bool {
// 	switch t {
// 	case WardrobeTwoSubTypeOne,
// 		WardrobeTwoSubTypeTwo,
// 		WardrobeTwoSubTypeThree,
// 		WardrobeTwoSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeTwoSubType(subType int32) WardrobeSubType {
// 	return WardrobeTwoSubType(subType)
// }

// //套装3子类型
// type WardrobeThreeSubType int32

// const (
// 	WardrobeThreeSubTypeOne WardrobeThreeSubType = iota
// 	WardrobeThreeSubTypeTwo
// 	WardrobeThreeSubTypeThree
// 	WardrobeThreeSubTypeFour
// )

// func (t WardrobeThreeSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeThreeSubType) Valid() bool {
// 	switch t {
// 	case WardrobeThreeSubTypeOne,
// 		WardrobeThreeSubTypeTwo,
// 		WardrobeThreeSubTypeThree,
// 		WardrobeThreeSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeThreeSubType(subType int32) WardrobeSubType {
// 	return WardrobeThreeSubType(subType)
// }

// //套装4子类型
// type WardrobeFourSubType int32

// const (
// 	WardrobeFourSubTypeOne WardrobeFourSubType = iota
// 	WardrobeFourSubTypeTwo
// 	WardrobeFourSubTypeThree
// 	WardrobeFourSubTypeFour
// )

// func (t WardrobeFourSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeFourSubType) Valid() bool {
// 	switch t {
// 	case WardrobeFourSubTypeOne,
// 		WardrobeFourSubTypeTwo,
// 		WardrobeFourSubTypeThree,
// 		WardrobeFourSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeFourSubType(subType int32) WardrobeSubType {
// 	return WardrobeFourSubType(subType)
// }

// //套装5子类型
// type WardrobeFiveSubType int32

// const (
// 	WardrobeFiveSubTypeOne WardrobeFiveSubType = iota
// 	WardrobeFiveSubTypeTwo
// 	WardrobeFiveSubTypeThree
// 	WardrobeFiveSubTypeFour
// )

// func (t WardrobeFiveSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeFiveSubType) Valid() bool {
// 	switch t {
// 	case WardrobeFiveSubTypeOne,
// 		WardrobeFiveSubTypeTwo,
// 		WardrobeFiveSubTypeThree,
// 		WardrobeFiveSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeFiveSubType(subType int32) WardrobeSubType {
// 	return WardrobeFiveSubType(subType)
// }

// //套装6子类型
// type WardrobeSixSubType int32

// const (
// 	WardrobeSixSubTypeOne WardrobeSixSubType = iota
// 	WardrobeSixSubTypeTwo
// 	WardrobeSixSubTypeThree
// 	WardrobeSixSubTypeFour
// )

// func (t WardrobeSixSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeSixSubType) Valid() bool {
// 	switch t {
// 	case WardrobeSixSubTypeOne,
// 		WardrobeSixSubTypeTwo,
// 		WardrobeSixSubTypeThree,
// 		WardrobeSixSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeSixSubType(subType int32) WardrobeSubType {
// 	return WardrobeSixSubType(subType)
// }

// //套装7子类型
// type WardrobeSevenSubType int32

// const (
// 	WardrobeSevenSubTypeOne WardrobeSevenSubType = iota
// 	WardrobeSevenSubTypeTwo
// 	WardrobeSevenSubTypeThree
// 	WardrobeSevenSubTypeFour
// )

// func (t WardrobeSevenSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeSevenSubType) Valid() bool {
// 	switch t {
// 	case WardrobeSevenSubTypeOne,
// 		WardrobeSevenSubTypeTwo,
// 		WardrobeSevenSubTypeThree,
// 		WardrobeSevenSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeSevenSubType(subType int32) WardrobeSubType {
// 	return WardrobeSevenSubType(subType)
// }

// //套装8子类型
// type WardrobeEightSubType int32

// const (
// 	WardrobeEightSubTypeOne WardrobeEightSubType = iota
// 	WardrobeEightSubTypeTwo
// 	WardrobeEightSubTypeThree
// 	WardrobeEightSubTypeFour
// )

// func (t WardrobeEightSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeEightSubType) Valid() bool {
// 	switch t {
// 	case WardrobeEightSubTypeOne,
// 		WardrobeEightSubTypeTwo,
// 		WardrobeEightSubTypeThree,
// 		WardrobeEightSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeEightSubType(subType int32) WardrobeSubType {
// 	return WardrobeEightSubType(subType)
// }

// //套装9子类型
// type WardrobeNineSubType int32

// const (
// 	WardrobeNineSubTypeOne WardrobeNineSubType = iota
// 	WardrobeNineSubTypeTwo
// 	WardrobeNineSubTypeThree
// 	WardrobeNineSubTypeFour
// )

// func (t WardrobeNineSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeNineSubType) Valid() bool {
// 	switch t {
// 	case WardrobeNineSubTypeOne,
// 		WardrobeNineSubTypeTwo,
// 		WardrobeNineSubTypeThree,
// 		WardrobeNineSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeNineSubType(subType int32) WardrobeSubType {
// 	return WardrobeNineSubType(subType)
// }

// //套装10子类型
// type WardrobeTenSubType int32

// const (
// 	WardrobeTenSubTypeOne WardrobeTenSubType = iota
// 	WardrobeTenSubTypeTwo
// 	WardrobeTenSubTypeThree
// 	WardrobeTenSubTypeFour
// )

// func (t WardrobeTenSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeTenSubType) Valid() bool {
// 	switch t {
// 	case WardrobeTenSubTypeOne,
// 		WardrobeTenSubTypeTwo,
// 		WardrobeTenSubTypeThree,
// 		WardrobeTenSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeTenSubType(subType int32) WardrobeSubType {
// 	return WardrobeTenSubType(subType)
// }

// //套装11子类型
// type WardrobeElevenSubType int32

// const (
// 	WardrobeElevenSubTypeOne WardrobeElevenSubType = iota
// 	WardrobeElevenSubTypeTwo
// 	WardrobeElevenSubTypeThree
// 	WardrobeElevenSubTypeFour
// )

// func (t WardrobeElevenSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeElevenSubType) Valid() bool {
// 	switch t {
// 	case WardrobeElevenSubTypeOne,
// 		WardrobeElevenSubTypeTwo,
// 		WardrobeElevenSubTypeThree,
// 		WardrobeElevenSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeElevenSubType(subType int32) WardrobeSubType {
// 	return WardrobeElevenSubType(subType)
// }

// //套装12子类型
// type WardrobeTwelveSubType int32

// const (
// 	WardrobeTwelveSubTypeOne WardrobeTwelveSubType = iota
// 	WardrobeTwelveSubTypeTwo
// 	WardrobeTwelveSubTypeThree
// 	WardrobeTwelveSubTypeFour
// )

// func (t WardrobeTwelveSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeTwelveSubType) Valid() bool {
// 	switch t {
// 	case WardrobeTwelveSubTypeOne,
// 		WardrobeTwelveSubTypeTwo,
// 		WardrobeTwelveSubTypeThree,
// 		WardrobeTwelveSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeTwelveSubType(subType int32) WardrobeSubType {
// 	return WardrobeTwelveSubType(subType)
// }

// //套装13子类型
// type WardrobeThirteenSubType int32

// const (
// 	WardrobeThirteenSubTypeOne WardrobeThirteenSubType = iota
// 	WardrobeThirteenSubTypeTwo
// 	WardrobeThirteenSubTypeThree
// 	WardrobeThirteenSubTypeFour
// )

// func (t WardrobeThirteenSubType) SubType() int32 {
// 	return int32(t)
// }

// func (t WardrobeThirteenSubType) Valid() bool {
// 	switch t {
// 	case WardrobeThirteenSubTypeOne,
// 		WardrobeThirteenSubTypeTwo,
// 		WardrobeThirteenSubTypeThree,
// 		WardrobeThirteenSubTypeFour:
// 		return true
// 	}
// 	return false
// }

// func CreateWardrobeThirteenSubType(subType int32) WardrobeSubType {
// 	return WardrobeThirteenSubType(subType)
// }

// var (
// 	wardrobeSubTypeFactoryMap = make(map[WardrobeType]WardrobeSubTypeFactory)
// )

// func CreateWardrobeSubType(typ WardrobeType, subType int32) WardrobeSubType {
// 	factory, ok := wardrobeSubTypeFactoryMap[typ]
// 	if !ok {
// 		return nil
// 	}
// 	return factory.CreateWardrobeSubType(subType)
// }

// func init() {
// 	wardrobeSubTypeFactoryMap[WardrobeTypeOne] = WardrobeSubTypeFactoryFunc(CreateWardrobeOneSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeTwo] = WardrobeSubTypeFactoryFunc(CreateWardrobeTwoSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeThree] = WardrobeSubTypeFactoryFunc(CreateWardrobeThreeSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeFour] = WardrobeSubTypeFactoryFunc(CreateWardrobeFourSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeFive] = WardrobeSubTypeFactoryFunc(CreateWardrobeFiveSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeSix] = WardrobeSubTypeFactoryFunc(CreateWardrobeSixSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeSeven] = WardrobeSubTypeFactoryFunc(CreateWardrobeSevenSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeEight] = WardrobeSubTypeFactoryFunc(CreateWardrobeEightSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeNine] = WardrobeSubTypeFactoryFunc(CreateWardrobeNineSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeTen] = WardrobeSubTypeFactoryFunc(CreateWardrobeTenSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeEleven] = WardrobeSubTypeFactoryFunc(CreateWardrobeElevenSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeTwelve] = WardrobeSubTypeFactoryFunc(CreateWardrobeTwelveSubType)
// 	wardrobeSubTypeFactoryMap[WardrobeTypeThirteen] = WardrobeSubTypeFactoryFunc(CreateWardrobeThirteenSubType)
// }
