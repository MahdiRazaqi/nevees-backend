package user

import (
	"context"
	"errors"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/jeyem/passwd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Fullname string             `bson:"fullname"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Created  time.Time          `bson:"created"`
}

func (u *User) collection() *mongo.Collection {
	return database.MongoDB.Collection("user")
}

// Mini user data
func (u *User) Mini() bson.M {
	return bson.M{
		"id":       u.ID,
		"username": u.Username,
		"fullname": u.Fullname,
		"email":    u.Email,
	}
}

// AuthByUserPass authenticate user with username and password
func AuthByUserPass(username, password string) (*User, error) {
	authError := errors.New("username or password not matched")

	u, err := LoadByUsername(username)
	if err != nil {
		return nil, authError
	}

	if !passwd.Check(password, u.Password) {
		return nil, authError
	}

	return u, nil
}

// LoadByUsername load user from username
func LoadByUsername(username string) (*User, error) {
	return FindOne(bson.M{"username": username})
}

// FindOne fine one user from database
func FindOne(filter bson.M) (*User, error) {
	u := new(User)
	if err := u.collection().FindOne(context.Background(), filter).Decode(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Insert user to database
func (u *User) Insert() error {
	u.ID = primitive.NewObjectID()
	u.Created = time.Now()
	_, err := u.collection().InsertOne(context.Background(), database.ConvertToBson(u))
	return err
}
