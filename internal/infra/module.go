package infra

type Module interface {
	Init(infra *Application) error
}
