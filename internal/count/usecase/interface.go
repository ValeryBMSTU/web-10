package usecase

type Provider interface {
	SelectCount() (int, error)
	CheckCountExistByMsg() (bool, error)
	UpdateCount(count int) error
}
