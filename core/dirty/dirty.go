package dirty

type DirtyService interface {
	IsLegal(name string) bool
}

type dirtyService struct {
}

func (ds *dirtyService) IsLegal(name string) bool {
	return true
}

var (
	ds = &dirtyService{}
)

func GetDirtyService() DirtyService {
	return ds
}
