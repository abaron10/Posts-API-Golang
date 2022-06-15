package models

import (
	"time"
)

type Post struct {
	Id        interface{} `json:"id"`
	Title     string      `json:"title"`
	Text      string      `json:"text"`
	CreatedBy string      `json:"created_by"`
	CreatedOn time.Time   `json:"created_on"`
}
