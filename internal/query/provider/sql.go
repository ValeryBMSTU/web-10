package provider

import (
	"database/sql"
	"errors"
	"fmt"
)

func (p *Provider) SelectNameQuery(name string) (string, error) {
	var msg string

	// Получаем одно сообщение из таблицы
	err := p.conn.QueryRow("SELECT name FROM greetings WHERE name = $1", name).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return "Hello, " + msg + "!", nil
}

func (p *Provider) CheckNameQueryExistByMsg(name string) (bool, error) {
	var msg string
	// Получаем одно сообщение из таблицы
	err := p.conn.QueryRow("SELECT name FROM greetings WHERE name = $1", name).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) InsertNameQuery(name string) error {
	_, err := p.conn.Exec("INSERT INTO greetings (name) VALUES ($1)", name)
	if err != nil {
		return err
	}

	fmt.Println("Inserted successfully!")

	return nil
}
