package user

import (
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database/mysql"
)

// User model
type User struct {
	ID       int       `json:"id" mysql:"id INT(255) NOT NULL UNIQUE PRIMARY KEY AUTO_INCREMENT"`
	Created  time.Time `json:"created" mysql:"created VARCHAR(255)"`
	Username string    `json:"username" mysql:"username VARCHAR(255) NOT NULL UNIQUE"`
	Fullname string    `json:"fullname" mysql:"fullname VARCHAR(255)"`
	Email    string    `json:"email" mysql:"email VARCHAR(255) NOT NULL UNIQUE"`
	Password string    `json:"password" mysql:"password VARCHAR(255) NOT NULL"`
	Role     string    `json:"role" mysql:"role VARCHAR(255)"`
}

func (u *User) table() *mysql.Database {
	return mysql.MySQL.Table("users", User{})
}

// Insert user to Database
func (u *User) Insert() error {
	u.Created = time.Now()
	return u.table().Insert(*u)
}

// Find user to Database
func (u *User) Find() error {
	return u.table().Select(u)
}

// // User model
// type User struct {
// 	ID       primitive.ObjectID `bson:"_id" json:"_id"`
// 	Username string             `bson:"username" json:"username"`
// 	Fullname string             `bson:"fullname" json:"fullname"`
// 	Email    string             `bson:"email" json:"email"`
// 	Password string             `bson:"password" json:"password"`
// 	Created  time.Time          `bson:"created" json:"created"`
// }

// // // User model
// // type User struct {
// // 	gorm.Model
// // 	Username string `gorm:"column:username"`
// // 	Fullname string `gorm:"column:fullname"`
// // 	Email    string `gorm:"column:email"`
// // 	Password string `gorm:"column:password"`
// // 	Role     string `gorm:"column:role"`
// // }

// // func table() *gorm.DB {
// // 	if !database.MySQL.HasTable(&User{}) {
// // 		return database.MySQL.Table("users").CreateTable(&User{})
// // 	}
// // 	return database.MySQL.Table("users")
// // }

// // User model
// type xx struct {
// 	ID       int
// 	Created  *time.Time
// 	Username string
// 	Fullname string
// 	Email    string
// 	Password string
// 	Role     string
// }

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

// // Mini user data
// func (u *User) Mini() bson.M {
// 	return bson.M{
// 		"id":       u.ID,
// 		"username": u.Username,
// 		"fullname": u.Fullname,
// 		"email":    u.Email,
// 	}
// }

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
