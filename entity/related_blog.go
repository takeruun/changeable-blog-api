package entity

import "time"

type RelatedBlog struct {
	ID          uint64     `json:"id" gorm:"primary_key"`
	Blog        Blog       `json:"blog"`
	RelatedBlog Blog       `json:"related_blog"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
