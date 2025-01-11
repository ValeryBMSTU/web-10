package usecase

func (u *Usecase) FetchHelloMessage(ip string) (string, error) {
	msg, err := u.p.SelectUser(ip)
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return "Hello, " + msg + "!", nil
}

func (u *Usecase) SetHelloMessage(msg, ip string) error {
	if u.p.UserIsExist(ip) {
		err := u.p.UpdateUser(msg, ip)
		if err != nil {
			return err
		}

		return nil
	}else{
		err := u.p.InsertUser(msg, ip)
		if err != nil {
			return err
		}

		return nil
	}

	
}
