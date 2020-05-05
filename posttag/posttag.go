package posttag

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Tag model
type Posttag struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	TagID     uint      `json:"tag_id" gorm:"type:varchar(255)"`
	PostID    uint      `json:"post_id" gorm:"type:varchar(255)"`
}

func (p *Posttag) table() *gorm.DB {
	if !database.MySQL.HasTable(p) {
		return database.MySQL.Model(p).CreateTable(p)
	}
	return database.MySQL.Model(p)
}

// Insert posttags to database
func (p *Posttag) Insert() error {
	return p.table().Create(p).Error
}
