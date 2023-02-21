package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //"-" means ignore this field
	// when model included gorm.DeletedAt field,  it will get soft delete ability automatically!
	// When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt‘s value to the current time, and the data is not findable with normal Query methods anymore.
}
