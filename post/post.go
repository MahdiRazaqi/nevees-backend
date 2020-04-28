package post

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Post model
type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Body      string    `json:"body" gorm:"type:text"`
	Thumbnail string    `json:"thumbnail" gorm:"type:varchar(255)"`
	UserID    uint      `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (p *Post) table() *gorm.DB {
	if !database.MySQL.HasTable(p) {
		return database.MySQL.Model(p).CreateTable(p)
	}
	return database.MySQL.Model(p)
}

// Insert post to database
func (p *Post) Insert() error {
	return p.table().Create(p).Error
}

// FindOne post from database
func (p *Post) FindOne(cond interface{}, args ...interface{}) error {
	return p.table().Where(cond, args...).First(p).Error
}

// Find posts from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Post, error) {
	p := &Post{}
	posts := &[]Post{}
	err := p.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(posts).Error
	return posts, err
}

// Update post from database
func (p *Post) Update(cond interface{}, args ...interface{}) error {
	query := p.table().Where(cond, args...).Updates(p)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}

// Delete post from database
func Delete(cond interface{}, args ...interface{}) error {
	p := &Post{}
	query := p.table().Where(cond, args...).Delete(p)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}
