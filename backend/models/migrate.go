package models
import (
	"gorm.io/gorm"
)
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Organization{},
		&User{},
		&Hostel{},
		&Floor{},
		&Room{},
		&Bed{},
		&Student{},
		&Allocation{},
		&Rent{},
		&Payment{},
		&NotificationLog{},
	)
}
