package models

import "time"

type Tab struct {
	Id         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Title      string    `json:"title"`
	Instrument string    `json:"instrument"`
	Tuning     string    `json:"tuning"`
	Content    string    `json:"content"`
	SongID     uint      `json:"songId"`
}
