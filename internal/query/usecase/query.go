package usecase

func (u *Usecase) SelectNameQuery(name string) (string, error) {
	msg, err := u.p.SelectNameQuery(name)
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) InsertNameQuery(name string) error {
	isExist, err := u.p.CheckNameQueryExistByMsg(name)
	if err != nil {
		return err
	}

	if isExist {
		return nil
	}

	err = u.p.InsertNameQuery(name)
	if err != nil {
		return err
	}

	return nil
}
