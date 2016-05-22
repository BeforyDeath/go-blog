package models

import (
    "time"
)

type Pages struct {
}

type Page struct {
    Id          int
    Alias       string
    Name        string
    Preview     string
    Description string
    Created_at  time.Time
    Visible     bool
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
