package feedbackfee

import "time"

const (
	expireTaskTime = time.Minute
)

type codeExpireTask struct {
	s FeedbackFeeService
}

func (t *codeExpireTask) Run() {
	t.s.CodeExpireCheck()
	return

}

func (t *codeExpireTask) ElapseTime() time.Duration {
	return expireTaskTime
}

func createCodeExpireTask(s FeedbackFeeService) *codeExpireTask {
	t := &codeExpireTask{
		s: s,
	}
	return t
}
