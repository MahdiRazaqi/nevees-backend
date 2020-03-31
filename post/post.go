package post

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post model
type Post struct {
	ID        primitive.ObjectID   `bson:"_id"`
	Title     string               `bson:"title"`
	Content   string               `bson:"content"`
	User      primitive.ObjectID   `bson:"user"`
	LikeCount int                  `bson:"like_count"`
	Tags      []primitive.ObjectID `bson:"tags"`
	Created   time.Time            `bson:"created"`
}

func (p *Post) collection() *mongo.Collection {
	return database.MongoDB.Collection("post")
}

// Insert post to database
func (p *Post) Insert() error {
	p.ID = primitive.NewObjectID()
	p.Created = time.Now()
	_, err := p.collection().InsertOne(context.Background(), database.ConvertToBson(p))
	return err
}
