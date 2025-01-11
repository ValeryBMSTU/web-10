package usecase

type Provider interface {
	SelectUser(string) (string, error)
	UserIsExist(string) bool
	InsertUser(string, string) error
	UpdateUser(string, string) error
}
