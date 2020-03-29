package user

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/connection"
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

func (u *User) collectionName() string {
	return "user"
}

func (u *User) collection() *mongo.Collection {
	return connection.MongoDB.Collection(u.collectionName())
}

func (u *User) bson() bson.M {
	val, _ := bson.Marshal(u)
	data := new(bson.M)
	bson.Unmarshal(val, data)
	return *data
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
	_, err := u.collection().InsertOne(context.Background(), u.bson())
	return err
}

// LoadByEmail for find user by email address
func LoadByEmail(email string) (*User, error) {
	return FindOne(bson.M{"email": email})
}
