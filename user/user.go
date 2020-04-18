package user

import (
	"time"

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

// // Insert user to database
// func (u *xx) Insert() error {
// 	q, err := database.MySQL.Prepare("INSERT INTO users (username) VALUES (?)")
// 	q.Exec("test")
// 	defer q.Close()

// 	return err
// }

// // // Insert user to database
// // func (u *User) Insert() error {
// // 	u.ID = primitive.NewObjectID()
// // 	u.Created = time.Now()
// // 	_, err := u.collection().InsertOne(context.Background(), database.ConvertToBson(u))
// // 	return err
// // }

// func (u *User) collection() *mongo.Collection {
// 	return database.MongoDB.Collection("user")
// }

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

// // AuthByUserPass authenticate user with username and password
// func AuthByUserPass(username, password string) (*User, error) {
// 	authError := errors.New("username or password not matched")

// 	u, err := LoadByUsername(username)
// 	if err != nil {
// 		return nil, authError
// 	}

// 	if !passwd.Check(password, u.Password) {
// 		return nil, authError
// 	}

// 	return u, nil
// }

// // LoadByUsername load user from username
// func LoadByUsername(username string) (*User, error) {
// 	return FindOne(bson.M{"username": username})
// }

// // FindOne fine one user from database
// func FindOne(filter bson.M) (*User, error) {
// 	u := new(User)
// 	if err := u.collection().FindOne(context.Background(), filter).Decode(u); err != nil {
// 		return nil, err
// 	}
// 	return u, nil
// }

// // Insert user to database
// func (u *User) Insert() error {
// 	u.ID = primitive.NewObjectID()
// 	u.Created = time.Now()
// 	_, err := u.collection().InsertOne(context.Background(), database.ConvertToBson(u))
// 	return err
// }

// // // Insert user to database
// // func (u *User) Insert() error {
// // 	test := &User{}
// // 	s := table().Create(u).Find(&test)
// // 	// _, err := u.collection().InsertOne(context.Background(), database.ConvertToBson(u))
// // 	// return err
// // 	fmt.Println(test)
// // 	fmt.Println(s)
// // 	return errors.New("sdfjsfdlkjsdkf")
// // }
