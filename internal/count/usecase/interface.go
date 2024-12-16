package usecase

type Provider interface {
	FetchCount() (int, error)
	UpdateCount(count int) (int, error)
	CheckCountExist() (bool, error)
}
