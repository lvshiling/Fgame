package chuangshi

// import "time"

// const (
// 	signTaskTime = time.Minute
// )

// type shenWangSignUpTask struct {
// 	s *chuangShiService
// }

// func (t *shenWangSignUpTask) Run() {
// 	for _, signUpObj := range t.s.shenWangSignUpMap {
// 		if !signUpObj.IfSigning() {
// 			continue
// 		}

// 		t.s.syncShenWangSignup(signUpObj)
// 	}
// 	return
// }

// func (t *shenWangSignUpTask) ElapseTime() time.Duration {
// 	return signTaskTime
// }

// func createShenWangSignUpTask(s *chuangShiService) *shenWangSignUpTask {
// 	t := &shenWangSignUpTask{
// 		s: s,
// 	}
// 	return t
// }
