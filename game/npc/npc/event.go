package npc

type NPCHPChangedEventData struct {
	oldHP    int64
	newHP    int64
	attackId int64
}

func (d *NPCHPChangedEventData) GetOldHP() int64 {
	return d.oldHP
}

func (d *NPCHPChangedEventData) GetNewHP() int64 {
	return d.newHP
}

func (d *NPCHPChangedEventData) GetAttackId() int64 {
	return d.attackId
}

func CreateNPCHPChangedEventData(oldHp, newHp, attackId int64) *NPCHPChangedEventData {
	d := &NPCHPChangedEventData{
		oldHP:    oldHp,
		newHP:    newHp,
		attackId: attackId,
	}
	return d
}
