package pg

type storage struct {
}

func NewStorage(dbURI string) (storage, error) {
	s := storage{}

	return s, nil
}
