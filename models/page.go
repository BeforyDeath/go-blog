package models

import (
	"fmt"
	"github.com/BeforyDeath/pagination"
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
	Pagination *pagination.Pagination
}

func (pm *Pages) GetTotal() (err error) {
	if pm.Pagination == nil {
		pm.Pagination = pagination.Create(0, 1, 5)

		fmt.Println("init pagination")

		rows, err := db.Query("SELECT count(*) as count FROM page")
		if err != nil {
			return err
		}
		defer rows.Close()

		var count int
		for rows.Next() {
			err := rows.Scan(&count)
			if err != nil {
				return err
			}
		}
		pm.Pagination.SetTotal(count)
		return nil
	}
	return nil
}

func (pm *Pages) GetList() ([]*Page, error) {
	rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM page ORDER BY id DESC")
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
