package api

type Usecase interface {
	SelectCount() (int, error)
	UpdateCount(count int) error
}
