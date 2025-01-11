package api

type Usecase interface {
	FetchHelloMessage(string) (string, error)
	SetHelloMessage(string, string) error
}
