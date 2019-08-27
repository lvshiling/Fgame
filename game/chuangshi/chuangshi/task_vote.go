package chuangshi

// import "time"

// const (
// 	voteTaskTime = time.Minute
// )

// type shenWangVoteTask struct {
// 	s *chuangShiService
// }

// func (t *shenWangVoteTask) Run() {
// 	for _, voteObj := range t.s.shenWangVoteMap {
// 		if !voteObj.IfVoting() {
// 			continue
// 		}

// 		t.s.syncShenWangVote(voteObj)
// 	}
// 	return

// }

// func (t *shenWangVoteTask) ElapseTime() time.Duration {
// 	return voteTaskTime
// }

// func createShenWangVoteTask(s *chuangShiService) *shenWangVoteTask {
// 	t := &shenWangVoteTask{
// 		s: s,
// 	}
// 	return t
// }
