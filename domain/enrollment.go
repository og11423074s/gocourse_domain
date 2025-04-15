package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Enrollment struct {
	ID        string           `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string           `json:"user_id,omitempty" gorm:"type:char(36);"`
	CourseID  string           `json:"course_id,omitempty" gorm:"type:char(36); not null"`
	Status    EnrollmentStatus `json:"status" gorm:"type:char(2);"`
	CreatedAt *time.Time       `json:"-"`
	UpdatedAt *time.Time       `json:"-"`
}

type EnrollmentStatus string

const (
	Pending  EnrollmentStatus = "P"
	Active   EnrollmentStatus = "A"
	Studying EnrollmentStatus = "S"
	Inactive EnrollmentStatus = "I"
)

func (e *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {

	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return
}
