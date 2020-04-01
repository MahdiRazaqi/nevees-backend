package tag

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tag model
type Tag struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Name    string             `bson:"name" json:"name"`
	Created time.Time          `bson:"created" json:"created"`
}

func (t *Tag) collection() *mongo.Collection {
	return database.MongoDB.Collection("tag")
}

// InsertOne tag to database
func (t *Tag) InsertOne() error {
	t.ID = primitive.NewObjectID()
	t.Created = time.Now()
	_, err := t.collection().InsertOne(context.Background(), database.ConvertToBson(t))
	return err
}

// FindOne tag
func FindOne(filter bson.M) (*Tag, error) {
	t := new(Tag)
	if err := t.collection().FindOne(context.Background(), filter).Decode(t); err != nil {
		return nil, err
	}
	return t, nil
}
