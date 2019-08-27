package template

type QuestTemplateVO struct {
	//id
	Id int `json:"id"`
	//任务名称
	Name string `json:"name"`
	//任务类型
	QuestType int32 `json:"quest_type"`
	//任务文本
	Objectives string `json:"objectives"`
	//前置任务
	PrevQuest string `json:"prev_quest"`
	//后置任务
	NextQuest string `json:"next_quest"`
}

func (m *QuestTemplateVO) IsOnceQuest() bool {
	return m.QuestType == int32(1)
}
