package feedbackfee

import "time"

const (
	codeTaskTime = time.Minute
)

type codeGenerateTask struct {
	s FeedbackFeeService
}

func (t *codeGenerateTask) Run() {
	t.s.CodeGenerate()
	return

}

func (t *codeGenerateTask) ElapseTime() time.Duration {
	return codeTaskTime
}

func createCodeGenerateTask(s FeedbackFeeService) *codeGenerateTask {
	t := &codeGenerateTask{
		s: s,
	}
	return t
}
