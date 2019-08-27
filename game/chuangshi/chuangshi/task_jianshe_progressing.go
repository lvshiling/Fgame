package chuangshi

// import "time"

// const (
// 	jianSheTaskTime = time.Minute
// )

// type chengFangJianSheTask struct {
// 	s *chuangShiService
// }

// func (t *chengFangJianSheTask) Run() {

// 	for _, jianSheObj := range t.s.chengFangJianSheMap {
// 		if !jianSheObj.IfProgressing() {
// 			continue
// 		}

// 		t.s.syncChengFangJianShe(jianSheObj)
// 	}
// 	return

// }

// func (t *chengFangJianSheTask) ElapseTime() time.Duration {
// 	return jianSheTaskTime
// }

// func createChengFangJianSheTask(s *chuangShiService) *chengFangJianSheTask {
// 	t := &chengFangJianSheTask{
// 		s: s,
// 	}
// 	return t
// }
