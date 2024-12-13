package usecase

func (u *Usecase) SelectCount() (int, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return 0, err
	}

	if msg == 0 {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) UpdateCount(count int) error {
	isExist, err := u.p.CheckCountExistByMsg()
	if err != nil {
		return err
	}

	if isExist {
		err = u.p.UpdateCount(count)
		if err != nil {
			return err
		}
	}

	/*if isExist {
		return nil
	}

	err = u.p.UpdateCount(count)
	if err != nil {
		return err
	}*/

	return nil
}
