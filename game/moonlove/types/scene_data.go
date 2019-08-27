package types

//排行榜数据
type RankData struct {
	PlayerId int64
	Name     string
	Number   int32
}

type RankDataList []*RankData

func (c RankDataList) Len() int {
	return len(c)
}

func (crl RankDataList) Less(i, j int) bool {
	if crl[i].Number == 0 {
		return true
	}
	if crl[j].Number == 0 {
		return false
	}
	return crl[i].Number < crl[j].Number
}

func (c RankDataList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//双人赏月数据
type MoonloveDoubleData struct {
	PlayerId      int64
	OhtherPayerId int64
}
