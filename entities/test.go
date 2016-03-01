package entities

import "time"

type Test struct {
    Id int64        `db:"id" json:"id"`
    Name string     `db:"name" json:"name"`
    Token string    `db:"token" json:"token"`
    Created int64   `db:"created_at" json:"created_at"`
    Updated int64   `db:"updated_at" json:"updated_at"`
}

func newTest(name string, token string) Test {
    return Test{
        Name: name,
        Token: token,
        Created: time.Now().UnixNano(),
    }
}