package models

import (
    "time"
)

type Elements struct {
}

type element struct {
    Id          int
    Alias       string
    Name        string
    Preview     string
    Description string
    Created_at  time.Time
    Visible     bool
}

func (m *Elements) GetByAlias(alias string) (*element, error) {
    var e element
    err := db.QueryRow("SELECT id, name, alias, description, created_at FROM element WHERE alias=? AND visible=?", alias, true).Scan(
        &e.Id, &e.Name, &e.Alias, &e.Description, &e.Created_at)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

func (m *Elements) GetList() ([]*element, error) {
    rows, err := db.Query("SELECT id, name, alias, preview, created_at FROM element WHERE visible=?", 1)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make([]*element, 0)

    for rows.Next() {
        e := new(element)
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
