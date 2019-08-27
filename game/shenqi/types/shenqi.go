package types

type ShenQiType int32

const (
	//轩辕剑
	ShenQiTypeXuanYuanJian ShenQiType = 1 + iota
	//盘古斧
	ShenQiTypePanGuFu
	//炼妖壶
	ShenQiTypeLianYaoHu
	//昊天塔
	ShenQiTypeHaoTianTa
	//伏羲琴
	ShenQiTypeFuXiQin
	//神农鼎
	ShenQiTypeShenNongDing
	//崆峒印
	ShenQiTypeKongTongYin
	//昆仑镜
	ShenQiTypeKunLunJing
	//女娲石
	ShenQiTypeNvWaShi
	//东皇钟
	ShenQiTypeDongHuangZhong
)

const (
	MinShenQiType = ShenQiTypeXuanYuanJian
	MaxShenQiType = ShenQiTypeDongHuangZhong
)

var (
	shenQiTypeMap = map[ShenQiType]string{
		ShenQiTypeXuanYuanJian:   "轩辕剑",
		ShenQiTypePanGuFu:        "盘古斧",
		ShenQiTypeLianYaoHu:      "炼妖壶",
		ShenQiTypeHaoTianTa:      "昊天塔",
		ShenQiTypeFuXiQin:        "伏羲琴",
		ShenQiTypeShenNongDing:   "神农鼎",
		ShenQiTypeKongTongYin:    "崆峒印",
		ShenQiTypeKunLunJing:     "昆仑镜",
		ShenQiTypeNvWaShi:        "女娲石",
		ShenQiTypeDongHuangZhong: "东皇钟",
	}
)

func (spt ShenQiType) Valid() bool {
	switch spt {
	case ShenQiTypeXuanYuanJian,
		ShenQiTypePanGuFu,
		ShenQiTypeLianYaoHu,
		ShenQiTypeHaoTianTa,
		ShenQiTypeFuXiQin,
		ShenQiTypeShenNongDing,
		ShenQiTypeKongTongYin,
		ShenQiTypeKunLunJing,
		ShenQiTypeNvWaShi,
		ShenQiTypeDongHuangZhong:
		return true
	}
	return false
}

func (spt ShenQiType) String() string {
	return shenQiTypeMap[spt]
}
