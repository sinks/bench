package bench

type BenchSQL interface {
	CreateTable() error
	SaveOrUpdate() error
	Save() error
	Update() error
	Id() string
}
