package post

import (
	"context"
	"errors"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post model
type Post struct {
	ID      primitive.ObjectID   `bson:"_id"`
	Title   string               `bson:"title"`
	Content string               `bson:"content"`
	User    primitive.ObjectID   `bson:"_user"`
	Tags    []primitive.ObjectID `bson:"_tags"`
	Created time.Time            `bson:"created"`
}

func (p *Post) collection() *mongo.Collection {
	return database.MongoDB.Collection("post")
}

// InsertOne post to database
func (p *Post) InsertOne() error {
	p.ID = primitive.NewObjectID()
	p.Created = time.Now()
	_, err := p.collection().InsertOne(context.Background(), database.ConvertToBson(p))
	return err
}

// DeleteOne post
func DeleteOne(filter bson.M) error {
	p := new(Post)
	count, err := p.collection().DeleteOne(context.Background(), filter)
	if count.DeletedCount == 0 && err == nil {
		return errors.New("document cloud not be deleted")
	}
	return err
}

// FindOne post
func FindOne(filter bson.M) (*Post, error) {
	p := new(Post)
	if err := p.collection().FindOne(context.Background(), filter).Decode(p); err != nil {
		return nil, err
	}
	return p, nil
}
