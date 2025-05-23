// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type BookCategory string

const (
	BookCategoryFiction    BookCategory = "fiction"
	BookCategoryNonFiction BookCategory = "non_fiction"
	BookCategoryScience    BookCategory = "science"
	BookCategoryTechnology BookCategory = "technology"
	BookCategoryArt        BookCategory = "art"
	BookCategoryHistory    BookCategory = "history"
	BookCategoryBiography  BookCategory = "biography"
	BookCategoryChildren   BookCategory = "children"
	BookCategoryEducation  BookCategory = "education"
	BookCategoryReference  BookCategory = "reference"
)

func (e *BookCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BookCategory(s)
	case string:
		*e = BookCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for BookCategory: %T", src)
	}
	return nil
}

type NullBookCategory struct {
	BookCategory BookCategory `json:"book_category"`
	Valid        bool         `json:"valid"` // Valid is true if BookCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBookCategory) Scan(value interface{}) error {
	if value == nil {
		ns.BookCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BookCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBookCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BookCategory), nil
}

type Book struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Metadata    []byte           `json:"metadata"`
	Category    BookCategory     `json:"category"`
	Price       pgtype.Numeric   `json:"price"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}
