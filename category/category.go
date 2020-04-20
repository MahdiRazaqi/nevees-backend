package category

import (
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jinzhu/gorm"
)

// Category model
type Category struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
}

func (c *Category) table() *gorm.DB {
	if !database.MySQL.HasTable(c) {
		return database.MySQL.CreateTable(c)
	}
	return database.MySQL
}

// Insert category to database
func (c *Category) Insert() error {
	return c.table().Create(c).Error
}

// FindOne category from database
func (c *Category) FindOne(order string, cond interface{}, args ...interface{}) error {
	return c.table().Where(cond, args...).Order(order).First(c).Error
}

// Find categorys from database
func Find(limit int, page int, order string, cond interface{}, args ...interface{}) (*[]Category, error) {
	c := &Category{}
	categorys := &[]Category{}
	err := c.table().Where(cond, args...).Order(order).Limit(limit).Offset(page - 1).Find(categorys).Error
	return categorys, err
}

// Save category from database
func (c *Category) Save() error {
	return c.table().Save(c).Error
}

// Delete category from database
func Delete(cond interface{}, args ...interface{}) error {
	c := &Category{}
	return c.table().Where(cond, args...).Delete(c).Error
}
