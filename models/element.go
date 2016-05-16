package models

type Elements struct {
}

type element struct {
    Id          int
    Alias       string
    Name        string
    Preview     string
    Description string
    Creates_at  string
    Visible     bool
}

func (m *Elements) GetList() ([]*element, error) {

    rows, err := db.Query("SELECT id, name FROM element")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make([]*element, 0)

    for rows.Next() {
        e := new(element)
        err := rows.Scan(&e.Id, &e.Name)
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
