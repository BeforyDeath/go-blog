package models

import (
	"time"
)

type Page struct {
	Id          int       `schema:"id"`
	Name        string    `schema:"name"`
	Alias       string    `schema:"alias"`
	Preview     string    `schema:"preview"`
	Description string    `schema:"description"`
	Created_at  time.Time `schema:"-"`
	Visible     bool      `schema:"visible"`
}

type Pages struct {
}

func (pm *Pages) GetList() ([]*Page, error) {
	rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM element WHERE visible=? ORDER BY id DESC", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*Page, 0)

	for rows.Next() {
		e := new(Page)
		err := rows.Scan(&e.Id, &e.Name, &e.Alias, &e.Preview, &e.Created_at)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (pm *Pages) GetByAlias(alias string) (*Page, error) {
	var p Page
	err := db.QueryRow("SELECT id, name, alias, description, created_at FROM element WHERE alias=? AND visible=?", alias, true).Scan(
		&p.Id, &p.Name, &p.Alias, &p.Description, &p.Created_at)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
