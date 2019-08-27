package types

// 路径类型
type PathType int32

const (
	PathTypeStill                  PathType = iota //不动
	PathTypeForward                                //前进
	PathTypeBack                                   //后退
	PathTypeBackTimesEnoughForward                 //因为后退次数足够触发的前进
)

func (t PathType) Valid() bool {
	switch t {
	case PathTypeStill,
		PathTypeForward,
		PathTypeBack,
		PathTypeBackTimesEnoughForward:
		return true
	default:
		return false
	}
}

// 奖池结构
type RewNode struct {
	level    int32   //等级
	rateList []int64 //概率列表
}

func (node *RewNode) GetLevel() int32 {
	return node.level
}

func (node *RewNode) GetRateList() []int64 {
	return node.rateList
}

func CreateRewNode(level, forwardRate, backRate, stillRate int32) *RewNode {
	ratelist := []int64{}
	if stillRate >= 0 {
		ratelist = append(ratelist, int64(stillRate))
	}
	if forwardRate >= 0 {
		ratelist = append(ratelist, int64(forwardRate))
	}
	if backRate >= 0 {
		ratelist = append(ratelist, int64(backRate))
	}
	d := &RewNode{
		level:    level,
		rateList: ratelist,
	}
	return d
}

type RewPools map[int32]*RewNode

func CreateRewPools(rewlist []*RewNode) (pool RewPools) {
	pool = make(map[int32]*RewNode)
	for _, node := range rewlist {
		pool[node.level] = node
	}
	return pool
}
