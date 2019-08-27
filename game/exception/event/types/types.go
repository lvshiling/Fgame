package types

type ExceptionEventType string

const (
	//异常
	ExceptionEventTypeException   ExceptionEventType = "Exception"
	ExceptionEventTypeDBException                    = "DBException"
)

type DBExceptionEventData struct {
	tableName string
	data      interface{}
	err       string
}

func (d *DBExceptionEventData) GetTableName() string {
	return d.tableName
}

func (d *DBExceptionEventData) GetData() interface{} {
	return d.data
}

func (d *DBExceptionEventData) GetError() string {
	return d.err
}

func CreateDBExceptionEventData(tableName string, data interface{}, err string) *DBExceptionEventData {
	eventData := &DBExceptionEventData{}
	eventData.tableName = tableName
	eventData.data = data
	eventData.err = err
	return eventData
}
