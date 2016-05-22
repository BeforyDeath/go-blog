package models

import (
    "time"
    "errors"
)

type Pages struct {
}

type Page struct {
    Id          int    `schema:"id"`
    Name        string `schema:"name"`
    Alias       string `schema:"alias"`
    Preview     string `schema:"preview"`
    Description string `schema:"description"`
    Created_at  time.Time `schema:"-"`
    Visible     bool `schema:"visible"`
}

func (self *Pages) Validate(p *Page) error {
    if p.Name == "" {
        return errors.New("Name not empty")
    }
    if p.Alias == "" {
        return errors.New("Alias not empty")
    }
    return nil
}

func (self *Pages) Create(p *Page) (int64, error) {

    err := self.Validate(p);
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

func (self *Pages) GetByAlias(alias string) (*Page, error) {
    var e Page
    err := db.QueryRow("SELECT id, name, alias, description, created_at FROM element WHERE alias=? AND visible=?", alias, true).Scan(
        &e.Id, &e.Name, &e.Alias, &e.Description, &e.Created_at)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

func (self *Pages) GetById(id int) (*Page, error) {
    var e Page
    err := db.QueryRow("SELECT id, name, alias, preview, description, visible, created_at FROM element WHERE id=?", id).Scan(
        &e.Id, &e.Name, &e.Alias, &e.Preview, &e.Description, &e.Visible, &e.Created_at)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

func (self *Pages) GetList() ([]*Page, error) {
    rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM element WHERE visible=?", 1)
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
