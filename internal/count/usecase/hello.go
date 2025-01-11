package usecase

func (u *Usecase) GetCountApi() (string, error) {
	msg, err := u.p.GetCountSql()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetCountApi(num int) error {
	err := u.p.SetCountSql(num)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) IncrementCountApi(num int) error {
	err := u.p.IncrementCountSql(num)
	if err != nil {
		return err
	}

	return nil
}

