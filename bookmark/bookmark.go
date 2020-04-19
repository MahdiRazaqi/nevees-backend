package bookmark

import (
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jinzhu/gorm"
)

// Bookmark model
type Bookmark struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	PostID    int       `json:"post_id" gorm:"type:int;foreignkey;not null"`
	UserID    int       `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (b *Bookmark) table() *gorm.DB {
	if !database.MySQL.HasTable(b) {
		return database.MySQL.CreateTable(b)
	}
	return database.MySQL
}

// Insert bookmark to database
func (b *Bookmark) Insert() error {
	return b.table().Create(b).Error
}

// FindOne bookmark from database
func (b *Bookmark) FindOne(order string, cond interface{}, args ...interface{}) error {
	return b.table().Where(cond, args...).Order(order).First(b).Error
}

// Find bookmarks from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Bookmark, error) {
	b := &Bookmark{}
	bookmarks := &[]Bookmark{}
	err := b.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(bookmarks).Error
	return bookmarks, err
}

// Save bookmark from database
func (b *Bookmark) Save() error {
	return b.table().Save(b).Error
}

// Delete bookmark from database
func (b *Bookmark) Delete() error {
	return b.table().Delete(b).Error
}
