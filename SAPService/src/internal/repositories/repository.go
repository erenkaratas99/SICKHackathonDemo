package repositories

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{collection: mc}
}

func (r *Repository) WriteKeyword(kw string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	id := uuid.New()
	timeNow := time.Now().Format(time.RFC3339)
	c := bson.M{
		"_id":       id.String(),
		"keyword":   kw,
		"createdAt": timeNow,
	}
	_, err := r.collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
