package types

//城战阵营
type ChuangShiCampType int32

const (
	ChuangShiCampTypeNone ChuangShiCampType = iota
	ChuangShiCampTypePanGu
	ChuangShiCampTypeNvWai
	ChuangShiCampTypeFuxi
)

var (
	ChuangShiCampTypeMap = map[ChuangShiCampType]string{
		ChuangShiCampTypeNone:  "中立",
		ChuangShiCampTypePanGu: "盘古",
		ChuangShiCampTypeNvWai: "女娲",
		ChuangShiCampTypeFuxi:  "伏羲",
	}
)

func (t ChuangShiCampType) Valid() bool {
	switch t {
	case ChuangShiCampTypeNone,
		ChuangShiCampTypePanGu,
		ChuangShiCampTypeNvWai,
		ChuangShiCampTypeFuxi:
		return true
	}
	return false
}

func (t ChuangShiCampType) String() string {
	return ChuangShiCampTypeMap[t]
}

func RandomChuangShiCamp() ChuangShiCampType {
	return ChuangShiCampTypeNone
	// index := mathutils.RandomWeights([]int64{1, 1, 1, 1})
	// return ChuangShiCampType(index)
}

// 城市类型
type ChuangShiCityType int32

const (
	ChuangShiCityTypeZhongli ChuangShiCityType = iota
	ChuangShiCityTypeMain
	ChuangShiCityTypeFushu
)

var (
	ChuangShiCityTypeMap = map[ChuangShiCityType]string{
		ChuangShiCityTypeZhongli: "中立城",
		ChuangShiCityTypeMain:    "主城",
		ChuangShiCityTypeFushu:   "附属城",
	}
)

func (t ChuangShiCityType) Valid() bool {
	switch t {
	case ChuangShiCityTypeMain,
		ChuangShiCityTypeFushu,
		ChuangShiCityTypeZhongli:
		return true
	}
	return false
}

func (t ChuangShiCityType) String() string {
	return ChuangShiCityTypeMap[t]
}

//官职
type ChuangShiGuanZhi int32

const (
	ChuangShiGuanZhiNone     ChuangShiGuanZhi = iota //默认
	ChuangShiGuanZhiPingMing                         //平民
	ChuangShiGuanZhiChengZhu                         //城主
	ChuangShiGuanZhiShenWang                         //神王
)

func (t ChuangShiGuanZhi) Valid() bool {
	switch t {
	case ChuangShiGuanZhiNone,
		ChuangShiGuanZhiPingMing,
		ChuangShiGuanZhiChengZhu,
		ChuangShiGuanZhiShenWang:
		return true
	}
	return false
}

func RandomChuangShiGuanZhi() ChuangShiGuanZhi {
	return ChuangShiGuanZhiNone
	// index := mathutils.RandomWeights([]int64{1, 1, 1, 1})
	// return ChuangShiGuanZhi(index)
}

//报名状态
type ShenWangSignUpType int32

const (
	ShenWangSignUpTypeNone    ShenWangSignUpType = iota //未报名
	ShenWangSignUpTypeSigning                           //报名中
	ShenWangSignUpTypeSuccess                           //投票成功
)

//投票状态
type ShenWangVoteType int32

const (
	ShenWangVoteTypeNone    ShenWangVoteType = iota //未投票
	ShenWangVoteTypeVoting                          //投票中
	ShenWangVoteTypeSuccess                         //投票成功
)

//神王选举状态
type ShenWangStatusType int32

const (
	ShenWangStatusTypeSign ShenWangStatusType = iota //报名阶段
	ShenWangStatusTypeVote                           //投票阶段
	ShenWangStatusTypeEnd                            //选举结束
)

//城防建设状态
type ChengFangStatusType int32

const (
	ChengFangStatusTypeNone        ChengFangStatusType = iota //未建设
	ChengFangStatusTypeProgressing                            //建设中
	ChengFangStatusTypeSuccess                                //建设成功
)

// 城池建设类型
type ChuangShiCityJianSheType int32

const (
	ChuangShiCityJianSheTypeTianQi    ChuangShiCityJianSheType = iota //天气台
	ChuangShiCityJianSheTypeTunTian                                   //屯田
	ChuangShiCityJianSheTypeTieJian                                   //铁匠铺
	ChuangShiCityJianSheTypeYiGuan                                    //医馆
	ChuangShiCityJianSheTypeYiZhan                                    //驿站
	ChuangShiCityJianSheTypeJiaoChang                                 //校场
	ChuangShiCityJianSheTypeJunYing                                   //军营
)

var (
	CityJianSheTypeMap = map[ChuangShiCityJianSheType]string{
		ChuangShiCityJianSheTypeTianQi:    "天气台",
		ChuangShiCityJianSheTypeTunTian:   "屯田",
		ChuangShiCityJianSheTypeTieJian:   "铁匠铺",
		ChuangShiCityJianSheTypeYiGuan:    "医馆",
		ChuangShiCityJianSheTypeYiZhan:    "驿站",
		ChuangShiCityJianSheTypeJiaoChang: "校场",
		ChuangShiCityJianSheTypeJunYing:   "军营",
	}
)

func (t ChuangShiCityJianSheType) Valid() bool {
	switch t {
	case ChuangShiCityJianSheTypeTianQi,
		ChuangShiCityJianSheTypeTunTian,
		ChuangShiCityJianSheTypeTieJian,
		ChuangShiCityJianSheTypeYiGuan,
		ChuangShiCityJianSheTypeYiZhan,
		ChuangShiCityJianSheTypeJiaoChang,
		ChuangShiCityJianSheTypeJunYing:
		return true
	}
	return false
}

func (t ChuangShiCityJianSheType) String() string {
	return CityJianSheTypeMap[t]
}
