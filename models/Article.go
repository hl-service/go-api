package models

import (
	"github.com/Kamva/mgm/v2"
)

// Article Model structure
type Article struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
}

// CreateArticle creates new instance of model
func CreateArticle(title, description string) *Article {
	return &Article{
		Title:       title,
		Description: description,
	}
}
