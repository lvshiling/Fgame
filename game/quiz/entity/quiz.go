package entity

//仙尊答题数据
type QuizEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	ServerId     int32  `gorm:"column:serverId"`
	LastQuizId   int32  `gorm:"column:lastQuizId"`
	LastQuizTime int64  `gorm:"column:lastQuizTime"`
	AnswerList   string `gorm:"column:answerList"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *QuizEntity) GetId() int64 {
	return e.Id
}

func (e *QuizEntity) TableName() string {
	return "t_quiz"
}
