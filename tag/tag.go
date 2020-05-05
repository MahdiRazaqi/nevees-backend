package tag

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Tag model
type Tag struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
}

func (t *Tag) table() *gorm.DB {
	if !database.MySQL.HasTable(t) {
		return database.MySQL.Model(t).CreateTable(t)
	}
	return database.MySQL.Model(t)
}

// Insert Tag to database
func (t *Tag) Insert() error {
	return t.table().Create(t).Error
}

// Find tags from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Tag, error) {
	t := &Tag{}
	tags := &[]Tag{}
	err := t.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(tags).Error
	return tags, err
}
