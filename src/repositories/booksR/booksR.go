package booksR

import (
	"github.com/jinzhu/gorm"

	"go-api-rest/src/models"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*").Group("books.id")

	if err := query.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func GetBookByID(id string, db *gorm.DB) (models.Book, bool, error) {
	books := models.Book{}

	query := db.Select("books.*").Group("books.id")

	err := query.Where("books.id = ?", id).First(&books).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return books, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return books, false, nil
	}

	return books, true, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var books models.Book

	if err := db.Where("id = ?", id).Delete(&books).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, book *models.Book) error {
	if err := db.Save(&book).Error; err != nil {
		return err
	}
	return nil
}
