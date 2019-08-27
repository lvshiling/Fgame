package types

type ZhenFaType int32

const (
	//蛇蟠
	ZhenFaTypeShePan ZhenFaType = 1 + iota
	//鸟翔
	ZhenFaTypeNiaoXiang
	//虎翼
	ZhenFaTypeHuYi
	//龙飞
	ZhenFaTypeLongFei
	//云垂
	ZhenFaTypeYunChui
	//风扬
	ZhenFaTypeFengYang
	//地载
	ZhenFaTypeDiZhai
	//天覆
	ZhenFaTypeTianFu
)

func (t ZhenFaType) Vaild() bool {
	switch t {
	case ZhenFaTypeShePan,
		ZhenFaTypeNiaoXiang,
		ZhenFaTypeHuYi,
		ZhenFaTypeLongFei,
		ZhenFaTypeYunChui,
		ZhenFaTypeFengYang,
		ZhenFaTypeDiZhai,
		ZhenFaTypeTianFu:
		return true
	}
	return false
}

const (
	ZhenFaTypeMin = ZhenFaTypeShePan
	ZhenFaTypeMax = ZhenFaTypeTianFu
)

type ZhenQiType int32

const (
	//阵基
	ZhenQiTypeZhenJi ZhenQiType = 1 + iota
	//阵心
	ZhenQiTypeZhenXin
	//阵旗
	ZhenQiTypeZhenQi
	//阵杆
	ZhenQiTypeZhenGan
	//阵纹
	ZhenQiTypeZhenWen
)

func (t ZhenQiType) Vaild() bool {
	switch t {
	case ZhenQiTypeZhenJi,
		ZhenQiTypeZhenXin,
		ZhenQiTypeZhenQi,
		ZhenQiTypeZhenGan,
		ZhenQiTypeZhenWen:
		return true
	}
	return false
}

const (
	ZhenQiTypeMin = ZhenQiTypeZhenJi
	ZhenQiTypeMax = ZhenQiTypeZhenWen
)
