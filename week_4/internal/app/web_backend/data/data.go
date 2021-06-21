package data

type DataRepo struct {
	db interface{}
}

func NewDataRepo() (*Data, error)  {
	d := &Data{
		db: "",
	}
	return d, nil
}
