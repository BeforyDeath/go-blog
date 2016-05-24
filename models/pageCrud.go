package models

import (
	"errors"
	"github.com/beforydeath/go-blog/core"
)

func (pm *Pages) Validate(p *Page) error {
	if p.Name == "" {
		return errors.New("Field `Name` required")
	}
	if p.Alias == "" {
		return errors.New("Field `Alias` required")
	}
	if len(p.Description) < 20 {
		return errors.New("Field `Description` small")
	}
	return nil
}

func (pm *Pages) GetListAll() ([]*Page, error) {
	rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM element ORDER BY id DESC")
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

func (pm *Pages) GetById(id int) (*Page, error) {
	var p Page
	err := db.QueryRow("SELECT id, name, alias, description, visible, created_at FROM element WHERE id=?", id).Scan(
		&p.Id, &p.Name, &p.Alias, &p.Description, &p.Visible, &p.Created_at)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (pm *Pages) Create(p *Page) (int64, error) {

	err := pm.Validate(p)
	if err != nil {
		return 0, err
	}

	p.Preview, err = core.SplitPreview(p.Description)
	if err != nil {
		return 0, err
	}

	res, err := db.Exec("INSERT INTO element (name, alias, preview, description, visible, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		p.Name, p.Alias, p.Preview, p.Description, p.Visible, p.Created_at)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
