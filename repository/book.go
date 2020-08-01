package repository

import (
	"go-cloud-native/model"

	"github.com/jinzhu/gorm"
)

// ListBooks ...
func ListBooks(db *gorm.DB) (model.Books, error) {
	books := make([]*model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// ReadBook ...
func ReadBook(db *gorm.DB, id uint) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// DeleteBook ...
func DeleteBook(db *gorm.DB, id uint) error {
	book := &model.Book{}
	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

// CreateBook ...
func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// UpdateBook ...
func UpdateBook(db *gorm.DB, book *model.Book) error {
	if err := db.First(&model.Book{}, book.ID).Update(book).Error; err != nil {
		return err
	}

	return nil
}
