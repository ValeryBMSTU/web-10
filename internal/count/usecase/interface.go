package usecase

type Provider interface {
	GetCountSql() (string, error)
	SetCountSql(int) error
	IncrementCountSql(int) error
}
