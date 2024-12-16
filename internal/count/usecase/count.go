package usecase

func (u *Usecase) FetchCount() (int, error) {
	msg, err := u.p.FetchCount()
	if err != nil {
		return 0, err
	}

	if msg == 0 {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) IncrementCount(count int) (int, error) {
	isExist, err := u.p.CheckCountExist()
	if err != nil {
		return 0, err
	}

	if !isExist {
		return 0, nil
	}

	newCount, err := u.p.UpdateCount(count)
	if err != nil {
		return 0, err
	}

	return newCount, nil
}
