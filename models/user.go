// models/task.go

package models

import (
	"time"
)

type User struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    Fullname   string    `json:"fullname"`
    NumberPhone       string    `json:"NumberPhone"`
    Birthday   time.Time `json:"birthday"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
