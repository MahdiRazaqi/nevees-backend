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

// Insert user to database
func (u *User) Insert() error {
	u.ID = primitive.NewObjectID()
	u.Created = time.Now()
	_, err := u.collection().InsertOne(context.TODO(), u.bson())
	return err
}
