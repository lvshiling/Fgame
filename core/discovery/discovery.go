package discovery

//服务发现
type Discovery interface {
	//注册服务
	RegisterService(service string, id string, target string) (err error)
	GetService(service string) (target string, err error)
	GetServiceWithId(service string, id string) (target string, err error)
}
