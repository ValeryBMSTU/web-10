package api

type Usecase interface {
	SelectNameQuery(name string) (string, error)
	InsertNameQuery(name string) error
}
