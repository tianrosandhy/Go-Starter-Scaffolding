package baseentity

import (
	"time"

	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

type Base struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`
}

// BeforeCreate Data
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	now := strfmt.DateTime(time.Now())

	b.UpdatedAt = now
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}

	return nil
}

// BeforeUpdate Data
func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	return b.BeforeCreate(tx)
}

// BeforeSave Data
func (b *Base) BeforeSave(tx *gorm.DB) error {
	return b.BeforeCreate(tx)
}
