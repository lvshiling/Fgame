package team

type TeamApplyData struct {
	applyId   int64
	applyTime int64
}

func NewTeamApplyData(applyId int64, applyTime int64) *TeamApplyData {
	data := &TeamApplyData{
		applyId:   applyId,
		applyTime: applyTime,
	}
	return data
}

func (tad *TeamApplyData) GetApplyId() int64 {
	return tad.applyId
}

//申请列表排序
type TeamApplyDataList []*TeamApplyData

func (tdl TeamApplyDataList) Len() int {
	return len(tdl)
}

func (tdl TeamApplyDataList) Less(i, j int) bool {
	return tdl[i].applyTime < tdl[j].applyTime
}

func (tdl TeamApplyDataList) Swap(i, j int) {
	tdl[i], tdl[j] = tdl[j], tdl[i]
}
