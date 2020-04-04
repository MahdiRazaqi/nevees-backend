package post

import (
	"context"
	"errors"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post model
type Post struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
	User    primitive.ObjectID `bson:"_user" json:"_user"`
	Tags    []string           `bson:"tags" json:"tags"`
	Created time.Time          `bson:"created" json:"created"`
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

// UpdateOne post
func (p *Post) UpdateOne(filter bson.M) error {
	_, err := p.collection().UpdateOne(context.Background(), filter, bson.M{"$set": database.ConvertToBson(p)})
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

// Find post
func Find(filter bson.M, page, limit int) ([]Post, error) {
	p := new(Post)
	ctx := context.Background()
	options := options.Find()
	options.SetLimit(int64(limit))
	if page > 0 {
		options.SetSkip(int64((page - 1) * limit))
	}

	cursor, err := p.collection().Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for cursor.Next(ctx) {
		p := new(Post)
		if err := cursor.Decode(p); err != nil {
			continue
		}
		posts = append(posts, *p)
	}

	return posts, nil
}
