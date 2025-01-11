package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) GetCountSql() (string, error) {
	var msg string


	err := p.conn.QueryRow("SELECT count FROM countdb").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
}

func (p *Provider) SetCountSql(num int) error {

	_, err := p.conn.Exec("UPDATE countdb SET count = $1", num)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) IncrementCountSql(num int) error {

	_, err := p.conn.Exec("UPDATE countdb SET count = count + $1", num)
	if err != nil {
		return err
	}

	return nil
}