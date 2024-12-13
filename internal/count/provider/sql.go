package provider

func (p *Provider) GetCount() (int, error) {
	var value int

	row := p.conn.QueryRow("SELECT COALESCE(count, 0) FROM count WHERE name=$1", "key1")
	err := row.Scan(&value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (p *Provider) AddCount(a int) error {
	_, err := p.conn.Exec("INSERT INTO count (name, count) VALUES ($2, $1) ON CONFLICT (name) DO UPDATE SET count = count.count + $1", a, "key1")
	if err != nil {
		return err
	}

	return nil
}
