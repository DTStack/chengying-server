package schema

const (
	PATCH_FILE = "patch"
)

type DiffService struct {
	ServiceName  string
	NewFiles     []string
	DiffFiles    []string
	DeletedFiles []string
}

type Patch struct {
	ProductName       string
	NewProductVersion string
	OldProductVersion string
	NewServices       []string
	DeletedServices   []string
	DiffServices      []*DiffService
}

func (p Patch) IsDeletedService(name string) bool {
	for _, deleteSvcName := range p.DeletedServices {
		if deleteSvcName == name {
			return true
		}
	}
	return false
}

func (p Patch) IsDiffService(name string) bool {
	for _, diffSvc := range p.DiffServices {
		if diffSvc.ServiceName == name {
			return true
		}
	}
	return false
}

func (p Patch) IsChangedService(name string) bool {
	return p.IsDeletedService(name) || p.IsDiffService(name)
}
