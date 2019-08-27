package scene

// 参与的仙盟信息
type allianceData struct {
	allianceId int64
	totalForce int64
}

type allianceDataSort []*allianceData

func (a allianceDataSort) Len() int           { return len(a) }
func (a allianceDataSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a allianceDataSort) Less(i, j int) bool { return a[i].totalForce < a[j].totalForce }
