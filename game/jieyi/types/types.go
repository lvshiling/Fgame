package types

// 结义道具类型
type JieYiDaoJuType int32

const (
	JieYiDaoJuTypeLow  JieYiDaoJuType = iota // 低级道具
	JieYiDaoJuTypeHigh                       // 高级道具
)

func (t JieYiDaoJuType) Valid() bool {
	switch t {
	case JieYiDaoJuTypeLow,
		JieYiDaoJuTypeHigh:
		return true
	default:
		return false

	}
}

var (
	jieyiDaoJuMap = map[JieYiDaoJuType]string{
		JieYiDaoJuTypeLow:  "低级道具",
		JieYiDaoJuTypeHigh: "高级道具",
	}
)

func (t JieYiDaoJuType) String() string {
	return jieyiDaoJuMap[t]
}

// 结义信物类型
type JieYiTokenType int32

const (
	JieYiTokenTypeInvalid JieYiTokenType = iota - 1
	JieYiTokenTypeLow                    // 低级信物
	JieYiTokenTypeMiddle                 // 中级信物
	JieYiTokenTypeHigh                   // 高级信物
)

func (t JieYiTokenType) Valid() bool {
	switch t {
	case JieYiTokenTypeLow,
		JieYiTokenTypeMiddle,
		JieYiTokenTypeHigh:
		return true
	default:
		return false
	}
}

var jieyiTokenMap = map[JieYiTokenType]string{
	JieYiTokenTypeInvalid: "无效信物",
	JieYiTokenTypeLow:     "低级信物",
	JieYiTokenTypeMiddle:  "中级信物",
	JieYiTokenTypeHigh:    "高级信物",
}

func (t JieYiTokenType) String() string {
	return jieyiTokenMap[t]
}

type JieYiTokenChangeMethod int32

const (
	JieYiTokenChangeMethodBuChaJia  JieYiTokenChangeMethod = iota // 补差价方式替换
	JieYiTokenChangeMethodHighToken                               // 激活高级道具方式替换
)

func (t JieYiTokenChangeMethod) Valid() bool {
	switch t {
	case JieYiTokenChangeMethodBuChaJia,
		JieYiTokenChangeMethodHighToken:
		return true
	default:
		return false
	}
}

type JieYiRank int32

const (
	JieYiRankLaoDa JieYiRank = iota + 1
	JieYiRankLaoEr
	JieYiRankLaoSan
	JieYiRankLaoSi
	JieYiRankLaoWu
)

func (t JieYiRank) Valid() bool {
	switch t {
	case JieYiRankLaoDa,
		JieYiRankLaoEr,
		JieYiRankLaoSan,
		JieYiRankLaoSi,
		JieYiRankLaoWu:
		return true
	default:
		return false
	}
}

var RankNameMap = map[JieYiRank]string{
	JieYiRankLaoDa:  "老大",
	JieYiRankLaoEr:  "老二",
	JieYiRankLaoSan: "老三",
	JieYiRankLaoSi:  "老四",
	JieYiRankLaoWu:  "老五",
}

func (t JieYiRank) GetRankString() string {
	return RankNameMap[t]
}

// 邀请状态
type InviteState int32

const (
	InviteStateInit    InviteState = iota // 初始化
	InviteStateFail                       // 失败
	InviteStateSuccess                    // 成功
)

func (s InviteState) Valid() bool {
	switch s {
	case InviteStateInit,
		InviteStateFail,
		InviteStateSuccess:
		return true
	default:
		return false
	}
}

// 消耗道具的方式
type JieYiItemUseType int32

const (
	JieYiItemUseTypeInvite JieYiItemUseType = iota
	JieYiItemUseTypeTiHuan
	JieYiItemUseTypeGiveXiongDi
)

var JieYiItemUseTypeMap = map[JieYiItemUseType]string{
	JieYiItemUseTypeInvite:      "邀请消耗",
	JieYiItemUseTypeTiHuan:      "自身替换消耗",
	JieYiItemUseTypeGiveXiongDi: "赠送兄弟消耗",
}

func (t JieYiItemUseType) String() string {
	return JieYiItemUseTypeMap[t]
}
