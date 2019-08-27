package job

type IJob interface {
	GetId() string        //作业标志
	Run() error           //作业运行的方法
	GetTickSecond() int64 //间隔时间
}
