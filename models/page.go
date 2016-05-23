package models

import (
    "time"
    "errors"
    "strings"
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

func (m *Pages) split(s string) (string, error) {
    str := strings.Split(s, "{preview}")
    if len(str) == 1 {
        return "", errors.New("No splin descriptions")
    }
    return str[0], nil
}

func (m *Pages) validate(p *Page) error {
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

func (m *Pages) Create(p *Page) (int64, error) {

    err := m.validate(p);
    if err != nil {
        return 0, err
    }

    p.Preview, err = m.split(p.Description)
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

func (m *Pages) GetByAlias(alias string) (*Page, error) {
    var p Page
    err := db.QueryRow("SELECT id, name, alias, description, created_at FROM element WHERE alias=? AND visible=?", alias, true).Scan(
        &p.Id, &p.Name, &p.Alias, &p.Description, &p.Created_at)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

func (m *Pages) GetById(id int) (*Page, error) {
    var p Page
    err := db.QueryRow("SELECT id, name, alias, description, visible, created_at FROM element WHERE id=?", id).Scan(
        &p.Id, &p.Name, &p.Alias, &p.Description, &p.Visible, &p.Created_at)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

func (m *Pages) GetList() ([]*Page, error) {
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
