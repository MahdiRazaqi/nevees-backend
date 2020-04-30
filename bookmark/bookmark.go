package bookmark

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Bookmark model
type Bookmark struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	PostID    uint      `json:"post_id" gorm:"type:int;foreignkey;not null"`
	UserID    uint      `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (b *Bookmark) table() *gorm.DB {
	if !database.MySQL.HasTable(b) {
		return database.MySQL.Model(b).CreateTable(b)
	}
	return database.MySQL.Model(b)
}

// Insert bookmark to database
func (b *Bookmark) Insert() error {
	return b.table().Create(b).Error
}

// Delete bookmark from database
func Delete(cond interface{}, args ...interface{}) error {
	b := &Bookmark{}
	query := b.table().Where(cond, args...).Delete(b)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}
