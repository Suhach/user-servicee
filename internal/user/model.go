package user

import "time"

type User struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(255);"`
	Pass  string `gorm:"type:varchar(255);"`
	CreatedAt time.Time  `gorm:"type:timestamptz;"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;"`
	DeletedAt *time.Time `gorm:"type:timestamptz;"`
}
