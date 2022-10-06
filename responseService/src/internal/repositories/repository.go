package repositories

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{collection: mc}
}
