package post

import (
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jinzhu/gorm"
)

// Post model
type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Body      string    `json:"body" gorm:"type:text"`
	Thumbnail string    `json:"thumbnail" gorm:"type:varchar(255)"`
	UserID    int       `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (p *Post) table() *gorm.DB {
	if !database.MySQL.HasTable(p) {
		return database.MySQL.CreateTable(p)
	}
	return database.MySQL
}

// Insert post to database
func (p *Post) Insert() error {
	return p.table().Create(p).Error
}

// FindOne post from database
func (p *Post) FindOne(order string, cond interface{}, args ...interface{}) error {
	return p.table().Where(cond, args...).Order(order).First(p).Error
}

// Find posts from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Post, error) {
	p := &Post{}
	posts := &[]Post{}
	err := p.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(posts).Error
	return posts, err
}

// Save post from database
func (p *Post) Save() error {
	return p.table().Save(p).Error
}

// Delete post from database
func (p *Post) Delete() error {
	return p.table().Delete(p).Error
}
