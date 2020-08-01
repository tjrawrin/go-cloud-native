package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Books ...
type Books []*Book

// Book ...
type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
}

// BookDtos ...
type BookDtos []*BookDto

// BookDto ...
type BookDto struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

// ToDto ...
func (b Book) ToDto() *BookDto {
	return &BookDto{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageURL,
		Description:   b.Description,
	}
}

// ToDto ...
func (bs Books) ToDto() BookDtos {
	dtos := make([]*BookDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

// BookForm ...
type BookForm struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

// ToModel ...
func (f *BookForm) ToModel() (*Book, error) {
	pubDate, err := time.Parse("2006-01-02", f.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}, nil
}
