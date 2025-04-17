package infra

type Module interface {
	Init(infra *Infra) error
}
