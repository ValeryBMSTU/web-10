package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectUser(ip string) (string, error) {
	var msg string

	// Получаем одно сообщение из таблицы hello, отсортированной в случайном порядке
	err := p.conn.QueryRow("SELECT name_user FROM users WHERE ip_address = $1", ip).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
}

func (p *Provider) UserIsExist(ip string) bool {
	// Получаем одно сообщение из таблицы hello

	err := p.conn.QueryRow("SELECT name_user FROM users WHERE ip_address = $1", ip).Scan(&ip)
	if err != nil {
		return false
	}

	return true
}

func (p *Provider) InsertUser(msg, ip string) error {
	_, err := p.conn.Exec("INSERT INTO users (name_user, ip_address) VALUES ($1, $2)", msg, ip)
	if err != nil {
		return err
	}

	return nil
}


func (p *Provider) UpdateUser(msg, ip string) error {
	_, err := p.conn.Exec("UPDATE users SET name_user = $1 WHERE ip_address = $2", msg, ip)
	if err != nil {
		return err
	}

	return nil
}