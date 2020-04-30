package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  m "github.com/choonggg/gotodo/models"
)

type Env struct {
  DB *gorm.DB
}

func New() *Env {
  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect DB")
  }

  env := Env{DB: db}

  db.AutoMigrate(&m.Todo{})

  return &env
}
