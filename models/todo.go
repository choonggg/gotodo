package models

import (
  "time"
)

type Todo struct {
  ID int64 `gorm:"primary_key;auto_increment" json:"id"`
  Slug string `gorm:"size:255;not null" json:"slug"`
  Title string `gorm:"size:255;not null" json:"title"`
  CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
  UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
