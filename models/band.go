package models

import "time"

type Band struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Albums    []Album   `json:"-"`
	Songs     []Song    `json:"-"`
}
