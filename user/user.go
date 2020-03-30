package user

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/connection"
	"github.com/dgrijalva/jwt-go"
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

// CreateToken generate new token
func (u *User) CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("secret-nevees"))
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

// LoadByUsername find user by username
func LoadByUsername(username string) (*User, error) {
	return FindOne(bson.M{"username": username})
}
