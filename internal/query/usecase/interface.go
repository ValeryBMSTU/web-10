package usecase

type Provider interface {
	SelectNameQuery(name string) (string, error)
	CheckNameQueryExistByMsg(name string) (bool, error)
	InsertNameQuery(name string) error
}
