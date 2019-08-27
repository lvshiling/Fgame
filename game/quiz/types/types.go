package types

type QuizAnswerType int32

const (
	//选项A
	QuizAnswerTypeA QuizAnswerType = iota + 1
	//选项B
	QuizAnswerTypeB
	//选项C
	QuizAnswerTypeC
	//选项D
	QuizAnswerTypeD
)

var (
	quizAnswerTypeMap = map[QuizAnswerType]string{
		QuizAnswerTypeA: "选项A",
		QuizAnswerTypeB: "选项B",
		QuizAnswerTypeC: "选项C",
		QuizAnswerTypeD: "选项D",
	}
)

func (spt QuizAnswerType) Valid() bool {
	switch spt {
	case QuizAnswerTypeA,
		QuizAnswerTypeB,
		QuizAnswerTypeC,
		QuizAnswerTypeD:
		return true
	}
	return false
}

func (spt QuizAnswerType) String() string {
	return quizAnswerTypeMap[spt]
}
