package models

import (
  "time"
  "log"
)

type Todo struct {
  ID int
  Content string
  UserID int
  CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
  cmd := `insert into todos (
    content,
    user_id,
    created_at) values (?, ?, ?)`

    _, err = Db.Exec(cmd, content, u.ID, time.Now())
    if err != nil {
      log.Fatalln(err)
    }
    return err
}

func GetTodo(id int) (todo Todo, err error) {
  cmd := `SELECT * FROM TODOS WHERE ID = ?`
  todo = Todo{}
  err = Db.QueryRow(cmd, id).Scan(
    &todo.ID,
    &todo.Content,
    &todo.UserID,
    &todo.CreatedAt,
  )

  return todo, err
}