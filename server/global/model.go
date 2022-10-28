package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //"-" means ignore this field
	// when model included gorm.DeletedAt field,  it will get soft delete ability automatically!
	// When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt‘s value to the current time, and the data is not findable with normal Query methods anymore.
}
