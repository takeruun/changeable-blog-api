package entity

import "time"

type Blog struct {
	ID                 uint64     `json:"id" gorm:"primary_key"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	Body               string     `json:"body"`
	NgihtBody          string     `json:"nigth_body"`
	MobileBody         string     `json:"mobile_body"`
	ThumbnailImagePath string     `json:"thumbnail_image_path"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
	DeletedAt          *time.Time `json:"deletedAt"`
}
