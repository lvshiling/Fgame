package storage

//数据库数据
type Entity interface {
	//获取主keyId
	GetId() int64
	TableName() string
}

//内存数据
type PersistanceObject interface {
	//数据库id
	GetDBId() int64
	//转换为数据库实体
	ToEntity() (e Entity, err error)
	//转换从数据库实体
	FromEntity(e Entity) (err error)
	//提交修改
	SetModified()
}
