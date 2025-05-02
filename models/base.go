package models

import "time"

type (
	BaseEntity struct {
		ID        string    `json:"id" gorm:"primary_key"`
		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	}

	FileEntity struct {
		BaseEntity
		FileName string `json:"file_name" gorm:"not null"`
		FileURL  string `json:"file_url" gorm:"not null"`
	}
)
