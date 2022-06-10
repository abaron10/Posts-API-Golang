package model

import "time"

type Post struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"created_by"`
	CreatedOn time.Time `json:"created_on"`
}
