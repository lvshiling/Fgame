package common

//结义信息
type PlayerJieYiObject interface {
	GetJieYiName() string
	GetJieYiRank() int32
	GetJieYiId() int64
}

type playerJieYiObject struct {
	jieYiId   int64
	jieYiName string
	rank      int32
}

func (o *playerJieYiObject) GetJieYiId() int64 {
	return o.jieYiId
}

func (o *playerJieYiObject) GetJieYiName() string {
	return o.jieYiName
}

func (o *playerJieYiObject) GetJieYiRank() int32 {
	return o.rank
}

func CreatePlayerJieYiObject(jieYiId int64, jieYiName string, rank int32) PlayerJieYiObject {
	obj := &playerJieYiObject{}
	obj.jieYiId = jieYiId
	obj.jieYiName = jieYiName
	obj.rank = rank
	return obj
}
