package tag

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Tag model
type Tag struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:datetime"`
	PostID     int       `json:"post_id" gorm:"type:int;foreignkey;not null"`
	CategoryID int       `json:"category_id" gorm:"type:int;foreignkey;not null"`
}

func (t *Tag) table() *gorm.DB {
	if !database.MySQL.HasTable(t) {
		return database.MySQL.CreateTable(t)
	}
	return database.MySQL
}

// Insert tag to database
func (t *Tag) Insert() error {
	return t.table().Create(t).Error
}

// FindOne tag from database
func (t *Tag) FindOne(order string, cond interface{}, args ...interface{}) error {
	return t.table().Where(cond, args...).Order(order).First(t).Error
}

// Find tags from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Tag, error) {
	t := &Tag{}
	tags := &[]Tag{}
	err := t.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(tags).Error
	return tags, err
}

// Save tag from database
func (t *Tag) Save() error {
	return t.table().Save(t).Error
}

// Delete tag from database
func Delete(cond interface{}, args ...interface{}) error {
	t := &Tag{}
	return t.table().Where(cond, args...).Delete(t).Error
}
