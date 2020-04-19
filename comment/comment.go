package comment

import (
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jinzhu/gorm"
)

// Comment model
type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Body      string    `json:"body" gorm:"type:text"`
	PostID    int       `json:"post_id" gorm:"type:int;foreignkey;not null"`
	UserID    int       `json:"user_id" gorm:"type:int;foreignkey;not null"`
}

func (c *Comment) table() *gorm.DB {
	if !database.MySQL.HasTable(c) {
		return database.MySQL.CreateTable(c)
	}
	return database.MySQL
}

// Insert comment to database
func (c *Comment) Insert() error {
	return c.table().Create(c).Error
}

// FindOne comment from database
func (c *Comment) FindOne(order string, cond interface{}, args ...interface{}) error {
	return c.table().Where(cond, args...).Order(order).First(c).Error
}

// Find comments from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Comment, error) {
	c := &Comment{}
	comments := &[]Comment{}
	err := c.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(comments).Error
	return comments, err
}

// Save comment from database
func (c *Comment) Save() error {
	return c.table().Save(c).Error
}

// Delete comment from database
func (c *Comment) Delete() error {
	return c.table().Delete(c).Error
}
