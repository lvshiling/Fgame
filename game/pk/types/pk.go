package types

type PkState int32

const (
	//和平模式
	PkStatePeach PkState = iota + 1
	//组队
	PkStateGroup
	//帮派
	PkStateBangPai
	//全体
	PkStateAll
	//攻防
	PkStateCamp
	//阵营
	PkStateZhenYing
)

func (pkState PkState) Valid() bool {
	switch pkState {
	case PkStatePeach,
		PkStateGroup,
		PkStateBangPai,
		PkStateAll,
		PkStateCamp,
		PkStateZhenYing:
		return true
	}
	return false
}

func (pkState PkState) Mask() int32 {
	val := uint(pkState) - 1
	return 1 << val
}

var pkStateMap = map[PkState]string{
	//和平模式
	PkStatePeach: "和平模式",
	//组队
	PkStateGroup: "组队",
	//帮派
	PkStateBangPai: "帮派",
	//全体
	PkStateAll: "全体",
	//攻防
	PkStateCamp: "攻防",
	//攻防
	PkStateZhenYing: "阵营",
}

func (pkState PkState) String() string {
	return pkStateMap[pkState]
}

//红名状态
type PkRedState int32

const (
	//初始状态
	PkRedStateInit PkRedState = iota
	//pk值大于等于1
	PkRedStateFirst
	//pk值大于等于10
	PkRedStateSecond
	//pk值大于等于20
	PkRedStateThird
	//pk值大于等于30
	PkRedStateFourth
	//pk值大于等于40
	PkRedStateFifth
)

var (
	pkReadStateMap = map[PkRedState]string{
		PkRedStateInit:   "pk初始化",
		PkRedStateFirst:  "pk第一阶段",
		PkRedStateSecond: "pk第二阶段",
		PkRedStateThird:  "pk第三阶段",
		PkRedStateFourth: "pk第四阶段",
		PkRedStateFifth:  "pk第五阶段",
	}
)

func (s PkRedState) String() string {
	return pkReadStateMap[s]
}

var (
	//万分比
	pkReadStatePenalizeMap = map[PkRedState]int32{
		PkRedStateInit:   0,
		PkRedStateFirst:  1000,
		PkRedStateSecond: 2000,
		PkRedStateThird:  3000,
		PkRedStateFourth: 4000,
		PkRedStateFifth:  5000,
	}
)

func (s PkRedState) Penalize() int32 {
	return pkReadStatePenalizeMap[s]
}

var (
	pkReadStateValueMap = map[PkRedState]int32{
		PkRedStateInit:   0,
		PkRedStateFirst:  1,
		PkRedStateSecond: 10,
		PkRedStateThird:  20,
		PkRedStateFourth: 30,
		PkRedStateFifth:  40,
	}
)

func PkRedStateFromValue(val int32) PkRedState {
	if val < 0 {
		return PkRedStateInit
	}
	currentState := PkRedStateInit
	for state := PkRedStateInit; state <= PkRedStateFifth; state++ {
		if val < pkReadStateValueMap[state] {
			break
		}
		currentState = state
	}
	return currentState
}

type PkCamp interface {
	Camp() int32
}

type PkCommonCamp int32

const (
	PkCommonCampDefault PkCommonCamp = 0
)

func (c PkCommonCamp) Camp() int32 {
	return int32(c)
}
