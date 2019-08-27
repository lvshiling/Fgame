package types

import (
	"fmt"
)

//活动模板错误
type WelfareError struct {
	err          error
	activityName string
	groupId      int32
}

func (te *WelfareError) Error() string {
	return fmt.Sprintf("welfare: activityName[%s] groupId[%d]  error[%s]", te.activityName, te.groupId, te.err.Error())
}

func NewWelfareError(activityName string, groupId int32, err error) *WelfareError {
	te := &WelfareError{
		activityName: activityName,
		groupId:      groupId,

		err: err,
	}
	return te
}

//活动数据错误
type WelfareRecordError struct {
	err error
	id  int
}

func (te *WelfareRecordError) Error() string {
	if te.err == nil {
		return fmt.Sprintf("welfare: id[%d]", te.id)
	}
	return fmt.Sprintf("welfare: id[%d]  error[%s]", te.id, te.err.Error())
}
func NewWelfareRecordError(id int, err error) *WelfareRecordError {
	te := &WelfareRecordError{
		id:  id,
		err: err,
	}
	return te
}
