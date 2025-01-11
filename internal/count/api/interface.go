package api

type Usecase interface {
	GetCountApi() (string, error)
	SetCountApi(int) error
	IncrementCountApi(int) error 
}
