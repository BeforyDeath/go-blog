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

func (pm *Pages) GetTotal() (count int, err error) {
	rows, err := db.Query("SELECT count(*) as count FROM page")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return
		}
	}
	return
}

func (pm *Pages) GetList(offset, limit int) ([]*Page, error) {
	rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM page ORDER BY id DESC LIMIT ?, ?", offset, limit)
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
	err := db.QueryRow("SELECT id, name, alias, description, created_at FROM page WHERE alias=?", alias).Scan(
		&p.Id, &p.Name, &p.Alias, &p.Description, &p.Created_at)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
