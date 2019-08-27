package worker

import "context"

//操作
type Operation interface {
	Context() context.Context
	Run() (result interface{}, err error)
	OnCallBack(result interface{}, err error)
}
type OperationFactory interface {
	CreateOperation(ctx context.Context, args ...interface{}) (op Operation)
}

type OperationFactoryFunc func(ctx context.Context, args ...interface{}) (op Operation)

func (off OperationFactoryFunc) CreateOperation(ctx context.Context, args ...interface{}) (op Operation) {
	return off(ctx, args...)
}

type BindOperation interface {
	Operation
	BindUUID() int64
}
