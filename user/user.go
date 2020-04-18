package user

import (
	"errors"
	"time"

	"github.com/jeyem/passwd"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
)

// User model
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
	Username  string    `json:"username" gorm:"type:varchar(255)"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Role      string    `json:"role" gorm:"type:varchar(255)"`
}

func (u *User) table() *gorm.DB {
	if !database.MySQL.HasTable(u) {
		return database.MySQL.CreateTable(u)
	}
	return database.MySQL
}

// Insert user to database
func (u *User) Insert() error {
	return u.table().Create(u).Error
}

// FindOne user from database
func (u *User) FindOne(order string, cond interface{}, args ...interface{}) error {
	return u.table().Where(cond, args...).Order(order).First(u).Error
}

// LoadByUsername load user from username
func (u *User) LoadByUsername(username string) error {
	return u.FindOne("", "username = ?", username)
}

// AuthByUserPass authenticate user with username and password
func (u *User) AuthByUserPass(username, password string) error {
	err := errors.New("username or password not matched")

	if u.LoadByUsername(username) != nil {
		return err
	}

	if !passwd.Check(password, u.Password) {
		return err
	}
	return nil
}

// Mini user data
func (u *User) Mini() bson.M {
	return bson.M{
		"id":         u.ID,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
		"username":   u.Username,
		"fullname":   u.Fullname,
		"email":      u.Email,
		"role":       u.Role,
	}
}
