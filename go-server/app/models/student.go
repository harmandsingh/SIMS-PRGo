package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID uuid.UUID `json:"id" gorm:"type:uuid;not null"`
	FirstName string `json:"firstname" gorm:"text;not null;default:null"`
	LastName string `json:"lastname" gorm:"text;not null;default:null"`
}