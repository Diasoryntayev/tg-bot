package files

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}
