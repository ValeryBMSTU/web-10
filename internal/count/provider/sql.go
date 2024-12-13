package provider

import (
	"database/sql"
	"errors"
	"fmt"
)

func (p *Provider) SelectCount() (int, error) {
	var count int

	err := p.conn.QueryRow("SELECT count FROM counters WHERE id = 1").Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

func (p *Provider) CheckCountExistByMsg() (bool, error) {
	var count int
	err := p.conn.QueryRow("SELECT count FROM counters WHERE id = 1").Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) UpdateCount(count int) error {
	_, err := p.conn.Exec("UPDATE counters SET count = count + $1 WHERE id = 1", count)
	if err != nil {
		return err
	}

	fmt.Println("Count updated successfully!")
	return nil
}
