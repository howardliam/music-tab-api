package models

import "time"

type Song struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Tabs      []Tab     `json:"-"`
	BandID    uint      `json:"bandId"`
	AlbumID   uint      `json:"albumId"`
}
