package models

type Posts struct {
}

type Post struct {
    id   int
    name string
}

func (p *Posts) GetList() ([]*Post, error) {

    rows, err := db.Query("SELECT id, name FROM post")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    posts := make([]*Post, 0)

    for rows.Next() {
        post := new(Post)
        err := rows.Scan(&post.id, &post.name)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }
    return posts, nil
}
