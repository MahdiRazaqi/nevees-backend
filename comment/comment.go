package comment

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neveesco/nevees-backend/database"
)

// Comment model
type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Body      string    `json:"body" gorm:"type:text"`
	PostID    uint      `json:"post_id" gorm:"type:int;foreignkey;not null"`
	UserID    uint      `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (c *Comment) table() *gorm.DB {
	if !database.MySQL.HasTable(c) {
		return database.MySQL.Model(c).CreateTable(c)
	}
	return database.MySQL.Model(c)
}

// Insert Comment to database
func (c *Comment) Insert() error {
	return c.table().Create(c).Error
}

// Find comments from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Comment, error) {
	c := &Comment{}
	comments := &[]Comment{}
	err := c.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(comments).Error
	return comments, err
}
