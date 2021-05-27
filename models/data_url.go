package models

import (
	_ "github.com/jinzhu/gorm"
)

type DataUrl struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	ShortUrl string `json:"short_url"`
	Alias    string `json:"alias"`
	Query    string `json:"query"`
}

type CreateDataUrl struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
	Query string `json:"query"`
}

type UpdateDataUrl struct {
	Title    string `json:"title"`
	Alias    string `json:"alias"`
	Query    string `json:"query"`
	ShortUrl string `json:"short_url"`
}
